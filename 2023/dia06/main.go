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
--- Day 6: Wait For It ---
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

	//Parseo archivo
	Time := []int{}
	Distancia := []int{}

	t := strings.Fields(strings.Split(input[0], ":")[1]) //Time:
	d := strings.Fields(strings.Split(input[1], ":")[1]) //Distance:

	for _, v := range t {
		num, _ := strconv.Atoi(v)
		Time = append(Time, num)
	}
	for _, v := range d {
		num, _ := strconv.Atoi(v)
		Distancia = append(Distancia, num)
	}

	//parte 1: itero
	sum_pt1 := 1
	for i := 0; i < len(Time); i++ {
		count := simulador(Time[i], Distancia[i])
		if count > 0 {
			sum_pt1 *= count
		}
	}
	fmt.Printf("Parte 1 = %d \n\n", sum_pt1)

	//Parte2:
	//Junto los numeros de Time y Distancia

	tt, dd := "", ""
	for _, n := range Time {
		s := strconv.Itoa(n)
		tt += s
	}
	for _, n := range Distancia {
		s := strconv.Itoa(n)
		dd += s
	}
	Time_pt2, _ := strconv.Atoi(tt)
	Distancia_pt2, _ := strconv.Atoi(dd)

	sum_pt2 := 1
	count := simulador(Time_pt2, Distancia_pt2)
	if count > 0 {
		sum_pt2 *= count
	}
	fmt.Printf("parte 2 = %v\n", sum_pt2)
}

func simulador(tiempo, distancia int) (cantidad int) {
	count := 0
	for vel := 0; vel <= tiempo; vel++ {
		timeRestante := tiempo - vel
		dist := timeRestante * vel

		if dist > distancia {
			count++
		}
	}
	return count
}
