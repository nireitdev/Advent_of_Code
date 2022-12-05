#!/bin/python3

# AOC 2022 - --- Day 5: Supply Stacks ---

#TODO: refact!!

ARCHIVO = "dia05.txt"

#Func para operaciones:
def stackBegin(stack :list, elem : str):
    stack.insert(0,elem)

def stackPush(stack :list, elem : str):
    stack.append(elem)

def stackPop(stack :list):
    if len(stack)>0:
        return stack.pop(-1)

def ejecuteMovs(from_stk :list, to_stk:list, cantidad:int):
    for ii in range(cantidad):
        stackPush(to_stk, stackPop(from_stk))

def ejecuteMovsMultipleAtOnce(from_stk :list, to_stk:list, cantidad:int):
    popped = []
    for ii in range(cantidad):
        if len(from_stk) > 0:
            popped.append(stackPop(from_stk))

    for ii in range(len(popped)-1,-1,-1) :
        stackPush(to_stk, popped[ii])


if __name__ == "__main__":
    f = open(ARCHIVO, "r")
    lines = f.readlines()

    #Colas
    count_colas = int( ( len( lines[0] )  )/4 )
    stacks = []
    movimientos = []
    
    for stk in range(count_colas):
        stacks.append([])
    
    parsing_stacks = True
    for l in lines:
        #no hago l.strip() porque me elimina los ' ' (espacios)
        if l.strip() == '' or l[1] == '1':
            parsing_stacks = False
            continue

        if parsing_stacks:
            stk = 0
            while stk*4 +1 < len( l ) -1 :
                letra = l[ stk*4 + 1 ]
                if letra != ' ':
                    stackBegin(stacks[stk], letra)
                stk = stk + 1
        else:
            #parsing movimientos:
            mov =  l.strip().split(' ')            
            movimientos.append( mov)

    #clone stacks
    stacks_parte_1 = [ x[:] for x in stacks]
    stacks_parte_2 = [ x[:] for x in stacks]

    for ii in range(len(movimientos)):
        #movs[1] = cantidad    movs[3] = FROM   movs[5] = TO
        ejecuteMovs( stacks_parte_1[ int( movimientos[ii][3] ) - 1 ],
                     stacks_parte_1[ int( movimientos[ii][5] ) - 1],
                     int( movimientos[ii][1] )
                    )
    
    parte_1 = ''
    for ii in range( len(stacks_parte_1) ):
        parte_1 = parte_1 +  stacks_parte_1[ii][-1]

    print("Parte 1 : Stack reordenado = ",parte_1)


    for ii in range(len(movimientos)):
        #movs[1] = cantidad    movs[3] = FROM   movs[5] = TO
        ejecuteMovsMultipleAtOnce( stacks_parte_2[ int( movimientos[ii][3] ) - 1 ],
                     stacks_parte_2[ int( movimientos[ii][5] ) - 1],
                     int( movimientos[ii][1] )
                    )
    parte_2 = ''
    for ii in range( len(stacks_parte_2) ):
        parte_2 = parte_2 +  stacks_parte_2[ii][-1]

    print("Parte 2 : Stack reordenado = ",parte_2)


    
