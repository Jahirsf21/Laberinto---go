package main

// Maze represents the structure of a maze
// Fields:
// - Name: The name of the maze.
// - Mode: The game mode (true for automatic, false for manual).
// - Matrix: The maze
// - End: The coordinates of the end point in the maze.
type Maze struct {
	Name     string  `json:"name"`
	Mode     bool    `json:"mode"`
	Matrix   [][]int `json:"matrix"`
	End      Coord   `json:"end"`
}


// BFS represents a node in the Breadth-First Search (BFS) algorithm.
// Fields:
// - path: The current path being explored.
// - steps: The number of steps taken to reach the current position.
type BFS struct {
	path  Camino
	steps int
}

// Coord represents a coordinate in the maze.
// Fields:
// - Fila: The row index of the coordinate.
// - Col: The column index of the coordinate.
type Coord struct {
	Fila int `json:"fila"`
	Col  int `json:"col"`
}

// Camino represents a path in the maze as a slice of coordinates.
type Camino []Coord

// MazeData represents a collection of mazes.
// Fields:
// - Mazes: A slice containing multiple Maze structures.
type MazeData struct {
	Mazes []Maze `json:"Mazes"`
}
