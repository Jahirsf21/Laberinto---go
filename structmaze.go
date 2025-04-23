package main

type Maze struct {
	Name     string  `json:"name"`
	Mode     bool    `json:"mode"`
	Matrix   [][]int `json:"matrix"`
	End      Coord   `json:"end"`
	Solution Camino  `json:"solution"`
}

type Coord struct {
	Fila int `json:"fila"`
	Col  int `json:"col"`
}

type Camino []Coord

type MazeData struct {
	Mazes []Maze `json:"Mazes"`
}
