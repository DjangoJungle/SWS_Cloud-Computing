import http.client
from http.server import BaseHTTPRequestHandler, HTTPServer
import json

class ProxyHandler(BaseHTTPRequestHandler):
    def do_POST(self):
        content_length = int(self.headers['Content-Length'])
        post_data = self.rfile.read(content_length)
        request_data = json.loads(post_data.decode())

        server_ip = request_data['server_ip']
        server_port = request_data['server_port']
        message = request_data['message']

        conn = http.client.HTTPConnection(server_ip, server_port)
        headers = {'Content-type': 'application/json'}
        conn.request('POST', '/', json.dumps({'message': message}), headers)
        response = conn.getresponse()
        data = response.read()
        conn.close()

        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()
        self.wfile.write(data)

def run(server_class=HTTPServer, handler_class=ProxyHandler, port=8888):
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)
    print(f'Starting proxy server on port {port}')
    httpd.serve_forever()

if __name__ == "__main__":
    run()
