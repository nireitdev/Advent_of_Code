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
--- Day 2: Cube Conundrum ---
*/

type Ball struct {
	Color    string
	Cantidad int
}

type Play struct {
	Balls []Ball
}

type Game struct {
	Numero int
	Plays  []Play
}

var Colores = map[string]int{"red": 12, "blue": 14, "green": 13}

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

	games := []Game{}

	//Parseo archivo
	for _, ss := range input {
		ll := strings.Split(ss, ":")
		nrogame, _ := strconv.Atoi(ll[0][5:])
		game := Game{Numero: nrogame}
		sp := strings.Split(ll[1], ";")
		for _, p := range sp {
			sp := strings.Split(p, ",")
			play := Play{}
			for _, b := range sp {
				sf := strings.Fields(b)
				cant, _ := strconv.Atoi(sf[0])
				ball := Ball{Color: sf[1], Cantidad: cant}
				play.Balls = append(play.Balls, ball)
			}
			game.Plays = append(game.Plays, play)
		}
		games = append(games, game)
	}

	//parte 1: verifico que ningun color supere el limite.
	suma_pt1 := 0

	for _, g := range games {
		es_posible := true
		for _, p := range g.Plays {
			for _, b := range p.Balls {
				if b.Cantidad > Colores[b.Color] {
					es_posible = false
					break
				}
			}
		}
		if es_posible {
			suma_pt1 += g.Numero
		}
	}

	fmt.Printf("Parte 1: suma de IDs de juegos validos= %d \n", suma_pt1)

	//Parte2:

	suma_pt2 := 0

	for _, g := range games {

		maxs := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, p := range g.Plays {
			for _, b := range p.Balls {
				if b.Cantidad > maxs[b.Color] {
					maxs[b.Color] = b.Cantidad
				}
			}
		}

		calc := maxs["red"] * maxs["green"] * maxs["blue"]
		suma_pt2 += calc

	}

	fmt.Printf("Parte 2: suma de la multiplicacion de max colores= %d \n", suma_pt2)
}
