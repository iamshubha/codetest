package main

import (
	util "codetest/config"
	f "codetest/pkg/internals"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

func main() {

	db := util.GetDB()
	util.InitDB(db)
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/uploadFile", f.UploadFile).Methods(http.MethodPost)
	router.HandleFunc("/uploadFileWithPath", f.UploadFileWithPath).Methods(http.MethodPost)
	router.HandleFunc("/getFilesList", f.GetFilesList).Methods(http.MethodGet)
	router.HandleFunc("/deleteFile", f.DeleteFile).Methods(http.MethodGet)
	router.HandleFunc("/getFilesByFolderName", f.GetFilesByFolderName).Methods(http.MethodGet)
	router.HandleFunc("/getFilesByFileName", f.GetFilesByFileName).Methods(http.MethodGet)
	router.HandleFunc("/getAllFiles", f.GetAllFiles).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)
}
