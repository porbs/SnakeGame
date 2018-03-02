package main

import (
	"container/list"
	"fmt"
)

// Direction type
type Direction int

// Direction enum
const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

// Point struct
type Point struct {
	X, Y int
}

// MovePoint moves point in passed direction
func MovePoint(direction Direction, point Point) Point {
	switch direction {
	case UP:
		point.Y++
	case RIGHT:
		point.X++
	case DOWN:
		point.Y--
	case LEFT:
		point.X--
	}

	return point
}

// OppositeDirection returns opposite to passed direction
func OppositeDirection(direction Direction) Direction {
	switch direction {
	case UP:
		return DOWN
	case RIGHT:
		return LEFT
	case DOWN:
		return UP
	case LEFT:
		return RIGHT
	}

	return direction
}

// Snake struct
type Snake struct {
	coordinates  *list.List
	direction    Direction
	lastPosition Point
}

// UpdatePosition updates position of snake
func (snake *Snake) UpdatePosition() {

	snake.lastPosition = snake.coordinates.Back().Value.(Point)
	snake.coordinates.Remove(snake.coordinates.Back())
	snake.coordinates.PushFront(MovePoint(snake.direction, snake.coordinates.Front().Value.(Point)))
}

// ChangeMovementDirection changes Snake object's movement direction
func (snake *Snake) ChangeMovementDirection(direction Direction) {
	if direction != OppositeDirection(snake.direction) {
		snake.direction = direction
	}
}

// Grow increases Snake object's size
func (snake *Snake) Grow() {
	snake.coordinates.PushBack(snake.lastPosition)
}

// PrintCoordinates prints Snake object's coordinates
func (snake *Snake) PrintCoordinates() {
	for el := snake.coordinates.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}
}

// CoordinatesToSlice translates coordinates list to slice
func (snake *Snake) CoordinatesToSlice() (result []Point) {
	for el := snake.coordinates.Front(); el != nil; el = el.Next() {
		result = append(result, el.Value.(Point))
	}
	return
}

// CreateSnake creates Snake object instance
func CreateSnake(initPosition Point, initDirection Direction, initLength int) *Snake {
	l := list.New()
	l.PushBack(initPosition)

	if initLength >= 2 {
		tailDirection := OppositeDirection(initDirection)
		tail := []Point{MovePoint(tailDirection, initPosition)}

		for i := 0; i < initLength-2; i++ {
			tail = append(tail, MovePoint(tailDirection, tail[i]))
		}

		for _, value := range tail {
			l.PushBack(value)
		}
	}

	return &Snake{
		coordinates: l,
		direction:   initDirection,
	}
}
