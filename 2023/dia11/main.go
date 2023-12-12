package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

/**
--- Day 11: Cosmic Expansion ---
*/

type Coord struct {
	X, Y int
}

//const EXPANSE = 1   		 	//Parte 1
const EXPANSE = 1000000 - 1 //Parte 2

func main() {

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
	Universo := []Coord{}

	//Genero un listado de expansion en las X
	expandirX := map[int]bool{}
	for x := 0; x < len(input[0]); x++ {
		vacio := true
		for y := 0; y < len(input); y++ {
			if input[y][x] == '#' {
				vacio = false
				break
			}
		}
		if vacio {
			expandirX[x] = true
		}
	}

	j, y := 0, 0
	for {

		if strings.Contains(input[j], "#") {
			for i, ch := range input[j] {
				x := i
				//desplazo en X:
				for k, _ := range expandirX {
					if i > k {
						x += EXPANSE
					}
				}
				if ch == '#' {
					Universo = append(Universo, Coord{x, y})
				}
			}

		} else {
			//expando el Universo en el sentido de las Y
			y += EXPANSE
		}

		j++
		y++
		if j == len(input) {
			break
		}
	}

	//fmt.Printf("%v", Universo)

	sum_pt1 := 0
	// calculo distancias
	for i := 0; i < len(Universo); i++ {
		for j = i + 1; j < len(Universo); j++ {
			sum_pt1 += dist(Universo[i], Universo[j])
		}
	}

	fmt.Printf("Suma Distancias %#d\n", sum_pt1)

}

func dist(pos1, pos2 Coord) int {
	difx := math.Abs(float64(pos1.X - pos2.X))
	dify := math.Abs(float64(pos1.Y - pos2.Y))
	return int(difx + dify)
}
