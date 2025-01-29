# Go by Example

[Go](https://go.dev/) is an open-source programming language designed for building scalable, secure, and reliable software. Please read the [official documentation](https://go.dev/doc/tutorial/getting-started) to learn more.

*Go by Example* is a hands-on introduction to Go using annotated example programs. Check out the [first example](hello-world.md) or browse the full list below.

Unless stated otherwise, the examples here assume the [latest major release Go](https://go.dev/doc/devel/release) and may use new language features. If something isn't working, try upgrading to the latest version.

{{< docstoc >}}

### FAQ

#### I found a problem with the examples; what do I do?

We're very happy to fix problem reports and accept contributions! Please submit
[an issue](https://github.com/jcbhmt/gobyexample.jcbhmr.com/issues) or send a Pull Request.
See `CONTRIBUTING.md` for more details.

#### What version of Go is required to run these examples?

Given Go's strong [backwards compatibility guarantees](https://go.dev/doc/go1compat),
we expect the vast majority of examples to work on the latest released version of Go
as well as many older releases going back years.

That said, some examples show off new features added in recent releases; therefore,
it's recommended to try running examples with the latest officially released Go version
(see Go's [release history](https://go.dev/doc/devel/release) for details).

#### I'm getting output in a different order from the example. Is the example wrong?

Some of the examples demonstrate concurrent code which has a non-deterministic
execution order. It depends on how the Go runtime schedules its goroutines and
may vary by operating system, CPU architecture, or even Go version.

Similarly, examples that iterate over maps may produce items in a different order
from what you're getting on your machine. This is because the order of iteration
over maps in Go is [not specified and is not guaranteed to be the same from one
iteration to the next](https://go.dev/ref/spec#RangeClause).

It doesn't mean anything is wrong with the example. Typically the code in these
examples will be insensitive to the actual order of the output; if the code is
sensitive to the order - that's probably a bug - so feel free to report it.
