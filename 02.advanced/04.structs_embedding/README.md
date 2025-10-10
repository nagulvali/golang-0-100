# Go Struct Embedding

## Table of Contents
1. [Introduction](#introduction)
2. [What is Struct Embedding?](#what-is-struct-embedding)
3. [Composition vs Inheritance](#composition-vs-inheritance)
4. [Basic Embedding](#basic-embedding)
5. [Anonymous vs Named Embedding](#anonymous-vs-named-embedding)
6. [Field Promotion](#field-promotion)
7. [Method Promotion](#method-promotion)
8. [Multiple Embedding](#multiple-embedding)
9. [Embedding Interfaces](#embedding-interfaces)
10. [Method Overriding](#method-overriding)
11. [Name Conflicts and Resolution](#name-conflicts-and-resolution)
12. [Embedded Pointers](#embedded-pointers)
13. [Common Patterns](#common-patterns)
14. [Real-World Examples](#real-world-examples)
15. [Best Practices](#best-practices)
16. [Embedding vs Other Approaches](#embedding-vs-other-approaches)

---

## Introduction

Struct embedding is Go's approach to code reuse and composition. Unlike traditional object-oriented languages that use inheritance hierarchies, Go uses embedding to build complex types from simpler ones. This approach leads to more flexible and maintainable code.

**Key Concepts:**
- Go doesn't have classes or inheritance
- Embedding allows one struct to include another
- Embedded fields and methods are "promoted" to the outer struct
- Composition is favored over inheritance
- Interface embedding enables powerful abstractions

---

## What is Struct Embedding?

**Struct embedding** is a mechanism where one struct type includes another struct type as an **anonymous field** (without a field name). When a struct is embedded:
- Its fields become accessible as if they were fields of the outer struct
- Its methods can be called on the outer struct
- The embedded type can still be accessed by its type name

Think of it as "has-a" relationship rather than "is-a":
- **Inheritance (OOP)**: Dog **is-a** Animal
- **Embedding (Go)**: Dog **has-a** Animal (more precisely, Dog embeds Animal's behavior)

---

## Composition vs Inheritance

### Traditional Inheritance (Other Languages)

```
// Pseudocode - NOT Go
class Animal {
    name: string
    speak(): void
}

class Dog extends Animal {
    breed: string
    // Inherits name and speak()
}
```

**Problems with Inheritance:**
- Tight coupling between parent and child
- Deep hierarchies become fragile
- Diamond problem (multiple inheritance)
- Hard to change base classes without affecting children

### Go's Composition Approach

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println("Some sound")
}

type Dog struct {
    Animal  // Embedding (composition)
    Breed   string
}

// Dog has access to Animal's fields and methods
// But it's composition, not inheritance
```

**Benefits:**
- Loose coupling
- More flexible
- Easier to understand and maintain
- No inheritance hierarchies
- Clear ownership of behavior

### Comparison Table

| Aspect | Inheritance (OOP) | Embedding (Go) |
|--------|------------------|----------------|
| **Relationship** | "is-a" | "has-a" |
| **Coupling** | Tight | Loose |
| **Flexibility** | Limited | High |
| **Multiple** | Complex/Problematic | Straightforward |
| **Override** | Implicit | Explicit |
| **Access** | Protected/Private keywords | Package visibility |

---

## Basic Embedding

### Named Field (Not Embedding)

This is **not** embedding - it's a regular nested struct:

```go
type Address struct {
    Street string
    City   string
}

type Person struct {
    Name    string
    Address Address  // Named field
}

func main() {
    p := Person{
        Name: "Alice",
        Address: Address{
            Street: "123 Main St",
            City:   "New York",
        },
    }
    
    // Must access through field name
    fmt.Println(p.Address.City)  // New York
    // fmt.Println(p.City)       // ERROR: undefined
}
```

### Anonymous Field (Embedding)

This **is** embedding:

```go
type Address struct {
    Street string
    City   string
}

type Person struct {
    Name    string
    Address  // Anonymous field (embedding)
}

func main() {
    p := Person{
        Name: "Bob",
        Address: Address{
            Street: "456 Oak Ave",
            City:   "Boston",
        },
    }
    
    // Can access directly (promoted)
    fmt.Println(p.City)          // Boston
    
    // Can also access through type name
    fmt.Println(p.Address.City)  // Boston
}
```

### Multiple Fields Example

```go
type Contact struct {
    Email string
    Phone string
}

type Person struct {
    Name string
    Age  int
    Contact  // Embedded
}

func main() {
    p := Person{
        Name: "Charlie",
        Age:  30,
        Contact: Contact{
            Email: "charlie@example.com",
            Phone: "555-0123",
        },
    }
    
    fmt.Println(p.Name)   // Charlie
    fmt.Println(p.Email)  // charlie@example.com (promoted)
    fmt.Println(p.Phone)  // 555-0123 (promoted)
}
```

---

## Anonymous vs Named Embedding

### Named Embedding (Composition)

```go
type Engine struct {
    Horsepower int
    Type       string
}

type Car struct {
    Brand  string
    Engine Engine  // Named field
}

func main() {
    car := Car{
        Brand: "Toyota",
        Engine: Engine{
            Horsepower: 200,
            Type:       "V6",
        },
    }
    
    // Must use field name
    fmt.Println(car.Engine.Horsepower)  // 200
}
```

### Anonymous Embedding (Promotion)

```go
type Engine struct {
    Horsepower int
    Type       string
}

type Car struct {
    Brand  string
    Engine  // Anonymous (embedded)
}

func main() {
    car := Car{
        Brand: "Honda",
        Engine: Engine{
            Horsepower: 180,
            Type:       "I4",
        },
    }
    
    // Direct access (promoted)
    fmt.Println(car.Horsepower)  // 180
    
    // Or through type name
    fmt.Println(car.Engine.Horsepower)  // 180
}
```

### When to Use Each

**Use Named Fields When:**
- The relationship is purely structural (has-a)
- You want explicit access paths
- The field name adds clarity
- You don't want promotion

**Use Embedding When:**
- You want to promote fields/methods
- Building type hierarchies
- Implementing interfaces by delegation
- Creating wrapper types

---

## Field Promotion

When you embed a struct, its fields are **promoted** to the outer struct, making them directly accessible.

### Basic Field Promotion

```go
type Dimensions struct {
    Width  float64
    Height float64
    Depth  float64
}

type Box struct {
    Dimensions  // Embedded
    Weight     float64
}

func main() {
    box := Box{
        Dimensions: Dimensions{
            Width:  10,
            Height: 20,
            Depth:  5,
        },
        Weight: 50,
    }
    
    // Promoted fields - direct access
    fmt.Println(box.Width)   // 10
    fmt.Println(box.Height)  // 20
    fmt.Println(box.Depth)   // 5
    
    // Own field
    fmt.Println(box.Weight)  // 50
}
```

### Nested Embedding

Promotion works through multiple levels:

```go
type Location struct {
    Latitude  float64
    Longitude float64
}

type Address struct {
    Location  // Embedded
    Street   string
    City     string
}

type Business struct {
    Address  // Embedded
    Name    string
}

func main() {
    biz := Business{
        Name: "Tech Corp",
        Address: Address{
            Street: "123 Main St",
            City:   "Seattle",
            Location: Location{
                Latitude:  47.6062,
                Longitude: -122.3321,
            },
        },
    }
    
    // All levels promoted
    fmt.Println(biz.Name)       // Tech Corp
    fmt.Println(biz.City)       // Seattle
    fmt.Println(biz.Latitude)   // 47.6062 (promoted from Location!)
}
```

### Promotion Rules

```go
type A struct {
    Field1 string
}

type B struct {
    A      // Embedded
    Field2 string
}

type C struct {
    B      // Embedded
    Field3 string
}

func main() {
    c := C{
        Field3: "C",
        B: B{
            Field2: "B",
            A: A{
                Field1: "A",
            },
        },
    }
    
    // All fields promoted to top level
    fmt.Println(c.Field1)  // A (from A)
    fmt.Println(c.Field2)  // B (from B)
    fmt.Println(c.Field3)  // C (from C)
    
    // Can still access through path
    fmt.Println(c.B.A.Field1)  // A
}
```

---

## Method Promotion

Just like fields, methods from embedded types are promoted to the outer type.

### Basic Method Promotion

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) Introduce() {
    fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.Name, p.Age)
}

func (p Person) GetAge() int {
    return p.Age
}

type Employee struct {
    Person    // Embedded
    EmployeeID int
    Department string
}

func main() {
    emp := Employee{
        Person: Person{
            Name: "Alice",
            Age:  30,
        },
        EmployeeID: 12345,
        Department: "Engineering",
    }
    
    // Promoted methods - call directly on Employee
    emp.Introduce()  // Hi, I'm Alice and I'm 30 years old.
    
    age := emp.GetAge()
    fmt.Println(age)  // 30
}
```

### Methods with Pointer Receivers

```go
type Counter struct {
    Count int
}

func (c *Counter) Increment() {
    c.Count++
}

func (c *Counter) Decrement() {
    c.Count--
}

func (c Counter) GetCount() int {
    return c.Count
}

type StatsTracker struct {
    Counter  // Embedded
    Name    string
}

func main() {
    tracker := StatsTracker{
        Counter: Counter{Count: 0},
        Name:    "Page Views",
    }
    
    // Promoted methods work even with pointer receivers
    tracker.Increment()
    tracker.Increment()
    tracker.Increment()
    
    fmt.Printf("%s: %d\n", tracker.Name, tracker.GetCount())
    // Page Views: 3
}
```

### Receiver Type in Promoted Methods

```go
type Base struct {
    Value int
}

// Value receiver
func (b Base) ShowValue() {
    fmt.Printf("Value: %d (address: %p)\n", b.Value, &b)
}

// Pointer receiver
func (b *Base) SetValue(v int) {
    b.Value = v
}

type Derived struct {
    Base  // Embedded
}

func main() {
    d := Derived{Base: Base{Value: 10}}
    
    // Value receiver method - creates copy
    d.ShowValue()
    
    // Pointer receiver method - modifies original
    d.SetValue(20)
    fmt.Println(d.Value)  // 20
}
```

---

## Multiple Embedding

You can embed multiple types in a single struct, and all their fields/methods are promoted.

### Basic Multiple Embedding

```go
type Name struct {
    First string
    Last  string
}

func (n Name) FullName() string {
    return n.First + " " + n.Last
}

type Contact struct {
    Email string
    Phone string
}

func (c Contact) HasEmail() bool {
    return c.Email != ""
}

type Person struct {
    Name     // Embedded
    Contact  // Embedded
    Age     int
}

func main() {
    p := Person{
        Name: Name{
            First: "John",
            Last:  "Doe",
        },
        Contact: Contact{
            Email: "john@example.com",
            Phone: "555-0123",
        },
        Age: 35,
    }
    
    // Fields promoted from both embedded types
    fmt.Println(p.First)     // John
    fmt.Println(p.Email)     // john@example.com
    
    // Methods promoted from both
    fmt.Println(p.FullName())  // John Doe
    fmt.Println(p.HasEmail())  // true
}
```

### Complex Multiple Embedding

```go
type Identifiable struct {
    ID string
}

func (i Identifiable) GetID() string {
    return i.ID
}

type Timestamped struct {
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (t Timestamped) Age() time.Duration {
    return time.Since(t.CreatedAt)
}

type Taggable struct {
    Tags []string
}

func (t *Taggable) AddTag(tag string) {
    t.Tags = append(t.Tags, tag)
}

type Article struct {
    Identifiable  // Embedded
    Timestamped   // Embedded
    Taggable      // Embedded
    Title        string
    Content      string
}

func main() {
    article := Article{
        Identifiable: Identifiable{ID: "article-123"},
        Timestamped: Timestamped{
            CreatedAt: time.Now().Add(-24 * time.Hour),
            UpdatedAt: time.Now(),
        },
        Taggable: Taggable{Tags: []string{"tech", "golang"}},
        Title:    "Understanding Embedding",
        Content:  "...",
    }
    
    // All promoted
    fmt.Println(article.GetID())    // article-123
    fmt.Println(article.Age())      // ~24h
    article.AddTag("tutorial")
    fmt.Println(article.Tags)       // [tech golang tutorial]
}
```

### Practical Use Case: Mixins

```go
// Mixin for soft delete
type SoftDeletable struct {
    DeletedAt *time.Time
}

func (s *SoftDeletable) Delete() {
    now := time.Now()
    s.DeletedAt = &now
}

func (s SoftDeletable) IsDeleted() bool {
    return s.DeletedAt != nil
}

// Mixin for auditing
type Auditable struct {
    CreatedBy string
    UpdatedBy string
}

// Models using mixins
type User struct {
    SoftDeletable  // Mixin
    Auditable      // Mixin
    Username      string
    Email         string
}

type Post struct {
    SoftDeletable  // Mixin
    Auditable      // Mixin
    Title         string
    Content       string
}

func main() {
    user := User{
        Username: "alice",
        Email:    "alice@example.com",
        Auditable: Auditable{
            CreatedBy: "system",
            UpdatedBy: "admin",
        },
    }
    
    // Use mixin methods
    user.Delete()
    fmt.Println(user.IsDeleted())  // true
}
```

---

## Embedding Interfaces

You can embed interfaces in structs and other interfaces to create powerful abstractions.

### Embedding Interface in Struct

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

// Struct embedding interfaces
type ReadWriter struct {
    Reader  // Embedded interface
    Writer  // Embedded interface
}

// Concrete implementations
type StringReader struct {
    data string
}

func (s StringReader) Read() string {
    return s.data
}

type ConsoleWriter struct{}

func (c ConsoleWriter) Write(s string) {
    fmt.Println("Writing:", s)
}

func main() {
    rw := ReadWriter{
        Reader: StringReader{data: "Hello, World!"},
        Writer: ConsoleWriter{},
    }
    
    // Use embedded interface methods
    data := rw.Read()
    rw.Write(data)
}
```

### Embedding Interface in Interface

This is very common in Go's standard library:

```go
// Compose interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// ReadWriter combines Reader and Writer
type ReadWriter interface {
    Reader  // Embedded
    Writer  // Embedded
}

// ReadWriteCloser combines all three
type ReadWriteCloser interface {
    Reader  // Embedded
    Writer  // Embedded
    Closer  // Embedded
}

// Or even simpler:
type ReadWriteCloser2 interface {
    ReadWriter  // Embedded
    Closer      // Embedded
}
```

### Standard Library Examples

```go
// io.ReadWriter
type ReadWriter interface {
    io.Reader  // Embedded
    io.Writer  // Embedded
}

// http.ResponseWriter (simplified)
type ResponseWriter interface {
    io.Writer           // Embedded
    Header() http.Header
    WriteHeader(int)
}

// database/sql.Rows (simplified)
type Rows interface {
    io.Closer  // Embedded
    Next() bool
    Scan(...interface{}) error
}
```

### Interface Embedding Pattern

```go
type Validator interface {
    Validate() error
}

type Saver interface {
    Save() error
}

type Loader interface {
    Load() error
}

// Composed interface
type Model interface {
    Validator  // Embedded
    Saver      // Embedded
    Loader     // Embedded
}

// Implementation
type User struct {
    ID   int
    Name string
}

func (u User) Validate() error {
    if u.Name == "" {
        return fmt.Errorf("name is required")
    }
    return nil
}

func (u User) Save() error {
    fmt.Println("Saving user:", u.Name)
    return nil
}

func (u User) Load() error {
    fmt.Println("Loading user")
    return nil
}

// User automatically satisfies Model interface
func processModel(m Model) error {
    if err := m.Validate(); err != nil {
        return err
    }
    if err := m.Load(); err != nil {
        return err
    }
    return m.Save()
}
```

---

## Method Overriding

In Go, you can "override" methods by defining a method with the same name on the outer struct.

### Basic Override

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return "Some generic sound"
}

func (a Animal) Eat() string {
    return a.Name + " is eating"
}

type Dog struct {
    Animal  // Embedded
    Breed  string
}

// Override Speak method
func (d Dog) Speak() string {
    return "Woof! Woof!"
}

// Eat is not overridden, so promoted method is used

func main() {
    dog := Dog{
        Animal: Animal{Name: "Buddy"},
        Breed:  "Golden Retriever",
    }
    
    fmt.Println(dog.Speak())  // Woof! Woof! (overridden)
    fmt.Println(dog.Eat())    // Buddy is eating (promoted)
    
    // Can still call original method
    fmt.Println(dog.Animal.Speak())  // Some generic sound
}
```

### Overriding with Extension

```go
type Base struct {
    Value int
}

func (b Base) Process() string {
    return fmt.Sprintf("Processing %d", b.Value)
}

type Enhanced struct {
    Base      // Embedded
    Multiplier int
}

// Override and extend
func (e Enhanced) Process() string {
    // Call base method
    baseResult := e.Base.Process()
    
    // Add enhancement
    enhanced := e.Value * e.Multiplier
    return fmt.Sprintf("%s, Enhanced: %d", baseResult, enhanced)
}

func main() {
    e := Enhanced{
        Base:       Base{Value: 10},
        Multiplier: 3,
    }
    
    fmt.Println(e.Process())
    // Processing 10, Enhanced: 30
}
```

### Selective Override

```go
type Logger struct{}

func (l Logger) Info(msg string) {
    fmt.Printf("[INFO] %s\n", msg)
}

func (l Logger) Error(msg string) {
    fmt.Printf("[ERROR] %s\n", msg)
}

func (l Logger) Debug(msg string) {
    fmt.Printf("[DEBUG] %s\n", msg)
}

type CustomLogger struct {
    Logger  // Embedded
    Prefix string
}

// Override only Error to add prefix
func (c CustomLogger) Error(msg string) {
    fmt.Printf("[ERROR] %s: %s\n", c.Prefix, msg)
}

func main() {
    logger := CustomLogger{
        Prefix: "MyApp",
    }
    
    logger.Info("Starting")           // [INFO] Starting (promoted)
    logger.Error("Something failed")  // [ERROR] MyApp: Something failed (overridden)
    logger.Debug("Debug info")        // [DEBUG] Debug info (promoted)
}
```

---

## Name Conflicts and Resolution

When embedding multiple types, field or method names can conflict. Go has specific rules for resolution.

### Field Name Conflicts

```go
type A struct {
    Name  string
    Value int
}

type B struct {
    Name  string
    Count int
}

type C struct {
    A      // Has Name
    B      // Also has Name - CONFLICT!
    Title string
}

func main() {
    c := C{
        A:     A{Name: "A's name", Value: 10},
        B:     B{Name: "B's name", Count: 5},
        Title: "C's title",
    }
    
    // Direct access causes ambiguity
    // fmt.Println(c.Name)  // ERROR: ambiguous selector c.Name
    
    // Must use qualified access
    fmt.Println(c.A.Name)  // A's name
    fmt.Println(c.B.Name)  // B's name
    
    // Unique fields work fine
    fmt.Println(c.Value)  // 10 (from A)
    fmt.Println(c.Count)  // 5 (from B)
}
```

### Resolving Conflicts with Outer Field

```go
type Base1 struct {
    ID int
}

type Base2 struct {
    ID int
}

type Derived struct {
    Base1  // Has ID
    Base2  // Has ID
    ID int // Outer field takes precedence
}

func main() {
    d := Derived{
        Base1: Base1{ID: 1},
        Base2: Base2{ID: 2},
        ID:    3,
    }
    
    fmt.Println(d.ID)        // 3 (outer field wins)
    fmt.Println(d.Base1.ID)  // 1
    fmt.Println(d.Base2.ID)  // 2
}
```

### Method Conflicts

```go
type Talker struct{}

func (t Talker) Speak() string {
    return "Talker speaking"
}

type Speaker struct{}

func (s Speaker) Speak() string {
    return "Speaker speaking"
}

type Person struct {
    Talker   // Has Speak
    Speaker  // Has Speak - CONFLICT!
}

// Resolve by defining method on Person
func (p Person) Speak() string {
    // Can choose which to call, or both
    return p.Talker.Speak() + " and " + p.Speaker.Speak()
}

func main() {
    p := Person{}
    
    fmt.Println(p.Speak())
    // Talker speaking and Speaker speaking
}
```

### Conflict Resolution Rules

```go
// Rule 1: Outer fields/methods shadow inner ones
type Inner struct {
    Value int
}

type Outer struct {
    Inner
    Value int  // Shadows Inner.Value
}

// Rule 2: Same level conflicts must be qualified
type A struct{ X int }
type B struct{ X int }
type C struct {
    A
    B
    // c.X is ambiguous, must use c.A.X or c.B.X
}

// Rule 3: Deeper nesting is shadowed by shallower
type Level1 struct {
    Value int
}

type Level2 struct {
    Level1
    Value int  // Shadows Level1.Value
}

type Level3 struct {
    Level2
    // Level3.Value refers to Level2.Value
    // Level1.Value is accessible as Level3.Level2.Level1.Value
}
```

---

## Embedded Pointers

You can embed pointer types, which behaves differently than embedding value types.

### Value vs Pointer Embedding

```go
type Config struct {
    Setting string
}

// Value embedding
type Service1 struct {
    Config  // Value
}

// Pointer embedding
type Service2 struct {
    *Config  // Pointer
}

func main() {
    config := Config{Setting: "enabled"}
    
    // Value embedding - copies the struct
    s1 := Service1{Config: config}
    s1.Setting = "disabled"
    fmt.Println(config.Setting)   // enabled (original unchanged)
    fmt.Println(s1.Setting)       // disabled
    
    // Pointer embedding - shares the struct
    s2 := Service2{Config: &config}
    s2.Setting = "modified"
    fmt.Println(config.Setting)   // modified (original changed!)
    fmt.Println(s2.Setting)       // modified
}
```

### Nil Pointer Embedded Field

```go
type Logger struct {
    Prefix string
}

func (l Logger) Log(msg string) {
    fmt.Printf("%s: %s\n", l.Prefix, msg)
}

type Service struct {
    *Logger  // Pointer embedding
    Name    string
}

func main() {
    // Nil embedded pointer
    s1 := Service{Name: "Service1"}
    
    // This will panic!
    // s1.Log("test")  // panic: runtime error: invalid memory address
    
    // Check before use
    if s1.Logger != nil {
        s1.Log("test")
    } else {
        fmt.Println("Logger not initialized")
    }
    
    // Proper initialization
    s2 := Service{
        Logger: &Logger{Prefix: "MyService"},
        Name:   "Service2",
    }
    s2.Log("test")  // MyService: test
}
```

### When to Use Pointer Embedding

```go
// Shared state between instances
type SharedConfig struct {
    MaxConnections int
    Timeout        time.Duration
}

type Server1 struct {
    *SharedConfig  // Pointer - shares config
    Name          string
}

type Server2 struct {
    *SharedConfig  // Pointer - shares same config
    Name          string
}

func main() {
    config := &SharedConfig{
        MaxConnections: 100,
        Timeout:        30 * time.Second,
    }
    
    server1 := Server1{SharedConfig: config, Name: "Server1"}
    server2 := Server2{SharedConfig: config, Name: "Server2"}
    
    // Modify through server1
    server1.MaxConnections = 200
    
    // server2 sees the change
    fmt.Println(server2.MaxConnections)  // 200
}
```

### Optional Embedded Behavior

```go
type Cache interface {
    Get(key string) (interface{}, bool)
    Set(key string, value interface{})
}

type Service struct {
    Cache  // Can be nil
    Name  string
}

func (s Service) GetData(key string) interface{} {
    // Check if cache is available
    if s.Cache != nil {
        if data, ok := s.Cache.Get(key); ok {
            return data
        }
    }
    
    // Fetch data from source
    data := s.fetchFromSource(key)
    
    // Cache if available
    if s.Cache != nil {
        s.Cache.Set(key, data)
    }
    
    return data
}

func (s Service) fetchFromSource(key string) interface{} {
    return "data from source"
}
```

---

## Common Patterns

### 1. Base Model Pattern

```go
// Base model with common fields
type Model struct {
    ID        uint
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (m *Model) BeforeSave() {
    now := time.Now()
    if m.ID == 0 {
        m.CreatedAt = now
    }
    m.UpdatedAt = now
}

// Domain models embedding base
type User struct {
    Model    // Embedded base
    Username string
    Email    string
}

type Post struct {
    Model    // Embedded base
    Title   string
    Content string
    UserID  uint
}

type Comment struct {
    Model    // Embedded base
    Content string
    UserID  uint
    PostID  uint
}
```

### 2. Decorator Pattern

```go
type HTTPClient interface {
    Do(req *http.Request) (*http.Response, error)
}

// Base client
type BaseClient struct {
    client *http.Client
}

func (b BaseClient) Do(req *http.Request) (*http.Response, error) {
    return b.client.Do(req)
}

// Decorator with logging
type LoggingClient struct {
    HTTPClient  // Embedded interface
}

func (l LoggingClient) Do(req *http.Request) (*http.Response, error) {
    fmt.Printf("Request: %s %s\n", req.Method, req.URL)
    resp, err := l.HTTPClient.Do(req)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Response: %d\n", resp.StatusCode)
    }
    return resp, err
}

// Decorator with retry
type RetryClient struct {
    HTTPClient  // Embedded interface
    MaxRetries int
}

func (r RetryClient) Do(req *http.Request) (*http.Response, error) {
    var resp *http.Response
    var err error
    
    for i := 0; i <= r.MaxRetries; i++ {
        resp, err = r.HTTPClient.Do(req)
        if err == nil {
            return resp, nil
        }
        fmt.Printf("Retry %d/%d\n", i+1, r.MaxRetries)
        time.Sleep(time.Second)
    }
    
    return resp, err
}

// Composition
func main() {
    base := BaseClient{client: &http.Client{}}
    
    // Wrap with logging
    logged := LoggingClient{HTTPClient: base}
    
    // Wrap with retry
    client := RetryClient{
        HTTPClient: logged,
        MaxRetries: 3,
    }
    
    // Use fully decorated client
    req, _ := http.NewRequest("GET", "https://example.com", nil)
    client.Do(req)
}
```

### 3. Mixin Pattern

```go
// Reusable mixins
type Timestamps struct {
    CreatedAt time.Time
    UpdatedAt time.Time
}

type SoftDelete struct {
    DeletedAt *time.Time
}

func (s *SoftDelete) Delete() {
    now := time.Now()
    s.DeletedAt = &now
}

func (s SoftDelete) IsDeleted() bool {
    return s.DeletedAt != nil
}

type Taggable struct {
    Tags []string
}

func (t *Taggable) AddTag(tag string) {
    t.Tags = append(t.Tags, tag)
}

// Models composing mixins
type Article struct {
    Timestamps  // Mixin
    SoftDelete  // Mixin
    Taggable    // Mixin
    Title      string
    Content    string
}

type Product struct {
    Timestamps  // Mixin
    SoftDelete  // Mixin
    Taggable    // Mixin
    Name       string
    Price      float64
}
```

### 4. Interface Segregation

```go
// Small, focused interfaces
type IDProvider interface {
    GetID() string
}

type NameProvider interface {
    GetName() string
}

type Saver interface {
    Save() error
}

type Deleter interface {
    Delete() error
}

// Compose as needed
type Repository interface {
    IDProvider
    Saver
    Deleter
}

type ReadOnlyRepository interface {
    IDProvider
    NameProvider
}
```

### 5. Builder with Embedding

```go
type HTTPRequest struct {
    Method  string
    URL     string
    Headers map[string]string
    Body    string
}

type RequestBuilder struct {
    *HTTPRequest
}

func NewRequestBuilder() *RequestBuilder {
    return &RequestBuilder{
        HTTPRequest: &HTTPRequest{
            Headers: make(map[string]string),
        },
    }
}

func (rb *RequestBuilder) WithMethod(method string) *RequestBuilder {
    rb.Method = method
    return rb
}

func (rb *RequestBuilder) WithURL(url string) *RequestBuilder {
    rb.URL = url
    return rb
}

func (rb *RequestBuilder) WithHeader(key, value string) *RequestBuilder {
    rb.Headers[key] = value
    return rb
}

func (rb *RequestBuilder) WithBody(body string) *RequestBuilder {
    rb.Body = body
    return rb
}

func (rb *RequestBuilder) Build() *HTTPRequest {
    return rb.HTTPRequest
}

func main() {
    req := NewRequestBuilder().
        WithMethod("POST").
        WithURL("https://api.example.com/users").
        WithHeader("Content-Type", "application/json").
        WithBody(`{"name":"Alice"}`).
        Build()
    
    fmt.Printf("%+v\n", req)
}
```

---

## Real-World Examples

### 1. HTTP Handler with Middleware

```go
type Handler interface {
    ServeHTTP(http.ResponseWriter, *http.Request)
}

// Base handler
type BaseHandler struct {
    db     *sql.DB
    logger *log.Logger
}

// Middleware handlers embedding base
type AuthHandler struct {
    BaseHandler  // Embedded
    secret      string
}

func (h AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    token := r.Header.Get("Authorization")
    if token != h.secret {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    h.logger.Println("Request authorized")
    // Continue to next handler
}

type LoggingHandler struct {
    BaseHandler  // Embedded
    Handler      // Next handler
}

func (h LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    h.logger.Printf("Started %s %s", r.Method, r.URL.Path)
    
    h.Handler.ServeHTTP(w, r)
    
    duration := time.Since(start)
    h.logger.Printf("Completed in %v", duration)
}
```

### 2. Database Model with CRUD

```go
// Base model
type BaseModel struct {
    ID        int
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}

func (m *BaseModel) SetID(id int) {
    m.ID = id
}

func (m *BaseModel) Touch() {
    m.UpdatedAt = time.Now()
}

func (m *BaseModel) SoftDelete() {
    now := time.Now()
    m.DeletedAt = &now
}

// Generic repository interface
type Repository interface {
    Create() error
    Update() error
    Delete() error
    Find(id int) error
}

// User model
type User struct {
    BaseModel  // Embedded
    Username  string
    Email     string
    Password  string
}

func (u *User) Create() error {
    u.CreatedAt = time.Now()
    u.UpdatedAt = time.Now()
    // Database insert logic
    fmt.Println("Creating user:", u.Username)
    return nil
}

func (u *User) Update() error {
    u.Touch()
    // Database update logic
    fmt.Println("Updating user:", u.Username)
    return nil
}

func (u *User) Delete() error {
    u.SoftDelete()
    // Database soft delete logic
    fmt.Println("Soft deleting user:", u.Username)
    return nil
}

func (u *User) Find(id int) error {
    // Database query logic
    fmt.Println("Finding user with ID:", id)
    return nil
}

// Product model with same pattern
type Product struct {
    BaseModel  // Embedded
    Name      string
    Price     float64
    Stock     int
}

func (p *Product) Create() error {
    p.CreatedAt = time.Now()
    p.UpdatedAt = time.Now()
    fmt.Println("Creating product:", p.Name)
    return nil
}

// ... similar CRUD methods
```

### 3. Event System with Embedding

```go
// Base event
type BaseEvent struct {
    ID        string
    Timestamp time.Time
    Source    string
}

func (e BaseEvent) GetID() string {
    return e.ID
}

func (e BaseEvent) GetTimestamp() time.Time {
    return e.Timestamp
}

// Specific events
type UserCreatedEvent struct {
    BaseEvent  // Embedded
    UserID    int
    Username  string
    Email     string
}

type OrderPlacedEvent struct {
    BaseEvent  // Embedded
    OrderID   int
    UserID    int
    Total     float64
}

type PaymentProcessedEvent struct {
    BaseEvent     // Embedded
    PaymentID    int
    OrderID      int
    Amount       float64
    Status       string
}

// Event handler
type EventHandler interface {
    Handle(event interface{}) error
}

// Handlers can access base event properties
type LoggingEventHandler struct{}

func (h LoggingEventHandler) Handle(event interface{}) error {
    switch e := event.(type) {
    case UserCreatedEvent:
        fmt.Printf("[%s] User created: %s (ID: %s)\n",
            e.GetTimestamp().Format(time.RFC3339),
            e.Username,
            e.GetID())
    case OrderPlacedEvent:
        fmt.Printf("[%s] Order placed: $%.2f (ID: %s)\n",
            e.GetTimestamp().Format(time.RFC3339),
            e.Total,
            e.GetID())
    }
    return nil
}
```

### 4. Plugin System

```go
// Plugin interface
type Plugin interface {
    Name() string
    Version() string
    Initialize() error
    Shutdown() error
}

// Base plugin implementation
type BasePlugin struct {
    name    string
    version string
}

func (p BasePlugin) Name() string {
    return p.name
}

func (p BasePlugin) Version() string {
    return p.version
}

func (p BasePlugin) Initialize() error {
    fmt.Printf("Initializing %s v%s\n", p.name, p.version)
    return nil
}

func (p BasePlugin) Shutdown() error {
    fmt.Printf("Shutting down %s\n", p.name)
    return nil
}

// Specific plugins
type LoggingPlugin struct {
    BasePlugin  // Embedded
    logFile    string
}

func NewLoggingPlugin() *LoggingPlugin {
    return &LoggingPlugin{
        BasePlugin: BasePlugin{
            name:    "Logger",
            version: "1.0.0",
        },
        logFile: "/var/log/app.log",
    }
}

func (p LoggingPlugin) Initialize() error {
    // Call base initialization
    if err := p.BasePlugin.Initialize(); err != nil {
        return err
    }
    
    // Plugin-specific initialization
    fmt.Printf("Opening log file: %s\n", p.logFile)
    return nil
}

type CachePlugin struct {
    BasePlugin  // Embedded
    cacheSize  int
}

func NewCachePlugin() *CachePlugin {
    return &CachePlugin{
        BasePlugin: BasePlugin{
            name:    "Cache",
            version: "2.0.0",
        },
        cacheSize: 1000,
    }
}

// Plugin manager
type PluginManager struct {
    plugins []Plugin
}

func (pm *PluginManager) Register(p Plugin) {
    pm.plugins = append(pm.plugins, p)
}

func (pm *PluginManager) InitializeAll() error {
    for _, p := range pm.plugins {
        if err := p.Initialize(); err != nil {
            return err
        }
    }
    return nil
}

func (pm *PluginManager) ShutdownAll() error {
    for _, p := range pm.plugins {
        if err := p.Shutdown(); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    manager := &PluginManager{}
    
    manager.Register(NewLoggingPlugin())
    manager.Register(NewCachePlugin())
    
    manager.InitializeAll()
    // ... use plugins ...
    manager.ShutdownAll()
}
```

---

## Best Practices

### 1. Prefer Composition Over Deep Nesting

```go
// Avoid: Deep nesting
type A struct {
    Value int
}

type B struct {
    A
    Extra string
}

type C struct {
    B
    More string
}

type D struct {
    C
    Another string
}

// Prefer: Flat composition
type Component1 struct {
    Value int
}

type Component2 struct {
    Extra string
}

type Entity struct {
    Component1  // Flat
    Component2  // Flat
    More       string
    Another    string
}
```

### 2. Use Embedding for "Is-Like-A" Relationships

```go
// Good: Shape behaviors
type Shape struct {
    X, Y float64
}

func (s *Shape) Move(dx, dy float64) {
    s.X += dx
    s.Y += dy
}

type Circle struct {
    Shape   // Circle is-like-a Shape
    Radius float64
}

type Rectangle struct {
    Shape   // Rectangle is-like-a Shape
    Width  float64
    Height float64
}
```

### 3. Be Explicit About Conflicts

```go
// When embedding multiple types with same methods,
// be explicit about which to use

type Printer struct{}
func (p Printer) Print() { fmt.Println("Printer") }

type Scanner struct{}
func (s Scanner) Print() { fmt.Println("Scanner") }

type MultiFunctionDevice struct {
    Printer
    Scanner
}

// Resolve ambiguity explicitly
func (m MultiFunctionDevice) Print() {
    fmt.Println("MultiFunction Device:")
    m.Printer.Print()
    m.Scanner.Print()
}
```

### 4. Use Pointer Embedding for Shared State

```go
// Use pointer embedding when you want shared state
type SharedConfig struct {
    MaxUsers int
}

type Service1 struct {
    *SharedConfig  // Pointer - shares state
}

type Service2 struct {
    *SharedConfig  // Pointer - shares state
}

// Use value embedding when you want independence
type Service3 struct {
    SharedConfig  // Value - independent copy
}
```

### 5. Document Embedded Types

```go
// User represents a system user with authentication and profile data.
// It embeds BaseModel for common database fields (ID, timestamps)
// and embeds Timestamps for automatic time tracking.
type User struct {
    BaseModel  // Provides ID, CreatedAt, UpdatedAt, DeletedAt
    Timestamps // Automatic timestamp management
    
    Username string
    Email    string
    password string
}
```

### 6. Don't Overuse Embedding

```go
// Avoid: Unnecessary embedding
type Name string

type Person struct {
    Name  // Embedded - but why?
    Age int
}

// Prefer: Simple field
type Person struct {
    Name string  // Clear and simple
    Age  int
}

// Use embedding when you need promoted methods
type Name struct {
    First string
    Last  string
}

func (n Name) Full() string {
    return n.First + " " + n.Last
}

type Person struct {
    Name  // Embedded - gives access to Full() method
    Age  int
}
```

### 7. Consider Interface Embedding

```go
// Good: Compose interfaces
type Reader interface {
    Read() ([]byte, error)
}

type Writer interface {
    Write([]byte) error
}

type ReadWriter interface {
    Reader  // Composed
    Writer  // Composed
}

// Better than:
type ReadWriter2 interface {
    Read() ([]byte, error)
    Write([]byte) error
}
// Because changes to Reader/Writer automatically propagate
```

---

## Embedding vs Other Approaches

### Embedding vs Regular Fields

```go
// Regular field - explicit access
type Car1 struct {
    engine Engine
}

func (c Car1) Start() {
    c.engine.Start()  // Must go through field
}

// Embedding - promoted access
type Car2 struct {
    Engine  // Embedded
}

func (c Car2) Start() {
    c.Start()  // Can call directly
}
```

### Embedding vs Interfaces

```go
// Interface - flexible but requires implementation
type Storage interface {
    Save(data []byte) error
    Load() ([]byte, error)
}

type Service1 struct {
    storage Storage  // Any implementation
}

// Embedding - concrete type
type FileStorage struct {
    path string
}

func (f FileStorage) Save(data []byte) error { return nil }
func (f FileStorage) Load() ([]byte, error) { return nil, nil }

type Service2 struct {
    FileStorage  // Specific implementation
}
```

### Embedding vs Wrapper Functions

```go
// Wrapper functions - explicit delegation
type Service1 struct {
    logger *Logger
}

func (s Service1) Log(msg string) {
    s.logger.Log(msg)
}

func (s Service1) Error(msg string) {
    s.logger.Error(msg)
}

// Embedding - automatic delegation
type Service2 struct {
    *Logger  // All methods promoted
}
// No wrapper methods needed!
```

### When to Use Each

| Use Case | Approach | Reason |
|----------|----------|--------|
| Need interface flexibility | Interface field | Can swap implementations |
| Reusing common behavior | Embedding | Auto-delegation of methods |
| "Has-a" relationship | Regular field | Explicit ownership |
| "Is-like-a" relationship | Embedding | Behavioral similarity |
| Multiple implementations | Interface | Polymorphism |
| Single concrete type | Embedding | Simplicity |
| Optional behavior | Pointer field/interface | Can be nil |
| Shared state | Pointer embedding | Multiple refs to same data |

---

## Summary

**Key Takeaways:**

1. **Composition Over Inheritance**: Go uses embedding for composition, not inheritance

2. **Promotion**: Embedded fields and methods are promoted to the outer struct

3. **Anonymous vs Named**: 
   - Anonymous embedding → promotion
   - Named fields → explicit access

4. **Multiple Embedding**: Can embed multiple types; all are promoted

5. **Interface Embedding**: Powerful for creating composed interfaces

6. **Method Override**: Outer methods shadow embedded ones

7. **Conflict Resolution**:
   - Same level → ambiguous, must qualify
   - Outer level → shadows inner

8. **Pointer Embedding**: Use for shared state and optional components

9. **Common Patterns**:
   - Base model pattern
   - Decorator pattern
   - Mixin pattern
   - Interface composition

10. **Best Practices**:
    - Prefer flat composition over deep nesting
    - Be explicit about conflicts
    - Document embedded types
    - Don't overuse embedding
    - Consider interface embedding for flexibility

**When to Use Embedding:**
- ✅ Reusing common behavior across types
- ✅ Building type hierarchies without inheritance
- ✅ Implementing interfaces by delegation
- ✅ Creating wrapper/decorator types
- ✅ Composing small, focused interfaces
- ❌ When simple fields would be clearer
- ❌ When you need interface flexibility
- ❌ When relationships are purely structural

Struct embedding is a powerful feature that enables flexible, maintainable code design. Use it wisely to build well-composed systems!

---

## Further Reading

- [Effective Go: Embedding](https://go.dev/doc/effective_go#embedding)
- [Go Spec: Struct Types](https://go.dev/ref/spec#Struct_types)
- [Go Blog: Embedding in Go](https://go.dev/blog/embedding)
- [Go by Example: Struct Embedding](https://gobyexample.com/struct-embedding)

## Practice Exercises

Try implementing these to master struct embedding:

1. **Shape System**: Create a shape system with common behaviors (Move, Rotate) using embedding
2. **Logger Decorators**: Build a logger with multiple decorators (timestamp, level, file) using embedding
3. **Model Base Class**: Create a base model with CRUD operations, then build specific models
4. **Plugin Architecture**: Implement a plugin system where plugins share common lifecycle methods
5. **Middleware Chain**: Build an HTTP middleware chain using embedding and interfaces
