FROM golang:1.16 as builder

WORKDIR /go/src/github.com/floppyzedolfin/square
COPY . /go/src/github.com/floppyzedolfin/square
RUN make compile
RUN chmod +x /go/src/github.com/floppyzedolfin/square/build/square.out

FROM alpine
COPY --from=builder /go/src/github.com/floppyzedolfin/square/build/square.out /app/square.out

ENTRYPOINT ["/app/square.out"]
