# RabbitMQ's  Routing  (`direct`)

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

### Example - FizBuz

All messaged from `n`th itteration that devidede by `5` go to buz channel, all dividede by `3` goto `fiz`

```go
done := make(chan struct{})

nameOfExchange := "fizbuz"

ch, err := conn.Channel()
if err != nil {
    log.Fatal("foo", err)
}

exchangeOptions := Exchange{
    Name: nameOfExchange,
    Type: "direct",
    Durable: true,
}
if err := CreateExchange(ch, exchangeOptions); err != nil {
    log.Fatal("Exchange Not Created", err)
}
```

```go
// --- Queues Createion --------------------------------------------------------------
var queues = []string{}
var fizbuz = map[int]string{3:"fiz", 5:"buz"}
for k, routingKey := range fizbuz {
    queue, err := QueueDeclare(ch, Queue{Exclusive:true})
    if err != nil {
        log.Fatal("queue", err)
    }

    fmt.Printf("Assign %#v to Routing Key(%s) Of exchange (%s)\n", queue.Name, routingKey, nameOfExchange)
    if err = ch.QueueBind(queue.Name, routingKey, nameOfExchange, false, nil); err != nil {
        log.Fatal("ch.QueueBind", err)
    }
    
    queues = append(queues, queue.Name)
}

stdout >>> Assign "amq.gen-KASByjxew-ZOYRHYk7oNBQ" to Routing Key(fiz) Of exchange (fizbuz)
stdout >>> Assign "amq.gen-D5mhpM-OD0QwrnyNES2GTg" to Routing Key(buz) Of exchange (fizbuz)
```

```go
for i := 0; i < 15; i++ {
    for k, routingKey := range fizbuz {
        if i%k != 0 {
            continue
        }

        Publish(ch, PublishOptions{
            Routing: routingKey,
            Exchange: nameOfExchange,
        }, amqp.Publishing{
            Body: []byte(fmt.Sprintf("Event (%s) #%d", routingKey, i)),
            DeliveryMode: 2, 
        }) 
    }  
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

stdout >>> Consumer, <nil>
stdout >>> Start Consumer: consumer: 0
stdout >>> consumer: 0 > Event (fiz) #0
stdout >>> Consumer, <nil>
stdout >>> Start Consumer: consumer: 1
stdout >>> consumer: 0 > Event (fiz) #3
stdout >>> consumer: 0 > Event (fiz) #6
stdout >>> consumer: 1 > Event (buz) #0
stdout >>> consumer: 0 > Event (fiz) #9
stdout >>> consumer: 1 > Event (buz) #5
stdout >>> consumer: 1 > Event (buz) #10
stdout >>> consumer: 0 > Event (fiz) #12
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

stdout >>> Cleanup...
stdout >>> Exiting consumer: consumer: 1
stdout >>> Queue To Be deleted: amqp.Queue{Name:"amq.gen-KASByjxew-ZOYRHYk7oNBQ", Messages:0, Consumers:1}, <nil>
stdout >>> Exiting consumer: consumer: 0
stdout >>> Queue To Be deleted: amqp.Queue{Name:"amq.gen-D5mhpM-OD0QwrnyNES2GTg", Messages:0, Consumers:1}, <nil>
```