# Section 1 - build the binary
FROM golang:1.16 as builder

WORKDIR /go/src/github.com/floppyzedolfin/square
COPY . /go/src/github.com/floppyzedolfin/square
# notice the .dockerignore prevented us from copying the vendors into the image
RUN make compile
RUN chmod +x /go/src/github.com/floppyzedolfin/square/build/square.out

# Section 2 - build the final image. All we need is the compiled binary
FROM scratch
COPY --from=builder /go/src/github.com/floppyzedolfin/square/build/square.out /app/square.out

ENTRYPOINT ["/app/square.out"]
