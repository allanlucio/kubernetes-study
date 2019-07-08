FROM golang AS builder

COPY . /go/
WORKDIR /go/src/webserver
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w" main.go
RUN chmod +x main.go



FROM busybox:1.30   
# WORKDIR /root/
COPY --from=builder /go/src/webserver/main .
ENTRYPOINT ["./main"] 

expose 8000