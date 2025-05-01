package main

type Maze struct {
	Name     string  `json:"name"`
	Mode     bool    `json:"mode"`
	Matrix   [][]int `json:"matrix"`
	End      Coord   `json:"end"`
	AllPaths Camino  `json:"allPaths"`
	BestPath Camino  `json:"bestPath"`
}


type BFS struct {
	path  Camino
	steps int
}

type Coord struct {
	Fila int `json:"fila"`
	Col  int `json:"col"`
}

type Camino []Coord

type MazeData struct {
	Mazes []Maze `json:"Mazes"`
}
