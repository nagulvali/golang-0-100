# Go Maps

Maps in Go are a built-in data structure that provides an unordered collection of key-value pairs. They are similar to dictionaries in Python, hash tables in other languages, or objects in JavaScript. Maps provide fast lookups, insertions, and deletions based on keys.

## What You'll Learn

- Map declaration and initialization
- Adding, updating, and deleting key-value pairs
- Map operations and methods
- Map iteration and range loops
- Map types and key constraints
- Nested maps and complex data structures
- Map performance characteristics
- Best practices and common patterns
- Common mistakes and troubleshooting

## Key Concepts

- **Map**: An unordered collection of key-value pairs
- **Key**: The identifier used to access values (must be comparable)
- **Value**: The data associated with a key (can be any type)
- **Hash Table**: The underlying data structure for fast lookups
- **Zero Value**: Maps have a zero value of `nil`

## Basic Map Declaration

### Syntax

```go
var <name> map[<key-type>]<value-type>
<name> := make(map[<key-type>]<value-type>)
<name> := map[<key-type>]<value-type>{}
```

### Examples

```go
// Declaration
var ages map[string]int

// Initialization with make
ages = make(map[string]int)

// Short declaration with make
ages := make(map[string]int)

// Literal initialization
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}
```

## Different Declaration Methods

### 1. Zero Value Declaration

```go
var ages map[string]int
// ages is nil, cannot be used until initialized
```

**Characteristics:**
- Map is `nil`
- Cannot add key-value pairs
- Must initialize with `make()` before use

### 2. Make Function

```go
ages := make(map[string]int)
// ages is an empty map, ready to use
```

**Characteristics:**
- Creates an empty map
- Ready for immediate use
- Can add key-value pairs

### 3. Literal Initialization

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}
```

**Characteristics:**
- Initializes with predefined values
- Most common for known data
- Clean and readable

### 4. Empty Literal

```go
ages := map[string]int{}
// Equivalent to make(map[string]int)
```

**Characteristics:**
- Creates an empty map
- Alternative to `make()`
- Shorter syntax

## Map Operations

### Adding and Updating Values

```go
ages := make(map[string]int)

// Add new key-value pairs
ages["Alice"] = 25
ages["Bob"] = 30
ages["Carol"] = 35

// Update existing values
ages["Alice"] = 26
ages["Bob"] = 31
```

### Accessing Values

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}

// Access values
aliceAge := ages["Alice"]  // 25
bobAge := ages["Bob"]      // 30

// Access non-existent key
daveAge := ages["Dave"]    // 0 (zero value for int)
```

### Checking Key Existence

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}

// Check if key exists
age, exists := ages["Alice"]
if exists {
    fmt.Printf("Alice is %d years old\n", age)
} else {
    fmt.Println("Alice not found")
}

// Shorter form
if age, ok := ages["Bob"]; ok {
    fmt.Printf("Bob is %d years old\n", age)
}
```

### Deleting Values

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}

// Delete a key-value pair
delete(ages, "Bob")

// Check if key exists after deletion
if _, exists := ages["Bob"]; !exists {
    fmt.Println("Bob has been removed")
}
```

## Map Types and Key Constraints

### Valid Key Types

Keys must be comparable types:

```go
// String keys
var stringMap map[string]int

// Integer keys
var intMap map[int]string

// Boolean keys
var boolMap map[bool]string

// Struct keys (if all fields are comparable)
type Point struct {
    X, Y int
}
var pointMap map[Point]string

// Array keys (if element type is comparable)
var arrayMap map[[2]int]string
```

### Invalid Key Types

```go
// Slice keys (not comparable)
// var sliceMap map[[]int]string  // Error

// Map keys (not comparable)
// var mapMap map[map[string]int]string  // Error

// Function keys (not comparable)
// var funcMap map[func()]string  // Error
```

### Value Types

Values can be any type:

```go
// Simple types
var stringMap map[string]string
var intMap map[string]int
var boolMap map[string]bool

// Complex types
var sliceMap map[string][]int
var mapMap map[string]map[string]int
var structMap map[string]Person

type Person struct {
    Name string
    Age  int
}
```

## Map Iteration

### Basic Iteration

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}

// Iterate over key-value pairs
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}

// Iterate over keys only
for name := range ages {
    fmt.Printf("Name: %s\n", name)
}

