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

	// creating struct with data type
	type processo struct {
		cnpj           interface{}
		entnome        interface{}
		numeroprocesso interface{}
		ramo           interface{}
		subramo        interface{}
		tipoproduto    interface{}
	}

	filteredData := []processo{}

	// creating map for tipo produto
	jsonValue := susepData["value"].([]interface{})

	// filtering data by tipo produto
	for _, data := range jsonValue {
		if data.(map[string]interface{})["tipoproduto"] == "PLANO DE PREVIDÊNCIA" {
			novoProcesso := processo{
				cnpj:           data.(map[string]interface{})["cnpj"],
				entnome:        data.(map[string]interface{})["entnome"],
				numeroprocesso: data.(map[string]interface{})["numeroprocesso"],
				ramo:           data.(map[string]interface{})["ramo"],
				subramo:        data.(map[string]interface{})["subramo"],
				tipoproduto:    data.(map[string]interface{})["tipoproduto"],
			}
			filteredData = append(filteredData, novoProcesso)
		}
	}

	processNumbers := []string{}

	// getting process numbers
	for _, data := range filteredData {
		processNumbers = append(processNumbers, data.numeroprocesso.(string))
	}

	fmt.Println(processNumbers)

}
