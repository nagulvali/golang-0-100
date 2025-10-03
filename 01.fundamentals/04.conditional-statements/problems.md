# Conditional Statements Practice Problems

This document contains practice problems for Go conditional statements (if/else, switch, select) to help you build logic and problem-solving skills. Problems are organized by difficulty level and topic.

## Table of Contents

1. [Basic If/Else Problems](#basic-ifelse-problems)
2. [Switch Statement Problems](#switch-statement-problems)
3. [Nested Conditional Problems](#nested-conditional-problems)
4. [Logical Operator Problems](#logical-operator-problems)
5. [Type Assertion Problems](#type-assertion-problems)
6. [Select Statement Problems](#select-statement-problems)
7. [Real-World Application Problems](#real-world-application-problems)
8. [Advanced Logic Problems](#advanced-logic-problems)
9. [Error Handling Problems](#error-handling-problems)
10. [Performance and Optimization Problems](#performance-and-optimization-problems)

---

## Basic If/Else Problems

### Problem 1: Number Classification
**Difficulty**: Easy
**Description**: Write a function that classifies a number as positive, negative, or zero.

### Problem 2: Age Category
**Difficulty**: Easy
**Description**: Determine age category based on age. Categories: "child" (0-12), "teen" (13-19), "adult" (20-64), "senior" (65+).

### Problem 3: Grade Calculator
**Difficulty**: Easy
**Description**: Convert numeric grade to letter grade. A: 90-100, B: 80-89, C: 70-79, D: 60-69, F: 0-59.

### Problem 4: Leap Year Checker
**Difficulty**: Easy
**Description**: Determine if a year is a leap year. Rules: divisible by 4, if divisible by 100 must also be divisible by 400.

### Problem 5: Triangle Type
**Difficulty**: Easy
**Description**: Determine the type of triangle based on side lengths. Types: "equilateral", "isosceles", "scalene", "invalid".

### Problem 6: Password Strength
**Difficulty**: Easy
**Description**: Check password strength based on length. "weak" (< 6 chars), "medium" (6-10 chars), "strong" (> 10 chars).

### Problem 7: Temperature Converter
**Difficulty**: Easy
**Description**: Convert between Celsius and Fahrenheit with direction parameter.

### Problem 8: Day of Week
**Difficulty**: Easy
**Description**: Get day name from day number (1-7). 1=Monday, 2=Tuesday, ..., 7=Sunday.

---

## Switch Statement Problems

### Problem 9: Calculator Operations
**Difficulty**: Easy
**Description**: Implement basic calculator operations using switch. Operations: "add", "subtract", "multiply", "divide".

### Problem 10: Month Days
**Difficulty**: Easy
**Description**: Get number of days in a month. Consider leap years for February.

### Problem 11: HTTP Status Code
**Difficulty**: Easy
**Description**: Get HTTP status message from status code. 200="OK", 404="Not Found", 500="Internal Server Error", etc.

### Problem 12: File Extension
**Difficulty**: Easy
**Description**: Get file type from file extension. ".txt"="text", ".jpg"="image", ".mp3"="audio", etc.

### Problem 13: Season Detector
**Difficulty**: Easy
**Description**: Determine season from month number. 12,1,2="winter", 3,4,5="spring", 6,7,8="summer", 9,10,11="autumn".

### Problem 14: Grade Point Average
**Difficulty**: Easy
**Description**: Convert letter grade to GPA points. A=4.0, B=3.0, C=2.0, D=1.0, F=0.0.

### Problem 15: Direction from Angle
**Difficulty**: Easy
**Description**: Get compass direction from angle in degrees. 0="North", 90="East", 180="South", 270="West".

---

## Nested Conditional Problems

### Problem 16: Complex Age Category
**Difficulty**: Medium
**Description**: Determine detailed age category with subcategories. "infant" (0-2), "toddler" (3-5), "child" (6-12), "teen" (13-19), "young adult" (20-35), "middle age" (36-64), "senior" (65+).

### Problem 17: Loan Eligibility
**Difficulty**: Medium
**Description**: Check loan eligibility based on multiple criteria. Age >= 18, income >= 30000, creditScore >= 650, employmentYears >= 2.

### Problem 18: Student Scholarship
**Difficulty**: Medium
**Description**: Determine scholarship amount based on multiple factors. GPA >= 3.5, income < 50000, extracurriculars >= 3.

### Problem 19: Weather Advisory
**Difficulty**: Medium
**Description**: Generate weather advisory based on multiple conditions. Consider heat index, wind chill, storm conditions.

### Problem 20: Restaurant Rating
**Difficulty**: Medium
**Description**: Rate restaurant based on multiple criteria. Each criteria 1-10, return "excellent", "good", "average", "poor".

---

## Logical Operator Problems

### Problem 21: Complex Password Validation
**Difficulty**: Medium
**Description**: Validate password with multiple requirements. Length >= 8, has uppercase, has lowercase, has digit, has special char.

### Problem 22: Access Control
**Difficulty**: Medium
**Description**: Check user access based on role and permissions. Roles: "admin", "user", "guest". Resources: "data", "settings", "reports". Actions: "read", "write", "delete".

### Problem 23: Discount Calculator
**Difficulty**: Medium
**Description**: Calculate discount based on multiple conditions. Member: 10%, Weekend: 5%, Holiday: 15%. Discounts can be combined.

### Problem 24: Game Score Multiplier
**Difficulty**: Medium
**Description**: Calculate score multiplier based on game conditions. Combo: 1.5x, Perfect: 2x, Streak: 1.2x, Level bonus.

### Problem 25: Insurance Premium
**Difficulty**: Medium
**Description**: Calculate insurance premium based on risk factors. Age < 25: +20%, Accidents: +30%, Violations: +15%, Credit < 600: +25%.

---

## Type Assertion Problems

### Problem 26: Type Checker
**Difficulty**: Medium
**Description**: Check and handle different types using type assertions. Handle int, string, bool, float64 types.

### Problem 27: Dynamic Calculator
**Difficulty**: Medium
**Description**: Perform operations on different numeric types. Handle int, float64, string (convert to number).

### Problem 28: JSON Type Handler
**Difficulty**: Medium
**Description**: Handle different JSON value types. Handle map, slice, string, number, bool, nil.

### Problem 29: Database Type Converter
**Difficulty**: Medium
**Description**: Convert database values to appropriate Go types. Target types: "int", "string", "bool", "float64".

### Problem 30: API Response Handler
**Difficulty**: Medium
**Description**: Handle different API response types. Handle success response, error response, data response.

---

## Select Statement Problems

### Problem 31: Channel Timeout
**Difficulty**: Medium
**Description**: Implement timeout for channel operations. Return value from channel or timeout error.

### Problem 32: Multiple Channel Reader
**Difficulty**: Medium
**Description**: Read from multiple channels with priority. Read from channels in priority order: ch1 > ch2 > ch3.

### Problem 33: Channel Selector
**Difficulty**: Medium
**Description**: Select from channels based on conditions. If condition is true, prefer ch1, else ch2.

### Problem 34: Non-blocking Channel Read
**Difficulty**: Medium
**Description**: Implement non-blocking channel read with default case. Return (value, true) if available, (empty, false) if not.

### Problem 35: Channel Multiplexer
**Difficulty**: Medium
**Description**: Combine multiple channels into one. Return channel that receives from both input channels.

---

## Real-World Application Problems

### Problem 36: E-commerce Order Processing
**Difficulty**: Hard
**Description**: Process e-commerce order with multiple validation steps. Validate: user exists, items in stock, payment method valid, address complete.

### Problem 37: User Authentication System
**Difficulty**: Hard
**Description**: Implement user authentication with multiple checks. Check: user exists, password correct, account active, not locked.

### Problem 38: Inventory Management
**Difficulty**: Hard
**Description**: Manage inventory with reorder points and alerts. Check each product: "in stock", "low stock", "out of stock", "overstocked".

### Problem 39: Banking Transaction Validator
**Difficulty**: Hard
**Description**: Validate banking transactions with multiple rules. Validate: sufficient funds, account type, transaction limits, time restrictions.

### Problem 40: Weather Station Data Processor
**Difficulty**: Hard
**Description**: Process weather station data with multiple sensors. Calculate: heat index, wind chill, weather conditions, alerts.

---

## Advanced Logic Problems

### Problem 41: Complex Decision Tree
**Difficulty**: Hard
**Description**: Implement a complex decision tree for loan approval. Complex decision tree with multiple conditions.

### Problem 42: Game State Machine
**Difficulty**: Hard
**Description**: Implement a game state machine with multiple states and transitions. Handle state transitions based on current state and input.

### Problem 43: Multi-Criteria Sorting
**Difficulty**: Hard
**Description**: Sort items based on multiple criteria with priority. Sort by multiple criteria in priority order, handle ties with secondary criteria.

### Problem 44: Dynamic Rule Engine
**Difficulty**: Hard
**Description**: Implement a rule engine that can evaluate complex conditions. Evaluate rules against data, return actions to execute.

### Problem 45: Complex Validation System
**Difficulty**: Hard
**Description**: Implement a validation system with multiple validation rules. Validate form data against rules, return field errors.

---

## Error Handling Problems

### Problem 46: Robust File Processor
**Difficulty**: Medium
**Description**: Process files with comprehensive error handling. Handle: file not found, permission denied, invalid format, disk full.

### Problem 47: Network Request Handler
**Difficulty**: Medium
**Description**: Handle network requests with timeout and retry logic. Handle: timeout, connection refused, invalid URL, server error.

### Problem 48: Database Operation Handler
**Difficulty**: Medium
**Description**: Handle database operations with error recovery. Handle: connection lost, syntax error, constraint violation, timeout.

### Problem 49: API Response Handler
**Difficulty**: Medium
**Description**: Handle API responses with different error scenarios. Handle: 4xx errors, 5xx errors, invalid JSON, network errors.

### Problem 50: Configuration Validator
**Difficulty**: Medium
**Description**: Validate configuration with detailed error reporting. Validate each field and return detailed error messages.

---

## Performance and Optimization Problems

### Problem 51: Efficient String Matching
**Difficulty**: Hard
**Description**: Implement efficient string matching with multiple patterns. Match text against multiple patterns efficiently.

### Problem 52: Cache Hit/Miss Optimizer
**Difficulty**: Hard
**Description**: Optimize cache operations with intelligent eviction. Implement LRU, LFU, or hybrid eviction strategy.

### Problem 53: Conditional Branch Optimizer
**Difficulty**: Hard
**Description**: Optimize conditional statements for better performance. Optimize conditional checks for large datasets.

### Problem 54: Memory-Efficient Data Processing
**Difficulty**: Hard
**Description**: Process large datasets with memory constraints. Process data stream efficiently without loading all into memory.

### Problem 55: Smart Retry Logic
**Difficulty**: Hard
**Description**: Implement intelligent retry logic with adaptive strategies. Implement exponential backoff, jitter, and circuit breaker.

---

## Bonus Challenges

### Problem 56: State Machine with Conditions
**Difficulty**: Expert
**Description**: Implement a complex state machine with conditional transitions. Update state machine, check transitions, execute actions.

### Problem 57: Rule-Based Expert System
**Difficulty**: Expert
**Description**: Implement a rule-based expert system for decision making. Implement forward chaining inference engine.

### Problem 58: Multi-Agent Decision System
**Difficulty**: Expert
**Description**: Implement a multi-agent system with distributed decision making. Implement decision making based on beliefs, goals, and messages.

### Problem 59: Adaptive Learning System
**Difficulty**: Expert
**Description**: Implement a system that learns and adapts its behavior. Implement learning algorithm that updates knowledge and rules.

### Problem 60: Quantum-Inspired Decision Tree
**Difficulty**: Expert
**Description**: Implement a decision tree with quantum-inspired superposition states. Implement quantum-inspired decision making.

---

## Tips for Solving These Problems

1. **Start Simple**: Begin with basic if/else problems and gradually move to complex ones
2. **Test Thoroughly**: Write test cases for edge cases and boundary conditions
3. **Consider Performance**: Think about time and space complexity
4. **Handle Errors**: Always consider error conditions and edge cases
5. **Use Go Idioms**: Leverage Go's features like multiple return values and error handling
6. **Practice Regularly**: Solve problems daily to build muscle memory
7. **Read Others' Solutions**: Learn from different approaches and implementations
8. **Focus on Logic**: Understand the problem before writing code
9. **Refactor**: Improve your solutions after getting them working
10. **Document**: Add comments explaining complex logic

## Additional Resources

- [Go Tour - Flow Control](https://tour.golang.org/flowcontrol/1)
- [Effective Go - Control Structures](https://golang.org/doc/effective_go.html#control-structures)
- [Go by Example - If/Else](https://gobyexample.com/if-else)
- [Go by Example - Switch](https://gobyexample.com/switch)
- [Go by Example - Select](https://gobyexample.com/select)

Happy coding and good luck with your practice! 🚀
