package service

import (
	"compress/bzip2"
	"encoding/csv"
	"expired-passport-checker/internal/domain"
	"fmt"
	"os"
)

type PassportIdService struct {
	matrix   *domain.Matrix
	dataFile string
}

//NewPassportIdService Конструктор сервиса
func NewPassportIdService() *PassportIdService {
	ps := new(PassportIdService)
	ps.matrix = domain.NewMatrix()
	return ps
}

// loadFromDataFile Загрузка с диска
func (ps *PassportIdService) loadFromDataFile() error {
	file, err := os.OpenFile(ps.dataFile, os.O_RDONLY, 0644)
	if os.IsNotExist(err) {
		return os.ErrNotExist
	}

	defer file.Close()
	err = ps.matrix.LoadMatrixFromDisk(ps.dataFile)
	return err
}

//Init Инициализация сервиса
func (ps *PassportIdService) Init(storageFile string) error {
	ps.dataFile = storageFile
	err := ps.loadFromDataFile()
	if err != nil {
		return err
	}
	return nil
}

// LoadDataFromArchive Загрузить данные из файла архива
func (ps *PassportIdService) LoadDataFromArchive(fileName string) (loaded int, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}

	defer f.Close()

	bzipFile := bzip2.NewReader(f)
	csvFile := csv.NewReader(bzipFile)

	loaded, err = ps.matrix.LoadFromCsvFile(csvFile, true)
	if err != nil {
		return
	}
	return
}

//SaveDataFile сохраняет матрицу данных в файл для последующего использования
func (ps *PassportIdService) SaveDataFile(fileName string) (err error) {
	err = ps.matrix.SaveMatrixToDisk(fileName)
	return
}

//CheckPassport проверяет паспорт
func (ps *PassportIdService) CheckPassport(PassportSeries, PassportNum string) bool {
	passportId := fmt.Sprintf("%s%s", PassportSeries, PassportNum)
	return ps.matrix.FindItemInMatrix(passportId)
}
