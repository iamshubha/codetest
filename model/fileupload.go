package model

type DataModel struct {
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
}

type GetData struct {
	Id         int    `json:"id"`
	FilePath   string `json:"filepath"`
	FileName   string `json:"filename"`
	FolderPath string `json:"folderpath"`
}
