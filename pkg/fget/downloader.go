package fget

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type FGet struct{}

func (fg *FGet) DownloadFile(URL string) (fileName string, err error) {
	/*Скачать файл по адресу URL во временную папку и вернуть имя файла */

	data, err := http.Get(URL)

	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := ioutil.TempFile(os.TempDir(), "tmp-")
	if err != nil {
		return
	}

	fileName = f.Name()

	defer data.Body.Close()
	defer f.Close()

	src := &ProgressReader{Reader: data.Body, total: data.ContentLength}
	_, err = io.Copy(f, src)
	if err != nil {
		return
	}

	return
}
