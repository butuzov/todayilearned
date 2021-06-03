# RabbitMQ's Request/Reply Pattern

### Connection

```go
import (
    "context"
    "fmt"
    "log"
    "time"
    "strings"
    
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
conn, err := amqp.Dial(addr)
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
import "context"

func Consumer(done chan struct{}, consumer ConsumeOptions, messages <-chan amqp.Delivery) {
    fmt.Printf("Start Consumer: %s\n", consumer.Name)
    defer fmt.Printf("Exiting consumer: %s\n", consumer.Name)
 
    for {
        
            
        select {
            case m, ok := <-messages: 
            if !ok {
                continue
            }   
            
            if len(m.Body) == 0 {
                time.Sleep(time.Second)
                continue
            } 
            
        
            body := string(m.Body)
            fmt.Printf("%s > %s\n", consumer.Name, body)
            
            time.Sleep(time.Duration(strings.Count(body, ".")) * time.Second / 10.0)
            
            
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

queue, err := QueueDeclare(ch, Queue{Name: "workerpool", Durable: true})
if err != nil {
    log.Fatal("bar", err)
}

queue

result >>> {workerpool 0 0}
```

```go
err := ch.Qos(
    1,     // prefetch count
    0,     // prefetch size
    false, // global
)

err
```

```go
// --- Publishing -------------------------------------------------------------
options := PublishOptions{
    Routing: queue.Name,
} 

for i := 0; i < 15; i++ {
    i := i
    go func(){ 
        Publish(ch, options, amqp.Publishing{
            Body: []byte(fmt.Sprintf("Event #%d %s", i, strings.Repeat(".", i))),
            DeliveryMode: 2, 
        })
    }()
}
```

```go
type conChan struct {
    Ch <-chan amqp.Delivery
    Options ConsumeOptions
}
var consumers = []conChan{}
// --- Consuming --------------------------------------------------------------
for i := 0; i < 2; i++ {
    cons := ConsumeOptions{
        Name:  fmt.Sprintf("reader-%d", i),
        NoAck: true,
        Queue: queue.Name,
    }
    
    chMsg, err := Consume(ch, cons)
    fmt.Printf("Consumer, %v\n", err)
    
    consumers = append(consumers, conChan{
        Ch: chMsg,
        Options: cons,
    })
}

fmt.Sprintln("Started...")

stdout >>> Consumer, <nil>
stdout >>> Consumer, <nil>
result >>> Started...
```

```go
for i := range consumers {
    go Consumer(done, consumers[i].Options, consumers[i].Ch)
}

time.Sleep(10*time.Second)

stdout >>> Start Consumer: reader-0
stdout >>> reader-0 > Event #2 ..
stdout >>> Start Consumer: reader-1
stdout >>> reader-0 > Event #13 .............
stdout >>> reader-0 > Event #3 ...
stdout >>> reader-0 > Event #11 ...........
stdout >>> reader-0 > Event #9 .........
stdout >>> reader-0 > Event #10 ..........
stdout >>> reader-0 > Event #12 ............
stdout >>> reader-0 > Event #5 .....
stdout >>> reader-0 > Event #4 ....
stdout >>> reader-0 > Event #14 ..............
stdout >>> reader-0 > Event #6 ......
stdout >>> reader-0 > Event #8 ........
stdout >>> reader-0 > Event #0 
stdout >>> reader-0 > Event #7 .......
```

```go
// --- Cleanup --------------------------------------------------------------
time.AfterFunc(6*time.Second, func(){
    fmt.Println("Cleanup...")
    close(done) 
    i, err := ch.QueueInspect(queue.Name)
    fmt.Printf("%#v, %v\n", i, err)
    DeleteQueue(ch, queue.Name)
    ch.Close()
})

<-done
// one more pause to get messages back to jupyter
time.Sleep(time.Second)
```

> Note: Somehow only `reader-0` picks cahnnel initially, `reader-1` also works but jsut after some time. I guess this is jupyter related issue.