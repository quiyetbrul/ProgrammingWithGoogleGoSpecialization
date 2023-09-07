# Once Synchronization

## Synchronous Initialization

### Initialization

- must happen once
- must happen before everything else
- how do you perform initialization with multiple goroutines
  - could perform initialization before starting the goroutines

### Sync.Once

- has one method, `once.Do(f)`
- fuction f is executed only one time
  - even if it is called in multiple goroutines
- all calls to `once.Do()` block until the first returns
  - ensures that initialization executes first

#### Sync.Once Example

- make two goroutines, initialization only once
- each goroutine executes `dostuff()`

```golang
var wg sync.WaitGroup
func main(){
  wg.Add(2)
  go dostuff()
  go dostuff()
  wg.Wait()
}
```

#### Using Sync.Once

- `setup()` should execute only once
- "hello" should not print until `setup()` returns

```golang
var on sync.Once
func setup(){
  fmt.Println("init")
}
func dostuff(){
  on.Do(setup)
  fmt.Println("hello")
  wg.Done()
}
```

```output
Init // Result of setup()
Hello // Result of one goroutine
hello // Result of the other goroutine
```

- `Init` appears only once
- `Init` appears before hello is printed

## Deadlock

- 
