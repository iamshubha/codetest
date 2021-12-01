package transaction

import (
	util "codetest/config"
	"codetest/model"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func AddFileInDB(fileName string, filePath string, folderPath string) {
	fmt.Println(fileName, "Data ----", filePath)
	db := util.GetDB()
	defer db.Close()
	if fileName == "" || filePath == "" {
		fmt.Println("Please send correct parameaters")
	}
	sqlQuery := `
	INSERT INTO filesys (filePath, fileName,folderPath)
	VALUES ($1,$2,$3)
	RETURNING id;
	`
	_, err := db.Exec(sqlQuery, filePath, fileName, folderPath)
	if err != nil {
		log.Println("message", "Fail")
		log.Print(err)
	}

	log.Println("message", "success")
}

func DeleteFileFormDatabase(Id int) {
	sqlQuery := `
	DELETE FROM filesys
	WHERE  id = $1;
	`
	db := util.GetDB()
	defer db.Close()

	_, err := db.Exec(sqlQuery, Id)
	if err != nil {
		panic(err)
	}
}

func SearchFileByFolder(folderPAth string) []model.GetData {
	sqlQuery := `
SELECT * FROM filesys WHERE folderPath =$1;
`

	db := util.GetDB()
	defer db.Close()
	dataRows, err := db.Query(sqlQuery, folderPAth)
	if err != nil {
		log.Fatalln(err)
	}
	defer dataRows.Close()
	data := make([]model.GetData, 0)
	for dataRows.Next() {
		detailsModel := model.GetData{}
		dataRows.Scan(&detailsModel.Id, &detailsModel.FileName, &detailsModel.FilePath, &detailsModel.FolderPath)
		log.Println(dataRows)
		data = append(data, detailsModel)
	}
	log.Println(data)
	return data

}

func SearchFileByName(Name string) []model.GetData {
	sqlQuery := `
SELECT * FROM filesys WHERE fileName =$1;
`

	db := util.GetDB()
	defer db.Close()
	dataRows, err := db.Query(sqlQuery, Name)
	if err != nil {
		log.Fatalln(err)
	}
	defer dataRows.Close()
	data := make([]model.GetData, 0)
	for dataRows.Next() {
		detailsModel := model.GetData{}
		dataRows.Scan(&detailsModel.Id, &detailsModel.FileName, &detailsModel.FilePath, &detailsModel.FolderPath)
		log.Println(dataRows)
		data = append(data, detailsModel)
	}
	log.Println(data)
	return data

}

func SearchFiles() []model.GetData {
	sqlQuery := `
SELECT * FROM filesys ;
`

	db := util.GetDB()
	defer db.Close()
	dataRows, err := db.Query(sqlQuery)
	if err != nil {
		log.Fatalln(err)
	}
	defer dataRows.Close()
	data := make([]model.GetData, 0)
	for dataRows.Next() {
		detailsModel := model.GetData{}
		dataRows.Scan(&detailsModel.Id, &detailsModel.FileName, &detailsModel.FilePath, &detailsModel.FolderPath)
		log.Println(dataRows)
		data = append(data, detailsModel)
	}
	log.Println(data)
	return data

}
