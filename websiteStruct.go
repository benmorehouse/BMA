package main

import(
	"net/http"
	"os"
)
//structure for the websites that we deal in 

type website struct{
	original string
	upper string
	lower string
}

func (temp website) all_occurances(input string) []string{
	var return_statement []string
	for i:=0;i<len(input);i++{
		if input[i] == temp.upper[0] || input[i] == temp.lower[0]{
			matched := true
			var appendedString string
			for j:=i;j<i+len(temp.original);j++{
				if(input[j] != temp.upper[j-i] && input[j] != temp.lower[j-i]){
					matched = false
					break
				}else{
					appendedString += string(input[j])// this is good 
				}
			}
			if matched == true{
				//gotta make sure appended string is not repeated
				is_matched := false
				for i:=0;i<len(return_statement);i++{
					if return_statement[i] == appendedString{
						is_matched = true
					}
				}
				if is_matched == false{
					return_statement = append(return_statement,appendedString)
				}
			}
		}
	}
	return return_statement // a string slice of all the possible cases of the word
}

// This is for you to check and make sure the website still functions
func httpCheck(file string)(error){
	_ ,err := http.Get("http://"+file+".com")
	return err
}

func txtToGo(filePath string){
	before,_:=os.Getwd()
	before += "/"
	before += filePath
	before +=".txt"
	after,_:=os.Getwd()
	after += "/"
	after += filePath
	after += ".go"
	os.Rename(before,after)
}
