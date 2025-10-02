## Go Arrays

Arrays in Go are fixed-length, contiguous sequences of elements of the same type. Unlike slices, an array's length is part of its type and cannot change after declaration. Arrays have value semantics (they copy on assignment and when passed to functions), which is an important difference from slices and maps.

### Key properties
- **Fixed length**: length is compile-time constant and part of the type (e.g., `[3]int`).
- **Value semantics**: assignment copies all elements; functions receive a full copy unless you pass a pointer.
- **Zero values**: arrays are zero-initialized; each element is the type's zero value.
- **Indexing**: 0-based indexing; out-of-range indexes panic at runtime.

### Declaring and initializing
```go
// Declaration with zero-values
var a [3]int           // [0 0 0]
var b [2]string        // ["" ""]

// Literal initialization
c := [3]int{1, 2, 3}
d := [3]int{1}         // [1 0 0]

// Indexed literals (sparse init)
e := [5]int{2: 10, 4: 20} // [0 0 10 0 20]

// Length inferred with ellipsis
f := [...]int{1, 1, 2, 3, 5}

// Arrays are different types if lengths differ
// var g [4]int = f // compile error: [4]int != [5]int
```

### Length and iteration
```go
arr := [...]string{"go", "is", "fun"}
n := len(arr) // 3

// Index-based iteration
for i := 0; i < len(arr); i++ {
    fmt.Println(i, arr[i])
}

// Range iteration (value is a copy of the element)
for i, v := range arr {
    fmt.Println(i, v)
}
```

### Value semantics (copies)
```go
x := [3]int{9, 8, 7}
y := x          // copies all elements
y[0] = 0        // x remains unchanged
fmt.Println(x)  // [9 8 7]
fmt.Println(y)  // [0 8 7]

// To modify the original, pass a pointer
func zeroFirst(a *[3]int) {
    a[0] = 0
}
zeroFirst(&x)
fmt.Println(x) // [0 8 7]
```

### Multidimensional arrays
```go
// 2D array: 3 rows, 2 columns
matrix := [3][2]int{
    {1, 2},
    {3, 4},
    {5, 6},
}

// Iterate
for r := 0; r < len(matrix); r++ {
    for c := 0; c < len(matrix[r]); c++ {
        fmt.Print(matrix[r][c], " ")
    }
    fmt.Println()
}
```

### Comparison and equality
Arrays are comparable if their element types are comparable. Two arrays are equal if they have the same length and all corresponding elements are equal.
```go
a1 := [3]int{1, 2, 3}
a2 := [3]int{1, 2, 3}
a3 := [3]int{1, 2, 4}
fmt.Println(a1 == a2) // true
fmt.Println(a1 == a3) // false
```

### Arrays vs. slices
- **Arrays**: fixed length, value semantics, type includes length; less common directly in APIs.
- **Slices**: dynamic length, reference semantics to an underlying array, much more common in Go code.

Create a slice view over an array (no copy):
```go
base := [...]int{0, 1, 2, 3, 4}
sub := base[1:4]     // slice referencing base: elements 1..3
sub[0] = 99          // modifies base[1]
fmt.Println(base)    // [0 99 2 3 4]

// Full array as slice
all := base[:]       // type: []int
```

Convert slice to array requires copying into a new array (or using array pointers with exact length):
```go
// Copy slice into array with matching length
s := []int{1, 2, 3}
var arr3 [3]int
copy(arr3[:], s)

// Or take a pointer to an array if slice length matches at compile time (Go 1.20+ with unsafe tricks omitted). Generally prefer copy.
```

### When to use arrays
- **Fixed-size data** known at compile time (e.g., byte sequences with constant length, lookup tables).
- **Interop** where a function explicitly requires an array or pointer to array.
- **As backing storage** for slices to avoid heap allocations in tight loops (allocate array on stack, slice it, pass the slice around).

Example of stack-allocating backing storage:
```go
func sum(nums []int) int {
    total := 0
    for _, v := range nums {
        total += v
    }
    return total
}

func main() {
    var buf [8]int        // stack array
    xs := buf[:0]         // slice backed by buf
    xs = append(xs, 1, 2, 3)
    fmt.Println(sum(xs))  // 6
}
```

### Common pitfalls
- **Confusing arrays with slices**: `[N]T` is an array; `[]T` is a slice. Length in brackets means array.
- **Unexpected copies**: assigning arrays or passing them to functions copies all elements; use pointers when mutation is intended.
- **Type mismatches by length**: `[3]int` and `[4]int` are different, not interchangeable.
- **Out-of-range panics**: always ensure indexes are within `0..len(a)-1`.

### Quick reference
```go
var a [3]int            // declare
b := [3]int{1, 2, 3}    // literal
c := [...]int{1, 2, 3}  // inferred length
lenA := len(a)          // length
for i, v := range a { _ = i; _ = v }
_ = b == [3]int{1, 2, 3} // comparable if element type is comparable
```
