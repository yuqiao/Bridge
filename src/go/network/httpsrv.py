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

import tornado.ioloop
import tornado.httpserver
import tornado.web

class IndexHandler(tornado.web.RequestHandler):
    def get(self):
        self.write("Hello, qiaoyun!");


def main():
    application = tornado.web.Application(
            handlers = [(r"/", IndexHandler)])
    http_server = tornado.httpserver.HTTPServer(application)
    http_server.listen(8081)
    tornado.ioloop.IOLoop.instance().start()

if __name__ == "__main__":
    main()

