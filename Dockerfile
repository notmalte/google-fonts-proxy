FROM golang:1.19

ENV EXTERNAL_URL="http://localhost:8080"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /google-fonts-proxy

EXPOSE 8080

CMD [ "/google-fonts-proxy" ]