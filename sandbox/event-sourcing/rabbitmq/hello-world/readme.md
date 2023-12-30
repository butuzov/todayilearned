<!--  weight: 20 -->
# Hello World

```go
```mermaid


```
```

```
# expectations
+----------+     +------------------+     +----------+
| producer | --> | [ queue:foobar ] | --> | consumer |
+----------+     +------------------+     +----------+

# reality
+----------+     +-------------------+     +-------------------+     +----------+
| producer | --> | (exchange:foobar) | --> | [ queue:foobar ]  | --> | consumer |
+----------+     +-------------------+     +-------------------+     +----------+
```

### Connection

```go
import (
    "context"
    "fmt"
    "log"
    "time"
    
    "github.com/streadway/amqp"
)
```

```go
var addr = "amqp://admin:admin@rabbitmq1:5672/"
```

```go
var (
    connect func(addr string) (*amqp.Connection, error) 
    handle func (addr string, ch chan *amqp.Error, conn *amqp.Connection) 
)

connect = func(addr string) (*amqp.Connection, error) {
    conn, err := amqp.Dial(addr);
    connErrorChan := conn.NotifyClose(make(chan *amqp.Error))
    go handle(addr, connErrorChan, conn)
    return conn, err
}

handle = func(addr string, ch chan *amqp.Error, conn *amqp.Connection) {
    err := <- ch
    fmt.Println("connection lost... ")
    t1 := time.Now()
    for {
        c, err := connect(addr)
        if err != nil {
          break
        }
        fmt.Printf("reconnected in %s...\n", time.Since(t1))
        conn = c
    }
}
```

```go
conn, err := connect(addr);
if err != nil {
  log.Fatal(err)
}
```

### Code

```go
type Queue struct {
    Name       string 
    Durable    bool
    Exclusive  bool
    AutoDelete bool
    NoWait     bool
    Args       amqp.Table
}
```

```go
type Exchange struct {
    Name       string
    Type       string
    Durable    bool
    AutoDelete bool
    Internal   bool
    NoWait     bool 
    Args       amqp.Table
}
```

```go
type PublishOptions struct {
    Exchange  string
    Routing   string
    Mandatory bool
    Immediate bool
}
```

```go
type ConsumeOptions struct {
    Queue     string
    Name      string
    NoLocal   bool     
    NoAck     bool      
    Exclusive bool
    NoWait    bool
    Args      amqp.Table
}
```

```go
func QueueDeclare(ch *amqp.Channel, q Queue) (amqp.Queue, error) {
    if q.Args == nil {
        q.Args = amqp.Table{}
    } 
    
    return ch.QueueDeclare(q.Name, q.Durable, q.AutoDelete, q.Exclusive, q.NoWait, q.Args)
}
```

```go
func DeleteQueue(ch *amqp.Channel, name string) (string, error) {
    cnt, err := ch.QueueDelete(name, false, false, true)
    return fmt.Sprintf("Deleted %d messages", cnt), err
}
```

```go
func CreateExchange(ch *amqp.Channel, e Exchange) error {
    if e.Args == nil {
        e.Args = amqp.Table{}
    } 
    
    return ch.ExchangeDeclare(e.Name, e.Type, e.Durable, e.AutoDelete, e.Internal, e.NoWait, e.Args)
}
```

```go
func Publish(ch *amqp.Channel, p PublishOptions, m amqp.Publishing) error {
    return ch.Publish(p.Exchange, p.Routing, p.Mandatory, p.Immediate, m)
}
```

```go
func Consume(ch *amqp.Channel, c ConsumeOptions) (<-chan amqp.Delivery, error) {
    if c.Args == nil {
        c.Args = amqp.Table{}
    } 
    
    return ch.Consume(c.Queue, c.Name, c.NoAck, c.Exclusive, c.NoLocal, c.NoWait, c.Args)
}
```

```go
import "math/rand"
import "context"

func Consumer(done chan struct{}, consumer ConsumeOptions, messages <-chan amqp.Delivery) {
    fmt.Printf("Start Consumer: %s\n", consumer.Name)
    defer fmt.Printf("Exiting consumer: %s\n", consumer.Name)
 
    for {
        // random deley helps to switch consumers to other goroutines
        time.Sleep(time.Duration(rand.Int63n(10)) * time.Millisecond)
            
        select {
            case m, ok := <-messages: 
            if !ok {
                continue
            }   
            
            if len(m.Body) == 0 {
                time.Sleep(time.Second)
                continue
            } 
            
            fmt.Printf("%s > %s\n", consumer.Name, string(m.Body))
            
            
            case <-done:
                return
            default:
        }       
    } 
}
```

### Example

```go
done := make(chan struct{})

ch, err := conn.Channel()
if err != nil {
    log.Fatal("foo", err)
}

queue, err := QueueDeclare(ch, Queue{Name: "hello-world"})
if err != nil {
    log.Fatal("bar", err)
}
queue

result >>> {hello-world 0 0}
```

```go
options := PublishOptions{
    Routing: queue.Name,
} 

for i := 0; i < 5; i++ {
    Publish(ch, options, amqp.Publishing{
        Body: []byte(fmt.Sprintf("Event #%d", i)),
    })  
}
```

```go
// --- Consuming --------------------------------------------------------------
cons := ConsumeOptions{
    Name:  fmt.Sprint("helloworld-reader"),
    Queue: queue.Name,
    NoAck: true,
}

chMsg, err := Consume(ch, cons)
fmt.Printf("Consumer, %v\n", err)

go Consumer(done, cons, chMsg)

time.Sleep(time.Second)

> Consumer, <nil>
> Start Consumer: helloworld-reader
> helloworld-reader > Event #0
> helloworld-reader > Event #1
> helloworld-reader > Event #2
> helloworld-reader > Event #3
> helloworld-reader > Event #4
```

```go
// --- Cleanup --------------------------------------------------------------
time.AfterFunc(1*time.Second, func(){
    fmt.Println("Cleanup...")
    close(done)
    
    i, err := ch.QueueInspect(queue.Name)
    fmt.Sprintf("%#v", i)
    
    DeleteQueue(ch, queue.Name)
    ch.Close()
})

<-done
time.Sleep(time.Second)

> Cleanup...
> Exiting consumer: helloworld-reader
```