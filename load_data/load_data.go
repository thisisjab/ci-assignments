package load_data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"xo-detection/models"
)

func UnmarshalTrainingDataFile(path string) []models.TrainingVectorJsonObject {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Cannot open file at given path:", path)
	}

	defer file.Close()

	bytesValue, _ := io.ReadAll(file)

	var data []models.TrainingVectorJsonObject

	json.Unmarshal(bytesValue, &data)

	return data
}

func PrepareData(data []models.TrainingVectorJsonObject) []models.TrainingVector {
	var results []models.TrainingVector

	for _, item := range data {
		var tempResult models.TrainingVector
		var tempWeights []float64

		for _, row := range item.Values {
			tempWeights = append(tempWeights, row...)
		}

		tempResult.Values = tempWeights
		tempResult.T = item.Label

		results = append(results, tempResult)
	}

	return results
}
