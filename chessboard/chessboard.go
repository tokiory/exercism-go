package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var occupied int
	for _, ok := range cb[file] {
		if ok {
			occupied++
		}
	}

	return occupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var occupied int

	if rank < 1 || rank > 8 {
		return occupied
	}

	for _, v := range cb {
		if v[rank-1] {
			occupied++
		}
	}

	return occupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var count int

	for _, r := range cb {
		count += len(r)
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var occupied int

	for char := byte('A'); char <= byte('H'); char++ {
		occupied += CountInFile(cb, string(char))
	}

	return occupied
}
