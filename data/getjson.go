package data

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetJson() string {
	body, err := ioutil.ReadFile("data.json")
	if err != nil {
		resp, err := http.Get("http://dados.susep.gov.br/olinda-ide/servico/produtos/versao/v1/odata/DadosProdutos?$format=json")
		if err != nil {
			log.Fatalln(err)
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
	}

	sb := string(body)
	return sb
}
