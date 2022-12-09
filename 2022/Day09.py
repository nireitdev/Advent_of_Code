#!/bin/python3

# AOC 2022 - --- Day 9: Rope Bridge ---


ARCHIVO = "dia09.txt"

def moverHead(head:dict, paso:str, cantidad:int):
    if paso == 'U':
        head['y'] += cantidad
    if paso == 'D':
        head['y'] -= cantidad
    if paso == 'L':
        head['x'] -= cantidad
    if paso == 'R':
        head['x'] += cantidad

def moverTail(head:dict, tail:dict):
    dist_X = head['x']-tail['x']
    dist_Y = head['y']-tail['y']

    if abs(dist_X)<=1 and abs(dist_Y)<=1:
        return
    # @refact
    if dist_X >= 1:
       tail['x'] += 1
    if dist_X <= -1:
       tail['x'] -= 1
    if dist_Y >= 1:
        tail['y'] += 1
    if dist_Y <= -1:
       tail['y'] -= 1

if __name__ == "__main__":
    f = open(ARCHIVO, "r")
    lines = f.readlines()

    # Parte 1: 1 "head" y 1 "tail" :
    
    posH = { 'x':0,'y':0 }   #(x,y)
    posT = { 'x':0,'y':0 }   #(x,y)
    visitas_unicas = set()
    for l in lines:
        paso, cantidad= l.strip().split(" ")
        for ii in range(int(cantidad)):
            moverHead(posH, paso, 1 )
            moverTail(posH,posT)
            visitas_unicas.add((posT['x'],posT['y']))
            # print("head=",posH)
            # print("tail=",posT)

    print("Parte 1 = visitas Tail : ", len(visitas_unicas))

    # Parte 2: 1 "head" y 9 "tails"

    TOTAL_NUDOS = 10  #nudo[0] = head
    knots = [ { 'x':0,'y':0 } for x in range(TOTAL_NUDOS)]
    visitas_unicas = set()
    for l in lines:
        paso, cantidad= l.strip().split(" ")
        for ii in range(int(cantidad)):
            moverHead(knots[0], paso, 1)
            for jj in range(1,TOTAL_NUDOS ):
                moverTail(knots[jj-1], knots[jj])
            visitas_unicas.add((knots[TOTAL_NUDOS-1]['x'], knots[TOTAL_NUDOS-1]['y']))

    print("Parte 2 = visitas Knot nro 9 : ", len(visitas_unicas))



