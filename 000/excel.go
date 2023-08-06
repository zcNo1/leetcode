package main

import (
	"github.com/xuri/excelize/v2"
	"log"
)

func main() {
	file, err := excelize.OpenFile("aa.xlsx")
	if err != nil {
		log.Println("open file fail: err: ", err)
		return
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Println("close err: ", err)
			return
		}
	}()

	sheets := file.GetSheetList()
	file.Cols()

}
