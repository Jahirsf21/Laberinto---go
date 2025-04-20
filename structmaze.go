package main

// Maze estructura para los laberintos
type Maze struct {
	Id              int     `json:"id"`
	Matriz          [][]int `json:"matriz"`
	Inicio          Coord   `json:"inicio"`
	Fin             Coord   `json:"fin"`
	Solucion        Camino  `json:"solucion"`
	CeldasVisitadas Camino  `json:"celdas_visitadas"`
}

// Coord estructura para la posición (Inicio, fin)
type Coord struct {
	Fila int `json:"fila"`
	Col  int `json:"col"`
}

// Camino estructura para el camino (conjunto de coordenadas o posiciones)
type Camino []Coord

//Estructura del laberinto
type MazeData struct {
	Mazes []Maze `json:"Mazes"`
}
