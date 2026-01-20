package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	data := []int{2, 3, 5, 7, 11}

	fileData, err := json.MarshalIndent(data, "", "   ")

	if err != nil {
		fmt.Println(err)
		return
	}

	os.WriteFile("dummy.json", fileData, 0644)

	fmt.Println(fileData)
}
