package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

const (
	Empty = 0
	Wall  = 3
	Start = 1
	End   = 2
)

func (a *App) CreateMaze(size int, mode bool) [][]int {
	var maze [][]int

	

	for i := 0; i < size; i++ {
		row := make([]int, size)
		for j := 0; j < size; j++ {
			row[j] = Wall
		}
		maze = append(maze, row)
	}

	rand.Seed(time.Now().UnixNano())

	if mode {

		endX, endY := rand.Intn(size), rand.Intn(size)
		maze[endX][endY] = End
		a.CarveMaze(maze, endX, endY)
		
	} else {

		startX := 1 + rand.Intn(size-2)
		startY := 1 + rand.Intn(size-2)
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

		a.ConnectPoints(maze, startX, startY, endX, endY)

		if startX == -1 || startY == -1 {
			maze = a.SetSPointRand(maze)
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

func (a *App) ConnectPoints(maze [][]int, startX, startY, endX, endY int) {
	a.CarveMaze(maze, endX, endY)
	if maze[startX][startY] == Wall {
		dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range dirs {
			nx, ny := startX+dir[0], startY+dir[1]
			if nx >= 0 && ny >= 0 && nx < len(maze) && ny < len(maze[0]) && maze[nx][ny] == Empty {
				maze[startX][startY] = Empty
				break
			}
		}
	}

	if !a.Backtracking(maze, startX, startY) {
		a.createDirectPath(maze, startX, startY, endX, endY)
	}
}

func (a *App) createDirectPath(maze [][]int, startX, startY, endX, endY int) {
	x, y := startX, startY

	for x != endX || y != endY {
		if x < endX {
			x++
		} else if x > endX {
			x--
		}

		if y < endY {
			y++
		} else if y > endY {
			y--
		}
		if maze[x][y] == Wall {
			maze[x][y] = Empty
		}
	}
}

func (a *App) Backtracking(maze [][]int, startX, startY int) bool {
	var visited [][]bool
	for range maze {
		var fila []bool
		for range maze[0] {
			fila = append(fila, false)
		}
		visited = append(visited, fila)
	}
	return a.backtrackHelper(maze, startX, startY, &visited)
}

func (a *App) backtrackHelper(maze [][]int, x, y int, visited *[][]bool) bool {
	if x < 0 || y < 0 || x >= len(maze) || y >= len(maze[0]) {
		return false
	}

	if maze[x][y] == Wall || (*visited)[x][y] {
		return false
	}

	if maze[x][y] == End {
		return true
	}

	(*visited)[x][y] = true

	if a.backtrackHelper(maze, x+1, y, visited) ||
		a.backtrackHelper(maze, x-1, y, visited) ||
		a.backtrackHelper(maze, x, y+1, visited) ||
		a.backtrackHelper(maze, x, y-1, visited) {
		return true
	}
	return false
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
