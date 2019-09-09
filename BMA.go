package main

import(
	"fmt"
	"os"
	"bufio"
	"flag"
	"strings"
	"log"
)

func main(){
	fmt.Println("                  ***Welcome to Autofill***\n")
	fmt.Println("/***********************************************************************/\n")
	fmt.Println("This is a tool that will be used to parse through files and change the \nnames of the file that we are working on to a new file which will be a new scraper/etc.\n")
	fmt.Println("/***********************************************************************/\n")
	input := "buzzfeed"
	currentWebsite := website{
		original: input,
		upper: strings.ToUpper(input),
		lower: strings.ToLower(input),
	}

	fmt.Print("Enter in the new website scraper : ")
	fmt.Scan(&input)
	err := httpCheck(input)

	if err !=nil{
		fmt.Println("err:",err)
		os.Exit(1)
	}

	f, err := os.Create(input + ".txt") // this is the file writer

	if err !=nil{
		fmt.Println("err:",err)
		os.Exit(1)
	}

	newWebsite := website{
		original: input,
	}

	fptr := flag.String("filepath","template.txt","") // the file pointer needed to do os.Open()
	flag.Parse()
	file, err := os.Open(*fptr) // file is the returned value of os.open. It is what we read through 
	if err != nil{
		fmt.Println("error: file not able to be opened")
	}

	scan := bufio.NewScanner(file)
	for scan.Scan(){ // we read in the string line by line  
		// could scan through and instead just use the ones that we found 
		temp := scan.Text() // this gives us each line
		newstrings := currentWebsite.all_occurances(temp) // newstrings is essentially anything that matches
		for i:=0;i<len(newstrings);i++{
			temp = strings.Replace(temp,newstrings[i],newWebsite.original,-1)
		}
		temp+=string("\n")
		// now we add old into new
		if _, err := f.Write([]byte(temp)); err != nil {
			log.Fatal(err)
		}
	}
		// essentially renames the file rename(old, new)
		//use os.Getwd to get the rooted path
	txtToGo(newWebsite.original)
	// at this point we need to run the file moving because now we have a file to work with 
	moveFile(newWebsite.original)
}
