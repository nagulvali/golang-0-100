## Go Operators & Type Conversions — Practice Problems

Build strong intuition for Go's operators, precedence, and type conversions during arithmetic and other operations. Unless specified, write idiomatic, safe Go. Prefer functions that return results rather than printing.

Conventions:
- Avoid panics for invalid input unless the task explicitly requests it.
- Demonstrate awareness of integer overflow, division behavior, and float pitfalls.
- Prefer explicit conversions where needed; don't rely on implicit widening (Go does not allow it).

### Arithmetic and Basic Conversions
1) SumInts: return the sum of `[]int` using `+` in a loop.
2) AverageInt: average of `[]int` as `float64` (convert to avoid integer division).
3) IntDiv: given `a, b int`, return integer quotient and remainder; handle `b == 0`.
4) FloatDiv: given `a, b int`, return `float64(a)/float64(b)`; handle `b == 0` by returning `math.Inf(+1)`.
5) ModNegatives: explore `%` with negatives. Implement `SafeMod(a, b int) (int, bool)` returning `(r, ok)`; `ok=false` if `b==0`.
6) PowerInt: compute `a^n` for non-negative `n` using fast exponentiation; beware overflow.
7) Distance2D: given `(x1,y1),(x2,y2) int`, return Euclidean distance as `float64`.
8) MeanAndStdDev: compute mean and standard deviation for `[]float64`; handle empty input.
9) FahrenheitToCelsius: conversion with `float64` math; format to 2 decimals as string.
10) RoundHalfAwayFromZero: implement rounding for `float64` to nearest integer as `int`.

### Operator Precedence and Parentheses
11) Eval1: compute `a + b*c - d/e` with correct precedence; verify with parentheses.
12) Eval2: rewrite `a - b - c` two ways to show non-associativity of subtraction.
13) PrecedenceQuiz: given `a=2, b=3, c=4`, evaluate `a+b<<c` and explain; then rewrite with explicit parentheses to change meaning.

### Increment/Decrement and Assignment Ops
14) Counter: given `n`, return slice `[0,1,2,...,n-1]` using `i++`.
15) CompoundAssign: demonstrate `+=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=` on a variable; return final value.
16) BalancedParenthesesScore: compute a score using `++/--` like counters for parentheses balance; invalid if negative.

### Comparisons and Booleans
17) NearlyEqualFloat: compare two floats with epsilon; return bool.
18) ClampInt: ensure `x` in `[lo, hi]` using comparisons; return clamped value.
19) IsBetween: inclusive range check `lo <= x && x <= hi`.
20) ThreeWayCompare: return -1, 0, 1 for `a < b`, `a == b`, `a > b`.
21) IsLeapYear: boolean logic with divisibility rules.
22) TruthTable: generate truth table for `&&, ||, !` for all combinations of two booleans.

### Bitwise Operations (on unsigned and signed with care)
23) BitGet: return whether bit k in `uint` x is set.
24) BitSet: set bit k; BitClear: clear bit k; BitToggle: toggle bit k.
25) CountBits: population count (Hamming weight) for `uint`.
26) IsPowerOfTwo: bit trick using `x & (x-1)`.
27) Parity: return 0/1 for even/odd number of set bits.
28) ByteSwap32: swap endianness of a `uint32` using shifts and masks.
29) RotateLeft/RotateRight for `uint32` using shifts and ors.
30) SignBit: for `int`, return true if negative using comparison; for `int32`, use bit check.

### Shifts and Edge Cases
31) SafeShiftLeft: shift `uint` left by k with guard against k >= word size; return `(v, ok)`.
32) ArithmeticVsLogical: show how right-shift behaves for signed vs unsigned; write small examples and document.
33) ScaleByPowerOfTwo: multiply/divide by `2^k` using shifts; verify with normal arithmetic.

### Type Conversion Between Numerics
34) IntToFloatAccurate: convert `int64` to `float64`; detect precision loss beyond 53 bits and return a flag.
35) FloatToIntClamped: convert `float64` to `int64` with clamping to min/max on overflow; handle NaN and +/-Inf.
36) MixedArithmetic: given `a int`, `b float64`, `c int8`, compute `a + int(b) - int(c)` as `int`; discuss precision trade-offs.
37) PercentOf: compute `(part/total)*100` with `float64` and return as `string` with `%.1f%%`.
38) AverageMixed: average of `[]int` and `[]float64` combined; result as `float64`.

