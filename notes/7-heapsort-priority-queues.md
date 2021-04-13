# 7 Heapsort / Priority Queues

## Introduction to Sorting

Importance of Sorting

Why don't CS profs ever stop talking about sorting?
- computers spend a lot of time sorting, historically 25% on mainframes
- sorting is the best studied problem in computer science, with many different algorithms known
- most of the interesting ideas we will encounter in the course can be taught in the context of sorting, such as divide and conquer, randomized algorithms, and lower bounds

Efficiency of Sorting
- sorting is important because that once a set of items is sorted, many other problems become easy
- further, using O(n log n) sorting algorithms leads natirally to sub quadratic algorithms for all these problems
- large scale data processing is impossible with omega(n^2) sorting

![](../assets/sorting-efficiency.png)

Pragmatics of Sorting: Comparison Functions
- explicitly controlling the order of keys is the job of the comparison function we apply to each pair of elements, including the question of increasing or decreasing order

Pragmatics of Sorting: Equal Elements
- elements with equal keys will all bunch together in any total order, but sometimes the relative order among these keys matters
- often there are secondary keys to test after the primary keys
- this is a job for the comparison function
- certain algorithms like quicksort require special care to run efficiently large numbers of equal elements

Pragmatics of Sorting: Library Functions
- any reasonable programming language has a built-in sort routine as a library function
- you are almost always better off using the system sort than writing your own routine

## Applications of Sorting

### Searching

- binary search lets you test whether an item is in a dictionary in O(log n) time
- search preprocessing is perhaps the single most important application of sorting

### Closest pair

- given n numbers, find the pair which are closest to each other
- once the numbers are sorted, the closest pair will be next to each other in sorted order, so an O(n) linear scan completes the job

### Element Uniqueness

- given a set of n items, are they all unique or are there any duplicates
- sort them and do a linear scan to check all adjacent pairs
- this is a special case of closest pair above

### Mode

- given a set of n items, which element occurs the largest number of times
- more generally, compute the frequency distribution
- sort them and do a linear scan to measure the length of all adjacent runs
- the number of instances of k in a sorted array can be found in O(log n) time by using binary search to look for the positions of both k - epsilon and k + epsilon

### Median and Selection

- what is the kth largest item in the set?
- once the keys are placed in sorted order in an array, the kth largest can be found in constant time by simply looking in the kth position of the array
- there is a linear time algorithm for this problem, but the idea comes from partial sorting

### Convex hulls

- given n points in two dimensions, find the smallest area polygon which contains them all
- convex hulls are the most important building block for more sophisticated geometric algorithms

Finding Convex Hulls

- once you have the points sorted by x coordinate, they can be inserted from left to right into the hull, since the rightmost point is always on the boundary
- sorting eliminates the need to check whether points are inside the current hull
- adding a new point might cause others to be deleted

## Selection Sort / Heapsort

### Selection Sort

- selection sort scans through the entire array, repeatedly finding the smallest remaining element

> for i = 1 to n
> 
> A: find the smallest of the first n - i + 1 items
> 
> B: pull it out of the array and put it first

- selection sort takes O(n(T(A) + T(B))) time

### The data structure matters

- using arrays or unsorted linked lists as the data structure, operation A takes O(n) time and operation B takes O(1), for an O(n^2) selection sort
- using balanced search trees or heaps, both of these operations can be done within O(log n) time, for an O(n log n) selection sort called heapsort
- balancing the work between the operations achieves a better tradeoff
- key question: can we use a different data structure?

### Priority Queues with Applications

Priority Queues
- priority queues are data structures which provide extra flexibility over sorting
- this is important because jobs often enter a system at arbitrary intervals
- it is more cost effective to insert a new job into a priority queue than to re sort everything on each new interval

Priority Queue Operations
- the basic priority queue supports three primary operations
  - Insert(Q, x, k): given an item x with key l, insert it into the priority queue Q
  - FindMin(Q) or FindMax(Q): returns a pointer to the item whose key is smaller (larger) than any other key in the priority queue Q
  - DelMin(Q) or DelMax(Q): removes the item from the priority queue Q whose key is minimum (maximum)
