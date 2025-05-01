package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"slices"
)

// Constants to represent the different types of cells in the maze.
const (
	Path  = 0 //Path
	Wall  = 3 //Wall
	Start = 1 //Start point
	End   = 2 //End point
)

// CreateMaze generates a maze with a given size and mode (manual or automatic)
// Params:
// - size int Size of the maze
// - mode bool The game mode
// Return:
// - Maze [][]int The maze with all configurations
func (a *App) CreateMaze(size int, mode bool) [][]int {
	maze := a.CreateMatrix(size)         //Create an initial matrix full of walls
	a.SetStartEndPoint(maze, size, mode) // Set the start and end points
	return maze
}

// CreateMatrix creates a square matrix filled with walls
// Params:
// - size int (size of the maze)
// Returns:
// - Matrix [][] The maze with the specified dimensions/size
func (a *App) CreateMatrix(size int) [][]int {
	var matrix [][]int
	// Loop through the size to create rows.
	for range size {
		var row []int
		// Fill each row with walls
		for range size {
			row = append(row, Wall)
		}
		// Add the row to the matrix
		matrix = append(matrix, row)
	}
	return matrix
}

// CarveMaze creates paths in the maze starting from a given point.
// Params:
// - maze [][]int: The maze where the paths will be generated.
// - row int: The row of the starting point.
// - col int: The column of the starting point.
// Returns:
// - None. The function modifies the maze in place.
func (a *App) CarveMaze(maze [][]int, row, col int) {
	// Define possible movement directions (up, down, left, right).
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// Randomize the order of directions to create varied maze patterns.
	rand.Shuffle(len(dirs), func(i, j int) { dirs[i], dirs[j] = dirs[j], dirs[i] })
	// First pass: Carve out main paths.
	for _, dir := range dirs {
		newRow, newCol := row+2*dir[0], col+2*dir[1]
		// Check if the new position is within bounds and is a wall.
		if newRow >= 0 && newCol >= 0 && newRow < len(maze) && newCol < len(maze[0]) && maze[newRow][newCol] == Wall {
			// Carve a path to the new position.
			maze[row+dir[0]][col+dir[1]] = Path
			maze[newRow][newCol] = Path
			// Recursively carve paths from the new position.
			a.CarveMaze(maze, newRow, newCol)
		}
	}
	// Second pass: Add some random connections for complexity.
	for _, dir := range dirs {
		newRow, newCol := row+2*dir[0], col+2*dir[1]
		midRow, midCol := row+dir[0], col+dir[1]
		if newRow >= 0 && newCol >= 0 && newRow < len(maze) && newCol < len(maze[0]) {
			if maze[newRow][newCol] == Path && maze[midRow][midCol] == Wall {
				if rand.Float64() < 0.05 { // 5% chance to carve out this middle wall.
					maze[midRow][midCol] = Path
				}
			}
		}
	}
}

// Setters

