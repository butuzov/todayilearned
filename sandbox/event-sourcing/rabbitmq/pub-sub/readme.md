<!-- menu: Pub-Sub -->
<!-- weight: 50 -->

# RabbitMQ's `Publications/Subscribe` Pattern (using `fanout`)

![](../images/exchanges-topic-fanout-direct.png)

![](images/exchanges-topic-fanout-direct.png)

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
conn, err := connect(addr)
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

nameOfExchange := "foobar"

ch, err := conn.Channel()
if err != nil {
    log.Fatal("foo", err)
}

exchangeOptions := Exchange{
    Name: nameOfExchange,
    Type: "fanout",
    Durable: true,
}
if err := CreateExchange(ch, exchangeOptions); err != nil {
    log.Fatal("Exchange Not Created", err)
}
```

```go
// --- Queues Createion --------------------------------------------------------------
var queues = []string{}
for i := 0; i < 2; i++ {
    queue, err := QueueDeclare(ch, Queue{Exclusive:true})
    if err != nil {
        log.Fatal("bar", err)
    }

    fmt.Printf("\n%#v\n", queue)
    if err = ch.QueueBind(queue.Name, "", nameOfExchange, false, nil); err != nil {
        log.Fatal("ch.QueueBind", err)
    }
    
    queues = append(queues, queue.Name)
}

queues

> 
> amqp.Queue{Name:"amq.gen-WyOom2DRJrUXtoWQF8S6YA", Messages:0, Consumers:0}
> 
> amqp.Queue{Name:"amq.gen-5PmIU6Jj6PckwAPpS3rpsw", Messages:0, Consumers:0}

result >>> [amq.gen-WyOom2DRJrUXtoWQF8S6YA amq.gen-5PmIU6Jj6PckwAPpS3rpsw]
```

```go
// ---- Producing messages -------------------------------------------------
options := PublishOptions{
    Exchange: nameOfExchange,
} 

for i := 0; i < 3; i++ {
    Publish(ch, options, amqp.Publishing{
        Body: []byte(fmt.Sprintf("Event #%d", i)),
        DeliveryMode: 1, 
    }) 
}
```

```go
// --- Queues Consuming --------------------------------------------------------------
for i := range queues {
    cons := ConsumeOptions{
        Name:  fmt.Sprintf("consumer: %d", i),
        NoAck: true,
        Queue: queues[i],
    }
    chMsg, err := Consume(ch, cons)
    fmt.Printf("Consumer, %v\n", err)
    go Consumer(done, cons, chMsg)
}


time.Sleep(time.Second)

> Consumer, <nil>
> Start Consumer: consumer: 0
> consumer: 0 > Event #0
> Consumer, <nil>
> Start Consumer: consumer: 1
> consumer: 0 > Event #1
> consumer: 1 > Event #0
> consumer: 0 > Event #2
> consumer: 1 > Event #1
> consumer: 1 > Event #2
```

```go
// --- Cleanup --------------------------------------------------------------
time.AfterFunc(6*time.Second, func(){
    fmt.Println("Cleanup...")
    close(done) 
    
    for i := range queues{
        info, err := ch.QueueInspect(queues[i])
        fmt.Printf("Queue To Be deleted: %#v, %v\n", info, err)
        DeleteQueue(ch, queues[i])

    }
    ch.Close()
})

<-done
// one more pause to get messages back to jupyter
time.Sleep(time.Second)

> Cleanup...
> Exiting consumer: consumer: 1
> Queue To Be deleted: amqp.Queue{Name:"amq.gen-WyOom2DRJrUXtoWQF8S6YA", Messages:0, Consumers:1}, <nil>
> Exiting consumer: consumer: 0
> Queue To Be deleted: amqp.Queue{Name:"amq.gen-5PmIU6Jj6PckwAPpS3rpsw", Messages:0, Consumers:1}, <nil>
```