package main

type Maze struct {
	Name   string  `json:"name"`
	Mode   bool    `json:"mode"`
	Matrix [][]int `json:"matrix"`
	Start  Coord   `json:"start"`
	End    Coord   `json:"end"`
}

type Coord struct {
	Fila int `json:"fila"`
	Col  int `json:"col"`
}

type Camino []Coord

type MazeData struct {
	Mazes []Maze `json:"Mazes"`
}
