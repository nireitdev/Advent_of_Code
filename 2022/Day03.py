#!/bin/python3

# AOC 2022 - --- Day 3: Rucksack Reorganization ---

ARCHIVO = "dia03.txt"

def prio(xx):
    letras = ('abcdefghijklmnopqrstuvwxyz')
    if xx in letras:
        return ord(xx) - 96
    else:
        return ord(xx) - 38

if __name__ == "__main__":

    f = open(ARCHIVO, "r")
    lines = f.readlines()

    #Primera Parte
    sum_prios = 0
    for ll in lines:
        l = ll.strip()
        half = len(l) // 2
        for xx in l[:half]:
            if xx in l[half:]:
                sum_prios = sum_prios + prio(xx)
                break #no seguir buscando

    print("Primera parte sum prios=", sum_prios)


    #Segunda Parte
    sum_prios = 0
    for ii in range(0,len(lines),3):
        X = lines[ii].strip()
        Y = lines[ii+1].strip()
        Z = lines[ii+2].strip()

        for xx in X:
            if xx in Y:
                if xx in Z:
                    sum_prios = sum_prios + prio(xx)
                    break  # no seguir buscando

    print("Segunda parte sum prios=", sum_prios)


            
    
