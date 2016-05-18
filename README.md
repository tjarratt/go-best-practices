# Best Practices

A collection of things I consider to be best practices in test-driven go development. This applies mainly to writing *applications* in Go, in particular JSON APIs, but can be extrapolated to command line tools, libraries and framework too.

This was inspired by Can Berk Guder's [iOS best practices repo](https://github.com/cbguder/bestpractices).

## Practices and Principles

Some of the practices and principles I tried to demonstrate in this sample project are:

* [Test-driven development][tdd] using [Ginkgo][] and [Gomega][]
* [Single responsibility principle][srp]
* [Composition over inheritance][coi]
* [Dependency injection][di] using hand-rolled constructor functions
* [Declaring interface for dependencies near the consumer][interface]
* [Vendored dependencies][dependencies] with a [vendoring] tool [you enjoy using](#notes)
* [Responsible usage of concurrency][concurrency]
* [Minimal integration tests][integrated] with [Gomega's gexec package][gexec]
* [Avoid named returns][named-returns] in functions whenever possible

[Ginkgo]: https://github.com/onsi/ginkgo
[Gomega]: https://github.com/onsi/gomega
[coi]: http://en.wikipedia.org/wiki/Composition_over_inheritance
[di]: http://en.wikipedia.org/wiki/Dependency_injection
[srp]: http://en.wikipedia.org/wiki/Single_responsibility_principle
[tdd]: http://en.wikipedia.org/wiki/Test-driven_development
[interface]: https://github.com/tjarratt/go-best-practices/blob/master/usecases/order_pizza_use_case.go#L24
[concurrency]: https://divan.github.io/posts/go_concurrency_visualize/
[integrated]: http://blog.thecodewhisperer.com/permalink/integrated-tests-are-a-scam
[dependencies]: https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit
[vendoring]: https://github.com/FiloSottile/gvt
[gexec]: https://onsi.github.io/gomega/#gexec-testing-external-processes
[named-returns]: https://github.com/cloudfoundry/cli/wiki/Coding-Style-Guide#named-return-args

### TODO
* [x] main.go
* [x] gvt
* [x] webserver with /pizza/order route
* [x] order pizza use case
* [x] pizza repository
* [x] wire up /pizza/order to usecase
* [x] minimal gexec test
* [x] gexec test asserts we can make a single http request
* [ ] middleware (composition)
* [ ] logging
* [ ] stats ?
* [ ] configuration (port, log server, stats server)
* [ ] panic handler (middleware?)
* [ ] find a better name for "domain" package (h/t to Dave Cheney)

#notes
* I prefer using `gvt` to manage my dependencies, but you may prefer `godep`, or just doing it by hand. So long as you keep your dependencies tracked, make it easy to setup new development environments and bump dependencies regularly, you should be fine.
