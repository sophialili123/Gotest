package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
)

func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error{
	path := request.URL.Path[len("/list/"):]
	panic(path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
