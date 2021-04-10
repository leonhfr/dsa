# 2 Asymptotic Notation

## Algorithmic Time Complexity

### The RAM Model of Computation

- algorithms are an important and durable part of computer science because they can be studied in a machine/language independent way
- this is because we use the RAM model of computation for all our analysis
  - each simple operation (+, -, =, if, call) takes 1 step
  - loops and subroutines calls are not simple operations; they depend upon the size of the data and the contents of a subroutine; "sort" is not a single step operation
  - each memory access takes exactly 1 step
- we measure the run time of an algorithm by counting the number of steps
- this model is useful and accurate in the same sense as the flat-earth model (which is useful)

### Worst-Case Complexity

- the worst case complexity of an algorithm is the function defined by the maximum number of steps taken on any instance of size n

### Best-Case and Average-Case Complexity

- the best case complexity of an algorithm is the function defined by the minimum number of steps taken on any instance of size n
- the average case complexity if an algorithm is the function defined by an average number of steps taken on any instance of size n
- each of these complexities defines a numerical function: time vs. size

### Complexity Analysis

- generally, we use the worst case complexity as our preferred measure of algorithm efficiency
- worst case analysis is generally easy to do, and usually reflects the average case
- randomized algorithms are of growing importance, and require an average-case type analysis to show off their merits

## The Big Oh Notation

### The Big Oh Notation

- best, worst, and average case are difficult to deal with because the precise function details are very complicated
- it's easier to talk about upper and lower bounds of the function
- asymptotic notation (O, Ω, Θ) are as well as we can practically deal with complexity functions

### Names of Bounding Functions

- g(n) = O(f(n)) means C * f(n) is an upper bound on g(n)
- g(n) = Ω(f(n)) means C * f(n) is a lower bound on g(n)
- g(n) = Θ(f(n)) means C1 * f(n) is an upper bound on g(n) and C2 * f(n) is a lower bound on g(n)
- C, C1, and C2 are all constants independent of n
- the definitions imply a constant n0 beyond which they are satisfied
- we do not care about small values of n

### Formal definitions

- f(n) = O(g(n)) if there are positive constants n0 and c such that to the right of n0, the value of f(n) always lies on or below c * g(n)
- f(n) = Ω(g(n)) if there are positive constants n0 and c such that to the right of n0, the value of f(n) always lies on or above c * g(n)
- f(n) = Θ(g(n)) if there are positive constants n0, c1, and c2 such that to the right of n0, the value of f(n) always lies between c1 * g(n) and c2 * g(n) inclusive

### Big Oh Addition/Subtraction

- standard rules