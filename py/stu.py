#!/bin/bash
' test file '
__auth__ = 'cui'

class Student(object):
    pass
    
class Stu1(object):
    def __init__(self,name,score):
	self.name = name
	self.score = score
    def print_score(self):
	print('%s: %s ' % (self.name,self.score))


class Stu2(object):
    def __init__(self,name,score):
	self.__name = name
	self.__score = score
    def print_score(self):
	print('%s %s ' % (self.__name,self.__score))

    def get_name(self):
	return self.__name
    def get_score(self):
	return self.__score
    def set_name(self):
	self.__name = name
    def set_score(self,score):
	if 0 <= score <=100:
	    self.__score = score
	else:
	    raise ValueError('bad score')

bart = Stu1('ccui',59)