- each of these operations can be easily supported using heaps or balanced binary trees in O(log n)

## Heaps

Heap Definition
- a binary heap is defined to be a binary tree with a key in each node such that
  - all leaves are on, at most, two adjacent levels
  - all leaves on the lowest level occur to the left, and all levels except the lowest one are completely filled
  - the key in root is <= all its children, and the left and right subtrees are again binary heaps
- conditions 1 and 2 specify shape of the tree, and condition 3 the labeling of the tree
- heaps maintain a partial order on the set of elements which is weaker than the sorted order (so it can be efficient to maintain) yet stronger than random order (so the minimum element can be quickly identified)
- the partial order defined by the heap structure is weaker than sorting, which explains why it is easier to build: linear time if we do it right

![](../assets/binary-heap.png)

Above: a heap-labeled tree of important years with the corresponding implicit heap representation

Heapsort animation

![](../assets/heapsort-example.gif)

Array-Based Heaps

- the most natural representation of this binary tree would involve storing each key in a node with pointers to its two children
- however, we can store a tree as an array of keys, using the position of the leys to implicitly satisfy the role of the pointers
- the left child of k sits in position 2k, and the right child in 2k+1
- the parent of k is in position |n/2|

Can we implicitly represent any binary tree>
- the implicit representation is only efficient if the tree is sparse, meaning that the number of nodes n < 2^h
- all missing internal nodes still take up space in our structure
- this is why we insist on heaps as being as balanced/full at each level as possible
- the array-based representation is also not as flexible to arbitrary modifications as a pointer-based tree

Constructing Heaps
- heaps can be constructed incrementally, by inserting new elements into the left-most open spot in the array
- if the new element is greater than its parent, swap their positions and recur
- since all but the last level is always filled, the height h of an n element heap is bounded because:

![](../assets/constructing-heaps.png)

so h = |log n|

- doing n such insertions really takes theta(n log n), because the last n/2 insertions require O(log n) time each

Heap Insertion

```c++
void pq_insert(priority_queue *q, item_type x) {
  if (q->n >= PQ_SIZE) {
    printf("Warning: priority queue overflow! \n");
  } else {
    q->n = (q->n) + 1;
    q->q[q->n] = x;
    bubble_up(q, q->n);
  }
}

void bubble_up(priority_queue *q, int p) {
  if (pq_parent(p) == -1) {
    return; /* at root of heap, no parent */
  }
  if (q->q[pq_parent(p)] > q->q[p]) {
    pq_swap(q, p, pq_parent(p));
    bubble_up(q, pq_parent(p));
  }
}
```

Extracting the Minimum Element

```c++
item_type extract_min(priority_queue *q) {
  int min = -1; /* minimum value */
  if (q->n <= 0) {
    printf("Warning: empty priority queue.\n");
  } else {
    min = q->q[1];
    q->q[1] = q->q[q->n];
    q->n = q->n - 1;
    bubble_down(q, 1);
  }
  return(min);
}

void bubble_down(priority_queue *q, int p) {
  int c; /* child index */
  int i; /* counter */
  int min_index; /* index of lightest child */
  c = pq_young_child(p);
  min_index = p;
  for (i = 0; i <= 1; i++) {
    if ((c + i) <= q->n) {
      if (q->q[min_index] > q->q[c + i]) {
        min_index = c + i;
      }
    }
  }
  if (min_index != p) {
    pq_swap(q, p, min_index);
    bubble_down(q, min_index);
  }
}
```

An Even Faster Way to Build a Heap

- given two heaps and a fresh element, they can be merged into one by making the new one the root and bubbling down (heapify)

```c++
void make_heap_fast(priority_queue *q, item_type s[], int n) {
  int i;  /* counter */
  q->n = n;
  for (i = 0; i < n; i++) {
    q->q[i + 1] = s[i];
  }
  for (i = q->n/2; i >= 1; i--) {
    bubble_down(q, i);
  }
}
```
