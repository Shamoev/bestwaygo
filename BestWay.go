/*
	Calculates and prints the cheapest way between al pairs of stations
	f there is more than one train with minimal price between two stations prints all the trains
 */
package main

import (
	"fmt"
)

func main() {
	// read records from xml
	var err error
	trainLegList, err = getTrainLegList("data.xml")
	if err != nil {
		fmt.Println("Error in opening or reading file:", err)
		return
	}
	// parse records
	trainLegParsedList = getTrainLegParsedList(trainLegList);
	// generate matrix of minimal prices and matrix of paths for further computing by Floyd-Warshall algorithm
	matrixOfPrices, matrixOfPaths = generateMatrices(trainLegParsedList)
	// calculate matrix of weights and matrix of history using Floyd-Warshal algorithm
	matrixOfWeights, matrixOfHistory = doFloydWarshall(matrixOfPrices, matrixOfPaths)
	// print all combinations of stations with minimal cost
	printAllPaths()
}
