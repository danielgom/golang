# Golang Learning 
​
## Topic 1: Go Fundamentals
​
### The Go platform
​
* [Docker is written in go](https://www.reddit.com/r/docker/comments/1w2aco/why_is_docker_written_in_go/)
* [C interop](https://programmer.ink/think/interoperability-between-go-and-c-language.html)
* [Go and garbage collection](https://tip.golang.org/doc/gc-guide)
* [Go Advantages](https://medium.com/@julienetienne/why-go-the-benefits-of-golang-6c39ea6cff7e)
* [Go build for the cloud especially gcp](https://cloud.google.com/go/home)
​
### Language Basics and Language structure
​
* [Start working with some exercises](https://go.dev/tour/welcome/1)
* Read this reference step by step, it's too much information [A must have!!!](https://go.dev/doc/effective_go)
* [Format your code please](https://pkg.go.dev/cmd/gofmt)
* [Packages and hello world](https://go.dev/tour/basics/1)
​
### Functions and methods
​
* [Functions by example](https://gobyexample.com/functions)
* Go to functions and method. Avoid recievers by now [A must have!!!](https://go.dev/doc/effective_go)
* [Alias typing functions](https://stackoverflow.com/questions/57108354/declare-function-with-type-alias-in-golang)
* [Defer](https://www.digitalocean.com/community/tutorials/understanding-defer-in-go-es)a
  * [defer io.NewObject().Close()](https://go.dev/tour/flowcontrol/12)
    * What is deferred ?
    * defer multiple prints?
    * defer sleep combination
​
### Data structures and Data types
* Default values? [Zero values](https://go.dev/tour/basics/12)
* A tricky notation [struct Inherit](https://sentry.io/answers/alias-type-definitions/)
* How to implement a set [Set](https://stackoverflow.com/questions/52231115/creating-map-with-empty-values)
* Grouping objects
  * Slices and arrays [w3schoo](https://www.w3schools.com/go/go_slices.php)
  * [Maps](https://go.dev/doc/effective_go#maps)
    * For each ? Iterate [for](https://go.dev/doc/effective_go#for)
* [Interfaces](https://www.digitalocean.com/community/tutorials/how-to-use-interfaces-in-go-es)
  * [Empty interface](https://go.dev/tour/methods/14)
    * `interface{}` ? == or != to `any`
    * [Alias](https://medium.com/a-journey-with-go/go-aliases-simple-and-efficient-8506d93b079e) typing and functions
​
### Pointers interfaces and functions types
* [Pointers and values](https://go.dev/doc/effective_go#pointers_vs_values)
* Why this function [signature](https://pkg.go.dev/net/http#HandleFunc) make sense ?
  * `func HandleFunc(pattern string, handler func(ResponseWriter, *Request))`
* Sentry example [interface and pointers](https://sentry.io/answers/interface-pointer-receiver/)
​
### Go modules and repo management
​
* Hope you never need this but: *Those Who Do Not Learn History Are Doomed To Repeat It.*
  * [Dep](https://github.com/golang/dep)
  * `go install` , everywhere
* [Go modules](https://go.dev/ref/mod)
* Practice: code two modules , publish the code on github and play with them.
  * What happen if you change the version
  * Tag a release
  * `go get yourpackage.com`
​
### Error handling
* [Error](https://go.dev/doc/effective_go#errors) everywhere
* Error [handling](https://earthly.dev/blog/golang-errors/)
* Recover
​
### Unit Testing and Benchmarking
[Test, test, test](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package)
[Optional BDD](https://github.com/go-bdd/gobdd)
​
### First list of exercises
* [Fizz buzz](https://en.wikipedia.org/wiki/Fizz_buzz)
* [Palindrome](https://en.wikipedia.org/wiki/Palindrome)
​
* Ask for input on a terminal , a file or hard code the value
## Topic 2: Concurrency and parallelism
​
* [CSP model](https://levelup.gitconnected.com/communicating-sequential-processes-csp-for-go-developer-in-a-nutshell-866795eb879d)
* [Patterns](https://go.dev/blog/pipelines)
* [Concurrency examples](https://go.dev/tour/concurrency/1)
* [It's all about the context](https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go)
​
## Topic 3: Go Back-end applications
* Write an API from scratch using [echo](https://echo.labstack.com)
  * Your endpoints should be able to crud an object
  * Your code needs 80 % of test coverage
    * You need to mock your db on testing
* Write a cmd application able to call a "public" API 
  * It should init a session and save requests and responses on a text file 
  * You can use this [API](https://rapidapi.com/BettingAPI/api/tennis-odds) if you are interested in
* Optional use grpc for [streaming data](https://grpc.io/docs/languages/go/quickstart/)
​
## Final Assessment
* Mock interview , look for some go expert
  * Use this questions to prepare your [interview](https://github.com/Devinterview-io/golang-interview-questions)
​
# Want to practice?
​
Coding to the next level [Gophercises](https://gophercises.com/)
[Exercism](https://exercism.org/tracks/go)
​
CheatSheets
[Everything you need](https://go.dev/doc/effective_go)
[GoRoadMap](https://github.com/Alikhll/golang-developer-roadmap)
[GoByExample](https://gobyexample.com/)