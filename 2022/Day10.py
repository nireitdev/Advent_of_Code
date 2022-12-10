#!/bin/python3

# AOC 2022 - --- Day 10: Cathode-Ray Tube ---


ARCHIVO = "dia10.txt"

def check_ciclo(ciclos: int,valor_X:int) -> int:
    div = (ciclos - 20) % 40  #multiplo de 40
    if div == 0:
        return ciclos * valor_X
    return 0

def dib_crt(crt:list, ciclo:int, pos_sprite:list):
    div_ciclo = (ciclo-1) % 40
    if div_ciclo in pos_sprite:
        crt[ciclo - 1] = '#'
    else:
        crt[ciclo - 1] = '.'

if __name__ == "__main__":
    f = open(ARCHIVO, "r")
    lines = f.readlines()

    ciclo = 1
    valor_X = 1
    sum = 0

    crt = [''] * (40 * 6 + 1 )  # +1 ultimo ciclo
    crt[0] = '#'      # +1 ciclo

    pos_sprite = [0,1,2] # = "###"

    for l in lines:

        instruc = l.strip().split(" ")

        if instruc[0] == 'noop':
            ciclo += 1
            
        elif instruc[0] == 'addx':
            ciclo += 1
            sum += check_ciclo(ciclo,valor_X)
            dib_crt(crt,ciclo,pos_sprite)

            ciclo += 1
            valor_X += int(instruc[1])
            pos_sprite = [valor_X-1, valor_X, valor_X+1]

        #+1 ciclo:
        sum += check_ciclo(ciclo, valor_X)
        dib_crt(crt, ciclo, pos_sprite)


    #Parte 1
    print("Parte 1: Total señal=" , sum)

    print("")


    #Parte 2
    #Mostrar letras:
    print("Parte 2: Letras del CTR:")
    for ii in range(6):
        print(''.join(crt[ ii*40 : (ii+1)*40]))  #crt[0:40], crt[40,80], ...

