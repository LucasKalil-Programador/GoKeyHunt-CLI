package output_results

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Save saves the ResultArray to a JSON file.
//
// This method serializes the ResultArray to JSON format and writes it to the specified file path.
// If an error occurs during serialization or file writing, the function logs the error and returns false.
//
// Parameters:
// - jsonPath: A string representing the path where the JSON file will be saved.
//
// Returns:
// - bool: True if the file was saved successfully, false otherwise.
func (rArray *ResultArray) Save(jsonPath string) bool {
	jsonData, err := json.MarshalIndent(rArray, "", "	")
	if err != nil {
		log.Println("Error on Marshal function:", err)
		return false
	}

	err = os.WriteFile(jsonPath, jsonData, 0644)
	if err != nil {
		log.Println("Error on write json file:", err)
		return false
	}
	return true
}

// Read reads a JSON file and returns a ResultArray instance.
//
// This function reads the content of the specified JSON file, deserializes it into a ResultArray instance,
// and sorts the results by WalletIndex. If an error occurs during file reading or deserialization, the error is returned.
//
// Parameters:
// - filePath: A string representing the path to the JSON file to be read.
//
// Returns:
// - *ResultArray: A pointer to the deserialized and sorted ResultArray instance.
// - error: An error if any occurred during file reading or deserialization, nil otherwise.
func Read(filePath string) (*ResultArray, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var resultsArray ResultArray
	if err := json.Unmarshal(bytes, &resultsArray); err != nil {
		return nil, err
	}

	resultsArray.Resuts = SortByWalletIndex(resultsArray.Resuts)
	return &resultsArray, nil
}

// ReadOrNew reads a JSON file and returns a ResultArray instance or a new empty ResultArray.
//
// This function attempts to read the specified JSON file and deserialize it into a ResultArray instance.
// If an error occurs during reading or deserialization, a new empty ResultArray is returned instead.
//
// Parameters:
// - filePath: A string representing the path to the JSON file to be read.
//
// Returns:
// - *ResultArray: A pointer to the deserialized ResultArray instance, or a new empty ResultArray if an error occurred.
func ReadOrNew(filePath string) *ResultArray {
	IntervalArr, err := Read(filePath)
	if err != nil {
		return NewEmptyResultArray()
	}
	return IntervalArr
}
