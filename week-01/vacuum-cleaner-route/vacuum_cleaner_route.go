package vacuumcleanerroute

func IsCompleteRoute(moves string) bool {
	var (
		x int
		y int
	)
	for _, c := range moves {
		switch c {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		}
	}
	return x == 0 && y == 0
}
