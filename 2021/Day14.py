#!/bin/python

#--- Day 14: Extended Polymerization ---

from os import read

ARCHIVO       = "day14.txt"

# ARCHIVO       = "day14-test.txt"

MAX_ITERACIONES = 40


compuesto = []
template = dict()
polimero = dict()
letras =  dict() 


f = open(ARCHIVO,"r")
lines_read = f.readlines()
for i in range(len(lines_read)):
    if lines_read[i] == "\n":   #linea vacia
        continue

    if i == 0:
        compuesto = list(lines_read[i].replace("\n",""))
    else:
        t =  lines_read[i].replace("\n","").replace(" ","").split("->") 
        template.update({t[0]:t[1]})
        

#Dadas el compuesto los separo en pares y los agrego a un dict como PAR:Cant {"AB":1}
for i in range(len(compuesto)):
    if i+1 == len(compuesto):
        break
    polimero.update( {"".join(compuesto[i:i+2]):1} )
    letras[compuesto[i]] = letras.get(compuesto[i],0) + 1


letras[compuesto[-1]] = letras.get(compuesto[-1],0) + 1


# print(polimero)        

for j in range(MAX_ITERACIONES):
    # print("iter:" + str(j+1))
    temp = []
    newpolimero = dict()
    for k in polimero.keys():
        #armo 2 nuevas cadenas:
        n1 = k[0] + template[k]
        n2 = template[k] + k[1]
         
        letras[template[k]] = letras.get(template[k],0) + polimero.get(k) 

        newpolimero[n1] = polimero.get(k) + newpolimero.get(n1,0)
        newpolimero[n2] = polimero.get(k) + newpolimero.get(n2,0)


    #el viejo polimero se "destruye" al hacer los splits:
    polimero = dict( newpolimero )
    
    # print(polimero)
    


most = max(letras, key=letras.get)
least = min(letras, key=letras.get)
max_ = letras[most]
min_ = letras[least]

print("Most common element = " + most + " " + str(max_) + "   Least common element=" + least +  " " + str(min_) )
print("Result: " + str( max_ - min_))

