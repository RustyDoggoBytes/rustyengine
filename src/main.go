package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var cellSize int32 = 30
var cellCount int32 = 25

var green = rl.Color{173, 204, 96, 255}
var darkGreen = rl.Color{43, 51, 24, 255}

type Food struct {
	pos rl.Vector2
}

func NewFood() Food {
	position := GetRandomVector()
	return Food{pos: position}
}

func GetRandomVector() rl.Vector2 {
	return rl.Vector2{
		X: float32(rl.GetRandomValue(0, cellCount-1)),
		Y: float32(rl.GetRandomValue(0, cellCount-1)),
	}
}

func (f Food) Draw() {
	rl.DrawRectangle(int32(f.pos.X)*cellSize, int32(f.pos.Y)*cellSize, cellSize, cellSize, darkGreen)
}

type Snake struct {
	body      []rl.Vector2
	Direction rl.Vector2
}

func NewSnake(startPos rl.Vector2) Snake {
	return Snake{
		body:      []rl.Vector2{startPos},
		Direction: rl.Vector2{X: 1, Y: 0},
	}
}

func (s *Snake) Draw() {
	for _, vector := range s.body {
		x := vector.X
		y := vector.Y
		floatCellSize := float32(cellSize)
		rec := rl.Rectangle{
			X:      x * floatCellSize,
			Y:      y * floatCellSize,
			Width:  floatCellSize,
			Height: floatCellSize,
		}
		rl.DrawRectangleRounded(rec, 0.5, 1, rl.Black)
	}
}

func (s *Snake) Add(v rl.Vector2) {
	s.body = append(s.body, v)
}

func (s *Snake) SetDirection(v rl.Vector2) {
	s.Direction = v
}

func (s *Snake) Update() {
	s.body = s.body[:len(s.body)-1]
	s.body = append([]rl.Vector2{rl.Vector2Add(s.body[0], s.Direction)}, s.body...)
}

func (s *Snake) Print() {
	fmt.Println("len", len(s.body), "dir", s.Direction, "body", s.body)
}

var lastUpdatedTime float64 = 0

func EventTriggered(interval float64) bool {
	currentTime := rl.GetTime()
	if currentTime-lastUpdatedTime >= interval {
		lastUpdatedTime = currentTime
		return true
	}

	return false
}

func main() {
	rl.InitWindow(cellCount*cellSize, cellCount*cellSize, "Retro Snake")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	food := NewFood()

	snake := NewSnake(rl.Vector2{X: 6, Y: 9})
	snake.Add(rl.Vector2{X: 5, Y: 9})
	snake.Add(rl.Vector2{X: 4, Y: 9})

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		if EventTriggered(float64(0.2)) {
			snake.Update()
			snake.Print()
		}

		if rl.IsKeyPressed(rl.KeyDown) {
			snake.Direction = rl.Vector2{X: 0, Y: 1}
			fmt.Println("down")
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			snake.Direction = rl.Vector2{X: -1, Y: 0}
			fmt.Println("left")
		}
		if rl.IsKeyPressed(rl.KeyUp) {
			snake.Direction = rl.Vector2{X: 0, Y: -1}
			fmt.Println("up")
		}
		if rl.IsKeyPressed(rl.KeyRight) {
			fmt.Println("right")
			snake.Direction = rl.Vector2{X: 1, Y: 0}
		}

		rl.ClearBackground(green)
		snake.Draw()
		food.Draw()

		rl.EndDrawing()
	}
}
