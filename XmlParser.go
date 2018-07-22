package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
	"time"
)

// main structure in xml
type TrainLegs struct {
	XMLName      xml.Name   `xml:"TrainLegs"`
	TrainLegList []TrainLeg `xml:"TrainLeg"`
}

// structure fo retrieving TrainLeg's from xml
type TrainLeg struct {
	XMLName            xml.Name `xml:"TrainLeg"`
	TrainId            string   `xml:"TrainId,attr"`
	DepartureStationId string   `xml:"DepartureStationId,attr"`
	ArrivalStationId   string   `xml:"ArrivalStationId,attr"`
	Price              string   `xml:"Price,attr"`
	ArrivalTime        string   `xml:"ArrivalTimeString,attr"`
	DepartureTime      string   `xml:"DepartureTimeString,attr"`
}

func (t TrainLeg) String() string {
	return fmt.Sprintf("TrainId = %s DepartureStationId = %s ArrivalStationId = %s Price = %s ArrivalTime = %s DepartureTime = %s",
		t.TrainId, t.DepartureStationId, t.ArrivalStationId, t.Price, t.ArrivalTime, t.DepartureTime)
}

// structure fo parsed TrainLeg's according to data types
// used for further computing
type TrainLegParsed struct {
	TrainId            int
	DepartureStationId int
	ArrivalStationId   int
	Price              int
	ArrivalTime        time.Time
	DepartureTime      time.Time
}

var (
	trainLegList       []TrainLeg
	trainLegParsedList []TrainLegParsed
)

func (t TrainLegParsed) String() string {
	return fmt.Sprintf("TrainId = %d DepartureStationId = %d ArrivalStationId = %d Price = %d ArrivalTime = %s DepartureTime = %s",
		t.TrainId, t.DepartureStationId, t.ArrivalStationId, t.Price, t.ArrivalTime, t.DepartureTime)
}

// returns an array (slice) of records according to TrainLeg structure (all fields are strings) and nil
// if there is an error occurs returns nil and error
func getTrainLegList(fileName string) (result []TrainLeg, err error) {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	byteArray, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}
	var trainLegs TrainLegs
	xml.Unmarshal(byteArray, &trainLegs)
	return trainLegs.TrainLegList, nil
}

// returns an array (slice) of parsed records according to TrainLegParsed structure. Price is int in cents
func getTrainLegParsedList(trainLegList []TrainLeg) []TrainLegParsed {
	result := make([]TrainLegParsed, 0, len(trainLegList))
	for _, trainLeg := range trainLegList {
		trainId, _ := strconv.Atoi(trainLeg.TrainId)
		departureStationId, _ := strconv.Atoi(trainLeg.DepartureStationId)
		arrivalStationId, _ := strconv.Atoi(trainLeg.ArrivalStationId)
		price, _ := strconv.ParseFloat(trainLeg.Price, 64)
		arrivalTime, _ := time.Parse("15:04:05", trainLeg.ArrivalTime)
		departureTime, _ := time.Parse("15:04:05", trainLeg.DepartureTime)
		trainLegParsed := TrainLegParsed{TrainId: trainId, DepartureStationId: departureStationId,
			ArrivalStationId: arrivalStationId, Price: int(price * 100), ArrivalTime: arrivalTime, DepartureTime: departureTime}
		result = append(result, trainLegParsed)
	}
	return result
}
