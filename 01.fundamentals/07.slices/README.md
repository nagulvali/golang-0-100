## Go Slices

Slices are dynamic, flexible views over arrays. They are the most common sequence type in Go. A slice has three components: pointer to an underlying array, length (number of elements), and capacity (maximum elements it can grow to before reallocating).

### Zero value and basics
```go
var nums []int      // zero value is nil
fmt.Println(nums)   // []
fmt.Println(nums == nil) // true
fmt.Println(len(nums))   // 0

// Literal initialization
nums = []int{1, 2, 3, 4, 5}
fmt.Println(nums)   // [1 2 3 4 5]
fmt.Println(len(nums))   // 5
```

### Creating with make
`make` allocates a slice with specified length and optional capacity.
```go
// make(type, length) -> length == capacity
nums1 := make([]int, 2)     // [0 0], len=2 cap=2

// make(type, length, capacity)
nums2 := make([]int, 2, 8)  // [0 0], len=2 cap=8
```

### Length vs capacity
- **len(s)**: number of accessible elements
- **cap(s)**: number of elements until reallocation is needed

### Appending and growth
`append` returns a new slice (may share or allocate a new array). Always use the returned value.
```go
s := []int{1, 2}
s = append(s, 3)           // [1 2 3]
more := []int{4, 5}
s = append(s, more...)     // [1 2 3 4 5]

// Capacity growth is implementation-dependent (often doubles),
// but you should never rely on an exact formula.
```

### Copying elements
`copy(dst, src)` copies `min(len(dst), len(src))` elements.
```go
src := []int{1, 2, 3}
dst := make([]int, 2)
n := copy(dst, src) // n == 2, dst = [1 2]
```

### Slicing (sub-slices)
```go
base := []int{0, 1, 2, 3, 4, 5}
a := base[1:4]   // elements at indexes 1..3 -> [1 2 3]
b := base[:3]    // [0 1 2]
c := base[3:]    // [3 4 5]
d := base[:]     // full slice
```

All sub-slices share the same underlying array; mutating one can affect the others if they overlap.
```go
a[0] = 99
fmt.Println(base) // [0 99 2 3 4 5]
```

### Full slice expression (three-index)
Control capacity of a sub-slice to prevent accidental growth into later elements.
```go
base := []int{0, 1, 2, 3, 4, 5}
sub := base[1:3:3] // len=2, cap=2 (cannot grow into base[3])
sub = append(sub, 9) // allocates a new backing array
```

### Nil vs empty slices
- `nil` slice: `var s []T` → `s == nil`, `len(s) == 0`, `cap(s) == 0`
- empty but non-nil: `[]T{}` or `make([]T, 0)` → `s != nil`, `len(s) == 0`

Both are interchangeable in most APIs; prefer returning `nil` when nothing to return, or follow your project's conventions.

### Iteration
```go
s := []string{"go", "is", "fun"}
for i := 0; i < len(s); i++ {
    fmt.Println(i, s[i])
}
for i, v := range s { // v is a copy of the element
    fmt.Println(i, v)
}
```

### Common idioms
Insert at index:
```go
// insert x at i
s = append(s[:i], append([]T{x}, s[i:]...)...)
```

Delete at index (preserve order):
```go
s = append(s[:i], s[i+1:]...)
```

Delete at index (do not preserve order):
```go
s[i] = s[len(s)-1]
s = s[:len(s)-1]
```

Clear slice (keep capacity):
```go
for i := range s {
    s[i] = zeroValue // e.g., 0, "", nil
}
s = s[:0]
```

Copy to new storage (detach from large backing array):
```go
t := append([]T(nil), s...) // or make([]T, len(s)); copy(t, s)
```

### Performance tips
- Pre-size with `make([]T, 0, n)` when you know approximate growth to minimize reallocations.
- Use the three-index slice to cap capacity when creating sub-slices that should not grow into the original.
- Avoid holding references to large arrays via tiny sub-slices; copy what you need.

### Arrays vs slices (quick recap)
- **Arrays**: fixed length, value semantics, rarely exposed directly.
- **Slices**: dynamic length, reference semantics to an underlying array, dominant in Go APIs.

### Example tying it together
```go
package main

import "fmt"

func main() {
    var nums []int
    fmt.Println(nums == nil, len(nums))

    nums = []int{1, 2, 3}
    nums = append(nums, 4, 5)

    head := nums[:2]      // share backing array
    tail := nums[2:4:4]   // capped; cannot grow into nums[4]
    tail = append(tail, 99) // allocates

    copyOf := append([]int(nil), nums...)
    fmt.Println(head, tail, copyOf)
}
```

