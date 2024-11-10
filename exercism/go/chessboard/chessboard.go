package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string][]bool

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	c := 0
	for k, v := range cb {
		if k == file {
			for _, x := range v {
				if x {
					c += 1
				}
			}
		}
	}
	return c
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	c := 0
	i := rank - 1
	if i > 7 || i < 0 {
		return 0
	}
	for _, v := range cb {
		if v[rank-1] {
			c += 1
		}
	}
	return c
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	c := 0
	for _, v := range cb {
		c += len(v)
	}
	return c
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	c := 0
	for _, v := range cb {
		for _, x := range v {
			if x {
				c += 1
			}
		}
	}
	return c
}
