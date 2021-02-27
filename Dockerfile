FROM golang:1.16 as builder

WORKDIR /go/src/github.com/floppyzedolfin/square
COPY . /go/src/github.com/floppyzedolfin/square
RUN go mod tidy
RUN go mod vendor
RUN make build

FROM scratch
COPY --from=builder /go/src/github.com/floppyzedolfin/square/build/square.out /app/square.out

CMD ["/app/square.out"]