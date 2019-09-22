package main

import(
	"log"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"github.com/golang-collections/collections/trie"
	"strings"
	"io/ioutil"
)

func parseFile(filePath string)(string,error){ // will take in content, then return the content
	if filePath ==""{
		err := errors.New("Entered file does not exist\nCheck to make sure that file still exists")
		return "0",err
	}else{
		returnVal , err := ioutil.ReadFile(filePath)
		return string(returnVal),err
	}
}

/*
	getExtension() crawls through wikipedia page with goquery to find
	the file extension that we want the user to have
*/

func getExtension()(string){
	url := "https://en.wikipedia.org/wiki/List_of_programming_languages"
	_ , err := http.Get(url)

	if err != nil{
		log.Fatal("website not valid")
	}

	doc , err := goquery.NewDocument(url)
	if err != nil{
		log.Fatal("couldnt create the document")
	}

	mytrie := trie.New()
	mytrie.Init()

	doc.Find("div.mw-parser-output div.div-col ul li").Each(func(i int, sel *goquery.Selection) {
		_name := sel.Find("a").Text()
		_url , _ := sel.Find("a").Attr("href")
		if _name != "" && _url != ""{
			url = "https://en.wikipedia.org" + _url
			mytrie.Insert(strings.ToLower(_name), _name + " " + url)
		}
	})
	fmt.Println("Which language do you want?")
	var input string
	fmt.Scan(&input)

	input = strings.ToLower(input)

	trieOutput := mytrie.Get(input)
	if trieOutput == nil{
		log.Fatal("we dont have that language")
	}

	output, ok := trieOutput.(string)
	if ok == false{
		log.Fatal("type error for trie interface")
	}

	getURL := strings.Fields(string(output))
	fileExtensionURL , err := getFileExtension(getURL[len(getURL)-1])

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("which file extension would you like")
	for _ , val := range strings.Fields(fileExtensionURL){
		fmt.Println(val)
	}
	found := false

	for found == false{
		fmt.Scan(&input)
		for _ , val := range strings.Fields(fileExtensionURL){
			if strings.ToLower(input) == strings.ToLower(val){
				return val
			}
		}
	}

	log.Fatal("End of program, no file extension found")
	return ""
}

func getFileExtension(_url string)(string, error){
	url := _url
	_ , err := http.Get(url)

	if err != nil{
		return "",err
	}

	doc , err := goquery.NewDocument(url)

	if err != nil{
		return "",err
	}

	extensions := ""
	doc.Find("table.infobox tbody tr").Each(func(i int, sel *goquery.Selection) {
		description := sel.Find("th a").Text()
		if strings.ToLower(string(description)) == "filename extensions"{
			extensions = sel.Find("td").Text()

		}
	})

	if extensions == ""{
		return "" , errors.New("couldnt find extensions")
	}

	trimmer := strings.Fields(extensions)

	for i , val := range trimmer{
		trimmer[i] = strings.Trim(val,",")
	}

	extensions = strings.Join(trimmer, " ")

	return extensions ,err
}

