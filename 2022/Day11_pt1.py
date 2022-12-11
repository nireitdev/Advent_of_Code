#!/bin/python3

# AOC 2022 - --- Day 11: Monkey in the Middle ---

#Resolucion parte 1, Parte 2 es una solucion muyyyy lenta. 
#Todo: buscar alguna alternativa a la parte 2

ARCHIVO = "dia11.txt"



class Mono:

    items = []
    count_inspeccion = 0
    habilitarDivBy3 = True  #parte 1

    def setItems(self, l:list):
        self.items = l

    def setOperacion(self,oper:str):
        self.operacion = oper

    def setTestDivision(self, div:int):
        self.div = div

    def setTestVerdadero(self,mono: 'Mono'):
        self.testMonoVerd = mono
    def setTestFalso(self,mono: 'Mono'):
        self.testMonoFalso = mono

    def setDivBy3(self,estado = True):
        self.habilitarDivBy3 = estado

    def tirarItem(self):
        '''Inspecciona el valor del item y le aplica la operacion'''
        self.count_inspeccion += 1
        old = self.items.pop(0)  #item = old "worry"
        worry = eval(self.operacion)
        if self.habilitarDivBy3:
            worry = worry // 3

        if worry % self.div == 0:
            self.testMonoVerd.agregarItem(worry)
        else:
            self.testMonoFalso.agregarItem(worry)


    def agregarItem(self, item:int):
        self.items.append(item)

    def getCantInspeccion(self):
        return self.count_inspeccion
    def getItems(self):
        return self.items




if __name__ == "__main__":
    f = open(ARCHIVO, "r")
    lines = f.readlines()

    CANT_MONOS = 8
    monos = [ Mono() for x in range(CANT_MONOS)]

    # Parseo la entrada:
    num_monos = -1
    for l in lines:
        ss = l.strip().split(":")
        if "Monkey" in ss[0]:
            num_monos += 1
            continue

        if "Starting" in ss[0]:
            monos[num_monos].setItems( [int(x) for x in  ss[1].split(",") ] )
            continue

        if "Operation" in ss[0]:
            monos[num_monos].setOperacion( ss[1].split("=")[1] )
            continue

        if "Test" in l:
            monos[num_monos].setTestDivision( int( ss[1].split("by")[1] ) )
            continue

        if "true:" in l:
            monos[num_monos].setTestVerdadero( monos[ int(ss[1].split("monkey")[1]) ] )
            continue

        if "false:" in l:
            monos[num_monos].setTestFalso( monos[ int(ss[1].split("monkey")[1]) ] )
            continue


    #Primer parte

    #Itero 20 veces:
    for xx in range(20):
        #Por cada item del mono:
        for ii in range(CANT_MONOS):
            for jj in range(len(monos[ii].getItems())):
                monos[ii].tirarItem()


        #Print por cada ronda:
        # for ii in range(CANT_MONOS):
            # print(monos[ii].items)


    count = []
    for ii in range(CANT_MONOS):
            count.append(monos[ii].getCantInspeccion())
    sorted_count = sorted(count)
    print("Parte 1 : total inspecciones de dos monos mas activos", sorted_count[-2]*sorted_count[-1])


