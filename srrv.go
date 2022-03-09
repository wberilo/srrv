package main

import (
	"errors"
	"os"
	"strconv"
	"path/filepath"
	"github.com/AlecAivazis/survey/v2"
	"net/http"
)

func main() {
	var files []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	var fileSelect = []*survey.Question{
		{
			Prompt: &survey.Select{
				Message: "Choose a file or a directory",
				Options: files,
			},
			Validate: survey.Required,
		},
	}
	fileSelected := ""
	survey.Ask(fileSelect, &fileSelected)
	//ask which port to use
	var portSelect = []*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Choose a port",
				Default: "8080",
			},
			Validate: func(val interface{}) error {
				port, err := strconv.Atoi(val.(string))
				if err != nil {
					return errors.New("port must be a number")
				}
				if port < 1 || port > 65535 {
					return errors.New("port must be between 1 and 65535")
				}
				return nil
			},
		},
	}

	portSelected := ""
	survey.Ask(portSelect, &portSelected)
	port, _ := strconv.Atoi(portSelected)

	serve_html_file(port, fileSelected)
}

func serve_html_file(port int, file_path string) {
	http.Handle("/", http.FileServer(http.Dir(file_path)))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
	println("Server started on port " + strconv.Itoa(port))
}

