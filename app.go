package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// Maze estructura para los laberintos
type Maze struct {
	ID           int      `json:"id"`
	Matriz       [][]int  `json:"matriz"`
	Inicio       Coord    `json:"inicio"`
	Fin          Coord    `json:"fin"`
	CaminoOptimo Camino   `json:"camino_optimo"`
	TodosCaminos []Camino `json:"todos_caminos"`
}

// Coord estructura para la posición (Inicio, fin)
type Coord struct {
	Fila int `json:"fila"`
	Col  int `json:"col"`
}

type Camino []Coord

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) CreateMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	fmt.Printf("Matriz creada: %dx%d\n", size, size)
	return matrix
}

func (a *App) MazeExist() bool {
	fileContent, err := os.ReadFile("Mazes.json")
	if err != nil {
		return false
	}
	var maze Maze
	if err := json.Unmarshal(fileContent, &maze); err != nil {
		return false
	}
	return len(maze.Matriz) > 0
}
