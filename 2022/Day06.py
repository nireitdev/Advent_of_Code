#!/bin/python3

# AOC 2022 - --- Day 6: Tuning Trouble ---

ARCHIVO = "dia06.txt"

def buscarDistinctChars(cadena:str, cant_distinct_chars:int):
    for ii in range(len(cadena)):
        if ii + cant_distinct_chars > len(cadena):
            #Nunca se deberia ejecutar esto:
            print("Marker no encontrado")
            return 0

        mark = list( cadena[ii: ii + cant_distinct_chars] )
        is_marker = True
        for jj in range(cant_distinct_chars - 1):
            char = mark.pop(0)
            if char in mark:
                is_marker = False
                break
        if is_marker:
            return ii+DISTINCT_CHARS
    return 0

if __name__ == "__main__":

    f = open(ARCHIVO, "r")
    lines = f.readlines()


    DISTINCT_CHARS = 4
    for ll in lines:
        l = ll.strip()
        pos = buscarDistinctChars(l, DISTINCT_CHARS)
        if pos >0:
            print("Primer parte: se encontro despues de = ", pos )

    DISTINCT_CHARS = 14
    for ll in lines:
        l = ll.strip()
        pos = buscarDistinctChars(l, DISTINCT_CHARS)
        if pos >0:
            print("Segunda parte: se encontro despues de = ", pos )

