package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

type rowsData struct {
	rows [][]string
}

func main() {

	excelFileName := "/Users/divya.p/Documents/Go_Workspace/src/MyXLSXFile.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				//rowone = append(rowone, text)
				fmt.Printf("%s\n", text)

			}
			fmt.Println()
		}
	}
	// dat, _ := ioutil.ReadFile("/Users/divya.p/Documents/Go_Workspace/src/MyXLSXFile.xlsx")
	// //check(err)
	// fmt.Print(string(dat))
	// buf := new(bytes.Buffer)
	// fmt.Println(encoder.Encoder(dat))
	// file, _ := os.Open("/Users/divya.p/Documents/Go_Workspace/src/MyXLSXFile.xlsx")
	// bufferedReader := bufio.NewReader(file)
	// str, _ := bufferedReader.ReadString('\n')
	// fmt.Println(string(str))

	//var mySlice [][][]string
	//var value string
	name := excelFileName
	mySlice, _ := xlsx.FileToSlice(name)
	//value = mySlice[0][0][0]
	fmt.Println(mySlice, "    value")

	mapvalue, err := xlFile.MarshallParts()
	if err != nil {
		log.Printf("   ", err)
	}
	//fmt.Println("mapvalueeeee    ", mapvalue)

	eee, err := json.Marshal(mapvalue)
	if err != nil {
		log.Printf("map marshalling  ", eee)
	}
	// encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	// encoder.Write(eee)
	// // Must close the encoder when finished to flush any partial blocks.
	// // If you comment out the following line, the last partial block "r"
	// // won't be encoded.
	// encoder.Close()

	encoded := base64.StdEncoding.EncodeToString(eee)
	//fmt.Println("encoded    ", encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	mapvaluedecode := make(map[string]string)

	json.Unmarshal(decoded, &mapvaluedecode)

	//fmt.Println(string(decoded))

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("mapvaluedecode  yes yes yes    ", mapvaluedecode)
}
