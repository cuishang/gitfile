#!/bin/bash

sum = 0
for x in list(range(100)):
    sum = sum + x
print(sum)

sum1 = 0
n = 99
while n > 0:
    sum1 = sum1 + n
    n = n - 2
print(sum1)

n1 = 1
while n1 <= 100:
    if n1 > 11:
    	break
    print(n1)
    n1 = n1 + 1
print('END')
