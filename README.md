# Best Practices

A collection of things I consider to be best practices in test-driven go development.

This was inspired by Can Berk Guder's [iOS best practices repo](https://github.com/cbguder/bestpractices).

## Practices and Principles

Some of the practices and principles I tried to demonstrate in this sample project are:

* [Test-driven development][tdd] using [Ginkgo][] and [Gomega][]
* [Dependency injection][di] using hand-rolled constructor functions.
* [Vendored dependencies][dependencies] with [gvt][gvt]
* [Single responsibility principle][srp]
* [Composition over inheritance][coi]
* [Responsible usage of concurrency][concurrency]
* [Minimal integration tests][integrated] with [Gomega's gexec package][gexec]

[Ginkgo]: https://github.com/onsi/ginkgo
[Gomega]: https://github.com/onsi/gomega
[coi]: http://en.wikipedia.org/wiki/Composition_over_inheritance
[di]: http://en.wikipedia.org/wiki/Dependency_injection
[srp]: http://en.wikipedia.org/wiki/Single_responsibility_principle
[tdd]: http://en.wikipedia.org/wiki/Test-driven_development
[concurrency]: https://divan.github.io/posts/go_concurrency_visualize/
[integrated]: http://blog.thecodewhisperer.com/permalink/integrated-tests-are-a-scam
[dependencies]: https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit
[gvt]: https://github.com/FiloSottile/gvt
[gexec]: https://onsi.github.io/gomega/#gexec-testing-external-processes

### TODO
* [x] main.go
* [x] gvt
* [x] webserver with /pizza/order route
* [x] order pizza use case
* [x] pizza repository
* [ ] wire up /pizza/order to usecase
* [ ] middleware (composition)
* [ ] logging
* [ ] stats ?
* [ ] configuration (port, log server, stats server)
* [ ] panic handler (middleware?)
