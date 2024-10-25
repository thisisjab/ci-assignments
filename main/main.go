package main

import (
	"fmt"
	"xo-detection/load_data"
	m "xo-detection/models"
	"xo-detection/training"
	"xo-detection/ui"
)

func main() {
	jsonData := load_data.UnmarshalTrainingDataFile("./training_data.json")
	cleanedData := load_data.PrepareData(jsonData)

	bias := 0.0
	weights := make([]m.Weights, 1)

	for i := 0 ; i < len(cleanedData[0].Values); i++ {
		weights[0] = append(weights[0], 0.0)
	}

	epochsToSucceed := training.Train(cleanedData, &weights, &bias, 0.2, 0.3)	

	fmt.Printf("Took %v epochs to succeed and final bias is %v", epochsToSucceed, bias)

	ui.RunUI(weights[len(weights) - 1], bias, 0.2)
}
