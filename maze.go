package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

const (
	Empty = 0
	Wall  = 3
	Start = 1
	End   = 2
)

func (a *App) CreateMaze(size int, mode bool) [][]int {
	maze := make([][]int, size)
	
	for i := range maze {
		maze[i] = make([]int, size)
		for j := range maze[i] {
			maze[i][j] = Empty
		}
	}

	
	for i := 0; i < size; i++ {
		maze[0][i] = Wall
		maze[size-1][i] = Wall 
		maze[i][0] = Wall      
		maze[i][size-1] = Wall 
	}

	rand.Seed(time.Now().UnixNano())

	
	startX := 1 + rand.Intn(size-2)
	startY := 1 + rand.Intn(size-2)
	maze[startX][startY] = Wall

	
	a.carveMaze(maze, startX, startY)

	for {
		endX := 1 + rand.Intn(size-2)
		endY := 1 + rand.Intn(size-2)
		if maze[endX][endY] == Empty {
			maze[endX][endY] = End
			break
		}
	}

	return maze
}

func (a *App) carveMaze(maze [][]int, x, y int) {
	type pos struct{ dx, dy int }
	dirs := []pos{{0, 2}, {2, 0}, {0, -2}, {-2, 0}}
	rand.Shuffle(len(dirs), func(i, j int) { dirs[i], dirs[j] = dirs[j], dirs[i] })

	inBounds := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < len(maze) && y < len(maze[0])
	}

	for _, d := range dirs {
		nx, ny := x+d.dx, y+d.dy
		if !inBounds(nx, ny) || maze[nx][ny] != Empty {
			continue
		}

		if rand.Float64() < 0.1 {
			maze[nx][ny] = Wall
			midX, midY := x+d.dx/2, y+d.dy/2
			maze[midX][midY] = Wall
		}

		a.carveMaze(maze, nx, ny)

	}
}

func (a *App) MazeExist() bool {
	fileContent, err := os.ReadFile("Mazes.json")
	if err != nil {
		return false
	}
	var data MazeData

	if err := json.Unmarshal(fileContent, &data); err != nil {
		return false
	}
	return len(data.Mazes) > 0
}

func (a *App) SaveMaze(maze [][]int, name string, mode bool) {
	var data MazeData
	fileContent, err := os.ReadFile("Mazes.json")
	if err != nil || len(fileContent) == 0 {
		data = MazeData{Mazes: []Maze{}}
	} else {
		if err := json.Unmarshal(fileContent, &data); err != nil {
			data = MazeData{Mazes: []Maze{}}
		}
	}

	endX, endY := a.GetEndPoint(maze)
	startX, startY := a.GetStartPoint(maze)

	newMaze := Maze{
		Name:   name,
		Mode:   mode,
		Matrix: maze,
		End:    Coord{Fila: endX, Col: endY},
		Start:  Coord{Fila: startX, Col: startY},
	}
	data.Mazes = append(data.Mazes, newMaze)

	updatedContent, err := json.Marshal(data)
	if err != nil {
		return
	}
	os.WriteFile("Mazes.json", updatedContent, 0644)
}

func (a *App) GetMazes() MazeData {
	fileContent, err := os.ReadFile("Mazes.json")
	if err != nil {
		return MazeData{Mazes: []Maze{}}
	}
	var data MazeData

	if err := json.Unmarshal(fileContent, &data); err != nil {
		return MazeData{Mazes: []Maze{}}
	}
	return data
}

func (a *App) GetStartPoint(maze [][]int) (int, int) {
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == Start {
				return i, j
			}
		}
	}
	return -1, -1
}

func (a *App) GetEndPoint(maze [][]int) (int, int) {
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == End {
				return i, j
			}
		}
	}
	return -1, -1
}

func (a *App) SetSPoint(maze [][]int, x, y int) [][]int {
	maze = a.DeleteStartPoint(maze)
	if maze[x][y] != Wall && maze[x][y] != End {
		maze[x][y] = Start
	}
	return maze
}

func (a *App) DeleteStartPoint(maze [][]int) [][]int {
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == Start {
				maze[i][j] = Empty
				return maze
			}
		}
	}
	return maze
}
