package training

import (
	"log"
	"xo-detection/data"
	"xo-detection/models"
	"xo-detection/utils"
)

func LoadWeightsOrTrain(savedWeightsPath string, saveToDisk, loadFromDisk bool, trainingFunc func() models.SavedWeightAndBiasJsonObject) models.SavedWeightAndBiasJsonObject {
	if !loadFromDisk || !utils.FileExists(savedWeightsPath) {
		log.Printf("Path to saved weights does not exist: %v\n", savedWeightsPath)
		log.Printf("Training right now.\n")

		objectToSave := trainingFunc()

		log.Printf("Seems like training is finished: %v\n", objectToSave)

		if saveToDisk {
			log.Printf("Saving weights to %v...\n", savedWeightsPath)
			success, err := data.SaveWeights(objectToSave, savedWeightsPath)

			if !success {
				log.Printf("Saving data faild.\n")
				log.Fatalln(err)
			}

			log.Printf("Saving weights to %v was successful.\n", savedWeightsPath)
		}

		return objectToSave
	}

	log.Printf("Loading weights from %v...\n", savedWeightsPath)
	return data.UnmarshalWeightsFile(savedWeightsPath)
}