### Bytes, Runes, and Strings
39) RuneDistance: given two `rune`, return numeric distance `int` after converting to `int32`.
40) ToUpperASCII: convert lowercase ASCII byte slice to uppercase in-place using bit trick (`b &^ 0x20`).
41) IsDigitASCII: check if `b` in '0'..'9' using comparisons; return bool.
42) HexValue: convert hex digit '0'..'9','a'..'f','A'..'F' to its integer value; else -1.
43) StringToBytesZeroCopyNote: implement safe `[]byte(s)` copy; add comment explaining why unsafe zero-copy is discouraged.
44) BytesToString: convert `[]byte` to `string` and explain allocation semantics.

### Constants and Untyped Constants
45) UntypedConstMix: show that `const x = 1` can be used as `int8`, `int`, `float64` in different contexts without explicit conversion; write three short uses.
46) IotaBitFlags: define bit flags with `iota` and test composition with bitwise ops.
47) TypedConstOverflow: illustrate that `const y int8 = 128` is a compile error; adjust to valid max.

### Division, Modulo, and Rounding
48) FloorDiv: implement floor division for integers that works with negatives (result <= true quotient); return `(q, r)` where `a = q*b + r` and `0 <= r < |b|`.
49) CeilDivPositive: for positive integers only, compute ceiling division `ceil(a/b)` without floats.
50) BankersRounding: implement round-half-to-even for `float64`.

### Operator Practice Mini-Katas
51) DotProduct: compute dot product of two equal-length `[]int`; result `int`.
52) NormalizeVector: given `[]float64`, return unit-vector; handle zero-vector.
53) ManhattanDistance: sum of absolute differences between two `[]int`.
54) ScalarAffine: compute `a*x + b` element-wise for `[]float64`.
55) MinMaxNormalize: transform `[]float64` to [0,1] using min-max scaling; constant slice yields zeros.

### Overflow and Limits
56) SafeAddInt32: add two `int32` with overflow detection; return `(sum, ok)`.
57) SafeMulUint64: multiply two `uint64` with overflow detection via division check; `(prod, ok)`.
58) AccumulateInt64: sum `[]int64` with overflow guard; stop and return error on overflow.

### Short-Circuiting and Side Effects
59) ShortCircuitEval: demonstrate that in `a != 0 && b/a > 1`, division is safe due to short-circuiting; write test-like function returning bool.
60) CoalesceInt: return first non-zero from `(a, b, c int)` using short-circuit-like logic with conditionals.

### Precedence Puzzles (Explain in comments)
61) Puzzle1: evaluate `a | b & c ^ d` under Go precedence rules; then compute the parenthesized variant `(a | b) & (c ^ d)` and compare.
62) Puzzle2: for `x := 1; y := x<<2 + 3`, show exact order of operations and resulting value.

### Conversions and APIs
63) ParseIntBase: parse a string in base 2..36 to `int64` using `strconv.ParseInt`; return error on invalid.
64) FormatIntBase: format an `int64` to a given base using `strconv.FormatInt`.
65) TimeDurationMath: given seconds as `int64`, convert to `time.Duration`, add 500ms, and return as `string` via `String()`.

### Mixed-Type Expression Drills
66) MixedExpr1: given `a int8`, `b uint16`, compute `int(b) - int(a)`; return `int`.
67) MixedExpr2: given `a uint`, `b int`, safely compute `a - uint(max(0, b))` guarding underflow; return `uint`.
68) MixedExpr3: given `x float32`, `y float64`, compute `float64(x)/y` as `float64`.

### Optional (Advanced)
69) ComplexArithmetic: implement addition, multiplication for `complex128`; return magnitude and phase (use `cmplx`).
70) IEEE754Inspect: decompose a `float64` into sign, exponent, mantissa using `math.Float64bits` and bit operations.

---

Hints and Notes:
- Go does not perform implicit numeric widening; always convert explicitly.
- Integer division truncates toward zero; `%` keeps the sign of the dividend.
- Use `math` and `strconv` packages for numeric helpers and conversions.
- Floating-point comparisons should use an epsilon; beware NaN (NaN != NaN).
- Bitwise operations are defined on unsigned and signed types, but shifting negative values requires care; prefer unsigned for bit ops.
- For performance-critical code, avoid unnecessary conversions in tight loops.


