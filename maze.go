package main

import (
	"encoding/json"
	"math/rand"
	"os"
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
	}
	var startX, startY, endX, endY int
	if mode {
		endX, endY = rand.Intn(size), rand.Intn(size)
		maze[endX][endY] = End
	} else {
		startX, startY = rand.Intn(size), rand.Intn(size)
		endX, endY = rand.Intn(size), rand.Intn(size)

		for startX == endX && startY == endY {
			endX, endY = rand.Intn(size), rand.Intn(size)
		}

		maze[startX][startY] = Start
		maze[endX][endY] = End
	}

	return maze
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
