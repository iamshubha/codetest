package internals

import (
	util "codetest/config"
	d "codetest/pkg/dbtransaction"

	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"

	_ "github.com/lib/pq"
)

func GetFilesList(w http.ResponseWriter, r *http.Request) {
	fPath := r.FormValue("path")
	finalPath := filepath.Join("./files" + fPath)

	err := ReadDir(finalPath)
	if err != nil {
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(map[string][]string{
		"message": listOfFilePath,
	})
	listOfFilePath = []string{}
	log.Println(listOfFilePath)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create(handler.Filename)
	/*
		// dir := "./files"
		// dst1, err := os.Create(filepath.Join(dir, filepath.Base(handler.Filename))) // dir is directory where you want to save file.
		// filepath.Abs(dst1.Name())
	*/
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	log.Println(path.Base(dst.Name()))

	// Copy the uploaded file to the created file on the filesys
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fullFilePath, _ := filepath.Abs(dst.Name())

	d.AddFileInDB(handler.Filename, fullFilePath, "./files")
	fmt.Fprintf(w, "Successfully Uploaded File\n", fullFilePath)
}

func UploadFileWithPath(w http.ResponseWriter, r *http.Request) {
	x := r.FormValue("file")
	creds := x
	log.Println(creds)
	if creds == "" {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Please send correct parameaters",
		})
		return
	}
	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	path := creds
	if err != nil {
		fmt.Println("Error Retrieving the File", err)
		return
	}

	defer file.Close()
	// Create file
	dir := "./files/" + path
	dst, err := os.Create(filepath.Join(dir, filepath.Base(handler.Filename))) // dir is directory where you want to save file.

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the filesys
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fullFilePath, _ := filepath.Abs(dst.Name())
	d.AddFileInDB(handler.Filename, fullFilePath, dir)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

var listOfFilePath []string

func ReadDir(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	names, _ := file.Readdirnames(0)
	for _, name := range names {
		filePath := fmt.Sprintf("%v/%v", path, name)
		log.Println(name, "data", path)
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}
		// fmt.Println(filePath)
		listOfFilePath = append(listOfFilePath, filePath)
		if fileInfo.IsDir() {
			ReadDir(filePath)
		}
	}
	// log.Println(listOfFilePath)
	return nil
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	fileId := r.FormValue("id")
	filePath := GetDetailsFromDatebase(fileId)
	log.Println("remove file path :", filePath, fileId)
	err := os.Remove(filePath)
	if err != nil {
		fmt.Fprintf(w, "No File Found \n")
		log.Fatal(err)
	}
	id, _ := strconv.Atoi(fileId)
	d.DeleteFileFormDatabase(id)

	fmt.Fprintf(w, "File Deleted Successfully \n")
}
func GetDetailsFromDatebase(Id string) (filePath string) {
	db := util.GetDB()
	defer db.Close()
	id, _ := strconv.Atoi(Id)
	sqlQuery := `
	SELECT filePath FROM filesys WHERE id = $1 ;
	`

	dataRow, err := db.Query(sqlQuery, id)
	if err != nil {
		log.Println(err)
		return
	}
	defer dataRow.Close()
	dataRow.Scan(&filePath)
	for dataRow.Next() {
		dataRow.Scan(&filePath)
	}
	log.Println("filepath /////", filePath)
	return
}

func GetFilesByFolderName(w http.ResponseWriter, r *http.Request) {

	folderPath := r.FormValue("path")
	folderPath = "./files/" + folderPath
	log.Println(folderPath)
	data := d.SearchFileByFolder(folderPath)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Data Found",
		"data":    data,
	})
}

func GetFilesByFileName(w http.ResponseWriter, r *http.Request) {

	fileName := r.FormValue("filename")
	log.Println(fileName)
	data := d.SearchFileByName(fileName)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Data Found",
		"data":    data,
	})
}

func GetAllFiles(w http.ResponseWriter, r *http.Request) {

	data := d.SearchFiles()

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Data Found",
		"data":    data,
	})
}
