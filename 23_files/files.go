package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("file name: ", fileInfo.Name())
	fmt.Println("file size: ", fileInfo.Size())
	fmt.Println("file or folder: ", fileInfo.IsDir())
	fmt.Println("file permissions: ", fileInfo.Mode())
	fmt.Println("file modified at: ", fileInfo.ModTime())

	fmt.Println("==============================================")

	//read file
	f1, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer f1.Close()

	fileData, err := f1.Stat()
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, fileData.Size())
	d, err := f1.Read(buffer)
	if err != nil {
		panic(err)
	}

	for i := range d {
		fmt.Println("data", i, string(buffer[i]))
	}

	fmt.Println("========================================================")

	//Alternate way to readFile

	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	fmt.Println("========================================================")

	//read folder

	dir, err := os.Open(".")
	//dir, err := os.Open("../") //one folder back
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	//fileInfo1, err := dir.ReadDir(2)
	fileInfo1, err := dir.ReadDir(-1) //-1 to get all files

	for _, fInfo := range fileInfo1 {
		fmt.Println(fInfo.Name(), fInfo.IsDir())
	}

	fmt.Println("========================================================")

	//write into file
	//1. create file
	f2, err := os.Create("example4.txt")
	if err != nil {
		panic(err)
	}

	defer f2.Close()
	f2.WriteString("Golang having built-in concurrency support!!")
	bytes := []byte(" Hello golang") //alternate way to write in file with bytes
	f2.Write(bytes)

	fmt.Println("========================================================")

	//read from file and write into another file (streaming fashion)

	sourceFile, err := os.Open("example4.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example5.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {

		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()

	fmt.Println("Written to new file successfully!!")

	fmt.Println("========================================================")

	//deleting file
	err1 := os.Remove("example3.txt")
	if err1 != nil {
		panic(err1)
	}

	fmt.Println("Deleted file successfully!!")

}
