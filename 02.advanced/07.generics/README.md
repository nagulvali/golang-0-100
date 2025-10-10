# Go Generics

## Table of Contents
1. [Introduction](#introduction)
2. [What are Generics?](#what-are-generics)
3. [Why Generics?](#why-generics)
4. [Basic Syntax](#basic-syntax)
5. [Type Parameters](#type-parameters)
6. [Type Constraints](#type-constraints)
7. [Built-in Constraints](#built-in-constraints)
8. [Generic Functions](#generic-functions)
9. [Generic Types](#generic-types)
10. [Generic Methods](#generic-methods)
11. [Type Inference](#type-inference)
12. [Multiple Type Parameters](#multiple-type-parameters)
13. [Custom Constraints](#custom-constraints)
14. [Comparable Constraint](#comparable-constraint)
15. [Union Types](#union-types)
16. [Common Patterns](#common-patterns)
17. [Real-World Examples](#real-world-examples)
18. [Limitations and Gotchas](#limitations-and-gotchas)
19. [Best Practices](#best-practices)
20. [Performance Considerations](#performance-considerations)

---

## Introduction

Generics (also called type parameters or parametric polymorphism) were introduced in Go 1.18 (released March 2022), marking one of the most significant additions to the language. Generics allow you to write functions and data structures that work with multiple types while maintaining type safety.

**Key Points:**
- Added in Go 1.18
- Enables type-safe, reusable code
- Compile-time type checking
- No runtime overhead
- Reduces code duplication

---

## What are Generics?

**Generics** allow you to write code that is parameterized by types. Instead of writing separate functions for different types, you write one generic function that works with any type that meets certain constraints.

### Before Generics (Go 1.17 and earlier)

```go
// Separate functions for each type
func SumInts(nums []int) int {
    var sum int
    for _, n := range nums {
        sum += n
    }
    return sum
}

func SumFloats(nums []float64) float64 {
    var sum float64
    for _, n := range nums {
        sum += n
    }
    return sum
}

// Using interface{} loses type safety
func SumInterface(nums []interface{}) interface{} {
    // Complex type assertions required
    // No compile-time type checking
}
```

### With Generics (Go 1.18+)

```go
// One function works for all numeric types
func Sum[T int | float64](nums []T) T {
    var sum T
    for _, n := range nums {
        sum += n
    }
    return sum
}

// Usage
ints := []int{1, 2, 3, 4, 5}
floats := []float64{1.5, 2.5, 3.5}

fmt.Println(Sum(ints))    // 15
fmt.Println(Sum(floats))  // 7.5
```

---

## Why Generics?

### Problems Solved by Generics

**1. Code Duplication**
```go
// Before: Duplicate code
type IntStack struct {
    items []int
}

type StringStack struct {
    items []string
}

// After: Single generic implementation
type Stack[T any] struct {
    items []T
}
```

**2. Type Safety**
```go
// Before: interface{} loses type safety
func Contains(slice []interface{}, item interface{}) bool {
    // Runtime type checks required
}

// After: Type-safe with generics
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}
```

**3. Performance**
```go
// Before: interface{} has boxing overhead
var x interface{} = 42  // Boxing: heap allocation

// After: Generics maintain native types
func Identity[T any](x T) T {
    return x  // No boxing, direct value
}
```

**4. Better API Design**
```go
// More expressive and safer APIs
type Result[T any, E error] struct {
    value T
    err   E
}
```

---

## Basic Syntax

### Generic Function Syntax

```go
func FunctionName[TypeParam Constraint](param Type) ReturnType {
    // function body
}
```

### Generic Type Syntax

```go
type TypeName[TypeParam Constraint] struct {
    // fields
}
```

### Components

1. **Type Parameter**: `T` (or any name)
2. **Constraint**: `any`, `comparable`, custom interface
3. **Square Brackets**: `[T any]` declares type parameters

### Simple Example

```go
// Generic function with 'any' constraint
func Print[T any](value T) {
    fmt.Println(value)
}

// Usage
Print(42)           // T is int
Print("hello")      // T is string
Print(3.14)         // T is float64
Print([]int{1, 2})  // T is []int
```

---

## Type Parameters

Type parameters are placeholders for actual types. They appear in square brackets.

### Single Type Parameter

```go
func First[T any](slice []T) T {
    return slice[0]
}

// Usage
numbers := []int{1, 2, 3}
first := First(numbers)  // first is int

words := []string{"a", "b", "c"}
firstWord := First(words)  // firstWord is string
```

### Naming Conventions

```go
// Common conventions
func Process[T any](item T) T          // T for Type
func Map[K comparable, V any](m map[K]V)  // K for Key, V for Value
func Transform[In, Out any](input In) Out  // Descriptive names
```

### Type Parameter Scope

```go
func Example[T any](x T) {
    // T is available throughout the function
    var y T  // y has type T
    
    inner := func(z T) T {
        return z  // T is also available in closures
    }
    
    return inner(x)
}
```

---

## Type Constraints

Constraints specify what operations are allowed on type parameters.

### The `any` Constraint

`any` is an alias for `interface{}` - allows any type:

```go
func Identity[T any](x T) T {
    return x  // Only operations: assignment, comparison with same type
}

// Can use with any type
Identity(42)
Identity("hello")
Identity(struct{}{})
```

### Constraints Limit Operations

```go
// Without constraint, can't perform operations
func Add[T any](a, b T) T {
    // return a + b  // ERROR: operator + not defined for T
    return a  // Can only return
}

// With constraint, can perform addition
func Add[T int | float64](a, b T) T {
    return a + b  // OK: + is valid for int and float64
}
```

---

## Built-in Constraints

Go provides built-in constraints in the `constraints` package.

### Importing Constraints

```go
import "golang.org/x/exp/constraints"
```

### Common Built-in Constraints

```go
// Ordered: types that support <, <=, >, >=
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// Integer: all integer types
func AbsInt[T constraints.Integer](x T) T {
    if x < 0 {
        return -x
    }
    return x
}

// Float: all float types
func AbsFloat[T constraints.Float](x T) T {
    if x < 0 {
        return -x
    }
    return x
}

// Signed: signed numeric types
func Negate[T constraints.Signed](x T) T {
    return -x
}

// Unsigned: unsigned integer types
func IsEven[T constraints.Unsigned](x T) bool {
    return x%2 == 0
}
```

### Available Built-in Constraints

```go
constraints.Signed       // int, int8, int16, int32, int64
constraints.Unsigned     // uint, uint8, uint16, uint32, uint64, uintptr
constraints.Integer      // Signed | Unsigned
constraints.Float        // float32, float64
constraints.Complex      // complex64, complex128
constraints.Ordered      // Integer | Float | string
```

---

## Generic Functions

Generic functions accept type parameters and work with multiple types.

### Basic Generic Function

```go
func Swap[T any](a, b T) (T, T) {
    return b, a
}

// Usage
x, y := Swap(1, 2)           // int, int
s1, s2 := Swap("a", "b")     // string, string
```

### Generic Function with Multiple Parameters

```go
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// Usage
nums := []int{1, 2, 3, 4, 5}
fmt.Println(Contains(nums, 3))  // true

words := []string{"go", "rust", "python"}
fmt.Println(Contains(words, "java"))  // false
```

### Generic Function with Constraints

```go
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

fmt.Println(Max(10, 20))        // 20
fmt.Println(Max(3.14, 2.71))    // 3.14
fmt.Println(Max("apple", "banana"))  // banana
```

### Generic Variadic Functions

```go
func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
    var sum T
    for _, n := range nums {
        sum += n
    }
    return sum
}

// Usage
fmt.Println(Sum(1, 2, 3, 4, 5))           // 15
fmt.Println(Sum(1.1, 2.2, 3.3, 4.4))      // 11.0
```

---

## Generic Types

You can create generic structs, interfaces, and type aliases.

### Generic Struct

```go
type Pair[T, U any] struct {
    First  T
    Second U
}

// Usage
p1 := Pair[int, string]{First: 1, Second: "one"}
p2 := Pair[string, bool]{First: "active", Second: true}

fmt.Println(p1.First, p1.Second)  // 1 one
fmt.Println(p2.First, p2.Second)  // active true
```

### Generic Stack

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

// Usage
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)
intStack.Push(3)

val, ok := intStack.Pop()
fmt.Println(val, ok)  // 3 true

stringStack := Stack[string]{}
stringStack.Push("hello")
stringStack.Push("world")
```

### Generic Queue

```go
type Queue[T any] struct {
    items []T
}

func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}

func (q *Queue[T]) Size() int {
    return len(q.items)
}
```

### Generic Map Type

```go
type Cache[K comparable, V any] struct {
    data map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
    return &Cache[K, V]{
        data: make(map[K]V),
    }
}

func (c *Cache[K, V]) Set(key K, value V) {
    c.data[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache[K, V]) Delete(key K) {
    delete(c.data, key)
}

// Usage
cache := NewCache[string, int]()
cache.Set("age", 30)
cache.Set("score", 95)

age, ok := cache.Get("age")
fmt.Println(age, ok)  // 30 true
```

---

## Generic Methods

Methods on generic types maintain the type parameters.

### Methods on Generic Types

```go
type List[T any] struct {
    items []T
}

func (l *List[T]) Add(item T) {
    l.items = append(l.items, item)
}

func (l *List[T]) Get(index int) (T, error) {
    if index < 0 || index >= len(l.items) {
        var zero T
        return zero, fmt.Errorf("index out of bounds")
    }
    return l.items[index], nil
}

func (l *List[T]) Len() int {
    return len(l.items)
}

func (l *List[T]) Filter(predicate func(T) bool) *List[T] {
    result := &List[T]{}
    for _, item := range l.items {
        if predicate(item) {
            result.Add(item)
        }
    }
    return result
}

// Usage
numbers := &List[int]{}
numbers.Add(1)
numbers.Add(2)
numbers.Add(3)
numbers.Add(4)
numbers.Add(5)

evens := numbers.Filter(func(n int) bool {
    return n%2 == 0
})

for i := 0; i < evens.Len(); i++ {
    val, _ := evens.Get(i)
    fmt.Println(val)  // 2, 4
}
```

### Methods with Additional Type Parameters

```go
type Container[T any] struct {
    value T
}

// Method with same type parameter
func (c *Container[T]) Get() T {
    return c.value
}

func (c *Container[T]) Set(value T) {
    c.value = value
}

// Method with additional type parameter
func (c *Container[T]) Transform[U any](fn func(T) U) *Container[U] {
    return &Container[U]{value: fn(c.value)}
}

// Usage
intContainer := &Container[int]{value: 42}
stringContainer := intContainer.Transform(func(i int) string {
    return fmt.Sprintf("Number: %d", i)
})

fmt.Println(stringContainer.Get())  // Number: 42
```

---

## Type Inference

Go can often infer type parameters from function arguments.

### Automatic Type Inference

```go
func Print[T any](value T) {
    fmt.Println(value)
}

// Type inferred from argument
Print(42)       // T inferred as int
Print("hello")  // T inferred as string

// Explicit type specification (optional)
Print[int](42)
Print[string]("hello")
```

### Inference with Multiple Parameters

```go
func Pair[T any](a, b T) [2]T {
    return [2]T{a, b}
}

// Inferred
result := Pair(1, 2)  // T is int

// Explicit
result := Pair[int](1, 2)
```

### When Inference Fails

```go
func Make[T any]() T {
    var zero T
    return zero
}

// Must specify type explicitly
val := Make[int]()      // OK
// val := Make()        // ERROR: cannot infer T
```

### Inference with Return Types

```go
func Convert[T any, U any](value T, converter func(T) U) U {
    return converter(value)
}

// Type inferred from function arguments
result := Convert(42, func(i int) string {
    return fmt.Sprintf("%d", i)
})
// T is int, U is string
```

---

## Multiple Type Parameters

Functions and types can have multiple type parameters.

### Multiple Parameters in Functions

```go
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// Usage
numbers := []int{1, 2, 3, 4, 5}
strings := Map(numbers, func(n int) string {
    return fmt.Sprintf("Number: %d", n)
})
// strings: ["Number: 1", "Number: 2", ...]
```

### Multiple Parameters in Types

```go
type Result[T any, E error] struct {
    Value T
    Err   E
}

func (r Result[T, E]) IsOk() bool {
    return r.Err == nil
}

func (r Result[T, E]) Unwrap() T {
    if r.Err != nil {
        panic(r.Err)
    }
    return r.Value
}

// Usage
func Divide(a, b float64) Result[float64, error] {
    if b == 0 {
        return Result[float64, error]{
            Err: fmt.Errorf("division by zero"),
        }
    }
    return Result[float64, error]{
        Value: a / b,
    }
}

result := Divide(10, 2)
if result.IsOk() {
    fmt.Println(result.Unwrap())  // 5
}
```

### Three or More Parameters

```go
type Triple[A, B, C any] struct {
    First  A
    Second B
    Third  C
}

func MakeTriple[A, B, C any](a A, b B, c C) Triple[A, B, C] {
    return Triple[A, B, C]{
        First:  a,
        Second: b,
        Third:  c,
    }
}

// Usage
t := MakeTriple(1, "two", 3.0)
fmt.Printf("%v\n", t)  // {1 two 3}
```

---

## Custom Constraints

You can define custom constraints using interfaces.

### Basic Custom Constraint

```go
type Number interface {
    int | int64 | float64
}

func Add[T Number](a, b T) T {
    return a + b
}

fmt.Println(Add(1, 2))        // 3
fmt.Println(Add(1.5, 2.5))    // 4.0
```

### Constraint with Methods

```go
type Stringer interface {
    String() string
}

func PrintAll[T Stringer](items []T) {
    for _, item := range items {
        fmt.Println(item.String())
    }
}

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

// Usage
people := []Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
}
PrintAll(people)
```

### Combining Constraints

```go
type Numeric interface {
    int | int64 | float64
}

type Addable interface {
    Numeric
    // Additional constraint
}

func Sum[T Addable](nums []T) T {
    var sum T
    for _, n := range nums {
        sum += n
    }
    return sum
}
```

### Constraint with Type and Method Requirements

```go
type Comparable[T any] interface {
    Compare(T) int  // Method requirement
}

func Max[T Comparable[T]](a, b T) T {
    if a.Compare(b) > 0 {
        return a
    }
    return b
}

type Version struct {
    Major, Minor, Patch int
}

func (v Version) Compare(other Version) int {
    if v.Major != other.Major {
        return v.Major - other.Major
    }
    if v.Minor != other.Minor {
        return v.Minor - other.Minor
    }
    return v.Patch - other.Patch
}

// Usage
v1 := Version{1, 2, 3}
v2 := Version{1, 3, 0}
max := Max(v1, v2)
fmt.Printf("%+v\n", max)  // {Major:1 Minor:3 Patch:0}
```

---

## Comparable Constraint

The `comparable` constraint allows types that can be compared with `==` and `!=`.

### Using comparable

```go
func Index[T comparable](slice []T, target T) int {
    for i, v := range slice {
        if v == target {
            return i
        }
    }
    return -1
}

// Works with any comparable type
fmt.Println(Index([]int{1, 2, 3}, 2))              // 1
fmt.Println(Index([]string{"a", "b", "c"}, "b"))   // 1
```

### comparable vs any

```go
// any: can't use == or !=
func PrintAny[T any](a, b T) {
    fmt.Println(a)
    // if a == b { }  // ERROR: can't compare
}

// comparable: can use == and !=
func Equal[T comparable](a, b T) bool {
    return a == b  // OK
}

// Works with comparable types
Equal(1, 1)        // true
Equal("a", "b")    // false

// Equal([]int{}, []int{})  // ERROR: slices not comparable
```

### Map Keys Must Be comparable

```go
type Set[T comparable] struct {
    data map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
    return &Set[T]{
        data: make(map[T]struct{}),
    }
}

func (s *Set[T]) Add(item T) {
    s.data[item] = struct{}{}
}

func (s *Set[T]) Contains(item T) bool {
    _, ok := s.data[item]
    return ok
}

func (s *Set[T]) Remove(item T) {
    delete(s.data, item)
}

// Usage
set := NewSet[string]()
set.Add("apple")
set.Add("banana")
fmt.Println(set.Contains("apple"))  // true
```

---

## Union Types

Union types (type sets) specify multiple allowed types using `|`.

### Basic Union

```go
type Integer interface {
    int | int8 | int16 | int32 | int64
}

func AddInt[T Integer](a, b T) T {
    return a + b
}
```

### Union with Multiple Types

```go
type Number interface {
    int | int64 | float32 | float64
}

func Abs[T Number](x T) T {
    if x < 0 {
        return -x
    }
    return x
}

fmt.Println(Abs(-5))     // 5
fmt.Println(Abs(-3.14))  // 3.14
```

### Approximate Type Constraints

Use `~` to include types with the same underlying type:

```go
type MyInt int

// Without ~: only exact types allowed
type ExactInt interface {
    int | int64
}

func UseExact[T ExactInt](x T) T {
    return x
}

// UseExact(MyInt(5))  // ERROR: MyInt not in type set

// With ~: includes derived types
type ApproximateInt interface {
    ~int | ~int64
}

func UseApproximate[T ApproximateInt](x T) T {
    return x
}

UseApproximate(MyInt(5))  // OK
UseApproximate(5)         // OK
```

### Union with Methods

```go
type StringLike interface {
    ~string
    Len() int
}

type MyString string

func (s MyString) Len() int {
    return len(s)
}

func Process[T StringLike](s T) {
    fmt.Println("Length:", s.Len())
    // Can also use string operations because of ~string
    fmt.Println("Value:", string(s))
}

Process(MyString("hello"))
```

---

## Common Patterns

### 1. Generic Container Types

```go
type Optional[T any] struct {
    value *T
}

func Some[T any](value T) Optional[T] {
    return Optional[T]{value: &value}
}

func None[T any]() Optional[T] {
    return Optional[T]{value: nil}
}

func (o Optional[T]) IsSome() bool {
    return o.value != nil
}

func (o Optional[T]) IsNone() bool {
    return o.value == nil
}

func (o Optional[T]) Unwrap() T {
    if o.IsNone() {
        panic("called Unwrap on None")
    }
    return *o.value
}

func (o Optional[T]) UnwrapOr(defaultValue T) T {
    if o.IsNone() {
        return defaultValue
    }
    return *o.value
}

// Usage
opt1 := Some(42)
opt2 := None[int]()

fmt.Println(opt1.Unwrap())        // 42
fmt.Println(opt2.UnwrapOr(0))     // 0
```

### 2. Generic Algorithms

```go
// Filter
func Filter[T any](slice []T, predicate func(T) bool) []T {
    result := make([]T, 0)
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

// Map
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// Reduce
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
    result := initial
    for _, v := range slice {
        result = fn(result, v)
    }
    return result
}

// Usage
numbers := []int{1, 2, 3, 4, 5}

evens := Filter(numbers, func(n int) bool {
    return n%2 == 0
})
// evens: [2, 4]

doubled := Map(numbers, func(n int) int {
    return n * 2
})
// doubled: [2, 4, 6, 8, 10]

sum := Reduce(numbers, 0, func(acc, n int) int {
    return acc + n
})
// sum: 15
```

### 3. Generic Linked List

```go
type Node[T any] struct {
    Value T
    Next  *Node[T]
}

type LinkedList[T any] struct {
    head *Node[T]
    tail *Node[T]
    size int
}

func (l *LinkedList[T]) Append(value T) {
    node := &Node[T]{Value: value}
    
    if l.head == nil {
        l.head = node
        l.tail = node
    } else {
        l.tail.Next = node
        l.tail = node
    }
    l.size++
}

func (l *LinkedList[T]) Prepend(value T) {
    node := &Node[T]{Value: value, Next: l.head}
    l.head = node
    if l.tail == nil {
        l.tail = node
    }
    l.size++
}

func (l *LinkedList[T]) ToSlice() []T {
    result := make([]T, 0, l.size)
    current := l.head
    for current != nil {
        result = append(result, current.Value)
        current = current.Next
    }
    return result
}

// Usage
list := &LinkedList[int]{}
list.Append(1)
list.Append(2)
list.Append(3)
fmt.Println(list.ToSlice())  // [1, 2, 3]
```

### 4. Generic Binary Tree

```go
type TreeNode[T constraints.Ordered] struct {
    Value T
    Left  *TreeNode[T]
    Right *TreeNode[T]
}

type BinarySearchTree[T constraints.Ordered] struct {
    root *TreeNode[T]
}

func (t *BinarySearchTree[T]) Insert(value T) {
    t.root = t.insertNode(t.root, value)
}

func (t *BinarySearchTree[T]) insertNode(node *TreeNode[T], value T) *TreeNode[T] {
    if node == nil {
        return &TreeNode[T]{Value: value}
    }
    
    if value < node.Value {
        node.Left = t.insertNode(node.Left, value)
    } else if value > node.Value {
        node.Right = t.insertNode(node.Right, value)
    }
    
    return node
}

func (t *BinarySearchTree[T]) Contains(value T) bool {
    return t.search(t.root, value)
}

func (t *BinarySearchTree[T]) search(node *TreeNode[T], value T) bool {
    if node == nil {
        return false
    }
    
    if value == node.Value {
        return true
    } else if value < node.Value {
        return t.search(node.Left, value)
    } else {
        return t.search(node.Right, value)
    }
}

func (t *BinarySearchTree[T]) InOrder() []T {
    result := []T{}
    t.inOrderTraversal(t.root, &result)
    return result
}

func (t *BinarySearchTree[T]) inOrderTraversal(node *TreeNode[T], result *[]T) {
    if node != nil {
        t.inOrderTraversal(node.Left, result)
        *result = append(*result, node.Value)
        t.inOrderTraversal(node.Right, result)
    }
}

// Usage
tree := &BinarySearchTree[int]{}
tree.Insert(5)
tree.Insert(3)
tree.Insert(7)
tree.Insert(1)
tree.Insert(9)

fmt.Println(tree.InOrder())      // [1, 3, 5, 7, 9]
fmt.Println(tree.Contains(3))    // true
fmt.Println(tree.Contains(10))   // false
```

---

## Real-World Examples

### 1. Generic Result Type (Error Handling)

```go
type Result[T any] struct {
    value T
    err   error
}

func Ok[T any](value T) Result[T] {
    return Result[T]{value: value}
}

func Err[T any](err error) Result[T] {
    var zero T
    return Result[T]{value: zero, err: err}
}

func (r Result[T]) IsOk() bool {
    return r.err == nil
}

func (r Result[T]) IsErr() bool {
    return r.err != nil
}

func (r Result[T]) Unwrap() T {
    if r.IsErr() {
        panic(fmt.Sprintf("called Unwrap on Err: %v", r.err))
    }
    return r.value
}

func (r Result[T]) UnwrapOr(defaultValue T) T {
    if r.IsErr() {
        return defaultValue
    }
    return r.value
}

func (r Result[T]) Map[U any](fn func(T) U) Result[U] {
    if r.IsErr() {
        return Err[U](r.err)
    }
    return Ok(fn(r.value))
}

// Usage
func ParseInt(s string) Result[int] {
    val, err := strconv.Atoi(s)
    if err != nil {
        return Err[int](err)
    }
    return Ok(val)
}

result := ParseInt("42")
if result.IsOk() {
    fmt.Println("Value:", result.Unwrap())  // Value: 42
}

doubled := result.Map(func(n int) int {
    return n * 2
})
fmt.Println("Doubled:", doubled.Unwrap())  // Doubled: 84
```

### 2. Generic Repository Pattern

```go
type Entity interface {
    GetID() string
}

type Repository[T Entity] struct {
    data map[string]T
    mu   sync.RWMutex
}

func NewRepository[T Entity]() *Repository[T] {
    return &Repository[T]{
        data: make(map[string]T),
    }
}

func (r *Repository[T]) Save(entity T) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    r.data[entity.GetID()] = entity
    return nil
}

func (r *Repository[T]) FindByID(id string) (T, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    entity, ok := r.data[id]
    if !ok {
        var zero T
        return zero, fmt.Errorf("entity not found: %s", id)
    }
    return entity, nil
}

func (r *Repository[T]) FindAll() []T {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    result := make([]T, 0, len(r.data))
    for _, entity := range r.data {
        result = append(result, entity)
    }
    return result
}

func (r *Repository[T]) Delete(id string) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    if _, ok := r.data[id]; !ok {
        return fmt.Errorf("entity not found: %s", id)
    }
    delete(r.data, id)
    return nil
}

// Usage
type User struct {
    ID       string
    Username string
    Email    string
}

func (u User) GetID() string {
    return u.ID
}

userRepo := NewRepository[User]()
userRepo.Save(User{ID: "1", Username: "alice", Email: "alice@example.com"})
userRepo.Save(User{ID: "2", Username: "bob", Email: "bob@example.com"})

user, _ := userRepo.FindByID("1")
fmt.Printf("%+v\n", user)

allUsers := userRepo.FindAll()
fmt.Printf("Total users: %d\n", len(allUsers))
```

### 3. Generic Event Bus

```go
type Event any

type Handler[E Event] func(E)

type EventBus[E Event] struct {
    handlers []Handler[E]
    mu       sync.RWMutex
}

func NewEventBus[E Event]() *EventBus[E] {
    return &EventBus[E]{
        handlers: make([]Handler[E], 0),
    }
}

func (eb *EventBus[E]) Subscribe(handler Handler[E]) {
    eb.mu.Lock()
    defer eb.mu.Unlock()
    eb.handlers = append(eb.handlers, handler)
}

func (eb *EventBus[E]) Publish(event E) {
    eb.mu.RLock()
    handlers := make([]Handler[E], len(eb.handlers))
    copy(handlers, eb.handlers)
    eb.mu.RUnlock()
    
    for _, handler := range handlers {
        go handler(event)
    }
}

// Usage
type UserCreatedEvent struct {
    UserID   string
    Username string
    Email    string
}

bus := NewEventBus[UserCreatedEvent]()

bus.Subscribe(func(event UserCreatedEvent) {
    fmt.Printf("Sending welcome email to %s\n", event.Email)
})

bus.Subscribe(func(event UserCreatedEvent) {
    fmt.Printf("Creating default settings for user %s\n", event.UserID)
})

bus.Publish(UserCreatedEvent{
    UserID:   "123",
    Username: "alice",
    Email:    "alice@example.com",
})
```

### 4. Generic Pool

```go
type Pool[T any] struct {
    items chan T
    new   func() T
}

func NewPool[T any](size int, factory func() T) *Pool[T] {
    return &Pool[T]{
        items: make(chan T, size),
        new:   factory,
    }
}

func (p *Pool[T]) Get() T {
    select {
    case item := <-p.items:
        return item
    default:
        return p.new()
    }
}

func (p *Pool[T]) Put(item T) {
    select {
    case p.items <- item:
    default:
        // Pool is full, discard item
    }
}

// Usage
type Connection struct {
    ID string
}

func createConnection() Connection {
    return Connection{ID: fmt.Sprintf("conn-%d", time.Now().UnixNano())}
}

pool := NewPool(5, createConnection)

// Get connection from pool
conn := pool.Get()
fmt.Printf("Got connection: %s\n", conn.ID)

// Use connection...

// Return to pool
pool.Put(conn)
```

---

## Limitations and Gotchas

### 1. No Method Type Parameters

```go
type Container[T any] struct {
    value T
}

// ERROR: Methods can't have their own type parameters independent of the type
// func (c Container[T]) Convert[U any]() U {
//     // ...
// }

// WORKAROUND: Add type parameter to method signature
func (c Container[T]) Convert[U any](fn func(T) U) U {
    return fn(c.value)
}
```

### 2. No Generic Types in Methods (Without Receiver Type Parameter)

```go
// ERROR: Can't declare generic method without type parameter on receiver
// func (s *MyStruct) GenericMethod[T any](x T) T {
//     return x
// }

// OK: Type parameter on the struct
type MyStruct[T any] struct{}

func (s *MyStruct[T]) Method(x T) T {
    return x
}
```

### 3. Type Switches Don't Work with Type Parameters

```go
func Process[T any](value T) {
    // ERROR: Can't use type switch on type parameter
    // switch value.(type) {
    // case int:
    //     ...
    // }
    
    // WORKAROUND: Use interface{} or any
    switch v := any(value).(type) {
    case int:
        fmt.Println("int:", v)
    case string:
        fmt.Println("string:", v)
    default:
        fmt.Println("other")
    }
}
```

### 4. Can't Use Operators Without Constraints

```go
// ERROR: Can't use operators without appropriate constraints
// func Add[T any](a, b T) T {
//     return a + b  // ERROR: operator + not defined for T
// }

// OK: With constraint
func Add[T constraints.Integer | constraints.Float](a, b T) T {
    return a + b  // OK
}
```

### 5. Zero Values

```go
func GetZero[T any]() T {
    var zero T  // Zero value of T
    return zero
}

fmt.Println(GetZero[int]())     // 0
fmt.Println(GetZero[string]())  // ""
fmt.Println(GetZero[*int]())    // <nil>
```

### 6. Type Inference Limitations

```go
func MakePair[T any]() (T, T) {
    var zero T
    return zero, zero
}

// ERROR: Can't infer T
// pair := MakePair()

// OK: Explicit type
pair := MakePair[int]()  // (0, 0)
```

### 7. Can't Create Instances Directly

```go
func Create[T any]() T {
    // ERROR: Can't use new(T) or make(T)
    // return new(T)
    
    // OK: Return zero value
    var zero T
    return zero
}

// WORKAROUND: Use factory function
func CreateWithFactory[T any](factory func() T) T {
    return factory()
}
```

---

## Best Practices

### 1. Use Generics When Type Safety Matters

```go
// Good: Type-safe generic stack
type Stack[T any] struct {
    items []T
}

// Avoid: Interface{} loses type safety
type UnsafeStack struct {
    items []interface{}
}
```

### 2. Prefer Simple Constraints

```go
// Good: Simple, clear constraint
func Sum[T constraints.Integer](nums []T) T {
    var sum T
    for _, n := range nums {
        sum += n
    }
    return sum
}

// Avoid: Overly complex constraints
type ComplexConstraint interface {
    ~int | ~int64 | ~float64
    String() string
    comparable
}
```

### 3. Use Descriptive Type Parameter Names

```go
// Good: Descriptive names
func Transform[Input, Output any](data Input, fn func(Input) Output) Output {
    return fn(data)
}

// Avoid: Unclear names
func Transform[T, U any](data T, fn func(T) U) U {
    return fn(data)
}

// OK for simple cases
func Identity[T any](x T) T {
    return x
}
```

### 4. Don't Overuse Generics

```go
// Avoid: Unnecessary generic
func PrintString[T string](s T) {
    fmt.Println(s)
}

// Better: Just use string
func PrintString(s string) {
    fmt.Println(s)
}

// Good use: Actually generic
func Print[T any](value T) {
    fmt.Println(value)
}
```

### 5. Document Generic Functions

```go
// Max returns the maximum of two Ordered values.
// It works with any type that supports comparison operators.
//
// Example:
//   Max(10, 20)        // 20
//   Max(3.14, 2.71)    // 3.14
//   Max("apple", "banana")  // "banana"
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

### 6. Consider Interface{} for Simple Cases

```go
// For simple delegation, interface{} might be simpler
func LogAny(v interface{}) {
    fmt.Println(v)
}

// Generics add value when type matters
func ProcessTyped[T any](v T) T {
    // Type information preserved
    return v
}
```

### 7. Use Built-in Constraints

```go
// Good: Use standard constraints
import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// Avoid: Reinventing constraints
type MyOrdered interface {
    int | int64 | float64 | string  // Missing many types
}
```

---

## Performance Considerations

### 1. Compile-Time Monomorphization

Go generates specialized code for each type used with generics:

```go
func Double[T int | float64](x T) T {
    return x * 2
}

// Compiler generates:
// - Double_int(x int) int       { return x * 2 }
// - Double_float64(x float64) float64 { return x * 2 }
```

**Implications:**
- ✅ No runtime overhead
- ✅ Fast execution
- ⚠️ Larger binary size
- ⚠️ Longer compile times

### 2. Generics vs Interface{}

```go
// Generic: No boxing, type-safe
func GenericSum[T constraints.Integer](nums []T) T {
    var sum T
    for _, n := range nums {
        sum += n
    }
    return sum
}

// Interface{}: Boxing overhead, type assertions
func InterfaceSum(nums []interface{}) int {
    sum := 0
    for _, n := range nums {
        sum += n.(int)  // Type assertion has cost
    }
    return sum
}
```

**Benchmark Results:**
- Generics: Faster, no allocation
- Interface{}: Slower, heap allocations

### 3. When to Avoid Generics

```go
// Avoid: Generic overkill for single use
func AddInts[T int](a, b T) T {
    return a + b
}

// Better: Simple function
func AddInts(a, b int) int {
    return a + b
}

// Good use: Multiple types
func Add[T constraints.Integer | constraints.Float](a, b T) T {
    return a + b
}
```

### 4. Memory Considerations

```go
// Each instantiation creates separate code
type Container[T any] struct {
    value T
}

// Using Container[int], Container[string], Container[float64]
// generates 3 separate type implementations
```

---

## Summary

**Key Takeaways:**

1. **Type Safety**: Generics provide compile-time type checking without sacrificing performance

2. **Code Reuse**: Write once, use with multiple types

3. **No Runtime Cost**: Monomorphization means no performance penalty

4. **Constraints**: Control what types can be used and what operations are allowed

5. **Type Inference**: Often don't need to specify types explicitly

6. **Common Use Cases**:
   - Container types (Stack, Queue, Optional)
   - Data structure algorithms (Map, Filter, Reduce)
   - Type-safe repositories and patterns
   - Generic utilities

7. **Limitations**:
   - No method type parameters independent of receiver
   - Type switches don't work directly
   - Can't create instances without factories
   - Some inference limitations

8. **Best Practices**:
   - Use when type safety matters
   - Keep constraints simple
   - Use descriptive names
   - Don't overuse
   - Document well

9. **Performance**:
   - No runtime overhead
   - Monomorphization at compile time
   - Larger binaries with many instantiations
   - Faster than interface{} approach

10. **When to Use**:
    - ✅ Container types
    - ✅ Algorithms on collections
    - ✅ Type-safe APIs
    - ✅ Reducing code duplication
    - ❌ Single-type functions
    - ❌ Simple delegations
    - ❌ When interface{} is clearer

Generics are a powerful addition to Go that enable writing flexible, type-safe, and reusable code. Use them wisely to create better abstractions while maintaining Go's simplicity!

---

## Further Reading

- [Go Blog: An Introduction to Generics](https://go.dev/blog/intro-generics)
- [Go Blog: When To Use Generics](https://go.dev/blog/when-generics)
- [Go Type Parameters Proposal](https://go.googlesource.com/proposal/+/HEAD/design/43651-type-parameters.md)
- [Go 1.18 Release Notes](https://go.dev/doc/go1.18)
- [Tutorial: Getting Started with Generics](https://go.dev/doc/tutorial/generics)

## Practice Exercises

Try implementing these to master generics:

1. **Generic LRU Cache**: Build a Least Recently Used cache with generic key and value types
2. **Generic Tree**: Implement a generic binary search tree with traversal methods
3. **Functional Utilities**: Create Map, Filter, Reduce, FlatMap for slices
4. **Generic Validation**: Build a generic validation framework with chainable validators
5. **Type-Safe Builder**: Create a type-safe builder pattern using generics

