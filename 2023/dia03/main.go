package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/**
--- Day 3: Gear Ratios ---
*/

func main() {

	f, err := os.Open("E:\\Datos\\Git\\adventofcode.com\\2023\\day03\\input.txt")
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

	//parsing
	board := []byte{}
	for _, s := range input {
		board = append(board, s...)
	}

	sum_pt1, sum_pt2 := 0, 0
	gears := map[int][]int{}
	largo, ancho := len(input[0]), len(input)

	PrintBoard(board, largo, ancho)

	for y := 0; y < largo; y++ {
		num := []byte{}
		is_partnumber := false

		//guarda las posiciones de "*" alrededor. Uso un set para guardar SOLO
		//una vez sus coordenadas y que no duplique
		gears_alrededor := map[int]bool{}

		fin_numero := false

		for x := 0; x < ancho; x++ {
			ch := board[y*ancho+x]
			if strings.Count("0123456789", string(ch)) > 0 {
				num = append(num, ch)

				//busco por simbolos en la adjacencia:
				for j := -1; j <= 1; j++ {
					for i := -1; i <= 1; i++ {
						u := x + i
						v := y + j
						if u >= 0 && v >= 0 && u < ancho && v < largo {
							ss := string(board[v*ancho+u])
							if ss != "." && !strings.ContainsAny("0123456789", ss) {
								is_partnumber = true
							}
							if ss == "*" {
								gears_alrededor[v*ancho+u] = true
							}
						}

					}
				}
				if x+1 == ancho {
					//fin del numero debido a que termina en el borde derecho
					fin_numero = true
				}
			} else {
				//fin del numero debido a que aparece un simbolo
				fin_numero = true
			}

			if fin_numero {
				//fin del numero
				if is_partnumber {
					n, _ := strconv.Atoi(string(num))
					sum_pt1 += n

				}

				for k, _ := range gears_alrededor {
					n, _ := strconv.Atoi(string(num))
					gears[k] = append(gears[k], n)
				}

				//clear
				num = []byte{}
				is_partnumber = false
				fin_numero = false
				gears_alrededor = map[int]bool{}
			}

		}
		//fmt.Println(string(num))
	}

	//parte 1:

	fmt.Printf("Parte 1: suma de los \"part numbers\"= %d \n", sum_pt1)

	//parte 2:

	for _, v := range gears {
		if len(v) == 2 {
			sum_pt2 += v[0] * v[1]
		}
	}

	fmt.Printf("Parte 2: suma de los \"gears ratios \"= %d \n", sum_pt2)
}

func PrintBoard(board []byte, ancho int, largo int) {
	fmt.Printf("\n")
	for y := 0; y < largo; y++ {
		fmt.Printf("%s\n", board[y*ancho:(y+1)*ancho])
	}
	fmt.Printf("\n")
}
