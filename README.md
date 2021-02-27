# square
Playing with Fiber v2

This project was aimed at playing with Fiber. It builds a microservice, inside a Docker image, that exposes an endpoint that squares an integer.

## Clone the repository
```bash
git clone https://github.com/floppyzedolfin/square.git 
```

## Make targets 
There are two options to build the image. Both options will generate the docker image, and should generate the same docker image. Read below for potential differences.

### make build
`make build` generates the image based on the deployed vendors, and on everything your environment has. This includes GOPATH values, overwritten commands - such as `go`, etc. It's useful for a dev point of view, as it provides an on-the-fly mechanism to build the current state of dev

`make stagedbuild` is a more generic build, as it happens inside a pristine environment. The overhead of this is that we'll need to download the dependencies each time. This command is used in CI for deliveries.

### make run
`make run` launches the service, exposing its port 3000 on the localhost's port 8530. The docker container will be removed when stopped.

### make stop
`make stop` ends the docker process (if you've forgotten wherein terminal tab it's running) 

## TODOs
- Add some UTs, over `square_impl.go` and `log.go`
- Refactor the `internal/*` files, I'm not sure I like the way they are
- Currently, the `square/build/Dockerfile` is an excerpt of the `square/Dockerfile`. I'd rather it wouldn't.
- Add a real CI