// Iterate over values only
for _, age := range ages {
    fmt.Printf("Age: %d\n", age)
}
```

### Ordered Iteration

```go
import "sort"

ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
    "Dave":  28,
}

// Get sorted keys
var names []string
for name := range ages {
    names = append(names, name)
}
sort.Strings(names)

// Iterate in sorted order
for _, name := range names {
    fmt.Printf("%s is %d years old\n", name, ages[name])
}
```

## Practical Examples

### Example 1: User Management

```go
package main

import "fmt"

type User struct {
    Name  string
    Email string
    Age   int
}

func main() {
    // Create a map of users
    users := make(map[string]User)
    
    // Add users
    users["alice"] = User{
        Name:  "Alice Johnson",
        Email: "alice@example.com",
        Age:   25,
    }
    
    users["bob"] = User{
        Name:  "Bob Smith",
        Email: "bob@example.com",
        Age:   30,
    }
    
    // Access user
    if user, exists := users["alice"]; exists {
        fmt.Printf("User: %s, Email: %s, Age: %d\n", 
                   user.Name, user.Email, user.Age)
    }
    
    // Update user
    users["alice"] = User{
        Name:  "Alice Johnson",
        Email: "alice.johnson@example.com", // Updated email
        Age:   26, // Updated age
    }
    
    // List all users
    fmt.Println("\nAll users:")
    for id, user := range users {
        fmt.Printf("ID: %s, Name: %s, Email: %s, Age: %d\n",
                   id, user.Name, user.Email, user.Age)
    }
    
    // Delete user
    delete(users, "bob")
    fmt.Printf("\nUsers after deletion: %d\n", len(users))
}
```

### Example 2: Word Frequency Counter

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "the quick brown fox jumps over the lazy dog the quick brown fox"
    words := strings.Fields(text)
    
    // Count word frequencies
    wordCount := make(map[string]int)
    for _, word := range words {
        wordCount[word]++
    }
    
    // Display results
    fmt.Println("Word frequencies:")
    for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }
    
    // Find most frequent word
    maxCount := 0
    var mostFrequent string
    for word, count := range wordCount {
        if count > maxCount {
            maxCount = count
            mostFrequent = word
        }
    }
    
    fmt.Printf("\nMost frequent word: '%s' (%d times)\n", mostFrequent, maxCount)
}
```

### Example 3: Configuration Management

```go
package main

import "fmt"

func main() {
    // Application configuration
    config := map[string]interface{}{
        "app_name":    "MyApp",
        "version":     "1.0.0",
        "debug_mode":  true,
        "max_users":   1000,
        "timeout":     30,
        "database": map[string]string{
            "host": "localhost",
            "port": "5432",
            "name": "myapp_db",
        },
    }
    
    // Access configuration values
    appName := config["app_name"].(string)
    debugMode := config["debug_mode"].(bool)
    maxUsers := config["max_users"].(int)
    
    fmt.Printf("App: %s\n", appName)
    fmt.Printf("Debug Mode: %t\n", debugMode)
    fmt.Printf("Max Users: %d\n", maxUsers)
    
    // Access nested map
    database := config["database"].(map[string]string)
    dbHost := database["host"]
    dbPort := database["port"]
    
    fmt.Printf("Database: %s:%s\n", dbHost, dbPort)
    
    // Update configuration
    config["max_users"] = 2000
    config["timeout"] = 60
    
    fmt.Printf("Updated Max Users: %d\n", config["max_users"])
}
```

### Example 4: Cache Implementation

```go
package main

import (
    "fmt"
    "time"
)

type CacheItem struct {
    Value     interface{}
    ExpiresAt time.Time
}

type Cache struct {
    items map[string]CacheItem
}

func NewCache() *Cache {
    return &Cache{
        items: make(map[string]CacheItem),
    }
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
    c.items[key] = CacheItem{
        Value:     value,
        ExpiresAt: time.Now().Add(ttl),
    }
}

func (c *Cache) Get(key string) (interface{}, bool) {
    item, exists := c.items[key]
    if !exists {
        return nil, false
    }
    
    // Check if expired
    if time.Now().After(item.ExpiresAt) {
        delete(c.items, key)
        return nil, false
    }
    
    return item.Value, true
}

func (c *Cache) Delete(key string) {
    delete(c.items, key)
}

func (c *Cache) Size() int {
    return len(c.items)
}

func main() {
    cache := NewCache()
    
    // Set values with TTL
    cache.Set("user:1", "Alice", 5*time.Second)
    cache.Set("user:2", "Bob", 10*time.Second)
    
    // Get values
    if value, exists := cache.Get("user:1"); exists {
        fmt.Printf("User 1: %s\n", value)
    }
    
    // Wait for expiration
    fmt.Println("Waiting for expiration...")
    time.Sleep(6 * time.Second)
    
    // Try to get expired value
    if value, exists := cache.Get("user:1"); exists {
        fmt.Printf("User 1: %s\n", value)
    } else {
        fmt.Println("User 1 expired")
    }
    
    // Get non-expired value
    if value, exists := cache.Get("user:2"); exists {
        fmt.Printf("User 2: %s\n", value)
    }
    
    fmt.Printf("Cache size: %d\n", cache.Size())
}
```

