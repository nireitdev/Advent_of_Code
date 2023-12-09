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
--- Day 9: Mirage Maintenance ---
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

	//Parseo
	secuencias := [][][]int{}
	for _, ss := range input {
		iters := [][]int{}
		iters = append(iters, AtoISlice(ss))

		//fmt.Printf("%v\n", iters)

		i := 1
		for {
			diffs := []int{}
			for j, v := range iters[i-1] {
				if j == 0 {
					continue
				}
				diffs = append(diffs, v-iters[i-1][j-1])
			}
			iters = append(iters, diffs)
			i++
			if SliceallCero(diffs) {
				//{0,0,0,0}
				break //for
			}
		}
		secuencias = append(secuencias, iters)
	}

	//Parte 1 : agrego al final un elemento
	sum_pt1 := 0

	for _, iters := range secuencias {
		//fmt.Printf("%#v \n", iters)
		//Una vez en cero vuelvo para atrás agregando un item en cada iters
		for i := len(iters) - 1; i >= 0; i-- {
			if i == len(iters)-1 {
				iters[i] = append(iters[i], 0)
				continue
			}
			last := iters[i][len(iters[i])-1]
			next_iter := iters[i+1][len(iters[i+1])-1]
			iters[i] = append(iters[i], last+next_iter)

		}
		//fmt.Printf("%#v \n", iters)
		sum_pt1 += iters[0][len(iters[0])-1]

	}

	fmt.Printf("Parte 1 = %d \n\n", sum_pt1)

	// Parte2: Agregar un nuevo elemento adelante del slice y
	//resta el primer elemento de cada slice:

	sum_pt2 := 0
	for _, iters := range secuencias {
		for i := len(iters) - 1; i >= 0; i-- {
			if i == len(iters)-1 {
				iters[i] = append([]int{0}, iters[i]...)
				continue
			}
			last := iters[i][0]
			past_iter := iters[i+1][0]
			iters[i] = append([]int{last - past_iter}, iters[i]...)

		}
		sum_pt2 += iters[0][0]
	}

	fmt.Printf("parte 2 = %v\n", sum_pt2)
}

func AtoISlice(st string) []int {
	strs := strings.Fields(st)
	res := []int{}
	for _, v := range strs {
		n, _ := strconv.Atoi(v)
		res = append(res, n)
	}
	return res
}

func SliceallCero(sl []int) bool {
	all_ceros := true
	for _, v := range sl {
		if v != 0 {
			all_ceros = false
		}
	}
	return all_ceros
}
