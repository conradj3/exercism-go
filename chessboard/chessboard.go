package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank = []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard = map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	return counter(cb, func(r string, f int, val bool) bool {
		return rank == r && val
	})
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	return counter(cb, func(rank string, f int, val bool) bool {
		return file == f && val
	})
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return counter(cb, func(rank string, file int, val bool) bool {
		return true
	})
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	return counter(cb, func(rank string, file int, val bool) bool {
		return val
	})
}
func counter(cb Chessboard, count func(rank string, file int, val bool) bool) int {
	var total int
	for key, rank := range cb {
		for i, file := range rank {
			if count(key, i+1, file) {
				total++
			}
		}
	}
	return total
}