## Nested Maps

### Simple Nested Maps

```go
// Map of maps
var nestedMap map[string]map[string]int

// Initialize
nestedMap = make(map[string]map[string]int)
nestedMap["group1"] = make(map[string]int)
nestedMap["group1"]["item1"] = 10
nestedMap["group1"]["item2"] = 20

// Access
value := nestedMap["group1"]["item1"] // 10
```

### Complex Nested Structures

```go
type Person struct {
    Name string
    Age  int
}

type Department struct {
    Name        string
    Manager     Person
    Employees   []Person
    Budget      float64
}

func main() {
    // Map of departments
    departments := make(map[string]Department)
    
    // Add departments
    departments["engineering"] = Department{
        Name: "Engineering",
        Manager: Person{Name: "Alice", Age: 35},
        Employees: []Person{
            {Name: "Bob", Age: 28},
            {Name: "Carol", Age: 32},
        },
        Budget: 100000.0,
    }
    
    departments["marketing"] = Department{
        Name: "Marketing",
        Manager: Person{Name: "Dave", Age: 40},
        Employees: []Person{
            {Name: "Eve", Age: 25},
        },
        Budget: 50000.0,
    }
    
    // Access nested data
    engDept := departments["engineering"]
    fmt.Printf("Engineering Manager: %s\n", engDept.Manager.Name)
    fmt.Printf("Engineering Budget: $%.2f\n", engDept.Budget)
    
    // List all departments
    for name, dept := range departments {
        fmt.Printf("\n%s Department:\n", name)
        fmt.Printf("  Manager: %s (age %d)\n", dept.Manager.Name, dept.Manager.Age)
        fmt.Printf("  Employees: %d\n", len(dept.Employees))
        fmt.Printf("  Budget: $%.2f\n", dept.Budget)
    }
}
```

## Map Performance

### Time Complexity

- **Insertion**: O(1) average, O(n) worst case
- **Lookup**: O(1) average, O(n) worst case
- **Deletion**: O(1) average, O(n) worst case
- **Iteration**: O(n)

### Memory Considerations

```go
// Pre-allocate map size if known
ages := make(map[string]int, 1000) // Hint for initial capacity

// Check map size
fmt.Printf("Map size: %d\n", len(ages))

// Check if map is nil
if ages == nil {
    fmt.Println("Map is nil")
}
```

## Best Practices

### 1. Initialize Maps Properly

```go
// Good
ages := make(map[string]int)
ages["Alice"] = 25

// Bad
var ages map[string]int
ages["Alice"] = 25 // Panic: assignment to entry in nil map
```

### 2. Check Key Existence

```go
// Good
if age, exists := ages["Alice"]; exists {
    fmt.Printf("Alice is %d years old\n", age)
} else {
    fmt.Println("Alice not found")
}

// Less preferred
age := ages["Alice"] // Returns 0 if key doesn't exist
if age != 0 {
    fmt.Printf("Alice is %d years old\n", age)
}
```

### 3. Use Appropriate Key Types

```go
// Good - string keys for user IDs
users := make(map[string]User)

// Good - int keys for sequential data
scores := make(map[int]int)

// Less preferred - complex struct keys
type ComplexKey struct {
    A, B, C string
}
complexMap := make(map[ComplexKey]string)
```

### 4. Handle Zero Values

```go
// Good - explicit zero value handling
if count, exists := wordCount["nonexistent"]; exists {
    fmt.Printf("Count: %d\n", count)
} else {
    fmt.Println("Word not found")
}

// Less preferred - relying on zero values
count := wordCount["nonexistent"] // 0
if count > 0 {
    fmt.Printf("Count: %d\n", count)
}
```

