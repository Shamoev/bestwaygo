package main

import (
	"sort"
	"fmt"
)

const INF = (1<<31 - 1) / 2 // infinity maxInt / 2

var (
	stationIds       []int
	numberOfStations int
	matrixOfPrices   [][]int
	matrixOfPaths    [][]int
)

// generates matrix of prices which represents weighted graph and matrix of paths for further retrieving shortest paths
// returns matrices
func generateMatrices(trainLegs []TrainLegParsed) (matrixOfPrices, matrixOfPaths [][]int) {
	stationIds = getStationIds(trainLegs)
	numberOfStations = len(stationIds)
	matrixOfPrices = make([][]int, numberOfStations)
	matrixOfPaths = make([][]int, numberOfStations)
	// initiate matrix of prices and matrix of paths with zeros
	for i := 0; i < numberOfStations; i++ {
		matrixOfPrices[i] = make([]int, numberOfStations)
		matrixOfPaths[i] = make([]int, numberOfStations)
	}
	// fill matrix of paths with -1. -1 means no further way - for further computing by Floyd Marshal algorithm
	for i := 0; i < len(matrixOfPaths); i++ {
		for j := 0; j < len(matrixOfPaths[i]); j++ {
			matrixOfPaths[i][j] = -1
		}
	}
	// fill matrix of prices with minimal costs between stations and prior fill matrix of paths
	for i := 0; i < numberOfStations; i++ {
		for j := 0; j < numberOfStations; j++ {
			departureStationId := stationIds[i]
			arrivalStationId := stationIds[j]
			minPrice := minPriceBetweenStations(departureStationId, arrivalStationId, trainLegs)
			matrixOfPrices[i][j] = minPrice
			if minPrice != INF {
				matrixOfPaths[i][j] = j
			}
		}
	}
	return matrixOfPrices, matrixOfPaths
}

// returns sorted array (slice) of different station ids
func getStationIds(trainLegs []TrainLegParsed) []int {
	// create map for choosing different stations
	mapOfStationIds := make(map[int]bool)
	for _, trainLeg := range trainLegs {
		mapOfStationIds[trainLeg.DepartureStationId] = true
	}
	result := make([]int, 0, len(mapOfStationIds))
	for k, _ := range mapOfStationIds {
		result = append(result, k)
	}
	sort.Ints(result)
	return result
}

// returns minimal price between two stations without link to certain train
func minPriceBetweenStations(departureStationId, arrivalStationId int, trainLegs []TrainLegParsed) int {
	// INF - no dirext connection
	result := INF
	for _, trainLeg := range trainLegs {
		if trainLeg.DepartureStationId == departureStationId && trainLeg.ArrivalStationId == arrivalStationId {
			if trainLeg.Price < result {
				result = trainLeg.Price
			}
		}
	}
	return result
}

// prints paths between all combinations of stations and minimal cost
//if there is more than one train with minimal price prints all the trains
// station [trains equal by price] station and so on | total cost
func printAllPaths() {
	fmt.Println("Paths between all pairs of stations")
	fmt.Println("station [trains equal by price] station ... and so on")
	for i := 0; i < numberOfStations; i++ {
		for j := 0; j < numberOfStations; j++ {
			if i != j {
				printPath(retrievePath(i, j, matrixOfHistory))
			}
		}
	}
}

// prints one path getting all the trains with equal minimal price and cost from matrixOfWeights
func printPath(path []int) {
	for i := 0; i < len(path); i++ {
		fmt.Print(getStationIdByIndex(path[i]), " ")
		if i != (len(path) - 1) {
			fmt.Print(getAllLowestPriceTrains(path[i], path[i+1]), " ")
		}
	}
	fmt.Print(" | Total cost: ", getCost(path[0], path[len(path)-1]))
	fmt.Println()
}

// returns stationId from stationIds by index
func getStationIdByIndex(index int) int {
	return stationIds[index]
}

// returns all cheapest trains with the same price. That price was used when matrixOfPrices was calculated
func getAllLowestPriceTrains(departureIndex, arrivalIndex int) []int {
	result := make([]int, 0, 10)
	departureId := stationIds[departureIndex]
	arrivalId := stationIds[arrivalIndex]
	price := matrixOfPrices[departureIndex][arrivalIndex]
	for _, trainLeg := range trainLegParsedList {
		if trainLeg.DepartureStationId == departureId && trainLeg.ArrivalStationId == arrivalId && trainLeg.Price == price {
			result = append(result, trainLeg.TrainId)
		}
	}
	return result
}
