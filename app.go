package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
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
	SetStartAndEnd(matrix)
	return matrix
}

func SetStartAndEnd(matrix [][]int) [][]int {
	rand.Seed(time.Now().UnixNano())

	size := len(matrix)

	// Coordenadas para salida (3)
	x1 := rand.Intn(size)
	y1 := rand.Intn(size)
	matrix[x1][y1] = 3

	// Coordenadas distintas para llegada (2)
	var x2, y2 int
	for {
		x2 = rand.Intn(size)
		y2 = rand.Intn(size)
		if x2 != x1 || y2 != y1 {
			break
		}
	}
	matrix[x2][y2] = 2

	fmt.Printf("Punto de salida (3): [%d][%d]\n", x1, y1)
	fmt.Printf("Punto de llegada (2): [%d][%d]\n", x2, y2)

	return matrix
}

func (a *App) MazeExist() bool {
	fileContent, err := os.ReadFile("Mazes.json")
	if err != nil {
		return false
	}
	var data struct {
		Mazes []Maze `json:"Mazes"`
	}

	if err := json.Unmarshal(fileContent, &data); err != nil {
		return false
	}
	return len(data.Mazes) > 0
}
