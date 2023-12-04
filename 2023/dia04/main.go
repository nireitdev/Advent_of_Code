package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

/**
--- Day 4: Scratchcards ---
*/

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

	win_cards := map[int][]int{}
	personal_cards := map[int][]int{}

	//Parseo archivo
	for _, ss := range input {
		//Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		r := strings.Split(ss, ":") //card
		s := strings.Split(r[1], "|")
		w := strings.Fields(s[0]) //win
		p := strings.Fields(s[1]) //personal
		num_card, err := strconv.Atoi(strings.Trim(r[0][4:], " "))
		if err != nil {
			log.Fatal(err)
		}

		for _, i := range w {
			n, _ := strconv.Atoi(i)
			win_cards[num_card] = append(win_cards[num_card], n)
		}
		for _, i := range p {
			n, _ := strconv.Atoi(i)
			personal_cards[num_card] = append(personal_cards[num_card], n)
		}

	}

	//parte 1: itero por cada set de "cards"
	suma_pt1 := 0

	for cardnum := 0; cardnum < len(win_cards); cardnum++ {
		inc_puntos := -1.0
		for _, w := range win_cards[cardnum] {
			for _, p := range personal_cards[cardnum] {
				if w == p {
					inc_puntos += 1.0
				}
			}
		}
		suma_pt1 += int(math.Pow(2.0, inc_puntos)) //POW() only floats!!!
	}

	fmt.Printf("Parte 1: suma total de puntos ganados= %d \n", int(suma_pt1))

	//Parte2: es lo mismo que la parte 1, solo que se van contando las cartas duplicadas
	suma_pt2 := 0

	//Cantidad total de "cards"  = originales + duplicados
	total_cards := map[int]int{}
	for c, _ := range win_cards {
		total_cards[c] = 1 //originales
	}

	for cardnum := 0; cardnum < len(win_cards); cardnum++ {
		nuevas_cards := 0
		for _, w := range win_cards[cardnum] {
			for _, p := range personal_cards[cardnum] {
				if w == p {
					nuevas_cards++
				}
			}
		}

		for i := 1; i <= nuevas_cards; i++ {

			//incremento nuevas_cards por la cantidad de total_cards que tengo duplicadas
			incr := 1
			if total_cards[cardnum] > 0 {
				incr = total_cards[cardnum]
			}

			//solo duplico "cards" validas
			if cardnum+i <= len(win_cards) {
				total_cards[cardnum+i] += incr
			}

		}

	}
	//fmt.Printf("%v \n", total_cards)

	for _, cant := range total_cards {
		suma_pt2 += cant
	}
	fmt.Printf("Parte 2: suma total de 'total scratchcards'= %d \n", suma_pt2)
}
