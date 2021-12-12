## `time` package - wall & monotoic clock

```go
import (
    "time"
    "fmt"
    "os"
)
```

## Time's

```go
time.Now()

result >>> 2021-06-21 03:00:48.7992596 +0000 UTC m=+13.928864001
```

```go
// UTC
time.Now().UTC()

result >>> 2021-06-21 03:00:50.5526857 +0000 UTC
```

```go
// unixtimestamp
time.Now().Unix()

result >>> 1624244451
```

## Layouts

```go
// Monday, January 2nd 2006
// at 15.04.05 in Rocky MOuntimtes (-07 timezone)
```

```go
// Go Provides different predefined layouts
const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700"
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // Handy time stamps.
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
    )

// RFC3339 for example!
time.Now().Format(time.RFC3339)

result >>> 2021-06-14T13:33:17Z
```

```go
time.Date(2019, 11, 7, 20, 34, 58, 651387237, time.UTC)

result >>> 2019-11-07 20:34:58.651387237 +0000 UTC
```

```go
t1 := time.Now()
```

```go
// tuple yar, month, date
t1.UTC().Date()

result >>> 2021 June 21
```

```go
hour, min, sec := t1.Clock()
fmt.Sprintf("%d:%d:%d", hour, min, sec)

result >>> 3:4:37
```

```go
// calling day, month year, minute directly...
fmt.Sprintf("%d %d", t1.Day(), t1.Second())

result >>> 21 37
```

```go
year, month, date := t1.UTC().Date()
time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

result >>> 2021-06-01 00:00:00 +0000 UTC
```

```go
r := func(){
    start := time.Now()
    defer func(){println( "Took me", time.Since(start) , "to execute")}()

    time.Sleep(2 * time.Second) 
}
r()

stderr >>> Took me 2.0014513s to execute
```

### Subtraction 

```go
year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

isYear3000AfterYear2000 := year3000.After(year2000) // True
isYear2000AfterYear3000 := year2000.After(year3000) // False

fmt.Printf("year3000.After(year2000) = %v\n", isYear3000AfterYear2000)
fmt.Printf("year2000.After(year3000) = %v\n", isYear2000AfterYear3000)

stdout >>> year3000.After(year2000) = true
stdout >>> year2000.After(year3000) = false
result >>> 33 <nil>
```

## TimeZones

```go
newYork, err := time.LoadLocation("America/New_York")
fmt.Sprintf("%s %v", newYork, err)

result >>> America/New_York <nil>
```

```go
t2 := time.Date(2019, 11, 7, 20, 34, 58, 651387237, newYork)
t2.Location()

result >>> America/New_York
```

```go
t2

result >>> 2019-11-07 20:34:58.651387237 -0500 EST
```

```go
t2.UTC()

result >>> 2019-11-08 01:34:58.651387237 +0000 UTC
```

```go
t2.Equal(time.Now())

result >>> false
```

```go
loc := time.FixedZone("Beijing Time", -60)
```

```go
t2.In(loc)

result >>> 2019-11-08 01:33:58.651387237 -0001 Beijing Time
```

### Gotchas

```go
time.Now().Format(time.UnixDate)

result >>> Fri Jun 25 15:25:52 UTC 2021
```

```go
os.Setenv("TZ", "Africa/Cairo")

time.Now().Format(time.UnixDate)

result >>> Fri Jun 25 15:25:55 UTC 2021
```

## Timeouts, Timers and Tickers

```go
// read from channel in 3 seconds
<-time.After(3 * time.Second)

result >>> 2021-06-23 02:39:31.1250852 +0000 UTC m=+37379.619555901 true
```

```go
ch := make(chan string)
time.AfterFunc(time.Second, func(){
    ch <- "foobar"
})

<-ch

result >>> foobartrue
```

```go
ticker := time.NewTicker(500 * time.Millisecond)
go func() {
   var f = "15:04:05.999999999"
   for t := range ticker.C {
      fmt.Println("ping @", t.Format(f))
   }
}()

time.Sleep(1600 * time.Millisecond)
ticker.Stop()
fmt.Sprintln("ticker stoped")

stdout >>> ping @ 02:47:15.6769849
stdout >>> ping @ 02:47:16.1734136
stdout >>> ping @ 02:47:16.673634
result >>> ticker stoped
```