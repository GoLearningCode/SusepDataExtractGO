package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/GoLearningCode/SusepDataExtractGO/data"
)

func main() {
	// getting the json data
	jsonData := data.GetJson()

	// saving the data in local file
	file, err := os.Create("data.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = file.WriteString(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	file.Sync()

	// parsing to maps

	var susepData map[string]interface{}

	json.Unmarshal([]byte(jsonData), &susepData)

	// printing odata context to confirm convertion
	fmt.Println(susepData["@odata.context"])

	// creating map for tipo produto
	jsonValue := susepData["value"].([]interface{})

	filtered_data := map[string]interface{}{}

	// filtering data by tipo produto
	for _, data := range jsonValue {
		if data.(map[string]interface{})["tipoproduto"] == "PLANO DE PREVIDÊNCIA" {
			filtered_data = map[string]interface{}{data, filtered_data}}
		}
	}

	fmt.Println(filtered_data)

}
