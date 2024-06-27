package main

import (
	"fmt"
	"lib"
	"os"
)

func main() {
	var asset_file_name string
	//check arguments
	if len(os.Args) == 3 {
		asset_file_name = os.Args[2]
	} else if len(os.Args) == 2 {
		asset_file_name = "standard.txt"
	} else {

		fmt.Println("Error : to much or missing arguments")
		return
	}

	sentence := os.Args[1]
	asset, thinkertoy := lib.GetFileContent("assets/" + asset_file_name)
	table_asset := lib.CreateTable(asset, thinkertoy)
	lib.PrintAscii(table_asset, sentence)
}
