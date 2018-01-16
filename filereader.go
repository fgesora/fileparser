package main

import (
		"os"
		"fmt"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}


func main(){
	file, err := os.Open("sampletxt.txt")
	check(err)

	defer file.Close()

	fileinfor, err := file.Stat()
	check(err)

	sizeoffile := fileinfor.Size()
	buffer := make([]byte, sizeoffile)

	bytesread, err := file.Read(buffer)
	check(err)

	fmt.Println("Bytes read in file: ", bytesread)
	fmt.Println("File size is: ", sizeoffile)
	fmt.Println("Bytes to string: ", string(buffer))
}