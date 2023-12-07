package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

/**
--- Day 7: Camel Cards ---

Nota: para resolver el problema aprovecho la libreria Sort de Go (bublesort,etc) y
se implementa la interface Sort.Interface con los metodos less(), len() y swap().

less() ordena primero por "type" y luego por "label" de la carta
rtfm: https://pkg.go.dev/sort#Interface
*/

const (
	High_card = iota
	One_pair
	Two_pair
	Three_of_a_kind
	Full_house
	Four_of_a_kind
	Five_of_a_kind
)

type Hand struct {
	Mano  string
	Monto int
	Tipo  int
}

//Global variable!! nice!! refaccctttt
var Labels = map[byte]int{'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1}

type Hands []Hand

func (h Hands) Len() int {
	return len(h)
}
func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h Hands) Less(i, j int) bool {
	//Veo el "Type" de cartas:
	if h[i].Tipo != h[j].Tipo {
		return h[i].Tipo < h[j].Tipo
	}
	//si son del mismo tipo, entonces debo ordenar por "Label":
	// true si es MENOR
	// 12345 < AAAAA  = True
	// QQ345 < QQAAA  = True
	// AA111 < TTTTT  = False
	carta1 := h[i].Mano
	carta2 := h[j].Mano

	for k := 0; k < 5; k++ {
		if carta1[k] == carta2[k] {
			continue
		}
		if Labels[carta1[k]] < Labels[carta2[k]] {
			//hand1 < hand2
			return true
		} else {
			return false
		}
	}
	return false

}

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

	//Parseo de las manos:
	hands := Hands{}
	for _, v := range input {
		h := Hand{}
		_, err := fmt.Sscanf(v, "%s %d", &h.Mano, &h.Monto)
		if err != nil {
			log.Fatal("error de parseo ", err)
		}

		h.Tipo = SetTipo(h.Mano, false) //6=Five of Kind, 5=Four of a kind,...,0 = High card
		hands = append(hands, h)
	}

	//Ordeno las cartas:
	sort.Sort(hands)

	//fmt.Printf("Pre: %v \n", hands)
	sum_pt1 := 0
	for i := 0; i < len(hands); i++ {
		sum_pt1 += (i + 1) * hands[i].Monto
	}

	fmt.Printf("Parte 1 = %d \n\n", sum_pt1)

	//Parte2: Cambio el comportamiento de la "J"
	Labels['J'] = 0

	//Parseo de las manos:

	for i, _ := range hands {
		hands[i].Tipo = SetTipo(hands[i].Mano, true) //6=Five of Kind, 5=Four of a kind,...,0 = High card
	}
	sort.Sort(hands)

	sum_pt2 := 0
	for i := 0; i < len(hands); i++ {
		sum_pt2 += (i + 1) * hands[i].Monto
	}

	//fmt.Printf("parte 2 = %v\n", hands)
	fmt.Printf("parte 2 = %v\n", sum_pt2)
}

func SetTipo(hand string, is_part2 bool) int {
	cartas := map[byte]int{}
	cant_J := 0
	for i := 0; i < 5; i++ {
		if is_part2 && hand[i] == 'J' {
			cant_J++
			continue
		}
		cartas[hand[i]]++

	}
	vals := []int{}
	for _, v := range cartas {
		vals = append(vals, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(vals)))

	if is_part2 {
		if cant_J == 5 {
			return Five_of_a_kind
		}
		vals[0] += cant_J
	}

	if vals[0] == 5 {
		return Five_of_a_kind
	}
	if vals[0] == 4 {
		return Four_of_a_kind
	}
	if vals[0] == 3 {
		if vals[1] == 2 {
			return Full_house
		}

		return Three_of_a_kind

	}

	if vals[0] == 2 {
		if vals[1] == 2 {
			return Two_pair
		}
		return One_pair
	}

	return High_card
}
