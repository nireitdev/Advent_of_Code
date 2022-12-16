#!/bin/python3

# AOC 2022 - --- Day 13: Distress Signal ---

ARCHIVO = "dia13.txt"

from enum import Enum
class Orden(Enum):
    RIGHT = 1
    NOT_RIGHT = 2
    UNDEF = 3

#Funcion recursiva para comparar listas:
def comparar( list_izq , list_der )->bool:
    ii=0
    while(True):

        #Verifico las longitudes de la listas:
        if len(list_izq) == len(list_der) == ii:
            return Orden.UNDEF

        if len(list_izq) == ii:
            return Orden.RIGHT
        if len(list_der) == ii:
            return Orden.NOT_RIGHT

        #Recursivo si son listas:
        if type(list_izq[ii])==type(list_der[ii])==list:
            res = comparar(list_izq[ii],list_der[ii])
            if res != Orden.UNDEF:
                return res

        #Comparo si son numeros
        if type(list_izq[ii]) == type(list_der[ii]) == int:
            if list_izq[ii] < list_der[ii]:
                return Orden.RIGHT
            elif list_izq[ii] > list_der[ii]:
                return Orden.NOT_RIGHT
            else:
                pass

        #Convierto a lista si es un numero
        if type(list_izq[ii])== list and type(list_der[ii])==int:
            res = comparar(list_izq[ii], [ list_der[ii] ] )
            if res != Orden.UNDEF:
                return res
        #Convierto a lista si es un numero
        if type(list_izq[ii])== int and type(list_der[ii])==list:
            res = comparar([ list_izq[ii] ], list_der[ii])
            if res != Orden.UNDEF:
                return res

        ii+=1
    #while(true)
#def comparar()


if __name__ == "__main__":
    f = open(ARCHIVO, "r")
    lines = f.readlines()

    ii = 0
    izq = list()
    der = list()
    for ll in lines:
        l = ll.strip()
        if l == "":
            continue
        if ii % 2 == 0 :
            izq.append(eval(l))
        else:
            der.append(eval(l))
        ii += 1

    count_rigth = 0
    for ii in range(len(izq)):
        result = comparar(izq[ii], der[ii])
        if  result== Orden.RIGHT:
            count_rigth += ii +1

    print("Primer parte: Suma Indices en Rigth: ",count_rigth)   #5292



    ii = 0
    lista_ordenada = list()

    for ll in lines:
        l = ll.strip()
        if l == "":
            continue

        # "achato" o "aliso" las listas:
        # asigno valor -1 cuando []
        l = l.replace('[]','-1')
        l = '[' + l.replace('[','').replace(']','') + ']'

        lista_ordenada.append(eval(l))

    #agrego los "divisores":
    lista_ordenada.append([2])
    lista_ordenada.append([6])

    # ordeno la lista segun los valores
    ls = sorted( lista_ordenada)

    res = 1
    for ii in range(len(ls)):
        if ls[ii] in [ [2], [6] ]:
            res *= ii+1

    print("Segunda parte: Decoder Key : ",res)   #23868

