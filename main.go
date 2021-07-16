package main

import (
	"log"
	"os"

	"github.com/GoLearningCode/SusepDataExtractGO/data"
)

func main() {
	jsonData := data.GetJson()
	file, err := os.Create("data.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err := file.WriteString(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	file.Sync()
}
