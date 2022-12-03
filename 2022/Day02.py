#!/bin/python3

# AOC 2022 - --- Day 2: Rock Paper Scissors ---

import sys

ARCHIVO = "dia02.txt"

def calc_part1(u, v):
    reslt = 0
    if u == 'A' and v == 'X':
        reslt = 1 + 3
    elif u == 'A' and v == 'Y':
        reslt = 2 + 6
    elif u == 'A' and v == 'Z':
        reslt = 3 + 0
    elif u == 'B' and v == 'X':
        reslt = 1 + 0
    elif u == 'B' and v == 'Y':
        reslt = 2 + 3
    elif u == 'B' and v == 'Z':
        reslt = 3 + 6
    elif u == 'C' and v == 'X':
        reslt = 1 + 6
    elif u == 'C' and v == 'Y':
        reslt = 2 + 0
    elif u == 'C' and v == 'Z':
        reslt = 3 + 3
    return reslt

def calc_part2(u, v):
    reslt = 0
    if u == 'A' and v == 'X':
        reslt = 3 + 0
    elif u == 'A' and v == 'Y':
        reslt = 1 + 3
    elif u == 'A' and v == 'Z':
        reslt = 2 + 6

    elif u == 'B' and v == 'X':
        reslt = 1 + 0
    elif u == 'B' and v == 'Y':
        reslt = 2 + 3
    elif u == 'B' and v == 'Z':
        reslt = 3 + 6

    elif u == 'C' and v == 'X':
        reslt = 2 + 0
    elif u == 'C' and v == 'Y':
        reslt = 3 + 3
    elif u == 'C' and v == 'Z':
        reslt = 1 + 6
    return reslt


if __name__ == "__main__":

    f = open(ARCHIVO, "r")
    lines = f.readlines()

    total_pt1 = 0
    total_pt2 = 0
    for ll in lines:
        l = ll.strip().split(' ')
        total_pt1 = total_pt1 + calc_part1(l[0], l[1])
        total_pt2 = total_pt2 + calc_part2(l[0], l[1])

    print("Result part 1 =" , total_pt1)
    print("Result part 2 =" , total_pt2)







            
    
