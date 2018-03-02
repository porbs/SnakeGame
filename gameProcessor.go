package main

// Game map configuration
const (
	gridWidth  int = 20
	gridHeight int = 20
)

// Tuple struct
type Tuple struct {
	A, B interface{}
}

// Check if apple's and snake's position are same
func isFeeding(s *Snake, a *Point) bool {
	return s.coordinates.Front().Value.(Point) == *a
}

// Generate valid position for apple
func getNewApplePos(s *Snake) Point {

	// Randomize point on the map
	point := Point{X: r.Intn(gridWidth), Y: r.Intn(gridHeight)}

	// While point is on snake: generate new point
	for el := s.coordinates.Front(); el != nil; el = el.Next() {
		if point == el.Value.(Point) {
			point = Point{X: r.Intn(gridWidth), Y: r.Intn(gridHeight)}
			el = s.coordinates.Front()
		}
	}

	return point
}

// Check if snake is alive
func isGameOver(s *Snake) bool {
	headValue := s.coordinates.Front().Value.(Point)

	// Check for wall smash
	if headValue.X <= -1 || headValue.X >= gridWidth+1 || headValue.Y <= -1 || headValue.Y >= gridHeight+1 {
		return true
	}

	// Check for self smash
	for el := s.coordinates.Front().Next(); el != nil; el = el.Next() {
		if headValue == el.Value.(Point) {
			return true
		}
	}

	return false
}

// Scale point coordinates to viewport coordinates
func scalePoint(point Point) (result Tuple) {
	nodeWidth := windowWidth / float64(gridWidth+1)
	nodeHeight := windowHeight / float64(gridHeight+1)

	result.A = Tuple{A: float64(point.X) * nodeWidth, B: float64(point.Y) * nodeHeight}
	result.B = Tuple{A: float64(point.X+1) * nodeWidth, B: float64(point.Y+1) * nodeHeight}

	return
}

// Scale snake coordinates to viewport coordinates
func scaleSnake(s Snake) (snake []Tuple) {
	for _, value := range s.CoordinatesToSlice() {
		snake = append(snake, scalePoint(value))
	}
	return
}
