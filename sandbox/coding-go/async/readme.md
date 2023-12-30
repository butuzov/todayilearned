<!-- weight: 3 -->

# Concurrency

## Reading

### Books

- [Concurrency in Go: Tools and Techniques for Developers](https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197)
- [Learn Concurrent Programming with Go](https://www.manning.com/books/learn-concurrent-programming-with-go)

## Articles

* [Concurrency is not parallelism](https://blog.golang.org/waza-talk)
* [Go Concurrency Patterns: Context](https://blog.golang.org/context)
* [Go Concurrency Patterns: Timing out, moving on](https://blog.golang.org/concurrency-timeouts)
* [Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)
* [Share Memory By Communicating](https://blog.golang.org/codelab-share)

### Packages
* https://github.com/golang/sync
* https://github.com/butuzov/harmony

### Talks

* [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#1) + [`video`](https://www.youtube.com/watch?v=f6kdp27TYZs)
* [Advanced Go Concurrency Patterns](https://blog.golang.org/advanced-go-concurrency-patterns) + [`video`](http://www.youtube.com/watch?v=QDDwwePbDtw)
* [@ronna-s Presentation: Semaphores, correctly](https://github.com/ronna-s/sema-presentation)


## Patterns, Examples, Usage, etc..

{{% list "atomic.go,cond.go,context_timeout.go,once.go,promise.go,race_cond_mutex.go,semaphores.go,tick.go,timeouts_2.go,workerpool_2.go,chans.go,context_cancel.go,generator.go,pool.go,race_cond_map.go,range_over_chan.go,signal.go,timeouts_1.go,workerpool_1.go" %}}
