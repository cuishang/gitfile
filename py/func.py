#!/bin/bash

def sum(*args):
    ax = 0
    for n in args:
	ax = ax + n
    return ax

s1 = sum(1,2)

def count(*args):
    def f(j):
	def g():
	    return j*j
	return g
    fs = []
    for i in range(1,4):
	fs.append(f(i))
    return fs
f1 = conut()
