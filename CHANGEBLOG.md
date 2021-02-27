# CHANGEBLOG

This is a personal project, with no intent whatsoever on making any money out of it. The aim is to play with [Fiber](https://gofiber.io/), as you've probably already guessed after checking the commit messages.

Why Fiber, and not something else? Well, I can't / won't write a microservice with several of them, so I had to pick. Fiber I hadn't worked with hitherto. The first major bump I had was when I faced was discovering Fiber v2 existed, after having written all my code. Well, that might be ~100 lines, but still. Updating the dependencies isn't always fun.

Then I added the logger, and tried to play with `runtime.Callers`. Fun thing, depending on the `build` command, the logs won't be the same - local path for a `make build`, and something in `/go/src/github.com/floppyzedolfin/square/..` for the `stagedbuild` command.

Oh, one fun thing I learnt as I was building my staged Dockerfile: build go without specifying `CGO_ENABLED=0` will somehow leave C dependencies in the produced binary, when using the `net` package - which, obviously, is what I was doing.
[This answer](https://stackoverflow.com/a/36308464/2106703) saved me countless hours of browsing the internet wondering why running my microservice on centos would work, but it wouldn't on alpine or scratch.  

For the sake of coverage, I've added an error case: when requesting the square of `0`, an error is returned.
