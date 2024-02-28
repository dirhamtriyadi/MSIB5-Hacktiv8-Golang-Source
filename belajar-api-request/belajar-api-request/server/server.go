package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, fileHeader, err := r.FormFile("image-file")

		if err != nil {
			fmt.Printf("error while creating FormFile instance: %s\n", err.Error())
			w.Write([]byte(err.Error()))
			return
		}

		newFile, err := os.Create(fileHeader.Filename)

		if err != nil {
			fmt.Printf("error while creating the imageFile: %s\n", err.Error())
			w.Write([]byte(err.Error()))
			return
		}

		bs, err := ioutil.ReadAll(file)

		if err != nil {
			fmt.Printf("error while reading the imageFile: %s\n", err.Error())
			w.Write([]byte(err.Error()))
			return
		}

		_, _ = bs, newFile

		_, err = newFile.Write(bs)
	})

	fmt.Println("Listening on PORT: 8080")
	http.ListenAndServe(":8080", nil)
}
