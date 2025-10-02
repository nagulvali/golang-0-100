## Go Slices — Practice Problems

A curated set of problems to build intuition and fluency with Go slices. Implement solutions in Go. Prefer in-place operations when appropriate, and pay attention to length vs capacity, allocations, and backing arrays.

Conventions:
- Unless specified, functions should not panic on invalid input; return a sensible result (or error) instead.
- Aim for O(n) time for linear scans, O(1) extra space when feasible.

### Basics and Construction
1) Initialize slice with N copies of value V (no loops in literal).
2) Return a slice of the first N natural numbers starting from 1.
3) Clone a slice without using `append(s, ...)` on itself.
4) Convert an array `[N]int` to `[]int` and back (copying required for back conversion).
5) Check if a slice is nil vs empty; normalize input so you always return an empty, non-nil slice.
6) Ensure capacity: given `[]int` and target capacity C, return a slice that shares data if possible; otherwise, reallocate.

### Indexing and Slicing
7) Safe sub-slice: implement `Sub(s []int, i, j int) []int` that clamps i, j into [0..len(s)] and returns a valid sub-slice.
8) Rightmost K elements: `LastK(s []T, k int) []T` without allocation; use full slice expression to cap capacity.
9) Middle window: return a view of length W centered around index I, clamped to bounds.
10) Split halves: split a slice into two nearly equal halves (left gets the extra when odd).

### Reading and Aggregation
11) Sum, min, max of `[]int`. Return zero values for empty input.
12) Prefix sums: return a slice p where `p[i] = sum(s[:i+1])`.
13) Count occurrences of a value V in `[]int` and first index of V (or -1).
14) Distinct elements preserving first occurrence order.
15) Frequency map of `[]string` values.

### In-place Mutation
16) Reverse a slice in-place.
17) Rotate right by K in-place (with O(1) extra space). Hint: triple-reverse.
18) Remove element at index I (preserve order). Return result slice.
19) Remove element at index I (do not preserve order). Return result slice.
20) Remove all occurrences of V in-place, returning the truncated slice (two-pointer write index).
21) Deduplicate consecutive duplicates in-place (like UNIX uniq).
22) Stable unique in-place with O(n) time and O(n) extra space allowed (map-based), returning truncated slice.

### Appending and Growth
23) Append many: append contents of `[][]int` into a single `[]int` with one allocation pre-sized by total length.
24) Bounded append: append up to K items from `src` to `dst` without reallocating `dst` (if insufficient cap, return `dst` unchanged and a bool).
25) Ensure no aliasing: return a copy of `s` that cannot affect future appends to `base` when appended to `base`.

### Copying and Detaching
26) Copy a sub-range `[i:j)` into a new slice `t` and return `t`.
27) Detach from large backing array: given a tiny sub-slice of a huge array-backed slice, return a compact copy.
28) Merge two sorted slices into a new sorted slice; stable.

### Searching and Two-Pointer Patterns
29) Binary search for target in a sorted `[]int` (return index or -1).
30) Lower bound and upper bound indices for target in sorted `[]int`.
31) Two-sum: given `[]int` and target, return indices of two numbers (any order). Provide O(n) hash-map solution.
32) Move zeros to the end in-place, preserving non-zero order.

### Partitioning and Filtering
33) Stable partition: move all elements satisfying predicate to front, preserving relative order; return split index.
34) Unstable partition: partition by predicate in-place with O(1) extra space, order not preserved; return split index.
35) Filter with allocation budget: filter `s` by predicate using at most one new allocation.

### Windows and Subarrays
36) Maximum subarray sum (Kadane). Return sum and `[i:j)` sub-slice indices.
37) Longest subarray with sum <= K (non-negative numbers) using sliding window; return indices.
38) All sub-slices of fixed length L (return a slice of views; discuss lifetime and aliasing).

### Strings, Bytes, and Runes
39) Reverse a `[]byte` in-place.
40) Reverse a UTF-8 string properly by converting to `[]rune` and back.
41) Deduplicate adjacent whitespace in a byte slice in-place, returning the truncated slice.
42) Tokenize a byte slice by a delimiter (single byte), returning sub-slice views; avoid allocations.

### Matrix (Slices of Slices)
43) Create an `R x C` matrix of `int` with independent rows (no shared backing rows).
44) Transpose a dense `[][]int`. Allocate a new matrix.
45) In-place transpose of a square matrix represented as `[][]int` with contiguous backing (first build representation), then transpose.
46) Spiral traversal of a 2D slice (return order of elements).

### Sorting and Rearrangement
47) Stable sort `[]struct{K int; V string}` by K then V using `slices.SortFunc` (Go 1.21+), or `sort.SliceStable` otherwise.
48) Group anagrams: given `[]string`, group words that are anagrams into `[][]string`.
49) Three-way partition (Dutch National Flag) for values -1, 0, 1 in-place.

### Concurrency-Safe Patterns (Read-Only Sharing)
50) Safely publish a large read-only slice to multiple goroutines: expose copies vs views; write an API that returns an immutable view (document contracts).
51) Lock-free snapshot: atomically swap a pointer to a slice used for reads while writers prepare a new slice (copy-on-write pattern). Implement the swap function.

### Memory and Capacity
52) Trim capacity: return `t := append([]T(nil), s...)` and explain why this detaches from future appends.
53) Cap a sub-slice using the three-index full slice expression to prevent append bleed-through into the original.
54) Grow strategy experiment: append 1..N into a slice, record capacity growth steps, and summarize the pattern.

### Error Handling and Validation
55) Validate indices `(i, j)` for slicing; return `(sub, ok)` without panicking.
56) Safe pop-front and pop-back for `[]T` returning `(value, rest, ok)`.

### Higher-Level Utilities
57) Map: apply `func(T) U` to `[]T` producing `[]U` with one allocation.
58) Reduce: fold `[]T` with `func(acc U, t T) U`.
59) Zip two slices into `[]struct{A, B}`; stop at min length.
60) Chunk a slice into fixed-size groups (last chunk may be smaller).

---

Hints and Notes:
- Prefer two-index slicing for views and three-index slicing to control capacity when you will append.
- When removing elements in-place, use a write index to avoid extra allocations.
- Be mindful of aliasing when returning views over shared backing arrays; document contracts.
- For stable operations that shrink the slice, prefer returning `s[:n]` rather than allocating a new slice.
- Use `copy` for overlapping ranges during insert/delete to prevent data clobbering.


