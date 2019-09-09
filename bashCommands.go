package main

import(
	"os"
	"log"
)

func moveFile(fileName string){ // this will take in a file and move it from Code to Overlooked
	wd , err := os.Getwd()
	// wd is /Users/benmorehouse/code/BMA

	folderName := fileName // we are gonna put this folder into the scraper folder
	fileName += ".go"
	err = os.Mkdir(folderName, 0700)

	if err != nil{
		log.Fatal(err)
	}

	before := wd + "/" + folderName
	after := "/Users/benmorehouse/Documents/Github/backend/scraper/" + folderName
	err = os.Rename(before, after)

	if err != nil{
		log.Fatal(err)
	}

	before = wd + "/" + fileName
	after =  "/Users/benmorehouse/Documents/Github/backend/scraper/" + folderName + "/" + fileName
	err = os.Rename(before, after)

	if err != nil{
		log.Fatal(err)
	}
}

