package dishHandler

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/utilities"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	//err = os.MkdirAll("./uploads", os.ModePerm)
	//
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		return
	}
	create, err := os.Create("tmp.csv")
	if err != nil {
		return
	}

	_, err = io.Copy(create, buf)
	if err != nil {
		return
	}
	csvFile, err := os.Open("./tmp.csv")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}(csvFile)
	data, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	dishList := utilities.CreateDishList(data)
	err = helper.CreateBulkDishes(dishList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
