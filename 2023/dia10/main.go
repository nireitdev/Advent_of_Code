package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/**
--- Day 10: Pipe Maze ---

Part2 : To-Be-Done!!!!
*/

const (
	EAST = iota
	WEST
	NORTH
	SOUTH

	MAX_X, MAX_Y = 140, 140

	//MAX_X, MAX_Y = 20, 10
)

type Coord struct {
	x, y int
}

type Elemento struct {
	char byte
	Coord
	origen int //desde donde viene N,E,W,S
}

func main() {
	Mapa := map[Coord]byte{} //rune??? byte?? string??

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	input := []string{}
	for s.Scan() {
		input = append(input, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	//Parseo

	start := Coord{-1, -1}
	for i, s := range input {
		for j, v := range s {
			Mapa[Coord{x: j, y: i}] = byte(v)
			//Visitado[Coord{x: j, y: i}] = false
			if byte(v) == 'S' {
				start = Coord{x: j, y: i}
			}
		}
	}

	//Parte 1 : Empezando Desde "S" recorro solo el bucle completo y la distancia es la mitad = Total /2 :
	counter := 0
	Visitados := map[Coord]bool{}
	for _, elem := range []Elemento{
		Elemento{Mapa[Coord{start.x, start.y + 1}], Coord{start.x, start.y + 1}, NORTH},
		Elemento{Mapa[Coord{start.x, start.y + 1}], Coord{start.x, start.y + 1}, SOUTH},
		Elemento{Mapa[Coord{start.x, start.y + 1}], Coord{start.x, start.y + 1}, EAST},
		Elemento{Mapa[Coord{start.x, start.y + 1}], Coord{start.x, start.y + 1}, WEST},
	} {
		count := 1
		visit := map[Coord]bool{}
		visit[Coord{elem.x, elem.y}] = true
		for {
			nxt, valido := next(elem)
			if !valido {
				//no es un bucle: break!
				break
			}

			elem = nxt
			elem.char = Mapa[Coord{elem.x, elem.y}]
			visit[Coord{elem.x, elem.y}] = true
			count++
			if elem.char == 'S' {
				//Fin bucle:
				if count > counter {
					Visitados = visit
					counter = count
				}

				break
			}

		}
	}

	fmt.Printf("Parte 1 %#d\n", counter/2)
	//fmt.Printf("Parte 1 %#v\n", Visitados)

	// Parte2:
	//TO-BE-DONE!!!
	sum_pt2 := 0
	
	fmt.Printf("parte 2 = %v\n", sum_pt2)
}
func next(from Elemento) (to Elemento, valido bool) {

	if from.char == '|' {
		if from.origen == SOUTH {
			to.x = from.x
			to.origen = SOUTH
			if from.y > 0 {
				to.y = from.y - 1
			} else {
				to.y = 0
			}
			return to, true
		}
		if from.origen == NORTH {
			to.x = from.x
			to.origen = NORTH
			if from.y+1 < MAX_Y {
				to.y = from.y + 1
			} else {
				to.y = from.y
			}
			return to, true
		}

		//W y E no validos
		return to, false
	}

	if from.char == '-' {

		if from.origen == WEST {
			to.y = from.y
			to.origen = WEST
			if from.x+1 < MAX_X {
				to.x = from.x + 1
			} else {
				to.x = from.x
			}
			return to, true
		}
		if from.origen == EAST {
			to.y = from.y
			to.origen = EAST
			if from.x > 0 {
				to.x = from.x - 1
			} else {
				to.x = 0
			}
			return to, true
		}

		//no validos: SOUTH NORTH
		return to, false
	}

	if from.char == 'L' {
		if from.origen == EAST {
			//hacia north
			to.x = from.x
			to.origen = SOUTH
			if from.y > 0 {
				to.y = from.y - 1
			} else {
				to.y = 0
			}
			return to, true
		}
		if from.origen == NORTH {
			//hacia east
			to.y = from.y
			to.origen = WEST
			if from.x+1 < MAX_X {
				to.x = from.x + 1
			} else {
				to.x = from.x
			}
			return to, true
		}
		return to, false
	}

	if from.char == 'J' {
		if from.origen == NORTH {
			to.y = from.y
			to.origen = EAST
			//hacia west
			if from.x > 0 {
				to.x = from.x - 1
			} else {
				to.x = 0
			}
			return to, true
		}
		if from.origen == WEST {
			//hacia north
			to.origen = SOUTH
			to.x = from.x
			if from.y > 0 {
				to.y = from.y - 1
			} else {
				to.y = 0
			}
			return to, true
		}
		return to, false
	}

	if from.char == '7' {
		if from.origen == WEST {
			//hacia sur
			to.origen = NORTH
			to.x = from.x
			if from.y+1 < MAX_Y {
				to.y = from.y + 1
			} else {
				to.y = from.y
			}
			return to, true
		}
		if from.origen == SOUTH {
			//hacia west
			to.origen = EAST
			to.y = from.y
			if from.x > 0 {
				to.x = from.x - 1
			} else {
				to.x = 0
			}
			return to, true
		}
		return to, false
	}

	if from.char == 'F' {
		if from.origen == EAST {
			to.x = from.x
			to.origen = NORTH
			if from.y+1 < MAX_Y {
				to.y = from.y + 1
			} else {
				to.y = from.y
			}
			return to, true
		}
		if from.origen == SOUTH {
			to.y = from.y
			to.origen = WEST
			if from.x+1 < MAX_X {
				to.x = from.x + 1
			} else {
				to.x = from.x
			}
			return to, true
		}
		return to, false
	}

	/*
		| is a vertical pipe connecting north and south.
		- is a horizontal pipe connecting east and west.
		L is a 90-degree bend connecting north and east.
		J is a 90-degree bend connecting north and west.
		7 is a 90-degree bend connecting south and west.
		F is a 90-degree bend connecting south and east.
		. is ground; there is no pipe in this tile.
		S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
	*/

	//No validos: S , '.'
	return to, false
}
