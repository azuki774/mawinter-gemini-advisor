#!/usr/bin/env python

import json
from http.server import BaseHTTPRequestHandler, HTTPServer


class MockHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()
        response = [
            {
                "category_id": 210,
                "category_name": "食費",
                "datetime": "2025-05-15T13:56:41+09:00",
                "from": "mawinter-fe",
                "id": 6995,
                "memo": "手動",
                "price": 1150,
                "type": ""
            },
            {
                "category_id": 220,
                "category_name": "電気代",
                "datetime": "2025-05-16T13:56:41+09:00",
                "from": "mawinter-fe",
                "id": 6996,
                "memo": "なんとか電気",
                "price": 7000,
                "type": ""
            },
            {
                "category_id": 220,
                "category_name": "通信費",
                "datetime": "2025-05-17T13:56:41+09:00",
                "from": "mawinter-fe",
                "id": 6997,
                "memo": "データ通信",
                "price": 450,
                "type": ""
            },
        ]
        responseBody = json.dumps(response, ensure_ascii=False)

        self.wfile.write(responseBody.encode('utf-8'))


def run(server_class=HTTPServer, handler_class=MockHandler, server_name='localhost', port=8080):

    server = server_class((server_name, port), handler_class)
    server.serve_forever()


def main():
    run()


if __name__ == '__main__':
    main()
