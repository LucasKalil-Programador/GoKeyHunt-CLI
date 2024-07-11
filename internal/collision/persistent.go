package collision

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const DefaultBase = 62

type intervalTemp struct {
	A string `json:"A"`
	B string `json:"B"`
}

type intervalsTemp struct {
	Data []intervalTemp `json:"Intervals"`
}

func (intArr *IntervalArray) Save(filePath string) bool {
	intervals := intArr.toTempIntervals()

	jsonData, err := json.Marshal(intervals)
	if err != nil {
		log.Println("Error on Marshal function:", err)
		return false
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		log.Println("Error on write json file:", err)
		return false
	}
	return true
}

func Read(filePath string) (*IntervalArray, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var intervalsTmp intervalsTemp
	if err := json.Unmarshal(bytes, &intervalsTmp); err != nil {
		return nil, err
	}

	intervalArr := toIntervalArray(intervalsTmp)
	return &intervalArr, nil
}

func ReadOrNew(filePath string) *IntervalArray {
	IntervalArr, err := Read(filePath)
	if err != nil {
		return NewEmptyIntervalArray()
	}
	return IntervalArr
}

func toIntervalArray(intervalsTmp intervalsTemp) IntervalArray {
	var intervals []Interval
	for _, intervalTemp := range intervalsTmp.Data {
		interval, success := new(Interval).SetString(intervalTemp.A, intervalTemp.B, DefaultBase)
		if success {
			intervals = append(intervals, *interval)
		}
	}
	SortByStart(intervals)
	return IntervalArray{data: intervals}
}

func (intArr *IntervalArray) toTempIntervals() intervalsTemp {
	intArr.Optimize()
	intervalsTmpArr := make([]intervalTemp, len(intArr.data))
	for i, interval := range intArr.data {
		intervalsTmpArr[i] = intervalTemp{A: interval.a.Text(DefaultBase), B: interval.b.Text(DefaultBase)}
	}
	intervals := intervalsTemp{Data: intervalsTmpArr}
	return intervals
}
