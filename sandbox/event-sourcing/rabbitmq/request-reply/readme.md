<!-- menu: "Request-Reply" -->
<!-- weight: 80 -->
# RabbitMQ's Request/Reply Pattern

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
import "github.com/google/uuid"
```

```go
import "github.com/davecgh/go-spew/spew"
```

```go
done := make(chan struct{})

ch, err := conn.Channel()
if err != nil {
    log.Fatal("foo", err)
}
```

```go
// create queue for server (one that will consume requests)
rpcQueue, err := QueueDeclare(ch, Queue{Name:"rpc_queue"})
if err != nil {
    log.Fatal("queue", err)
}
                                        
rpcQueue

result >>> {rpc_queue 0 0}
```

```go
// create unnamed queue for responses.
queue, err := QueueDeclare(ch, Queue{Exclusive:true})
if err != nil {
    log.Fatal("queue", err)
}
queue

result >>> {amq.gen-ICLLgRGJhr833OkwmS0RJA 0 0}
```

```go
ch.Qos(
    1,     // prefetch count
    0,     // prefetch size
    false, // global
)
```

```go
// creae server client to accept request
chInput, err := Consume(ch, ConsumeOptions{
    Name:  "server",
    NoAck: true,
    Queue: rpcQueue.Name,
})

chOutput, err := Consume(ch, ConsumeOptions{
    Name:  "client",
    NoAck: true,
    Queue: queue.Name,
})
```

```go
Publish(ch, PublishOptions{Routing:"rpc_queue"}, amqp.Publishing{
    ContentType:   "text/plain",
    CorrelationId: uuid.New().String(),
    ReplyTo:       queue.Name,
    Body:          []byte("ping"),
})
```

```go
// Server Consuming Events
go func(channel <-chan amqp.Delivery){
    for event := range channel {
        
        // response to other queue
        Publish(ch, PublishOptions{Routing: event.ReplyTo}, amqp.Publishing{
            ContentType:   "text/plain",
            CorrelationId: event.CorrelationId,
            ReplyTo:       queue.Name,
            Body:          []byte(fmt.Sprintf("%s > pong", string(event.Body))),
        })
        
    }    
}(chInput)
```

```go
// Server Consuming Events
event := <-chOutput 

fmt.Printf("Got Answer (body) - %s \n", event.Body)
fmt.Printf("Got Answer (correlation id) - %s\n", event.CorrelationId)
// do not acknoledge recieved messge
event.Ack(false)

> Got Answer (body) - ping > pong
> Got Answer (correlation id) - 978063f1-8edd-4ead-ac2e-973c52c57860
```

```go
// pproccess request at server (customer_)
```

```go
// --- Cleanup --------------------------------------------------------------
time.AfterFunc(6*time.Second, func(){
    fmt.Println("Cleanup...")
    close(done) 
    
  
    info, err := ch.QueueInspect(rpcQueue.Name)
    fmt.Printf("Queue To Be deleted: %#v, %v\n", info, err)
    DeleteQueue(ch, rpcQueue.Name)

    info, err := ch.QueueInspect(queue.Name)
    fmt.Printf("Queue To Be deleted: %#v, %v\n", info, err)
    DeleteQueue(ch, queue.Name)
     
    
    ch.Close()
})

<-done
// one more pause to get messages back to jupyter
time.Sleep(time.Second)

> Cleanup...
> Queue To Be deleted: amqp.Queue{Name:"rpc_queue", Messages:0, Consumers:0}, Exception (504) Reason: "channel/connection is not open"
> Queue To Be deleted: amqp.Queue{Name:"amq.gen-ICLLgRGJhr833OkwmS0RJA", Messages:0, Consumers:0}, Exception (504) Reason: "channel/connection is not open"
```