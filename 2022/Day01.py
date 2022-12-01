#!/bin/python3

# AOC 2022 - --- Day 1: Calorie Counting ---

import sys

ARCHIVO = "dia01.txt"

if __name__ == "__main__":
    f = open(ARCHIVO, "r")
    lines = f.readlines()
    calorias = []
    total = 0
    for ll in lines:
        if ll.strip() == '':
            calorias.append(total)
            total = 0
            continue
        total = total + int( ll.strip() )

    sorted_cals = sorted(calorias)

    #print(sorted_cals[-3:])
    print("Parte 1: max calorias)=", max(calorias))  #sorted_cals[-1]
    print("Parte 2: sum 3 max calorias)=", sum(sorted_cals[-3:]))







            
    
