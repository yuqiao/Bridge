#!/usr/bin/env python
#coding=utf-8
#
# Brief:
# Author: rongqiao.yurq (rongqiao.yurq@taobao.com)
#
# Updates:
# 0.1.0 --------------------
#   Create it.
#

from ctypes import *

helloworld = CDLL('helloworld.so')

def hello( name ):
    p = c_char_p(name)
    helloworld.hello(p)

def main():
    hello('ctypes')

if __name__ == "__main__":
    main()

