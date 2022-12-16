#!/bin/python3

# AOC 2022 - --- Day 12: Hill Climbing Algorithm ---

ARCHIVO = "dia12.txt"

INFINITO = 99999


def visitar(arbol:dict(), nodo:set(), distancia = 0):
    for n in arbol[nodo]['hijos']:
        #solo actualiza si la distancia actual es menor que una anterior
        #Alg Dijstra
        if distancia + 1 < arbol[n]['distancia']:
            arbol[n]['distancia'] = distancia +1
            visitar(arbol, n, distancia + 1 )


import sys
sys.setrecursionlimit(10000)   #limites de la recursion

if __name__ == "__main__":
    f = open(ARCHIVO, "r")
    lines = f.readlines()


    tablero = list()
    arbol = dict()
    nodo_S = dict()
    nodo_E = dict()

    for ll in lines:
        l = list( ll.strip() )
        tablero.append(l)


    for yy in range(len(tablero)):
        for xx in range(len(tablero[yy])):
            letra = tablero[yy][xx]
            coord = (yy,xx)
            if letra == "S":
                nodo_S = coord
                letra = "a"

            if letra == 'E':
                nodo_E = coord
                letra ="z"

            arbol[coord] = {'valor':ord(letra), 'hijos': list(), 'visitado': False, 'distancia':INFINITO }

    #Completo los hijos:
    for key in arbol.keys():
        yy,xx = key
        #Valores permitidos: desde la 'a' hasta (valor + 1 )
        # range('a', (valor +1) +1 )   range no incluye el ultimo por eso +2
        valores_permitidos= list(range(ord('a'),arbol[key]['valor']+ 2))

        #@refactor refactorizar en algun momento!!!

        #Miro en cada direccion para detectar los nodos hijos
        if (yy+1,xx) in arbol and arbol[(yy+1,xx)]['valor'] in valores_permitidos:
            arbol[key]['hijos'].append((yy+1,xx))

        if (yy-1,xx) in arbol and arbol[(yy-1,xx)]['valor'] in valores_permitidos:
            arbol[key]['hijos'].append((yy-1,xx))

        if (yy,xx-1) in arbol and arbol[(yy,xx-1)]['valor'] in  valores_permitidos:
            arbol[key]['hijos'].append((yy, xx - 1))

        if (yy,xx+1) in arbol and arbol[(yy,xx+1)]['valor'] in valores_permitidos:
            arbol[key]['hijos'].append((yy,xx+1))

    visitar(arbol,nodo_S, 0)

    print("Parte 1 : Minimo camino desde 'S' -> 'E' = ", arbol[nodo_E]['distancia'])  #447


    #Parte 2 - empezar por cualquier 'a' del borde  y minimizar hasta llegar a "E"


    lista_distancias = list()
    for nodo in arbol.keys():
        #reseteo distancias
        for key in arbol.keys():
            arbol[key]['distancia'] = INFINITO

        y,x = nodo
        if x in [ 0 , len(tablero[yy])- 1 ] or  y in [ 0, len(tablero) -1]:
            if arbol[nodo]['valor'] == ord('a'):
                visitar(arbol, nodo, 0)

                lista_distancias.append(arbol[nodo_E]['distancia'])

                print("distancia desde el nodo ", nodo, " = ", arbol[nodo_E]['distancia'])

    print("Parte 2: Minimo camino desde 'a' -> 'E' = ", min(lista_distancias) )  #446