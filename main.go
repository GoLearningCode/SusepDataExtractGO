package main

import (
	"fmt"

	"github.com/GoLearningCode/SusepDataExtractGO/data"
)

func main() {
	sb := data.GetJson()
	fmt.Println(sb)
}
