#!/usr/bin/env python

from http.server import BaseHTTPRequestHandler, HTTPServer


class MockHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.end_headers()
        response = """
        これはテスト用のGeminiレスポンスです
        ここに任意のテキストを記述できます
        """

        self.wfile.write(response.encode('utf-8'))


def run(server_class=HTTPServer, handler_class=MockHandler, server_name='localhost', port=8000):

    server = server_class((server_name, port), handler_class)
    server.serve_forever()


def main():
    run(server_name='0.0.0.0', port=8000)


if __name__ == '__main__':
    main()
