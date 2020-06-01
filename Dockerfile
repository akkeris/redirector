FROM golang:1.7
ADD . .
RUN go build -o redirector main.go
CMD ["./redirector"]

EXPOSE 9000

