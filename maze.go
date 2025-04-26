package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"slices"
)

const (
	Path  = 0
	Wall  = 3
	Start = 1
	End   = 2
)

//Functions about the maze generation

func (a *App) CreateMaze(size int, mode bool) [][]int {
	maze := a.CreateMatrix(size)
	a.SetStartEndPoint(maze, size, mode)
	return maze
}

func (a *App) CreateMatrix(size int) [][]int {
	var matrix [][]int
	for range size {
		var row []int
		for range size {
			row = append(row, Wall)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func (a *App) CarveMaze(maze [][]int, row, col int) {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rand.Shuffle(len(dirs), func(i, j int) { dirs[i], dirs[j] = dirs[j], dirs[i] })
	for _, dir := range dirs {
		newRow, newCol := row+2*dir[0], col+2*dir[1]
		if newRow >= 0 && newCol >= 0 && newRow < len(maze) && newCol < len(maze[0]) && maze[newRow][newCol] == Wall {
			maze[row+dir[0]][col+dir[1]] = Path
			maze[newRow][newCol] = Path
			a.CarveMaze(maze, newRow, newCol)
		}
	}
	for _, dir := range dirs {
		newRow, newCol := row+2*dir[0], col+2*dir[1]
		midRow, midCol := row+dir[0], col+dir[1]

		if newRow >= 0 && newCol >= 0 && newRow < len(maze) && newCol < len(maze[0]) {
			if maze[newRow][newCol] == Path && maze[midRow][midCol] == Wall {
				if rand.Float64() < 0.1 {
					maze[midRow][midCol] = Path
				}
			}
		}
	}
}

// Setters
func (a *App) SetStartEndPoint(maze [][]int, size int, mode bool) {
	var endRow, endCol int
	if mode {
		endRow, endCol := rand.Intn(size), rand.Intn(size)
		maze[endRow][endCol] = End
		a.CarveMaze(maze, endRow, endCol)

	} else {
		startRow, startCol := rand.Intn(size), rand.Intn(size)
		a.CarveMaze(maze, startRow, startCol)
		for {
			startRow, startCol = rand.Intn(size), rand.Intn(size)
			if maze[startRow][startCol] == Path {
				maze[startRow][startCol] = Start
				break
			}
		}
		for {
			endRow, endCol = rand.Intn(size), rand.Intn(size)
			if (endRow != startRow || endCol != startCol) && maze[endRow][endCol] == Path {
				maze[endRow][endCol] = End
				break
			}
		}
	}
}

func (a *App) SetStartPoint(maze [][]int, row, col int) [][]int {
	maze = a.DeleteStartPoint(maze)
	if maze[row][col] != Wall && maze[row][col] != End {
		maze[row][col] = Start
	}
	return maze
}

func (a *App) SetStartPointRandom(maze [][]int) [][]int {
	for row := range maze {
		for col := range maze[row] {
			if maze[row][col] == Start {
				maze[row][col] = Path
			}
		}
	}
	for {
		startRow, startCol := rand.Intn(len(maze)), rand.Intn(len(maze[0]))
		if maze[startRow][startCol] == Path {
			maze[startRow][startCol] = Start
			break
		}
	}
	return maze
}

func (a *App) DeleteStartPoint(maze [][]int) [][]int {
	for row := range maze {
		for col := range maze[row] {
			if maze[row][col] == Start {
				maze[row][col] = Path
			}
		}
	}
	return maze
}

// Getters
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

// Load and save maze functions
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

// Save maze and load maze functions
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

// Validations
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

// GetPaths function
func (a *App) GetPaths(maze [][]int, row, col int) []Camino {
	var paths []Camino
	var currentPath Camino
	visited := a.CreateBooleanMatrix(len(maze))
	a.GetPathsAux(maze, row, col, &visited, &currentPath, &paths)
	return paths
}

func (a *App) GetPathsAux(maze [][]int, row, col int, visited *[][]bool, currentPath *Camino, paths *[]Camino) {
	pathCount := len(*paths)
	if pathCount >= 5 {
		return
	}
	if row < 0 || col < 0 || row >= len(maze) || col >= len(maze[0]) {
		return
	}
	*currentPath = append(*currentPath, Coord{Fila: row, Col: col})
	if maze[row][col] == End {
		var newPath Camino
		for _, coord := range *currentPath {
			newPath = append(newPath, Coord{Fila: coord.Fila, Col: coord.Col})
		}
		*paths = append(*paths, newPath)
		*currentPath = (*currentPath)[:len(*currentPath)-1]
		return
	}

	if (*visited)[row][col] || maze[row][col] == Wall {
		*currentPath = (*currentPath)[:len(*currentPath)-1]
		return
	}

	(*visited)[row][col] = true
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range dirs {
		newRow, newCol := row+dir[0], col+dir[1]
		if newRow >= 0 && newCol >= 0 && newRow < len(maze) && newCol < len(maze[0]) {
			a.GetPathsAux(maze, newRow, newCol, visited, currentPath, paths)
		}
	}

	(*visited)[row][col] = false
	*currentPath = (*currentPath)[:len(*currentPath)-1]
}

func (a *App) GetBestPath(maze [][]int, row, col int) Camino {
	queue := []Camino{{Coord{Fila: row, Col: col}}}
	visited := a.CreateBooleanMatrix(len(maze))
	visited[row][col] = true
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]
		lastPos := currentPath[len(currentPath)-1]
		row, col := lastPos.Fila, lastPos.Col
		if maze[row][col] == End {
			return currentPath
		}
		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newCol >= 0 && newRow < len(maze) && newCol < len(maze[0]) && maze[newRow][newCol] != Wall && !visited[newRow][newCol] {
				visited[newRow][newCol] = true
				newPath := slices.Clone(currentPath)
				newPath = append(newPath, Coord{Fila: newRow, Col: newCol})
				queue = append(queue, newPath)
			}
		}
	}
	return Camino{}
}

func (a *App) CreateBooleanMatrix(size int) [][]bool {
	var matrix [][]bool
	for range size {
		var row []bool
		for range size {
			row = append(row, false)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