// SetStartEndPoint sets the start and end points in the maze based on the mode.
// Params:
// - maze [][]int: The maze where the points will be set.
// - size int: The size of the maze.
// - mode bool: The game mode (true for automatic, false for manual).
// Returns:
// - None. The function modifies the maze in place.
func (a *App) SetStartEndPoint(maze [][]int, size int, mode bool) {
	var endRow, endCol int
	if mode {
		// Automatic mode: Only set the end point.
		endRow, endCol = rand.Intn(size), rand.Intn(size)
		maze[endRow][endCol] = End
		a.CarveMaze(maze, endRow, endCol)
	} else {
		// Manual mode: Set both start and end points.
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

// SetStartPoint sets a new starting point in the maze.
//
// Parameters:
// - maze [][]int: The maze where the start point going to be placed
// - row int: The row where you want to set the starting point.
// - col int: The column where you want to set the starting point.
//
// Returns:
// - A copy of the maze with the new starting point set,
// provided the specified cell is not a Wall or an End.
// If the cell is invalid, the maze is returned unchanged at that position.
func (a *App) SetStartPoint(maze [][]int, row, col int) [][]int {
	maze = a.DeleteStartPoint(maze)
	if maze[row][col] != Wall && maze[row][col] != End {
		maze[row][col] = Start
	}
	return maze
}

// SetStartPointRandom randomly selects a cell in the maze to set a new starting point
// Params:
// - maze [][]int: The maze where the start point going to be placed
// Returns:
// - The maze with the new random starting point set to a Path cell.
// If a starting point already existed, it is removed.
// The random selection is repeated until a valid Path cell is found.
func (a *App) SetStartPointRandom(maze [][]int) [][]int {
	a.DeleteStartPoint(maze)
	for {
		startRow, startCol := rand.Intn(len(maze)), rand.Intn(len(maze[0]))
		if maze[startRow][startCol] == Path {
			maze[startRow][startCol] = Start
			break
		}
	}
	return maze
}

// DeleteStartPoint removes the starting point from the maze by converting it back to a path cell.
// Params:
// - maze [][]int: The maze where the starting point will be removed.
// Returns:
// - [][]int: The updated maze with the starting point replaced by a path.
func (a *App) DeleteStartPoint(maze [][]int) [][]int {
	for row := range maze {
		for col := range maze[row] {
			// If the current cell is the starting point, replace it with a path.
			if maze[row][col] == Start {
				maze[row][col] = Path
			}
		}
	}
	return maze
}

// Getters

// GetStartPoint retrieves the coordinates of the starting point in the maze.
// Params:
// - maze [][]int: The maze to search for the starting point.
// Returns:
// - int, int: The row and column of the starting point, or (-1, -1) if not found.
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

// GetEndPoint retrieves the coordinates of the ending point in the maze.
// Params:
// - maze [][]int: The maze to search for the ending point.
// Returns:
// - int, int: The row and column of the ending point, or (-1, -1) if not found.
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

// SaveMaze saves the maze to a JSON file.
// Params:
// - maze [][]int: The maze to save.
// - name string: The name of the maze.
// - mode bool: The game mode (manual or automatic).
// Returns:
// - None.
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

// Validation
// MazeExist checks if there are any saved mazes in the JSON file.
// Returns:
// - bool: True if there are saved mazes, false otherwise.
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

// GetPaths finds all possible paths from a starting point to the end point in the maze.
// Params:
// - maze [][]int: The maze to search for paths.
// - row int: The starting row.
// - col int: The starting column.
// Returns:
// - []Camino: A list of paths found in the maze.
func (a *App) GetPaths(maze [][]int, row, col int) []Camino {
	var paths []Camino
	var currentPath Camino
	visited := a.CreateBooleanMatrix(len(maze))
	a.GetPathsAux(maze, row, col, &visited, &currentPath, &paths)
	return paths
}

// GetPathsAux is a helper function to recursively find paths in the maze.
// Params:
// - maze [][]int: The maze to search for paths.
// - row int: The current row.
// - col int: The current column.
// - visited *[][]bool: A matrix to track visited cells.
// - currentPath *Camino: The current path being explored.
// - paths *[]Camino: The list of paths found so far.
// Returns:
// - None.
func (a *App) GetPathsAux(maze [][]int, row, col int, visited *[][]bool, currentPath *Camino, paths *[]Camino) {
	pathCount := len(*paths)
	if pathCount >= 5 { // Limit the number of paths to 5.
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

// GetPath finds a single path from the starting point to the end point in the maze using BFS.
// Params:
// - maze [][]int: The maze to search for a path.
// - startRow int: The row of the starting point.
// - startCol int: The column of the starting point.
// Returns:
// - Camino: A path as a list of coordinates. Returns an empty path if no path is found.
func (a *App) GetPath(maze [][]int, startRow, startCol int) Camino {
	queue := []BFS{{path: Camino{{Fila: startRow, Col: startCol}}, steps: 0}}
	visited := a.CreateBooleanMatrix(len(maze))
	visited[startRow][startCol] = true
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // Directions: up, down, left, right
	var path Camino
	for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			lastPos := current.path[len(current.path)-1]
			path = append(path, lastPos)
			if maze[lastPos.Fila][lastPos.Col] == End {
					return path
			}
			for _, dir := range dirs {
					newRow, newCol := lastPos.Fila+dir[0], lastPos.Col+dir[1]
					if newRow >= 0 && newCol >= 0 && newRow < len(maze) && newCol < len(maze[0]) &&
							maze[newRow][newCol] != Wall && !visited[newRow][newCol] {
							visited[newRow][newCol] = true
							newPath := slices.Clone(current.path)
							newPath = append(newPath, Coord{Fila: newRow, Col: newCol})
							queue = append(queue, BFS{path: newPath, steps: current.steps + 1})
					}
			}
	}
	return path
}


// GetBestPath finds the shortest path from the starting point to the end point in the maze using BFS.
// Params:
// - maze [][]int: The maze to search for the shortest path.
// - startRow int: The row of the starting point.
// - startCol int: The column of the starting point.
// Returns:
// - Camino: The shortest path as a list of coordinates. Returns an empty path if no path is found.
func (a *App) GetBestPath(maze [][]int, startRow, startCol int) Camino {
	queue := []BFS{{path: Camino{{Fila: startRow, Col: startCol}}, steps: 0}}
	visited := a.CreateBooleanMatrix(len(maze))
	visited[startRow][startCol] = true
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			lastPos := current.path[len(current.path)-1]
			if maze[lastPos.Fila][lastPos.Col] == End {
					return current.path
			}
			for _, dir := range dirs {
					newRow, newCol := lastPos.Fila+dir[0], lastPos.Col+dir[1]
					if newRow >= 0 && newCol >= 0 && newRow < len(maze) && newCol < len(maze[0]) &&
							maze[newRow][newCol] != Wall && !visited[newRow][newCol] {
							visited[newRow][newCol] = true
							newPath := slices.Clone(current.path)
							newPath = append(newPath, Coord{Fila: newRow, Col: newCol})
							queue = append(queue, BFS{path: newPath, steps: current.steps + 1})
					}
			}
	}
	return Camino{}
}

// GetWorstPath finds the longest path from the starting point to the end point in the maze.
// Params:
// - maze [][]int: The maze to search for the longest path.
// - row int: The row of the starting point.
// - col int: The column of the starting point.
// Returns:
// - Camino: The longest path as a list of coordinates. Returns an empty path if no path is found.
func (a *App) GetWorstPath(maze [][]int, row, col int) Camino {
	var worstPath Camino
	visited := a.CreateBooleanMatrix(len(maze))
	var currentPath Camino
	a.GetWorstPathAux(maze, row, col, visited, &currentPath, &worstPath)
	return worstPath
}

// GetWorstPathAux is a helper function to recursively find the longest path in the maze.
// Params:
// - maze [][]int: The maze to search for the longest path.
// - row int: The current row.
// - col int: The current column.
// - visited [][]bool: A matrix to track visited cells.
// - currentPath *Camino: The current path being explored.
// - worstPath *Camino: The longest path found so far.
// Returns:
// - None. The function modifies the worstPath in place.
func (a *App) GetWorstPathAux(maze [][]int, row, col int, visited [][]bool, currentPath *Camino, worstPath *Camino) {
	if row < 0 || col < 0 || row >= len(maze) || col >= len(maze[0]) || visited[row][col] || maze[row][col] == Wall {
			return
	}
	*currentPath = append(*currentPath, Coord{Fila: row, Col: col})
	visited[row][col] = true
	if maze[row][col] == End {
			if len(*currentPath) > len(*worstPath) {
					*worstPath = (*worstPath)[:0]
					*worstPath = append(*worstPath, *currentPath...)
			}
	} else {
			dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // Directions: up, down, left, right
			for _, dir := range dirs {
					newRow, newCol := row+dir[0], col+dir[1]
					a.GetWorstPathAux(maze, newRow, newCol, visited, currentPath, worstPath)
			}
	}
	*currentPath = (*currentPath)[:len(*currentPath)-1]
	visited[row][col] = false
}

// CreateBooleanMatrix creates a boolean matrix to track visited cells in the maze.
// Params:
// - size int: The size of the maze (number of rows and columns).
// Returns:
// - [][]bool: A boolean matrix initialized to false.
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