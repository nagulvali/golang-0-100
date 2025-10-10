# Go Structs

## Table of Contents
1. [Introduction](#introduction)
2. [What is a Struct?](#what-is-a-struct)
3. [Defining Structs](#defining-structs)
4. [Creating Struct Instances](#creating-struct-instances)
5. [Accessing and Modifying Fields](#accessing-and-modifying-fields)
6. [Methods and Receivers](#methods-and-receivers)
7. [Pointer vs Value Receivers](#pointer-vs-value-receivers)
8. [Constructor Functions](#constructor-functions)
9. [Embedded Structs (Composition)](#embedded-structs-composition)
10. [Struct Tags](#struct-tags)
11. [Anonymous Structs](#anonymous-structs)
12. [Struct Comparison and Zero Values](#struct-comparison-and-zero-values)
13. [Visibility and Exported Fields](#visibility-and-exported-fields)
14. [Common Patterns](#common-patterns)
15. [Best Practices](#best-practices)
16. [Real-World Examples](#real-world-examples)

---

## Introduction

Structs are one of the fundamental building blocks of Go programming. They provide a way to create custom data types by grouping together related data fields. Unlike many other languages, Go doesn't have classes, but structs combined with methods provide similar object-oriented programming capabilities.

Structs in Go are:
- Type-safe composite data types
- Value types (passed by value by default)
- The primary way to organize and encapsulate data
- The foundation for implementing OOP principles in Go

---

## What is a Struct?

A **struct** (short for structure) is a user-defined composite data type that groups together zero or more values of different types under a single name. Each value in a struct is called a **field**.

**Key Characteristics:**
- **Composite Type**: Can contain multiple fields of different types
- **Value Type**: Structs are copied when assigned or passed to functions
- **Type Safety**: Fields are strongly typed
- **Zero Value**: All fields are initialized to their zero values
- **No Inheritance**: Go uses composition over inheritance

**Think of structs as:**
- Blueprints for creating objects
- Records or entities in your application
- A way to model real-world concepts

---

## Defining Structs

### Basic Syntax

```go
type StructName struct {
    field1 type1
    field2 type2
    field3 type3
}
```

### Simple Example

```go
type Person struct {
    Name string
    Age  int
}
```

### Example with Multiple Types

```go
type Book struct {
    Title       string
    Author      string
    Pages       int
    Price       float64
    ISBN        string
    IsAvailable bool
}
```

### Example with Various Field Types

```go
type User struct {
    ID        int
    Username  string
    Email     string
    IsActive  bool
    Tags      []string           // Slice field
    Metadata  map[string]string  // Map field
    CreatedAt time.Time          // Time field
}
```

### Grouping Fields of Same Type

You can group fields of the same type on one line:

```go
type Rectangle struct {
    Width, Height float64
    X, Y          int
}
```

---

## Creating Struct Instances

Go provides multiple ways to create and initialize struct instances.

### 1. Zero Value Initialization

```go
var p Person
// p.Name = "" (zero value for string)
// p.Age = 0 (zero value for int)
```

### 2. Struct Literal (With Field Names)

**Recommended approach** - explicit and clear:

```go
p := Person{
    Name: "Alice",
    Age:  30,
}
```

### 3. Struct Literal (Without Field Names)

Must provide values in order (not recommended for maintainability):

```go
p := Person{"Bob", 25}
```

### 4. Partial Initialization

Unspecified fields get zero values:

```go
p := Person{
    Name: "Charlie",
    // Age will be 0
}
```

### 5. Using `new` Keyword

Returns a pointer to the struct:

```go
p := new(Person)
// p is *Person, all fields are zero values
p.Name = "David"
p.Age = 35
```

### 6. Pointer to Struct Literal

```go
p := &Person{
    Name: "Eve",
    Age:  28,
}
// p is *Person
```

### Example: Multiple Initialization Methods

```go
func main() {
    // Method 1: Zero value
    var person1 Person
    fmt.Printf("%+v\n", person1)  // {Name: Age:0}
    
    // Method 2: Named fields
    person2 := Person{
        Name: "Alice",
        Age:  30,
    }
    fmt.Printf("%+v\n", person2)  // {Name:Alice Age:30}
    
    // Method 3: Positional (not recommended)
    person3 := Person{"Bob", 25}
    fmt.Printf("%+v\n", person3)  // {Name:Bob Age:25}
    
    // Method 4: Pointer with new
    person4 := new(Person)
    person4.Name = "Charlie"
    person4.Age = 35
    fmt.Printf("%+v\n", person4)  // &{Name:Charlie Age:35}
    
    // Method 5: Pointer to literal
    person5 := &Person{Name: "David", Age: 40}
    fmt.Printf("%+v\n", person5)  // &{Name:David Age:40}
}
```

---

## Accessing and Modifying Fields

### Accessing Fields

Use dot notation to access struct fields:

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    
    // Access fields
    fmt.Println(p.Name)  // Alice
    fmt.Println(p.Age)   // 30
}
```

### Modifying Fields

```go
func main() {
    p := Person{Name: "Bob", Age: 25}
    
    // Modify fields
    p.Name = "Robert"
    p.Age = 26
    
    fmt.Println(p)  // {Robert 26}
}
```

### With Pointers

Go automatically dereferences pointers to structs:

```go
func main() {
    p := &Person{Name: "Charlie", Age: 35}
    
    // No need to write (*p).Name
    p.Name = "Charles"
    p.Age = 36
    
    fmt.Println(p)  // &{Charles 36}
}
```

### Nested Struct Access

```go
type Address struct {
    Street string
    City   string
}

type Person struct {
    Name    string
    Address Address
}

func main() {
    p := Person{
        Name: "Alice",
        Address: Address{
            Street: "123 Main St",
            City:   "New York",
        },
    }
    
    // Access nested fields
    fmt.Println(p.Address.City)      // New York
    
    // Modify nested fields
    p.Address.Street = "456 Oak Ave"
    fmt.Println(p.Address.Street)    // 456 Oak Ave
}
```

---

## Methods and Receivers

Methods are functions with a special receiver argument. The receiver appears between the `func` keyword and the method name.

### Basic Method Syntax

```go
func (receiver ReceiverType) MethodName(parameters) returnType {
    // method body
}
```

### Example: Value Receiver Methods

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with value receiver
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    
    fmt.Println(rect.Area())       // 50
    fmt.Println(rect.Perimeter())  // 30
}
```

### Example: Pointer Receiver Methods

```go
type Counter struct {
    Count int
}

// Method with pointer receiver (can modify the receiver)
func (c *Counter) Increment() {
    c.Count++
}

// Method with pointer receiver
func (c *Counter) Decrement() {
    c.Count--
}

// Method with value receiver (read-only)
func (c Counter) GetCount() int {
    return c.Count
}

func main() {
    counter := Counter{Count: 0}
    
    counter.Increment()
    counter.Increment()
    fmt.Println(counter.GetCount())  // 2
    
    counter.Decrement()
    fmt.Println(counter.GetCount())  // 1
}
```

### Methods vs Functions

```go
// Function
func CalculateArea(r Rectangle) float64 {
    return r.Width * r.Height
}

// Method
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    
    // Function call
    area1 := CalculateArea(rect)
    
    // Method call (more object-oriented)
    area2 := rect.Area()
    
    fmt.Println(area1, area2)  // 50 50
}
```

---

## Pointer vs Value Receivers

Understanding when to use pointer vs value receivers is crucial for writing efficient Go code.

### Value Receiver

**Use when:**
- The method doesn't need to modify the receiver
- The struct is small (a few fields)
- The struct doesn't contain fields that shouldn't be copied (like sync.Mutex)

```go
type Point struct {
    X, Y int
}

// Value receiver - doesn't modify original
func (p Point) Move(dx, dy int) Point {
    p.X += dx
    p.Y += dy
    return p  // Returns new Point
}

func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func main() {
    p1 := Point{X: 1, Y: 2}
    p2 := p1.Move(3, 4)
    
    fmt.Println(p1)  // (1, 2) - unchanged
    fmt.Println(p2)  // (4, 6) - new point
}
```

### Pointer Receiver

**Use when:**
- The method needs to modify the receiver
- The struct is large (copying would be expensive)
- Consistency (if some methods need pointers, use pointers for all)

```go
type BankAccount struct {
    Balance float64
}

// Pointer receiver - modifies original
func (b *BankAccount) Deposit(amount float64) {
    b.Balance += amount
}

// Pointer receiver - modifies original
func (b *BankAccount) Withdraw(amount float64) error {
    if b.Balance < amount {
        return fmt.Errorf("insufficient funds")
    }
    b.Balance -= amount
    return nil
}

// Pointer receiver for consistency
func (b *BankAccount) GetBalance() float64 {
    return b.Balance
}

func main() {
    account := BankAccount{Balance: 100}
    
    account.Deposit(50)
    fmt.Println(account.GetBalance())  // 150
    
    account.Withdraw(30)
    fmt.Println(account.GetBalance())  // 120
}
```

### Comparison Table

| Aspect | Value Receiver | Pointer Receiver |
|--------|---------------|------------------|
| **Modification** | Cannot modify original | Can modify original |
| **Memory** | Creates copy | Uses original |
| **Performance** | Slower for large structs | Faster for large structs |
| **Nil Safety** | Cannot be nil | Can be nil (need to check) |
| **Use Case** | Read-only operations | Modifying operations |
| **Concurrency** | Safe to copy | Need synchronization |

### Decision Guide

```go
type SmallStruct struct {
    Value int
}

type LargeStruct struct {
    Data [1000000]int
}

// Value receiver for small struct (read-only)
func (s SmallStruct) GetValue() int {
    return s.Value
}

// Pointer receiver for small struct (modify)
func (s *SmallStruct) SetValue(v int) {
    s.Value = v
}

// Pointer receiver for large struct (avoid copying)
func (l *LargeStruct) Process() {
    // Even if not modifying, use pointer to avoid copying
}
```

---

## Constructor Functions

Go doesn't have constructors like traditional OOP languages. Instead, we use constructor functions (factory functions) by convention.

### Basic Constructor

```go
type Person struct {
    Name string
    Age  int
}

// Constructor function (naming convention: NewTypeName)
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

func main() {
    person := NewPerson("Alice", 30)
    fmt.Printf("%+v\n", person)  // &{Name:Alice Age:30}
}
```

### Constructor with Validation

```go
type User struct {
    Username string
    Email    string
    Age      int
}

func NewUser(username, email string, age int) (*User, error) {
    if username == "" {
        return nil, fmt.Errorf("username cannot be empty")
    }
    if age < 0 {
        return nil, fmt.Errorf("age cannot be negative")
    }
    if !strings.Contains(email, "@") {
        return nil, fmt.Errorf("invalid email")
    }
    
    return &User{
        Username: username,
        Email:    email,
        Age:      age,
    }, nil
}

func main() {
    user, err := NewUser("alice", "alice@example.com", 25)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("%+v\n", user)
}
```

### Constructor with Default Values

```go
type Config struct {
    Host    string
    Port    int
    Timeout time.Duration
}

func NewConfig(host string) *Config {
    return &Config{
        Host:    host,
        Port:    8080,           // Default value
        Timeout: 30 * time.Second, // Default value
    }
}

func main() {
    config := NewConfig("localhost")
    fmt.Printf("%+v\n", config)
    // &{Host:localhost Port:8080 Timeout:30s}
}
```

### Constructor with Options Pattern

```go
type Server struct {
    Host    string
    Port    int
    Timeout time.Duration
    MaxConn int
}

// Option function type
type ServerOption func(*Server)

// Option functions
func WithPort(port int) ServerOption {
    return func(s *Server) {
        s.Port = port
    }
}

func WithTimeout(timeout time.Duration) ServerOption {
    return func(s *Server) {
        s.Timeout = timeout
    }
}

func WithMaxConn(maxConn int) ServerOption {
    return func(s *Server) {
        s.MaxConn = maxConn
    }
}

// Constructor with options
func NewServer(host string, opts ...ServerOption) *Server {
    // Default values
    server := &Server{
        Host:    host,
        Port:    8080,
        Timeout: 30 * time.Second,
        MaxConn: 100,
    }
    
    // Apply options
    for _, opt := range opts {
        opt(server)
    }
    
    return server
}

func main() {
    // Use defaults
    server1 := NewServer("localhost")
    fmt.Printf("%+v\n", server1)
    
    // Customize with options
    server2 := NewServer(
        "example.com",
        WithPort(9000),
        WithTimeout(60*time.Second),
        WithMaxConn(200),
    )
    fmt.Printf("%+v\n", server2)
}
```

---

## Embedded Structs (Composition)

Go uses composition over inheritance. You can embed one struct inside another to reuse fields and methods.

### Basic Embedding

```go
type Address struct {
    Street string
    City   string
    State  string
}

type Person struct {
    Name    string
    Age     int
    Address Address  // Embedded struct
}

func main() {
    person := Person{
        Name: "Alice",
        Age:  30,
        Address: Address{
            Street: "123 Main St",
            City:   "New York",
            State:  "NY",
        },
    }
    
    // Access nested fields
    fmt.Println(person.Address.City)  // New York
}
```

### Anonymous Embedding (Promotion)

When you embed a struct without a field name, its fields are "promoted":

```go
type Address struct {
    Street string
    City   string
}

type Person struct {
    Name string
    Age  int
    Address  // Anonymous embedding (no field name)
}

func main() {
    person := Person{
        Name: "Bob",
        Age:  25,
        Address: Address{
            Street: "456 Oak Ave",
            City:   "Boston",
        },
    }
    
    // Can access embedded fields directly
    fmt.Println(person.City)          // Boston (promoted)
    fmt.Println(person.Address.City)  // Boston (also works)
}
```

### Method Promotion

Methods from embedded structs are promoted:

```go
type Engine struct {
    Horsepower int
}

func (e Engine) Start() {
    fmt.Println("Engine started")
}

func (e Engine) Stop() {
    fmt.Println("Engine stopped")
}

type Car struct {
    Brand string
    Model string
    Engine  // Anonymous embedding
}

func main() {
    car := Car{
        Brand: "Toyota",
        Model: "Camry",
        Engine: Engine{Horsepower: 200},
    }
    
    // Call promoted methods
    car.Start()  // Engine started
    car.Stop()   // Engine stopped
}
```

### Multiple Embedding

```go
type Person struct {
    Name string
}

type Employee struct {
    EmployeeID int
}

type Manager struct {
    Person     // Embed Person
    Employee   // Embed Employee
    Department string
}

func main() {
    manager := Manager{
        Person:     Person{Name: "Alice"},
        Employee:   Employee{EmployeeID: 12345},
        Department: "Engineering",
    }
    
    // Access promoted fields
    fmt.Println(manager.Name)       // Alice
    fmt.Println(manager.EmployeeID) // 12345
}
```

### Overriding Methods

You can override methods from embedded structs:

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println("Some generic sound")
}

type Dog struct {
    Animal
}

// Override Speak method
func (d Dog) Speak() {
    fmt.Printf("%s says: Woof!\n", d.Name)
}

func main() {
    dog := Dog{Animal: Animal{Name: "Buddy"}}
    
    dog.Speak()         // Buddy says: Woof! (overridden)
    dog.Animal.Speak()  // Some generic sound (original)
}
```

---

## Struct Tags

Struct tags are metadata attached to struct fields, commonly used for serialization, validation, and ORM mapping.

### Basic Syntax

```go
type StructName struct {
    FieldName type `key:"value" key2:"value2"`
}
```

### JSON Serialization Tags

```go
type User struct {
    ID        int    `json:"id"`
    Username  string `json:"username"`
    Email     string `json:"email"`
    Password  string `json:"-"`              // Omit from JSON
    IsActive  bool   `json:"is_active"`
    CreatedAt time.Time `json:"created_at,omitempty"`  // Omit if empty
}

func main() {
    user := User{
        ID:       1,
        Username: "alice",
        Email:    "alice@example.com",
        Password: "secret",
        IsActive: true,
    }
    
    // Marshal to JSON
    jsonData, _ := json.Marshal(user)
    fmt.Println(string(jsonData))
    // {"id":1,"username":"alice","email":"alice@example.com","is_active":true}
    
    // Note: Password is omitted, created_at is omitted (empty)
}
```

### Multiple Tags

```go
type Product struct {
    ID          int     `json:"id" db:"product_id" validate:"required"`
    Name        string  `json:"name" db:"product_name" validate:"required,min=3"`
    Price       float64 `json:"price" db:"price" validate:"required,min=0"`
    Description string  `json:"description,omitempty" db:"description"`
}
```

### Reading Struct Tags with Reflection

```go
import "reflect"

type User struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
}

func main() {
    u := User{}
    t := reflect.TypeOf(u)
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Field: %s\n", field.Name)
        fmt.Printf("  JSON tag: %s\n", field.Tag.Get("json"))
        fmt.Printf("  Validate tag: %s\n", field.Tag.Get("validate"))
    }
}
```

### Common Tag Uses

```go
// Database ORM (GORM example)
type User struct {
    ID        uint   `gorm:"primaryKey"`
    Username  string `gorm:"uniqueIndex;not null"`
    Email     string `gorm:"uniqueIndex;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}

// Validation
type SignupForm struct {
    Username string `validate:"required,min=3,max=20"`
    Email    string `validate:"required,email"`
    Age      int    `validate:"required,gte=18,lte=100"`
}

// XML serialization
type Book struct {
    XMLName xml.Name `xml:"book"`
    Title   string   `xml:"title"`
    Author  string   `xml:"author"`
    Year    int      `xml:"year,attr"`
}
```

---

## Anonymous Structs

Anonymous structs are structs without a named type, useful for one-time use cases.

### Basic Anonymous Struct

```go
func main() {
    // Define and initialize in one step
    person := struct {
        Name string
        Age  int
    }{
        Name: "Alice",
        Age:  30,
    }
    
    fmt.Printf("%+v\n", person)  // {Name:Alice Age:30}
}
```

### Use Case: Table-Driven Tests

```go
func TestAddition(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"mixed numbers", 5, -3, 2},
        {"with zero", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("add(%d, %d) = %d; want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

### Use Case: JSON Response

```go
func getUser(w http.ResponseWriter, r *http.Request) {
    // Anonymous struct for response
    response := struct {
        Success bool   `json:"success"`
        Message string `json:"message"`
        Data    struct {
            ID       int    `json:"id"`
            Username string `json:"username"`
        } `json:"data"`
    }{
        Success: true,
        Message: "User retrieved successfully",
    }
    
    response.Data.ID = 123
    response.Data.Username = "alice"
    
    json.NewEncoder(w).Encode(response)
}
```

### Use Case: Configuration

```go
func main() {
    config := struct {
        Server struct {
            Host string
            Port int
        }
        Database struct {
            Name string
            User string
        }
    }{
        Server: struct {
            Host string
            Port int
        }{
            Host: "localhost",
            Port: 8080,
        },
        Database: struct {
            Name string
            User string
        }{
            Name: "mydb",
            User: "admin",
        },
    }
    
    fmt.Printf("%+v\n", config)
}
```

---

## Struct Comparison and Zero Values

### Zero Values

Each struct field is initialized to its type's zero value:

```go
type User struct {
    Name     string    // ""
    Age      int       // 0
    IsActive bool      // false
    Balance  float64   // 0.0
    Tags     []string  // nil
    Meta     map[string]string  // nil
}

func main() {
    var user User
    fmt.Printf("%+v\n", user)
    // {Name: Age:0 IsActive:false Balance:0 Tags:[] Meta:map[]}
}
```

### Checking for Zero Value

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) IsZero() bool {
    return p == Person{}
}

func main() {
    var p1 Person
    fmt.Println(p1.IsZero())  // true
    
    p2 := Person{Name: "Alice"}
    fmt.Println(p2.IsZero())  // false
}
```

### Struct Comparison

Structs are comparable if all their fields are comparable:

```go
type Point struct {
    X, Y int
}

func main() {
    p1 := Point{X: 1, Y: 2}
    p2 := Point{X: 1, Y: 2}
    p3 := Point{X: 2, Y: 3}
    
    fmt.Println(p1 == p2)  // true
    fmt.Println(p1 == p3)  // false
}
```

### Non-Comparable Structs

Structs with slices, maps, or functions cannot be compared with `==`:

```go
type User struct {
    Name string
    Tags []string  // slice - not comparable
}

func main() {
    u1 := User{Name: "Alice", Tags: []string{"go", "rust"}}
    u2 := User{Name: "Alice", Tags: []string{"go", "rust"}}
    
    // This will NOT compile:
    // fmt.Println(u1 == u2)  // invalid operation
    
    // Use reflect.DeepEqual instead:
    fmt.Println(reflect.DeepEqual(u1, u2))  // true
}
```

---

## Visibility and Exported Fields

In Go, visibility is determined by capitalization:

### Exported vs Unexported

```go
package models

// Exported struct (can be used outside package)
type User struct {
    // Exported fields (can be accessed outside package)
    ID       int
    Username string
    Email    string
    
    // Unexported field (only accessible within package)
    password string
}

// Exported method
func (u *User) SetPassword(password string) {
    u.password = hashPassword(password)
}

// Unexported method
func (u *User) validatePassword() bool {
    return len(u.password) > 0
}

// Unexported helper function
func hashPassword(password string) string {
    // Hash implementation
    return password  // Simplified
}
```

### Accessing from Another Package

```go
package main

import "myapp/models"

func main() {
    user := models.User{
        ID:       1,
        Username: "alice",
        Email:    "alice@example.com",
        // password: "secret",  // ERROR: unexported field
    }
    
    user.SetPassword("secret")  // OK: use exported method
}
```

---

## Common Patterns

### 1. Builder Pattern

```go
type User struct {
    ID       int
    Username string
    Email    string
    IsActive bool
}

type UserBuilder struct {
    user User
}

func NewUserBuilder() *UserBuilder {
    return &UserBuilder{}
}

func (b *UserBuilder) WithID(id int) *UserBuilder {
    b.user.ID = id
    return b
}

func (b *UserBuilder) WithUsername(username string) *UserBuilder {
    b.user.Username = username
    return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
    b.user.Email = email
    return b
}

func (b *UserBuilder) SetActive(active bool) *UserBuilder {
    b.user.IsActive = active
    return b
}

func (b *UserBuilder) Build() User {
    return b.user
}

func main() {
    user := NewUserBuilder().
        WithID(1).
        WithUsername("alice").
        WithEmail("alice@example.com").
        SetActive(true).
        Build()
    
    fmt.Printf("%+v\n", user)
}
```

### 2. Fluent Interface

```go
type Query struct {
    table  string
    fields []string
    where  string
    limit  int
}

func NewQuery() *Query {
    return &Query{}
}

func (q *Query) From(table string) *Query {
    q.table = table
    return q
}

func (q *Query) Select(fields ...string) *Query {
    q.fields = fields
    return q
}

func (q *Query) Where(condition string) *Query {
    q.where = condition
    return q
}

func (q *Query) Limit(n int) *Query {
    q.limit = n
    return q
}

func (q *Query) Build() string {
    // Build SQL query
    return fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d",
        strings.Join(q.fields, ", "), q.table, q.where, q.limit)
}

func main() {
    query := NewQuery().
        From("users").
        Select("id", "name", "email").
        Where("age > 18").
        Limit(10).
        Build()
    
    fmt.Println(query)
}
```

### 3. Singleton Pattern

```go
type Database struct {
    Connection string
}

var instance *Database
var once sync.Once

func GetDatabase() *Database {
    once.Do(func() {
        instance = &Database{
            Connection: "connected",
        }
        fmt.Println("Database instance created")
    })
    return instance
}

func main() {
    db1 := GetDatabase()
    db2 := GetDatabase()
    
    fmt.Println(db1 == db2)  // true (same instance)
}
```

### 4. State Pattern

```go
type OrderState interface {
    Next(*Order)
    Prev(*Order)
    String() string
}

type Order struct {
    state OrderState
}

type PendingState struct{}
type ProcessingState struct{}
type CompletedState struct{}

func (s *PendingState) Next(o *Order) {
    o.state = &ProcessingState{}
}
func (s *PendingState) Prev(o *Order) {}
func (s *PendingState) String() string { return "Pending" }

func (s *ProcessingState) Next(o *Order) {
    o.state = &CompletedState{}
}
func (s *ProcessingState) Prev(o *Order) {
    o.state = &PendingState{}
}
func (s *ProcessingState) String() string { return "Processing" }

func (s *CompletedState) Next(o *Order) {}
func (s *CompletedState) Prev(o *Order) {
    o.state = &ProcessingState{}
}
func (s *CompletedState) String() string { return "Completed" }

func NewOrder() *Order {
    return &Order{state: &PendingState{}}
}

func (o *Order) Next() {
    o.state.Next(o)
}

func (o *Order) Prev() {
    o.state.Prev(o)
}

func (o *Order) Status() string {
    return o.state.String()
}
```

---

## Best Practices

### 1. Keep Structs Focused

```go
// Good: Single responsibility
type User struct {
    ID       int
    Username string
    Email    string
}

type UserProfile struct {
    Bio       string
    Website   string
    AvatarURL string
}

// Avoid: Too many responsibilities
type User struct {
    ID            int
    Username      string
    Email         string
    Bio           string
    Website       string
    AvatarURL     string
    Orders        []Order
    PaymentInfo   PaymentInfo
    // Too many concerns mixed together
}
```

### 2. Use Pointer Receivers Consistently

```go
// Good: Consistent use of pointer receivers
type User struct {
    Name string
}

func (u *User) SetName(name string) {
    u.Name = name
}

func (u *User) GetName() string {
    return u.Name  // Use pointer for consistency
}
```

### 3. Initialize Structs Properly

```go
// Good: Use constructor for complex initialization
func NewUser(username, email string) (*User, error) {
    if username == "" || email == "" {
        return nil, errors.New("username and email required")
    }
    
    return &User{
        Username:  username,
        Email:     email,
        CreatedAt: time.Now(),
        IsActive:  true,
    }, nil
}

// Avoid: Relying on manual initialization
user := User{}
user.Username = "alice"
user.Email = "alice@example.com"
// Forgot to set CreatedAt and IsActive!
```

### 4. Use Composition Over Complex Hierarchies

```go
// Good: Composition
type Logger interface {
    Log(string)
}

type Service struct {
    logger Logger
    db     Database
}

// Avoid: Deep embedding hierarchies that become confusing
```

### 5. Document Exported Structs

```go
// User represents a system user with authentication credentials.
// All fields except password are exported for JSON serialization.
type User struct {
    // ID is the unique identifier for the user.
    ID int `json:"id"`
    
    // Username is the user's display name.
    Username string `json:"username"`
    
    // Email is the user's email address used for notifications.
    Email string `json:"email"`
    
    // password is stored in hashed form and never exported.
    password string
}
```

### 6. Use Zero Values Wisely

```go
// Good: Zero value is useful
type Config struct {
    Host    string // Empty string is fine default
    Port    int    // 0 is fine default
    Verbose bool   // false is fine default
}

// Consider: When zero value isn't appropriate
type Server struct {
    config *Config  // Use pointer, nil indicates "not configured"
}
```

### 7. Avoid Premature Optimization

```go
// Start simple
type Person struct {
    Name string
    Age  int
}

// Only optimize when needed (e.g., for large structs passed frequently)
func ProcessPerson(p *Person) {  // Use pointer only if necessary
    // ...
}
```

---

## Real-World Examples

### 1. HTTP Server with Dependency Injection

```go
type Server struct {
    router   *http.ServeMux
    db       *sql.DB
    logger   *log.Logger
    config   Config
}

type Config struct {
    Host         string
    Port         int
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
}

func NewServer(db *sql.DB, logger *log.Logger, config Config) *Server {
    s := &Server{
        router: http.NewServeMux(),
        db:     db,
        logger: logger,
        config: config,
    }
    s.routes()
    return s
}

func (s *Server) routes() {
    s.router.HandleFunc("/users", s.handleUsers())
    s.router.HandleFunc("/orders", s.handleOrders())
}

func (s *Server) handleUsers() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Use s.db, s.logger, etc.
        s.logger.Println("Handling users request")
        // ...
    }
}

func (s *Server) Start() error {
    addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
    s.logger.Printf("Starting server on %s\n", addr)
    
    server := &http.Server{
        Addr:         addr,
        Handler:      s.router,
        ReadTimeout:  s.config.ReadTimeout,
        WriteTimeout: s.config.WriteTimeout,
    }
    
    return server.ListenAndServe()
}
```

### 2. Repository Pattern

```go
type User struct {
    ID        int       `db:"id"`
    Username  string    `db:"username"`
    Email     string    `db:"email"`
    CreatedAt time.Time `db:"created_at"`
}

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *User) error {
    query := `
        INSERT INTO users (username, email, created_at)
        VALUES ($1, $2, $3)
        RETURNING id
    `
    return r.db.QueryRow(query, user.Username, user.Email, time.Now()).
        Scan(&user.ID)
}

func (r *UserRepository) FindByID(id int) (*User, error) {
    query := `SELECT id, username, email, created_at FROM users WHERE id = $1`
    
    user := &User{}
    err := r.db.QueryRow(query, id).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.CreatedAt,
    )
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (r *UserRepository) Update(user *User) error {
    query := `
        UPDATE users
        SET username = $1, email = $2
        WHERE id = $3
    `
    _, err := r.db.Exec(query, user.Username, user.Email, user.ID)
    return err
}

func (r *UserRepository) Delete(id int) error {
    query := `DELETE FROM users WHERE id = $1`
    _, err := r.db.Exec(query, id)
    return err
}
```

### 3. Service Layer with Business Logic

```go
type OrderService struct {
    orderRepo   *OrderRepository
    productRepo *ProductRepository
    userRepo    *UserRepository
    logger      *log.Logger
}

type Order struct {
    ID         int
    UserID     int
    Items      []OrderItem
    Total      float64
    Status     OrderStatus
    CreatedAt  time.Time
}

type OrderItem struct {
    ProductID int
    Quantity  int
    Price     float64
}

type OrderStatus string

const (
    OrderPending    OrderStatus = "pending"
    OrderProcessing OrderStatus = "processing"
    OrderCompleted  OrderStatus = "completed"
    OrderCancelled  OrderStatus = "cancelled"
)

func NewOrderService(
    orderRepo *OrderRepository,
    productRepo *ProductRepository,
    userRepo *UserRepository,
    logger *log.Logger,
) *OrderService {
    return &OrderService{
        orderRepo:   orderRepo,
        productRepo: productRepo,
        userRepo:    userRepo,
        logger:      logger,
    }
}

func (s *OrderService) CreateOrder(userID int, items []OrderItem) (*Order, error) {
    // Validate user exists
    user, err := s.userRepo.FindByID(userID)
    if err != nil {
        return nil, fmt.Errorf("user not found: %w", err)
    }
    
    // Validate and calculate total
    var total float64
    for _, item := range items {
        product, err := s.productRepo.FindByID(item.ProductID)
        if err != nil {
            return nil, fmt.Errorf("product not found: %w", err)
        }
        item.Price = product.Price
        total += item.Price * float64(item.Quantity)
    }
    
    // Create order
    order := &Order{
        UserID:    userID,
        Items:     items,
        Total:     total,
        Status:    OrderPending,
        CreatedAt: time.Now(),
    }
    
    if err := s.orderRepo.Create(order); err != nil {
        return nil, fmt.Errorf("failed to create order: %w", err)
    }
    
    s.logger.Printf("Order %d created for user %s\n", order.ID, user.Username)
    return order, nil
}

func (s *OrderService) ProcessOrder(orderID int) error {
    order, err := s.orderRepo.FindByID(orderID)
    if err != nil {
        return err
    }
    
    if order.Status != OrderPending {
        return fmt.Errorf("order not in pending status")
    }
    
    order.Status = OrderProcessing
    return s.orderRepo.Update(order)
}
```

### 4. Event-Driven Architecture

```go
type Event struct {
    Type      string
    Payload   interface{}
    Timestamp time.Time
}

type EventHandler func(Event)

type EventBus struct {
    handlers map[string][]EventHandler
    mu       sync.RWMutex
}

func NewEventBus() *EventBus {
    return &EventBus{
        handlers: make(map[string][]EventHandler),
    }
}

func (eb *EventBus) Subscribe(eventType string, handler EventHandler) {
    eb.mu.Lock()
    defer eb.mu.Unlock()
    eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

func (eb *EventBus) Publish(event Event) {
    eb.mu.RLock()
    handlers := eb.handlers[event.Type]
    eb.mu.RUnlock()
    
    for _, handler := range handlers {
        go handler(event)
    }
}

// Usage
type UserCreatedPayload struct {
    UserID   int
    Username string
    Email    string
}

func main() {
    bus := NewEventBus()
    
    // Subscribe to events
    bus.Subscribe("user.created", func(e Event) {
        payload := e.Payload.(UserCreatedPayload)
        fmt.Printf("Sending welcome email to %s\n", payload.Email)
    })
    
    bus.Subscribe("user.created", func(e Event) {
        payload := e.Payload.(UserCreatedPayload)
        fmt.Printf("Creating default settings for user %d\n", payload.UserID)
    })
    
    // Publish event
    bus.Publish(Event{
        Type: "user.created",
        Payload: UserCreatedPayload{
            UserID:   123,
            Username: "alice",
            Email:    "alice@example.com",
        },
        Timestamp: time.Now(),
    })
}
```

---

## Summary

**Key Takeaways:**

1. **Structs are Fundamental**: They're the primary way to create custom types and organize data in Go

2. **No Classes**: Go uses structs + methods instead of traditional classes

3. **Value vs Pointer**:
   - Value receivers for small, read-only operations
   - Pointer receivers for modifications or large structs

4. **Composition Over Inheritance**: Use embedded structs to build complex types

5. **Constructor Functions**: Use `NewTypeName` pattern for initialization and validation

6. **Struct Tags**: Powerful metadata for serialization, validation, and ORM mapping

7. **Visibility**: Capital letters export structs and fields to other packages

8. **Common Patterns**:
   - Builder pattern for complex initialization
   - Repository pattern for data access
   - Service pattern for business logic
   - Dependency injection for testability

9. **Best Practices**:
   - Keep structs focused on single responsibility
   - Use constructors for complex initialization
   - Document exported structs and fields
   - Consider zero values in design
   - Use composition to build complex behavior

10. **Performance Considerations**:
    - Structs are value types (copied by default)
    - Use pointers for large structs or when mutation is needed
    - Embedded structs have no runtime cost

Structs are the backbone of Go programming. Master them, and you'll be able to design clean, efficient, and maintainable code!

---

## Further Reading

- [Go by Example: Structs](https://gobyexample.com/structs)
- [Effective Go: Data](https://go.dev/doc/effective_go#data)
- [Go Spec: Struct Types](https://go.dev/ref/spec#Struct_types)
- [Go Blog: Go's Declaration Syntax](https://go.dev/blog/declaration-syntax)

## Practice Exercises

Try implementing these to master structs:

1. **Task Manager**: Create a task management system with `Task`, `Project`, and `User` structs
2. **E-commerce System**: Build an order system with `Product`, `Cart`, `Order`, and `Customer` structs
3. **Banking System**: Implement a banking system with `Account`, `Transaction`, and `Customer` structs using proper encapsulation
4. **API Client**: Create an HTTP API client using structs for request/response models
5. **Configuration System**: Build a configuration system with nested structs and validation
