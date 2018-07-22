package main

var (
	matrixOfWeights [][]int
	matrixOfHistory [][]int
)

// calculates matrix of weights which contains minimal weight of path between nodes
// and matrix of path using Floyd-Warshall algorithm
func doFloydWarshall(matrixOfPrices, MatrixOfPaths [][]int) (matrixOfWeights, matrixOfHistory [][]int) {
	size := len(matrixOfPrices)
	// create copies of matrices
	matrixOfWeights = make([][]int, size)
	matrixOfHistory = make([][]int, size)
	for i := 0; i < size; i++ {
		matrixOfWeights[i] = make([]int, size)
		matrixOfHistory[i] = make([]int, size)
		copy(matrixOfWeights[i], matrixOfPrices[i])
		copy(matrixOfHistory[i], MatrixOfPaths[i])

	}
	// do Floyd-Warshall algorithm
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrixOfWeights[i][j] != INF {
				for k := 0; k < size; k++ {
					if matrixOfWeights[i][k] > matrixOfWeights[i][j]+matrixOfWeights[j][k] {
						matrixOfWeights[i][k] = matrixOfWeights[i][j] + matrixOfWeights[j][k]
						matrixOfHistory[i][k] = matrixOfHistory[i][j]
					}
				}
			}
		}
	}

	return matrixOfWeights, matrixOfHistory
}

// retrieves path from matrix of history
// returns an array of ints which represents path [from - to]
// every int represents index of StationId in stationIds array
func retrievePath(from int, to int, matrixOfHistory [][]int) []int {
	result := make([]int, 0, len(matrixOfHistory))
	if matrixOfHistory[from][to] != -1 {
		result = append(result, from)
		for matrixOfHistory[from][to] != to {
			result = append(result, matrixOfHistory[from][to])
			from = matrixOfHistory[from][to]
		}
		result = append(result, to)
	}
	return result
}

// returns cost between two nodes
func getCost(from, to int) int {
	return matrixOfWeights[from][to]
}
