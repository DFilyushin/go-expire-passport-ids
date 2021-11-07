package domain

import (
	"encoding/csv"
	protoMessage "expired-passport-checker/internal/domain/proto"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

const (
	MaxPassportIdLen = 10
	BytesArrayLen = 125
)

type Matrix struct {
	items       map[int32][]byte
	stringItems map[string]bool
}

func (m *Matrix) getStringItems() []string {
	/*Получить список строк*/
	var result []string
	for k, _ := range m.stringItems {
		result = append(result, k)
	}
	return result
}

func NewMatrix() *Matrix {
	/*Конструктор*/
	m := new(Matrix)
	m.items = make(map[int32][]byte)
	m.stringItems = make(map[string]bool, 0)
	return m
}

func (m *Matrix) isNumeric(item string) bool {
	/*Проверка числа*/
	for _, char := range item {
		if char < 48 || char > 57 {
			return false
		}
	}
	return true
}

func (m *Matrix) addStringItem(item string) {
	/*Добавить строку*/
	m.stringItems[item] = true
}

func (m *Matrix) clearStringItems() {
	/*Очистка списка элементов*/
	for k := range m.stringItems {
		delete(m.stringItems, k)
	}
}

func (m *Matrix) clearByteMatrix() {
	/*Очистка матрицы*/
	for k := range m.items {
		delete(m.items, k)
	}
}

func (m *Matrix) splitItemByTwoInt(item string) (int32, int16) {
	/*Разделение строки на два числа*/
	value, _ := strconv.Atoi(item[:7])
	firstItem := int32(value)

	value, _ = strconv.Atoi(item[7:])
	secondItem := int16(value)
	return firstItem, secondItem
}

func (m *Matrix) isCorrectPassportId(item string) bool {
	return m.isNumeric(item) && len(item) == MaxPassportIdLen
}

func (m *Matrix) setArrayBit(array []byte, value int16) []byte {
	/*Установить бит в массиве байтов*/
	numByte := value / 8
	numBite := value % 8
	array[numByte] = array[numByte] + byte(1<<numBite)
	return array
}

func (m *Matrix) checkArrayBit(array []byte, value int16) bool {
	/*Проверить наличие бита в массиве*/
	numByte := value / 8
	numBite := value % 8
	currentByte := byte(1 << numBite)
	return currentByte == currentByte&array[numByte]
}

func (m *Matrix) AddItemToMatrix(item string) {
	/*Добавление элемента в матрицу*/
	if !m.isCorrectPassportId(item) {
		m.addStringItem(item)
		return
	}
	firstItem, secondItem := m.splitItemByTwoInt(item)

	bitArray, ok := m.items[firstItem]
	if !ok {
		bitArray = make([]byte, BytesArrayLen)
	}

	m.items[firstItem] = m.setArrayBit(bitArray, secondItem)
}

func (m *Matrix) FindItemInMatrix(item string) bool {
	/*Поиск элемента в матрице*/
	if !m.isCorrectPassportId(item) {
		_, isOk := m.stringItems[item]
		return isOk
	}

	firstItem, secondItem := m.splitItemByTwoInt(item)
	bitArray, isOk := m.items[firstItem]
	if !isOk {
		return false
	}

	return m.checkArrayBit(bitArray, secondItem)
}

func (m *Matrix) LoadFromCsvFile(f *csv.Reader, skipHeader bool) (int, error) {
	/*Чтение из csv файла данных*/
	var processedLines int = 0

	for {
		line, err := f.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}

		if processedLines == 0 && skipHeader {
			processedLines++
			continue
		}

		fullLine := line[0] + line[1]
		m.AddItemToMatrix(fullLine)

		processedLines++
	}
	return processedLines, nil
}

func (m *Matrix) SaveMatrixToDisk(fileName string) error {
	/*Сохранить матрицу в файл*/

	numbersMap := make([]*protoMessage.NumbersMap, 0)

	for k, v := range m.items {
		numberMapItem := protoMessage.NumbersMap{
			SevenDigitsKey:       k,
			ThreeDigitsBitsValue: v}
		numbersMap = append(numbersMap, &numberMapItem)
	}

	var pdm = protoMessage.PassportDataMessage{
		CsvHeader:      "Matrix",
		OtherLines:     m.getStringItems(),
		NumbersOnlyMap: numbersMap,
	}

	out, err := proto.Marshal(&pdm)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(fileName, out, 0644); err != nil {
		return err
	}

	return nil
}

func (m *Matrix) LoadMatrixFromDisk(fileName string) error {
	/*Загрузить матрицу из файла*/

	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	passportData := &protoMessage.PassportDataMessage{}
	if err := proto.Unmarshal(in, passportData); err != nil {
		log.Fatalln("Failed to parse passport data:", err)
	}

	m.clearStringItems()
	m.clearByteMatrix()

	for _, item := range passportData.OtherLines {
		m.stringItems[item] = true
	}

	for _, item := range passportData.NumbersOnlyMap {
		m.items[item.SevenDigitsKey] = item.ThreeDigitsBitsValue
	}

	return nil
}
