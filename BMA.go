package main

import(
	"fmt"
	"strings"
	"log"
	"io/ioutil"
	"flag"
)

type word struct{
	original string
	upper string
	lower string
}

func BMA(){
	filePath := flag.String("path","", "Path of file to be changed")
	fileExtension := flag.String("FE","","File extension for new file to be changed")
	flag.Parse()

	if *filePath == ""{
		log.Fatal("No filepath entered")
	}

	fileText , err := parseFile(*filePath)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Print("Enter desired word to replace:")
	var input string
	fmt.Scan(&input)

	currentWord := word{
		original: input,
		upper: strings.ToUpper(input),
		lower: strings.ToLower(input),
	}

	fmt.Print("Entered new desired word:")
	fmt.Scan(&input)

	newWord := word{
		original: input,
	}

	tempFileText := strings.Split(fileText, "\n") // will splitup into array of lines
	instanceCount := 0
	for i , val := range tempFileText{
		line := strings.Split(val, " ")
		for j , val2 := range line{
			//make function that will scan all three instances of word
			if strings.ToLower(val2) == currentWord.lower{
				line[j] = newWord.original
				instanceCount ++
			}
		}
		newLine := strings.Join(line, " ")
		tempFileText[i] = newLine
	}
	newFile := strings.Join(tempFileText, "\n")

	if instanceCount == 0{
		log.Fatal("No instance of the word \"",currentWord.original,"\" found")
	}else{
		fmt.Println("There were ",instanceCount," times the word ",currentWord.original," came up")
	}

	// need to trim the filepath 

	newFilePath := trimFilePath(*filePath) + *fileExtension

	if *fileExtension == ""{
		newFilePath = trimFilePath(*filePath) + getExtension()
	}

	err = ioutil.WriteFile(newFilePath,[]byte(newFile),0644)

	if err != nil{
		log.Fatal("couldnt do final write in path")
	}
}



