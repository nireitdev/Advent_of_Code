package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

/**
--- Day 8: Haunted Wasteland ---

RTFM: https://nickymeuleman.netlify.app/garden/aoc2023-day08

*/

type Node struct {
	L string
	R string
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

	//parseo
	comms := ""
	nodes := map[string]Node{}

	for i, v := range input {
		if i == 0 {
			comms = v
			continue
		}

		n := Node{}
		c := ""
		_, err := fmt.Sscanf(v, "%s = (%s %s)", &c, &n.L, &n.R) //AAA = (BBB, CCC)
		if err != nil {
			//continue
		}
		n.L = strings.Replace(n.L, ",", "", 1) //me queda una "," despues del sscanf
		n.R = strings.Replace(n.R, ")", "", 1) //me queda una "," despues del sscanf

		nodes[c] = n

	}
	sum_pt1 := 0

	next := "AAA"
	i := 0
	for {
		if next == "ZZZ" {
			break
		}
		if comms[i] == 'L' {
			next = nodes[next].L
		} else {
			next = nodes[next].R
		}

		sum_pt1++

		i++
		if i >= len(comms) {
			i = 0
		}

	}

	fmt.Printf("Parte 1 = %d \n\n", sum_pt1)

	//Parte2: Utilizo LCM()
	// por cada __A existe UN SOLO __Z que se repite por siempre.
	// RTFM: https://nickymeuleman.netlify.app/garden/aoc2023-day08

	nexts := []string{}
	for k, _ := range nodes {
		if len(k) == 3 && k[2] == 'A' {
			nexts = append(nexts, k)
		}
	}
	//fmt.Printf(" %#v \n", nexts)

	resp := make([]int, len(nexts))

	var wg sync.WaitGroup
	wg.Add(len(nexts))

	//Busco el primer __Z por cada __A (Go-coroutines)
	for j, nxt := range nexts {
		go func(nro int, start string) {
			defer wg.Done()
			found := 0
			next := start
			i := 0
			for {
				if comms[i] == 'L' {
					next = nodes[next].L
				} else {
					next = nodes[next].R
				}
				found++
				if next[2] == 'Z' {
					fmt.Printf("Go_func() nro: %d  Inicio: %s Fin:%s  Iters:%d \n", nro, start, next, found)
					resp[nro] = found //shared memory!!!
					return
				}
				i++
				if i >= len(comms) {
					i = 0
				}
			}

		}(j, nxt)
	}

	wg.Wait()

	fmt.Printf("parte 2 = %v\n", lcd(resp))
}

func lcd(nums []int) int {
	a := 1
	for _, n := range nums {
		a = (a * n) / gcd(a, n)
	}
	return a
}
func gcd(a int, b int) int {
	for b > 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
