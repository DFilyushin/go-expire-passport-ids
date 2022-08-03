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
	BytesArrayLen    = 125
)

type Matrix struct {
	items       map[int32][]byte
	stringItems map[string]bool
}

// getStringItems Получить список строк
func (m *Matrix) getStringItems() []string {
	var result []string
	for k, _ := range m.stringItems {
		result = append(result, k)
	}
	return result
}

// NewMatrix Конструктор
func NewMatrix() *Matrix {
	m := new(Matrix)
	m.items = make(map[int32][]byte)
	m.stringItems = make(map[string]bool, 0)
	return m
}

// isNumeric Проверка числа
func (m *Matrix) isNumeric(item string) bool {
	for _, char := range item {
		if char < 48 || char > 57 {
			return false
		}
	}
	return true
}

//addStringItem Добавление строки
func (m *Matrix) addStringItem(item string) {
	m.stringItems[item] = true
}

//clearStringItems Очистка списка элементов
func (m *Matrix) clearStringItems() {
	for k := range m.stringItems {
		delete(m.stringItems, k)
	}
}

//clearByteMatrix Очистка матрицы
func (m *Matrix) clearByteMatrix() {
	for k := range m.items {
		delete(m.items, k)
	}
}

//splitItemByTwoInt Разделение строки на два числа
func (m *Matrix) splitItemByTwoInt(item string) (int32, int16) {
	value, _ := strconv.Atoi(item[:7])
	firstItem := int32(value)

	value, _ = strconv.Atoi(item[7:])
	secondItem := int16(value)
	return firstItem, secondItem
}

//isCorrectPassportId Проверка корректности пасппорта
func (m *Matrix) isCorrectPassportId(item string) bool {
	return m.isNumeric(item) && len(item) == MaxPassportIdLen
}

//setArrayBit Установить бит в массиве байтов
func (m *Matrix) setArrayBit(array []byte, value int16) []byte {
	numByte := value / 8
	numBite := value % 8
	array[numByte] = array[numByte] + byte(1<<numBite)
	return array
}

//checkArrayBit Проверить наличие бита в массиве
func (m *Matrix) checkArrayBit(array []byte, value int16) bool {
	numByte := value / 8
	numBite := value % 8
	currentByte := byte(1 << numBite)
	return currentByte == currentByte&array[numByte]
}

// AddItemToMatrix Добавление элемента в матрицу
func (m *Matrix) AddItemToMatrix(item string) {
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

//FindItemInMatrix Поиск элемента в матрице
func (m *Matrix) FindItemInMatrix(item string) bool {
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

//LoadFromCsvFile Чтение из csv файла данных
func (m *Matrix) LoadFromCsvFile(f *csv.Reader, skipHeader bool) (int, error) {
	var processedLines = 0

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

//SaveMatrixToDisk Сохранить матрицу в файл
func (m *Matrix) SaveMatrixToDisk(fileName string) error {
	numbersMap := make([]*protoMessage.NumbersMap, 0)

	for k, v := range m.items {
		numberMapItem := protoMessage.NumbersMap{
			SevenDigitsKey:       k,
			ThreeDigitsBitsValue: v}
		numbersMap = append(numbersMap, &numberMapItem)
	}

	var message = protoMessage.PassportDataMessage{
		CsvHeader:      "Matrix",
		OtherLines:     m.getStringItems(),
		NumbersOnlyMap: numbersMap,
	}

	out, err := proto.Marshal(&message)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(fileName, out, 0644); err != nil {
		return err
	}

	return nil
}

//LoadMatrixFromDisk  Загрузить матрицу из файла
func (m *Matrix) LoadMatrixFromDisk(fileName string) error {
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
