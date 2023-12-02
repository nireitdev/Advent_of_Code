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
--- Day 1: Trebuchet?! ---
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

	//parte 1
	suma := 0

	for _, v := range input {
		numero := BuscaNumeros(v)
		suma += numero
	}
	fmt.Printf("Parte 1: suma calibraciones= %d \n", suma)

	//Parte2: Los numeros pueden estar escritos en letras
	//mismo que la parte 1 , pero antes reemplazo strings
	//Trampa del problema:
	//		The right calibration values for string "eighthree" is 83 and for "sevenine" is 79.

	letras := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	suma_pt2 := 0

	for _, ss := range input {
		ss := ss

		//Busco el primer numero mas a la izquierda:
		idmin, min := math.MaxInt32, ""
		for k, v := range letras {
			pos := strings.Index(ss, k)
			if pos == -1 {
				continue
			}
			if pos < idmin {
				idmin = pos
				min = v
			}
		}

		//Busco el ultimo numero mas a la derecha:
		idmax, max := -1, ""
		for k, v := range letras {
			pos := strings.LastIndex(ss, k)
			if pos == -1 {
				continue
			}
			if pos > idmax {
				idmax = pos
				max = v
			}
		}
		fmt.Printf("%s => %s%s \n", ss, min, max)
		num_pt2, _ := strconv.Atoi(fmt.Sprintf("%s%s", min, max))

		suma_pt2 += num_pt2
	}
	fmt.Printf("Parte 2: suma calibraciones= %d \n", suma_pt2)
}

func BuscaNumeros(input string) (numero int) {

	min, max := "", ""

	for _, ch := range input {
		if _, err := strconv.Atoi(string(ch)); err == nil {
			if min == "" {
				min = string(ch)
			} else {
				max = string(ch)
			}
		}
	}
	if max == "" {
		max = min
	}
	num, _ := strconv.Atoi(fmt.Sprintf("%s%s", min, max))

	return num
}
