# บทที่ 9: Goroutines & Channels

> Go ทำ concurrency ง่ายกว่า Java มาก — goroutine เบากว่า thread 1000x, channel แทนที่ shared memory

## Java Thread vs Go Goroutine

```java
// Java — Thread หนัก (~1MB stack per thread)
ExecutorService executor = Executors.newFixedThreadPool(10);
executor.submit(() -> {
    System.out.println("running in thread: " + Thread.currentThread().getName());
});
executor.shutdown();
```

```go
// Go — Goroutine เบามาก (~2KB stack, โตได้)
go func() {
    fmt.Println("running in goroutine")
}()
```

> **goroutine ≠ thread** — Go runtime จัดการ multiplex goroutine ลง OS thread ให้อัตโนมัติ (M:N scheduling)

## Channel — สื่อสารระหว่าง goroutine

```java
// Java — ใช้ shared state + synchronization
BlockingQueue<String> queue = new LinkedBlockingQueue<>();
queue.put("hello");           // producer
String msg = queue.take();    // consumer
```

```go
// Go — channel เป็น built-in type
ch := make(chan string)       // unbuffered channel
ch <- "hello"                 // ส่ง (block จนกว่าจะมีคนรับ)
msg := <-ch                   // รับ (block จนกว่าจะมีคนส่ง)

ch := make(chan string, 10)   // buffered channel (เหมือน BlockingQueue ขนาด 10)
```

> **ปรัชญา Go:** *"Don't communicate by sharing memory; share memory by communicating."* — ส่งข้อมูลผ่าน channel แทนที่จะ lock + share variable

## WaitGroup — รอให้ goroutine ทำเสร็จ

```java
// Java
CountDownLatch latch = new CountDownLatch(3);
executor.submit(() -> { doWork(); latch.countDown(); });
latch.await();
```

```go
// Go
var wg sync.WaitGroup
wg.Add(3)
for i := 0; i < 3; i++ {
    go func(id int) {
        defer wg.Done()
        doWork(id)
    }(i)
}
wg.Wait()
```

## Select — multiplex channels

```go
select {
case msg := <-ch1:
    fmt.Println("from ch1:", msg)
case msg := <-ch2:
    fmt.Println("from ch2:", msg)
case <-time.After(time.Second):
    fmt.Println("timeout!")
default:
    fmt.Println("nothing ready")
}
```

> **Java เทียบ:** `select` ≈ `CompletableFuture.anyOf()` แต่อ่านง่ายกว่ามาก

## ไฟล์ในบทนี้

- `main.go` — goroutine, channel, WaitGroup, select, buffered channel
