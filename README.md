![](https://github.com/floppyzedolfin/square/workflows/build/badge.svg)  ![](https://github.com/floppyzedolfin/square/workflows/coverage/badge.svg)

# square

Playing with Fiber v2

This project was aimed at playing with Fiber. It builds a microservice, inside a Docker image, that exposes an endpoint
that squares an integer.

## Clone the repository

```bash
git clone https://github.com/floppyzedolfin/square.git 
```

## Make targets

There are two options to build the image. Both options will generate the docker image, and should generate the same
docker image. Read below for potential differences.

### Compilation

- `make compile` compiles the service's binary locally. This includes GOPATH values, overwritten commands - such as `go`
  , etc. It's useful for a dev point of view, as it provides an on-the-fly mechanism to build the current state of dev

### Docker image build

- `make localbuild` generates the image based on the deployed vendors, and on everything your environment has.
- `make stagedbuild` is a more generic build, as it happens inside a pristine environment. The overhead of this is that
  we'll need to download the dependencies each time. This is the command to be used in CI for deliveries.

### Tests and coverage

- `make test` runs all the Unit Tests.
- `make checkcoverage` runs all the unit tests and ensures the coverage is at least 80%.

### Execution

- `make run` launches the service, exposing its port 3000 on the localhost's port 8530. The docker container will be
  removed when stopped.

### Cleanup

- `make clean` removes local build and test caches, and removes the docker image
- `make stop` ends the docker process (if you've forgotten wherein terminal tab it's running)

## Playing with it

Once you've ran `make run`, the service is up and running and you can shoot requests :

```bash
> curl -X POST -H "content-type:application/json" localhost:8530/square -d '{"value":4}'
{"value":16}
>
```

An error scenario has been built-in, for testing an example purposes, if the input value is `0`.

## Limitations

- So far, I haven't found a json parser that will cause invalid fields to raise an error. In our example here, the following request is valid:
```json
{
  "value": -4,
  "foo": "bar"
}
```



## TODOs

- Refactor the `internal/*` files, I'm not sure I like the way they are
- Currently, the `square/build/Dockerfile` is an excerpt of the `square/Dockerfile`. I'd rather it wouldn't.

