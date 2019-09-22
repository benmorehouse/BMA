# BMA

## Tool to sift through and change key words in files and filename

This is an application that I built which will:
	- sift through files and find the text you indicated
	- return a new file with the indicated text you want it replaced
		 The new file name will be based off this input
	- the type of file is up to you at the end

The file will be placed in the same path that you input.

# Installation

If you do not have homebrew, go here: https://brew.sh

Once you have homebrew, if you do not have go installed then enter
	
	brew install go

into your desired directory.
Then install the following two packages:
	
	Goquery: go get github.com/PuerkitoBio/goquery
	Golang Collections (not native): go get github.com/golang-collections/collections/trie

Then pull this project into your desired directory.

	git pull git@github.com:benmorehouse/BMA.git

Then build
	
	go build 

Then you are all set to run BMA!
		
# Instructions:

BMA has two required flags the user must input
	path
	FE (File Extension)

Here, you passed arguments for both flags
	./BMA -path=/Users/benmorehouse/documents/data.csv -FE=.xml
Here, you passed arguments for just the one path flag
	./BMA -path=/Users/benmorehouse/repositories/example.txt
It will prompt you for an input on what the new file extension should be

Here, you passed no arguments 
	./BMA
It will return an error, as no file was input thus you cannot use BMA

# Thanks and Welcome to BMA


