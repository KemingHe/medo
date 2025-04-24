# Goroutines and Concurrency in Go

Go's concurrency model is built around goroutines and channels, providing a simpler approach to concurrent programming compared to thread-based models.

## Concurrency vs. Parallelism (Not the Same)

- **Concurrency**: Handling multiple tasks in overlapping time periods (not necessarily simultaneously)
- **Parallelism**: Executing multiple tasks simultaneously (requires multiple cores)
- Go's runtime scheduler automatically distributes goroutines across available CPU cores

## Goroutines

Goroutines are lightweight threads managed by the Go runtime. They enable concurrent execution with minimal overhead.

```go
// Basic goroutine
go func() {
    // Code runs concurrently
    fmt.Println("Running in goroutine")
}()

// With arguments
go func(message string) {
    fmt.Println(message)
}("Hello from goroutine")
```

**Key Properties:**

- **Lightweight**: ~2KB initial stack (vs. MB for OS threads)
- **Multiplexed**: Many goroutines execute on fewer OS threads
- **Fast creation**: Much faster than thread creation
- **Communicate via channels**, not shared memory

## Synchronization with `sync` Package

### WaitGroup

WaitGroup is used to wait for a collection of goroutines to finish.

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)  // Increment counter
    go func(id int) {
        defer wg.Done()  // Decrement counter when goroutine completes
        fmt.Printf("Worker %d done\n", id)
    }(i)
}

wg.Wait()  // Block until counter reaches zero
fmt.Println("All workers done")
```

### Mutexes

Mutexes protect access to shared resources.

```go
var (
    counter int
    mutex   sync.Mutex
)

func increment() {
    mutex.Lock()
    counter++  // Safe access to shared variable
    mutex.Unlock()
}

// RWMutex allows multiple readers but exclusive writers
var rwMutex sync.RWMutex

func read() {
    rwMutex.RLock()  // Shared read lock
    // Read operations
    rwMutex.RUnlock()
}

func write() {
    rwMutex.Lock()  // Exclusive write lock
    // Write operations
    rwMutex.Unlock()
}
```

## Channels

Channels provide a way for goroutines to communicate and synchronize.

```go
// Unbuffered channel
ch := make(chan int)

// Buffered channel with capacity 5
buffered := make(chan string, 5)

// Send value (blocks until received)
ch <- 42

// Receive value (blocks until sent)
value := <-ch

// Non-blocking operations with select
select {
case v := <-ch:
    fmt.Println("Received:", v)
case ch <- 10:
    fmt.Println("Sent 10")
default:
    fmt.Println("No communication")
}
```

### Channel Patterns

#### Worker Pool

Worker pools distribute work across multiple goroutines to achieve parallel processing, improving performance for CPU-intensive or I/O-bound tasks. This pattern prevents overwhelming system resources while maintaining controlled concurrency.

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- j * 2  // Process job and send result
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)  // No more jobs

    // Collect results
    for a := 1; a <= 9; a++ {
        <-results
    }
}
```

#### Timeout

```go
select {
case result := <-ch:
    fmt.Println("Result:", result)
case <-time.After(500 * time.Millisecond):
    fmt.Println("Timeout")
}
```

## Channel Directionality

```go
func send(ch chan<- int) {  // Send-only channel
    ch <- 42
}

func receive(ch <-chan int) {  // Receive-only channel
    val := <-ch
    fmt.Println(val)
}
```

## Common Goroutine Patterns

### Fan-out, Fan-in

Fan-out distributes work to multiple goroutines for parallel processing (scaling horizontally), while fan-in consolidates results from multiple sources into a single channel. This pattern is ideal for computationally intensive tasks that can be divided into independent subtasks.

```go
func fanOut(input <-chan int, n int) []<-chan int {
    channels := make([]<-chan int, n)
    for i := 0; i < n; i++ {
        channels[i] = worker(input)
    }
    return channels
}

func fanIn(channels []<-chan int) <-chan int {
    merged := make(chan int)
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for v := range ch {
                merged <- v
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(merged)
    }()
    
    return merged
}
```

## Best Practices

1. **Never leak goroutines** - ensure they can exit properly
2. **Prefer channels** for communication between goroutines
3. **Use mutexes** only when shared memory access is required
4. **Pass data explicitly** to goroutines using parameters
5. **Close channels** from the sender side, never the receiver
6. **Check for nil channels** - they block forever
7. **Use context** for cancellation and timeouts

## Gotchas

1. **Loop variable capture**

    ```go
    // Incorrect - all goroutines share same 'i'
    for i := 0; i < 5; i++ {
        go func() {
            fmt.Println(i)  // May print unexpected values
        }()
    }

    // Correct - pass 'i' as parameter
    for i := 0; i < 5; i++ {
        go func(id int) {
            fmt.Println(id)  // Prints 0,1,2,3,4 (order not guaranteed)
        }(i)
    }
    ```

2. **Premature program exit**:
   Main doesn't wait for goroutines to finish unless synchronized

3. **Channel deadlocks**
   - Sending to a channel with no receivers
   - Receiving from a channel with no senders
   - Full deadlocks detected at runtime: "fatal error: all goroutines are asleep"
