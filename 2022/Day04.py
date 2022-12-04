#!/bin/python3

# AOC 2022 - --- Day 4: Camp Cleanup ---


ARCHIVO = "dia04.txt"

def contenido_total(secc_1, secc_2):
    X = secc_1.split('-')
    Y = secc_2.split('-')
    if int(X[0])<=int(Y[0]) and int(X[1])>=int(Y[1]):
        return 1
    if int(Y[0])<=int(X[0]) and int(Y[1])>=int(X[1]):
        return 1
    return 0

def contenido_parcial(secc_1, secc_2):

    if contenido_total(secc_1,secc_2) == 1:
        return 1

    X = [ int(ii) for ii in secc_1.split('-') ]
    Y = [ int(ii) for ii in secc_2.split('-') ]

    for ii in range(X[0],X[1]+1,1):
        if ii in range(Y[0],Y[1]+1,1):
            return 1

    return 0


if __name__ == "__main__":

    f = open(ARCHIVO, "r")
    lines = f.readlines()

    #Primera Parte
    count = 0
    for ll in lines:
        l = ll.strip().split(',')
        count = count + contenido_total(l[0],l[1])

    print("Parte 1 : Secciones Totalmente re-asignadas ",count)

    #Segunda Parte
    count = 0
    for ll in lines:
        l = ll.strip().split(',')
        count = count + contenido_parcial(l[0],l[1])

    print("Parte 2 : Secciones Parcialemnte re-asignadas ",count)

