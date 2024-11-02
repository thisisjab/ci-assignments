package data

import (
	"encoding/json"
	"os"
	"xo-detection/models"
)

func SaveWeights(data models.SavedWeightAndBiasJsonObject, path string) (success bool, err error) {
	jsonString, _ := json.MarshalIndent(data, "", "\t")
	err = os.WriteFile(path, jsonString, os.ModePerm)

	if err != nil {
		return false, err
	}

	return true, nil
}
