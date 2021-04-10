# Introduction to Algorithms

## What is an algorithm?

- algorithms are the ideas behind computer programs
- procedure, way of solving a general problem
- defining the algorithmic problem: describing the inputs the problem has, and describing what desired properties the output must have
- we seek algorithms which are _correct_ and _efficient_

## Efficiency

- a faster algorithm running on a slower computer will _always_ win for sufficiently large instances
- usually, problems don't have to get that large before the faster algorithm wins

## Correctness

- for any algorithm, we must prove that it _always_ returns the desired output for all legal instances of the problem
- for sorting, this means even if (1) the input is already sorted, or (2) it contains repeated elements
- algorithm correctness is not obvious in many optimization problems
- algorithm problems must be carefully specified to allow a provably correct algorithm to exist; we can find the "shortest your", but nor the "best tour"

## Expressing algorithms

- we need some way to express the sequence of steps comprising an algorithm
- in order of increasing precision, we have English, pseudocode, and real programming languages; unfortunately, ease of expression moves in the reverse order
- describe the ideas of an algorithm in ENglish, moving to pseudocode to clarify sufficiently tricky details of the algorithm

## Demonstrating incorrectness

- searching for counterexamples is the best way to disprove the correctness of a heuristic
- think about all small examples
- think about examples with ties on your decision criteria
- this about examples with extremes of big and small

## Induction and recursion

- failure to find a counterexample to a given algorithm does not mean "it is obvious" that the algorithm is correct
- mathematical induction is a very useful method for proving the correctness of recursive algorithms
- recursion and induction are the same basic idea: (1) basis case, (2) general assumption, (3) general case