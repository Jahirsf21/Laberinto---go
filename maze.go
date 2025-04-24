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
	var maze [][]int

	for range size {
		row := make([]int, size)
		for j := range size {
			row[j] = Wall
		}
		maze = append(maze, row)
	}

	if mode {

		endX, endY := rand.Intn(size), rand.Intn(size)
		maze[endX][endY] = End
		a.CarveMaze(maze, endX, endY)

	} else {
		startX, startY := rand.Intn(size), rand.Intn(size)
		maze[startX][startY] = Empty
		a.CarveMaze(maze, startX, startY)
		for {
			startX, startY = rand.Intn(size), rand.Intn(size)
			if maze[startX][startY] == Empty {
				maze[startX][startY] = Start
				break
			}
		}
		var endX, endY int
		for {
			endX, endY = rand.Intn(size), rand.Intn(size)
			if (endX != startX || endY != startY) && maze[endX][endY] == Empty {
				maze[endX][endY] = End
				break
			}
		}
	}
	return maze
}

func (a *App) CarveMaze(maze [][]int, x, y int) {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rand.Shuffle(len(dirs), func(i, j int) { dirs[i], dirs[j] = dirs[j], dirs[i] })

	for _, dir := range dirs {
		nx, ny := x+2*dir[0], y+2*dir[1]

		if nx >= 0 && ny >= 0 && nx < len(maze) && ny < len(maze[0]) && maze[nx][ny] == Wall {

			maze[x+dir[0]][y+dir[1]] = Empty
			maze[nx][ny] = Empty
			a.CarveMaze(maze, nx, ny)
		}
	}
	for _, dir := range dirs {
		nx, ny := x+2*dir[0], y+2*dir[1]
		mx, my := x+dir[0], y+dir[1]
		if nx >= 0 && ny >= 0 && nx < len(maze) && ny < len(maze[0]) {
			if maze[nx][ny] == Empty && maze[mx][my] == Wall {
				if rand.Float64() < 0.1 {
					maze[mx][my] = Empty
				}
			}
		}
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
	newMaze := Maze{
		Name:   name,
		Mode:   mode,
		Matrix: maze,
		End:    Coord{Fila: endX, Col: endY},
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

func (a *App) SetSPointRand(maze [][]int) [][]int {
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == Start {
				maze[i][j] = Empty
			}
		}
	}
	startX, startY := rand.Intn(len(maze)), rand.Intn(len(maze))
	if maze[startX][startY] != Wall && maze[startX][startY] != End {
		maze[startX][startY] = Start
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

func (a *App) GetAllPaths(maze [][]int, startX, startY int) []Camino {
	var paths []Camino
	var currentPath Camino
	var visited [][]bool
	for range maze {
		var fila []bool
		for range maze[0] {
			fila = append(fila, false)
		}
		visited = append(visited, fila)
	}
	a.FindPaths(maze, startX, startY, &visited, &currentPath, &paths, 5)
	return paths
}

func (a *App) GetBestPath(maze [][]int, startX, startY int) Camino {
	allPaths := a.GetAllPaths(maze, startX, startY)
	if len(allPaths) == 0 {
		return Camino{}
	}
	bestPath := allPaths[0]
	minLength := len(bestPath)
	for _, path := range allPaths {
		if len(path) < minLength {
			bestPath = path
			minLength = len(path)
		}
	}
	return bestPath
}

func (a *App) FindPaths(maze [][]int, x, y int, visited *[][]bool, currentPath *Camino, paths *[]Camino, maxPaths int) {
	pathCount := len(*paths)
	if pathCount >= maxPaths {
		return
	}
	if x < 0 || y < 0 || x >= len(maze) || y >= len(maze[0]) || maze[x][y] == Wall || (*visited)[x][y] {
		return
	}
	*currentPath = append(*currentPath, Coord{Fila: x, Col: y})
	if maze[x][y] == End {
		var pathCopy Camino
		for _, coord := range *currentPath {
			pathCopy = append(pathCopy, Coord{Fila: coord.Fila, Col: coord.Col})
		}
		*paths = append(*paths, pathCopy)
		*currentPath = (*currentPath)[:len(*currentPath)-1]
		return
	}
	(*visited)[x][y] = true
	a.FindPaths(maze, x+1, y, visited, currentPath, paths, maxPaths)
	a.FindPaths(maze, x-1, y, visited, currentPath, paths, maxPaths)
	a.FindPaths(maze, x, y+1, visited, currentPath, paths, maxPaths)
	a.FindPaths(maze, x, y-1, visited, currentPath, paths, maxPaths)
	(*visited)[x][y] = false
	*currentPath = (*currentPath)[:len(*currentPath)-1]
}
