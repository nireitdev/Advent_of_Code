#!/bin/python3

# AOC 2022 - --- Day 7: No Space Left On Device ---

#Lo resolvi utilizando diccionarios. 
#También se podría resolver con OOP y haciendo una funcion recursiva para el parseo.

ARCHIVO = "dia07.txt"

if __name__ == "__main__":

    f = open(ARCHIVO, "r")
    lines = f.readlines()

    #defaults:
    ii = 0
    current_dir = ''
    tamanio = 0
    path = { '//': {'tamanio': 0} }

    while ii<len(lines):

        l = lines[ii].strip()

        #comandos:
        if l == '$ ls' :
            pass
        elif l == '$ cd ..' :
            #Sumo al directorio "padre" el actual tamaño del directorio "hijo"
            #y finalmente cambio de directorio al "padre"
            last_tamanio = path[current_dir]['tamanio']
            dir = current_dir.split('/')
            del dir[-1]            
            current_dir = '/'.join(dir)
            path[current_dir]['tamanio'] = path[current_dir]['tamanio'] + last_tamanio


        elif l[0:4] == '$ cd':
            #cambio de path:
            current_dir = current_dir +'/'+ l[5:]

        elif l[0:3] == 'dir':
            #agrego nuevo path en el diccionario:
            path[current_dir +'/'+ l[4:]] = {'tamanio': 0}

        else:
            #parseo de archivos y tamaños:
            file = l.split(' ');
            path[current_dir][file[1]] = int(file[0])
            path[current_dir]['tamanio'] = path[current_dir]['tamanio'] + int(file[0])

        ii = ii + 1


    
    #El ultimo "cd" no vuelve al directorio raiz '//' 
    #y por lo tanto me queda sin sumar su tamaño al total:    
    while current_dir != '//' :
        last_tamanio = path[current_dir]['tamanio']
        dir = current_dir.split('/')
        del dir[-1]
        current_dir = '/'.join(dir)
        path[current_dir]['tamanio'] = path[current_dir]['tamanio'] + last_tamanio



    #print(path)

    #Parte 1
    #Busco los directorios con tamanio > 10000
    # Y SE SUMAN TODOS, aun los que estan incluidos dentro de otros!!
    suma = 0
    for keys in path.keys():
        #if len( keys.split('/') ) == 4:
        if path[keys]['tamanio'] < 100000:
                #print(keys, path[keys]['tamanio'])
                suma = suma + path[keys]['tamanio']

    #print(path)
    print("Parte 1 suma = ", suma)



    #Parte 2
    print('Parte 2 , ocupado=',  path['//']['tamanio'])
    tamano_borrar =   path['//']['tamanio'] - 40_000_000

    #defaults:
    minimo = 1_000_000_000
    path_borrar = {'path':'//', 'tamanio':  path['//']['tamanio']}

    for keys in path.keys():
        diff = path[keys]['tamanio'] - tamano_borrar
        if  diff >= 0 and diff < minimo:
                path_borrar['path'] = keys
                path_borrar['tamanio'] = path[keys]['tamanio']
                minimo = diff

    print("Se necesita borrar : " , path_borrar['path'])
    print("Tamaño del directorio: " , path_borrar['tamanio'])

