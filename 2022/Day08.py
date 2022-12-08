#!/bin/python3

# AOC 2022 - --- Day 8: Treetop Tree House ---


ARCHIVO = "dia08.txt"

def vision(grid:list , pos_x, pos_y):
    altura = grid[pos_y][pos_x]
    horizontal = grid[pos_y]

    #visible desde Rigth:
    if max( horizontal[:pos_x] ) < altura:
        return True
    # visible desde Left:
    if max( horizontal[pos_x+1:] ) < altura:
        return True

    vertical = []
    for ii in range(len(grid[0])):
        vertical.append(grid[ii][pos_x])

    # visible desde Top:
    if max(  vertical[:pos_y] ) < altura:
        return True
    # visible desde Bottom:
    if max(  vertical[pos_y+1:] ) < altura:
        return True

    return False
    #end vision()


def escenario(grid:list , pos_x, pos_y):

    def score(valor:int, lista:list):
        for ii in range(len(lista)):
            if lista[ii] >= valor:
                return ii+1
        return len(lista)

    altura = grid[pos_y][pos_x]
    horizontal = grid[pos_y]
    vertical = []
    for ii in range(len(grid[0])):
        vertical.append(grid[ii][pos_x])

    #visible desde Rigth:
    valor = 1
    izq = horizontal[:pos_x][::-1]
    der = horizontal[pos_x+1:]
    arriba = vertical[:pos_y][::-1]
    abajo = vertical[pos_y+1:]

    valor *= score(altura, der)
    valor *= score(altura, izq)
    valor *= score(altura, arriba)
    valor *= score(altura, abajo)

    return valor
    #end escenario()


if __name__ == "__main__":

    f = open(ARCHIVO, "r")
    lines = f.readlines()

    grid = [[] for x in range(len(lines[0].strip()))]

    for ii in range(len(lines)):
        l = lines[ii].strip()
        grid[ii] = list(l)


    #Borde:
    visible = 4 * (len(grid[0])-1)
    max_score = 0

    for ii in range(1,len(grid[0])-1):
        for jj in range(1,len(grid[0])-1):
            if vision(grid,ii,jj):
                visible += 1
            score = escenario(grid,ii,jj)
            if score > max_score:
                max_score = score

    print("Parte 1: " , visible)
    print("Parte 2: " , max_score)






