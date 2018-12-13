#!/bin/bash

import os
print('process (%s) start...' % os.getpid())

pid = os.fork()
if pid ==0:
    print('i am child process (%s) and my parentid is %s' %(os.getpid,os.getppid))
else:
    print('i (%s) create a childid is %s' %(os.getpid,pid))
