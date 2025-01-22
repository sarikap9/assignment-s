from flask import Flask

app = Flask(__name__)

@app.route('/ping', methods=['GET'])
def ping():
    return "pong from Python"

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=3001)
