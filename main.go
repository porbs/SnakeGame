package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Window configuration
const (
	windowWidth  float64 = 768
	windowHeight float64 = 768
	updateSpeed  float64 = 1
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Init game data
var (
	direction = UP
	snake     = CreateSnake(Point{X: int(float64(gridWidth)/2) - 2, Y: int(float64(gridHeight) / 2)}, direction, 3)
	apple     = getNewApplePos(snake)
	last      = time.Now()
)

// Draw a snake with specified colours
func drawSnake(snake *Snake, imd *imdraw.IMDraw, colorScheme []color.RGBA) {
	if len(colorScheme) < 2 {
		panic("Color scheme is not full")
	}
	imd.Color = colorScheme[0]
	imd.EndShape = imdraw.RoundEndShape
	for i, value := range scaleSnake(*snake) {
		if i == 0 {
			imd.Color = colorScheme[1]
		}
		imd.Push(pixel.V(value.A.(Tuple).A.(float64), value.A.(Tuple).B.(float64)),
			pixel.V(value.B.(Tuple).A.(float64), value.B.(Tuple).B.(float64)))
		if i == 0 {
			imd.Color = colorScheme[0]
		}
		imd.Rectangle(0)
	}
}

// Draw an apple
func drawApple(apple *Point, imd *imdraw.IMDraw) {
	appleData := scalePoint(*apple)
	imd.Color = colornames.Red
	imd.Push(pixel.V(appleData.A.(Tuple).A.(float64), appleData.A.(Tuple).B.(float64)),
		pixel.V(appleData.B.(Tuple).A.(float64), appleData.B.(Tuple).B.(float64)))
	imd.Rectangle(0)
}

// WASD keys listener
func wasdControll(win *pixelgl.Window, direction *Direction) bool {
	if win.JustPressed(pixelgl.KeyW) {
		*direction = UP
		return true
	} else if win.JustPressed(pixelgl.KeyS) {
		*direction = DOWN
		return true
	} else if win.JustPressed(pixelgl.KeyA) {
		*direction = LEFT
		return true
	} else if win.JustPressed(pixelgl.KeyD) {
		*direction = RIGHT
		return true
	}
	return false
}

// Arrow keys listener
func arrowControll(win *pixelgl.Window, direction *Direction) bool {
	if win.JustPressed(pixelgl.KeyUp) {
		*direction = UP
		return true
	} else if win.JustPressed(pixelgl.KeyDown) {
		*direction = DOWN
		return true
	} else if win.JustPressed(pixelgl.KeyLeft) {
		*direction = LEFT
		return true
	} else if win.JustPressed(pixelgl.KeyRight) {
		*direction = RIGHT
		return true
	}
	return false
}

// Draw gameover scene
func drawGameOverScene(imd *imdraw.IMDraw) {
	imd.Color = color.RGBA{255, 0, 0, 128}
	imd.Push(pixel.V(0, 0), pixel.V(windowWidth, windowHeight))
	imd.Rectangle(0)
}

// Update snake state
func snakeIteration(s *Snake, a *Point, direction Direction, updateTime *time.Time, lastUpdate float64) {
	if lastUpdate >= updateSpeed {
		if !isGameOver(s) {
			s.ChangeMovementDirection(direction)
			s.UpdatePosition()
			if isFeeding(s, a) {
				s.Grow()
				*a = getNewApplePos(snake)
			}
		}
		*updateTime = time.Now()
	}
}

func run() {

	// Init window
	cfg := pixelgl.WindowConfig{
		Title:     "Tizen 5.0",
		Bounds:    pixel.R(0, 0, windowWidth, windowHeight),
		VSync:     true,
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Init painter
	imd := imdraw.New(nil)

	for !win.Closed() {

		// calculate delta time from last update
		dt := time.Since(last).Seconds()

		// key listener
		if wasdControll(win, &direction) || arrowControll(win, &direction) {
			dt = updateSpeed
		}

		// Restart the game if ESC key pressed
		if win.JustPressed(pixelgl.KeyEscape) {
			direction = UP
			snake = CreateSnake(Point{X: int(float64(gridWidth)/2) - 2, Y: int(float64(gridHeight) / 2)}, direction, 3)
			apple = getNewApplePos(snake)
		}

		// Game loop
		snakeIteration(snake, &apple, direction, &last, dt)

		// Draw scene
		imd.Clear()
		drawSnake(snake, imd, []color.RGBA{colornames.Indigo, colornames.Black})
		drawApple(&apple, imd)

		// Background colour
		win.Clear(colornames.Lightskyblue)

		// Draw gameover scene in case of gameover
		if isGameOver(snake) {
			drawGameOverScene(imd)
		}

		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
