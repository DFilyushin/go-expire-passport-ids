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

func NewPassportIdService() *PassportIdService {
	ps := new(PassportIdService)
	ps.matrix = domain.NewMatrix()
	return ps
}

func (ps *PassportIdService) loadFromDataFile() error {
	/*Загрузка с диска*/
	file, err := os.OpenFile(ps.dataFile, os.O_RDONLY, 0644)
	defer file.Close()
	if os.IsNotExist(err) {
		return os.ErrNotExist
	}
	err = ps.matrix.LoadMatrixFromDisk(ps.dataFile)
	return err
}

func (ps *PassportIdService) Init(storageFile string) error {
	/*Инициализация*/
	ps.dataFile = storageFile
	err := ps.loadFromDataFile()
	if err != nil {
		return err
	}
	return nil
}

func (ps *PassportIdService) LoadDataFromArchive(fileName string) (loaded int, err error) {
	/*Загрузить данные из файла csv*/
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

func (ps *PassportIdService) SaveDataFile(fileName string) (err error) {
	/*Сохранить матрицу данных в файл для дальнейшей загрузки*/
	err = ps.matrix.SaveMatrixToDisk(fileName)
	return
}

func (ps *PassportIdService) CheckPassport(PassportSeries, PassportNum string) bool {
	/*Проверка паспорта*/
	passportId := fmt.Sprintf("%s%s", PassportSeries, PassportNum)
	return ps.matrix.FindItemInMatrix(passportId)
}
