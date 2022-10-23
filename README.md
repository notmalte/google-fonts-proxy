# Google Fonts Proxy

Simple drop-in replacement for Google Fonts, e.g. for GDPR reasons - just run the proxy and change your Google Fonts links to it.

## Try it out:

[https://googlefontsproxy.fly.dev/css2?family=Roboto&display=swap](https://googlefontsproxy.fly.dev/css2?family=Roboto&display=swap)

## How-to:

**Option A: Run Docker image**

```sh
$ docker run \
    -p "8080:8080" \
    -e "EXTERNAL_URL=[url to reach proxy]" \
    notmalte/google-fonts-proxy
```

Now the proxy is listening on `:8080`.

**Option B: Build and run Go binary locally**

```sh
$ git clone https://github.com/notmalte/google-fonts-proxy.git && cd google-fonts-proxy
$ go build
$ EXTERNAL_URL="[url to reach proxy]" ./google-fonts-proxy
```

Now the proxy is listening on `:8080`.

**Option C: Build and run Docker image locally**

```sh
$ git clone https://github.com/notmalte/google-fonts-proxy.git && cd google-fonts-proxy
$ docker build -t google-fonts-proxy .
$ docker run \
    -p "8080:8080" \
    -e "EXTERNAL_URL=[url to reach proxy]" \
    google-fonts-proxy
```

Now the proxy is listening on `:8080`.

**Adjust your `<link>` or `@import` url:**

```
https://fonts.googleapis.com/css2?family=Roboto&display=swap

->

https://[your url]/css2?family=Roboto&display=swap
```

## Configuration

**`EXTERNAL_URL`** (default: `http://localhost:8080`)

Set this to the public URL of the proxy, including protocol (and port if necessary). This is necessary as all `https://fonts.gstatic.com` references need to be replaced in the CSS files and absolute URLs are required.
The default value is only for local testing.

**`PORT`** (default: `8080`)

The proxy listens on this port. If you are using Docker, you should adjust the port mapping instead (`... -p "1234:8080" ...`).
