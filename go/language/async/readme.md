# Sync / Concurrency

* [x] sync.Once: _running functionality only once_ ([`once.go`](once.go) )
* [x] sync.Pool: _heavy resources allocation_ ([`pool.go`](pool.go) )
* [x] sync.WaitGroup: using waitgroups ([`once.go`](once.go), [`race_cond_mutex.go`](race_cond_mutex.go), [`race_cond_map.go`](race_cond_map.go), etc...  )
* [x] sync.Map: _Concurrent data structure - map_ ([`race_cond_map.go`](race_cond_map.go) )
* [x] sync.Cond ([`cond.go`](cond.go))
* [x] channels ( [`chans.go`](chans.go), [`generator.go`](generator.go), [`range_over_chan.go`](range_over_chan.go), etc... )
* [x] go routines (almost all of the provided files)
* [x] sync.Mutex: _Mutual Exclusion_ [`race_cond_mutex.go`](race_cond_mutex.go)
* [ ] TODO(butuzov): sync.RWMutex
* [x] atomic operations: ( [`atomic.go`](atomic.go) )
* [x] data races: conccurent access to shared state ( [`atomic.go`](atomic.go), [`race_cond_map.go`](race_cond_map.go) )
* [x] timers / timeouts ([`timeouts_1.go`](timeouts_1.go), [`timeouts_2.go`](timeouts_2.go))
* [x] ticker: ticker workflow ([`tick.go`](tick.go))
* [x] patterns: semaphore (x/sync/semaphore)
* [x] patterns: promise ([`range_over_chan.go`](range_over_chan.go))
* [x] patterns: worker pools ( [`workerpool_1.go`](workerpool_1.go), [`workerpool_2.go`](workerpool_2.go))
* [x] signals: OS signals [`signal.go`](signal.go)
* [x] context

### Reading List
* https://github.com/golang/sync
* [Concurrency is not parallelism](https://blog.golang.org/waza-talk)
* [Share Memory By Communicating](https://blog.golang.org/codelab-share)
* [Go Concurrency Patterns: Context](https://blog.golang.org/context)
* [Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)
* [Go Concurrency Patterns: Timing out, moving on](https://blog.golang.org/concurrency-timeouts)
* [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#1) + [`video`](https://www.youtube.com/watch?v=f6kdp27TYZs)
* [Advanced Go Concurrency Patterns](https://blog.golang.org/advanced-go-concurrency-patterns) + [`video`](http://www.youtube.com/watch?v=QDDwwePbDtw)
* [@ronna-s Presentation: Semaphores, correctly](https://github.com/ronna-s/sema-presentation)
