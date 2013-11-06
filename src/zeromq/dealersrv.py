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

import zmq
import sys
import time

def main():
    if len(sys.argv) == 2:
        index = sys.argv[1]
    else:
        index = '0'
    context = zmq.Context()

    server = context.socket(zmq.DEALER)
    server.setsockopt(zmq.IDENTITY,index)
    server.bind("tcp://*:555%s" % index)

    

     

if __name__ == "__main__":
    main()

