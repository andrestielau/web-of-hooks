FROM golang:1.21.6

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o woh .

EXPOSE 3000
EXPOSE 3001

CMD ["/app/woh", "serve"]