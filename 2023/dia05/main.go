package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
--- Day 5: If You Give A Seed A Fertilizer ---
*/

//Mapeos:
const (
	SeedToSoil = iota
	SoilToFertilizer
	FertilizerToWater
	WaterToLight
	LightToTemperature
	TemperatureToHumidity
	HumidityToLocation
)

//Almacena un rango
type Range struct {
	Destination int
	Source      int
	Length      int
}

//almacena todos los mapeos
type Mapper struct {
	Ranges []Range
}

type SubRange struct {
	Inicio int
	Fin    int
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

	Mapeos := []Mapper{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}}
	Seeds := []int{}

	//Parseo archivo
	count := -1
	for i, ss := range input {

		if i == 0 {
			//Primer linea es de "seeds"
			x := strings.Split(ss, ":") //seeds: 79 14 55 13
			y := strings.Fields(x[1])   //79 14 55 13
			for _, v := range y {
				n, _ := strconv.Atoi(strings.Trim(v, " "))
				Seeds = append(Seeds, n)
			}
		} else {
			//Mapeos:
			if strings.Contains(ss, ":") {
				count++
				continue
			}
			if ss == "" {
				continue
			}
			r := Range{}
			n, _ := fmt.Sscanf(ss, "%d %d %d", &r.Destination, &r.Source, &r.Length)
			if n != 3 || err != nil {
				log.Fatal("No se pudo parsear strings a rango!", err)
			}
			Mapeos[count].Ranges = append(Mapeos[count].Ranges, r)
		}

	}

	//fmt.Printf("%#v \n", Mapeos[SeedToSoil].Ranges)
	//source := 82
	//fmt.Printf("Source: %d  Dest: %d \n", source, BusquedaDestino(Mapeos[HumidityToLocation].Ranges, source))

	//parte 1: itero por cada semilla
	min_loc := math.MaxInt
	for _, seed := range Seeds {
		soil := BusquedaDestino(Mapeos[SeedToSoil].Ranges, seed)
		fert := BusquedaDestino(Mapeos[SoilToFertilizer].Ranges, soil)
		watr := BusquedaDestino(Mapeos[FertilizerToWater].Ranges, fert)
		ligh := BusquedaDestino(Mapeos[WaterToLight].Ranges, watr)
		temp := BusquedaDestino(Mapeos[LightToTemperature].Ranges, ligh)
		humd := BusquedaDestino(Mapeos[TemperatureToHumidity].Ranges, temp)
		loct := BusquedaDestino(Mapeos[HumidityToLocation].Ranges, humd)
		if loct < min_loc {
			min_loc = loct
		}
	}
	//Tambien funciona:
	//for _, seed := range Seeds {
	//	valor := seed
	//	for _, m := range Mapeos {
	//		valor = BusquedaDestino(m.Ranges, valor)
	//	}
	//	if valor < min_loc {
	//		min_loc = valor
	//	}
	//}

	fmt.Printf("Parte 1: minimo location= %d \n\n", min_loc)


	
	//Parte2: es lo mismo que la parte 1, cambia la forma de procesar los seeds y con GOroutines!!:
	// 2 minutos en un Ryzen 7 bitch!

	i := 0
	resultado := make(chan int, 2)

	startPt2 := time.Now()

	wg := sync.WaitGroup{}

	for {
		sMin := Seeds[i]
		sMax := sMin + Seeds[i+1]

		wg.Add(1)
		go func(seedMin int, seedMax int) {
			defer wg.Done()
			min_loc := math.MaxInt
			start := time.Now()
			fmt.Printf("Inicio rango : %d, %d \n", seedMin, seedMax)
			for seed := seedMin; seed <= seedMax; s++ {
				valor := seed
				for _, m := range Mapeos {
					valor = BusquedaDestino(m.Ranges, valor)
				}
				if valor < min_loc {
					min_loc = valor
				}
			}
			resultado <- min_loc
			fmt.Printf("Fin rango : %d, %d => min: %d , tiempo %v\n", seedMin, seedMax, min_loc, time.Since(start))

		}(sMin, sMax)

		i += 2
		if i >= len(Seeds) {
			break
		}
	}

	go func() {
		min := math.MaxInt
		for v := range resultado {
			if v < min {
				min = v
			}
		}
		fmt.Printf("Parte 2: minimo location= %d \n", min)
	}()

	wg.Wait()
	close(resultado)
	time.Sleep(20 * time.Second)
	fmt.Printf("Total Ejecucion %v\n", time.Since(startPt2))
}

// Funcion que hace un lookup en la tabla de mapeos
// y devuelvo el valor destino
func BusquedaDestino(tabla []Range, source int) (destination int) {
	// Esta en la tabla:
	for _, r := range tabla {
		if source >= r.Source && source <= r.Source+r.Length {
			return source + r.Destination - r.Source
		}
	}
	//NO esta en la tabla:
	return source
}
