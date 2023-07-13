package domain

type Board struct {
	Size    int
	Snakes  map[int]int
	Ladders map[int]int
}

func pairs2Map(pairs [][2]int) map[int]int {
	res := map[int]int{}
	for _, x := range pairs {
		res[x[0]] = x[1]
	}

	return res
}

func NewBoard(size int, snakes [][2]int, ladders [][2]int) Board {
	snakeMap := pairs2Map(snakes)
	ladderMap := pairs2Map(ladders)
	return Board{
		Size:    size,
		Snakes:  snakeMap,
		Ladders: ladderMap,
	}
}
