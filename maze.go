package main

import (
	"encoding/json"
	"math/rand"
	"os"
)

func (a *App) CreateMaze(size int, mode bool) [][]int {
	var maze [][]int
	for range size {
		var row []int
		for range size {
			row = append(row, 0)
		}
		maze = append(maze, row)
	}
	if mode {
		endX, endY := rand.Intn(size), rand.Intn(size)
		maze[endX][endY] = 2
	} else {
		startX, startY := rand.Intn(size), rand.Intn(size)
		endX, endY := rand.Intn(size), rand.Intn(size)

		for startX == endX && startY == endY {
			endX, endY = rand.Intn(size), rand.Intn(size)
		}

		maze[startX][startY] = 1
		maze[endX][endY] = 2
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

func (a *App) SetSPoint(maze [][]int, x, y int) [][]int {
	maze[x][y] = 1
	return maze
}

func (a *App) DeleteStartPoint(maze [][]int) [][]int {
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == 1 {
				maze[i][j] = 0
				return maze // Salir después de encontrar y actualizar el punto de inicio
			}
		}
	}
	return maze // Retornar la matriz sin cambios si no se encuentra un punto de inicio
}
