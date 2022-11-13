from io import BytesIO
import base64
from flask import Flask, request, jsonify
from sd import txt2img_pipe

app = Flask(__name__)
app.config["MAX_CONTENT_LENGTH"] = 16 * 1024 * 1024


@app.route("/api/v1/txt2img", methods=["POST"])
def txt2img():
    prompt = request.form.get("prompt")
    if prompt == None:
        return jsonify("Missing parameter 'prompt'"), 400

    image = txt2img_pipe(prompt).images[0]

    buff = BytesIO()
    image.save(buff, format="JPEG")
    data = base64.b64encode(buff.getvalue())

    return jsonify({
        "data": data.decode("utf-8"),
    }), 200


app.run(host="0.0.0.0")
