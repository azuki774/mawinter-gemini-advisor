#!/usr/bin/env python

from http.server import BaseHTTPRequestHandler, HTTPServer


class MockHandler(BaseHTTPRequestHandler):
    def do_POST(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json; charset=utf-8')
        self.end_headers()
        response = """
{
  "candidates": [
    {
      "content": {
        "parts": [
          {"text": "ここに生成されたテキストが入ります"}
        ],
        "role": "model"
      },
      "finishReason": "STOP",
      "index": 0,
      "safetyRatings": [
        {
          "category": "HARM_CATEGORY_SEXUALLY_EXPLICIT",
          "probability": "NEGLIGIBLE"
        }
      ]
    }
  ],
  "promptFeedback": {
    "safetyRatings": []
  }
}
        """
        self.wfile.write(response.encode('utf-8'))


def run(server_class=HTTPServer, handler_class=MockHandler, server_name='localhost', port=8000):

    server = server_class((server_name, port), handler_class)
    server.serve_forever()


def main():
    run(server_name='0.0.0.0', port=8000)


if __name__ == '__main__':
    main()
