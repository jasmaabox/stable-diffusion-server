## Setup

## Setup worker

Acquire a machine with a GPU that can run Stable Diffusion in areasonable amount
of time.

Install dependencies for `worker`:

```
cd worker

conda create -n sd python=3.9
pip install -r requirements.txt
```

Set desired model in `MODEL_NAME` in `sd.py` and accept model license on
HuggingFace.

Authenticate with HuggingFace:

```
huggingface-cli login
```

Start server:

```
python server.py
```

Install Tailscale on the machine. Find the Tailscale IP of the machine running
your worker. This will be used to form `WORKER_URL`.

## Setup proxy

Acquire a Linux server.

Follow [Tailscale Linux installation
instructions](https://tailscale.com/download/linux).

In `/opt`, clone the repo and build the proxy server:

```
cd proxy
make
```

Create `.env` from `sample.env` and fill environment variables. `WORKER_URL`
will be `http://<SD MACHINE IP>:5000`.

Copy `sd-proxy.service` to `/etc/systemd/system`. Enable and start the service:

```
sudo systemctl daemon-reload
sudo systemctl enable sd-proxy.service
sudo systemctl start sd-proxy.service
```


## Test server

```
curl -X POST http://<PROXY ENDPOINT>:8080/api/v1/txt2img \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "prompt=a blue haired girl in a raincoat on a stormy evening, pixiv, artstation"
```