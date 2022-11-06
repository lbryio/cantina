from flask import Flask, request
import binascii
app = Flask(__name__)

import explore

@app.route('/')
def hello_world():
    if 'u' in request.args:
        link = request.args['u']
    if 'l' in request.args:
        link = binascii.unhexlify(request.args['l'])

    return explore.get(link)

if __name__ == '__main__':
   app.run()