## Common Patterns

### 1. Set Implementation

```go
// Set using map[type]bool
type Set map[string]bool

func NewSet() Set {
    return make(Set)
}

func (s Set) Add(item string) {
    s[item] = true
}

func (s Set) Contains(item string) bool {
    return s[item]
}

func (s Set) Remove(item string) {
    delete(s, item)
}

func (s Set) Size() int {
    return len(s)
}

func main() {
    set := NewSet()
    set.Add("apple")
    set.Add("banana")
    set.Add("apple") // Duplicate
    
    fmt.Printf("Set size: %d\n", set.Size())
    fmt.Printf("Contains 'apple': %t\n", set.Contains("apple"))
    fmt.Printf("Contains 'orange': %t\n", set.Contains("orange"))
}
```

### 2. Grouping Data

```go
func main() {
    // Group people by age
    people := []struct {
        Name string
        Age  int
    }{
        {"Alice", 25},
        {"Bob", 30},
        {"Carol", 25},
        {"Dave", 30},
        {"Eve", 25},
    }
    
    // Group by age
    ageGroups := make(map[int][]string)
    for _, person := range people {
        ageGroups[person.Age] = append(ageGroups[person.Age], person.Name)
    }
    
    // Display groups
    for age, names := range ageGroups {
        fmt.Printf("Age %d: %v\n", age, names)
    }
}
```

### 3. Counting Frequencies

```go
func main() {
    // Count character frequencies
    text := "hello world"
    charCount := make(map[rune]int)
    
    for _, char := range text {
        charCount[char]++
    }
    
    // Display frequencies
    for char, count := range charCount {
        if char != ' ' {
            fmt.Printf("'%c': %d\n", char, count)
        }
    }
}
```

## Common Mistakes

### 1. Using Nil Maps

```go
// Error
var ages map[string]int
ages["Alice"] = 25 // Panic: assignment to entry in nil map

// Correct
ages := make(map[string]int)
ages["Alice"] = 25
```

### 2. Not Checking Key Existence

```go
// Problematic
age := ages["Alice"] // Returns 0 if key doesn't exist
if age > 0 {
    fmt.Printf("Alice is %d years old\n", age)
}

// Correct
if age, exists := ages["Alice"]; exists {
    fmt.Printf("Alice is %d years old\n", age)
}
```

### 3. Modifying Maps During Iteration

```go
// Problematic
for key, value := range ages {
    if value < 18 {
        delete(ages, key) // May cause unpredictable behavior
    }
}

// Correct
var toDelete []string
for key, value := range ages {
    if value < 18 {
        toDelete = append(toDelete, key)
    }
}
for _, key := range toDelete {
    delete(ages, key)
}
```

### 4. Using Incomparable Key Types

```go
// Error
var sliceMap map[[]int]string // Slices are not comparable

// Correct
var stringMap map[string]string
var intMap map[int]string
```

## Exercises

### Exercise 1: Student Grades
Create a map to store student grades and implement functions to add, update, and calculate average grades.

### Exercise 2: Word Counter
Implement a word counter that counts the frequency of each word in a text.

### Exercise 3: Cache System
Create a simple cache system with TTL (time-to-live) functionality.

### Exercise 4: Grouping Data
Group a list of people by their department and calculate statistics for each department.

### Exercise 5: Set Operations
Implement set operations (union, intersection, difference) using maps.

## Next Steps

Now that you understand maps, explore:

1. **[Pointers](../09.pointers)** - Learn about memory addresses and pointers
2. **[Functions](../10.functions)** - Create reusable code blocks
3. **[Structs](../11.structs)** - Define custom data types
4. **[Interfaces](../12.interfaces)** - Define behavior contracts

## Key Takeaways

- Maps are unordered key-value collections
- Keys must be comparable types
- Values can be any type
- Maps have O(1) average time complexity for operations
- Always check key existence when accessing values
- Use `make()` to initialize maps before use
- Maps are reference types
- Zero value of a map is `nil`

## Additional Resources

- [Go Tour - Maps](https://tour.golang.org/moretypes/19)
- [Effective Go - Maps](https://golang.org/doc/effective_go.html#maps)
- [Go by Example - Maps](https://gobyexample.com/maps)
- [Go Language Specification - Maps](https://golang.org/ref/spec#Map_types)
