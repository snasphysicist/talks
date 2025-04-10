---
marp: true
class: invert
style: |
  section {
    align-content: start;    
  }
  .two-columns {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 1rem;
  }
  .four-columns {
    display: grid;
    grid-template-columns: repeat(4, minmax(0, 1fr));
    gap: 1rem;
  }
---

# Don't Be An Iter-Hater!

## Iterators in Go

## Scott Nicholas Allan Smith

## London Gophers

## 2025-04-16

---

# About Me

// TODO

---

# Where Are We Going?

- Motivation
    - What Are Iterators? 
    - In Other Languages
- Go Iterator Basics
    - Range Over Functions, `iter.Seq` & `iter.Seq2`
    - Reimplementing Range Over Slices, Maps
    - Anatomy Of/Patterns In A Go Iterator
- Examples
    - Iterating Over a Custom Type (`Set`)
    - Searching in a Huge Text File
    - Querying JSON in a Database Without Database Level JSON Support
    - Infinite Sequences

---

# Iterators

A standard way of "going through" a sequence of values

Usually all they provide is
- some way of getting the next element (if any)
- some way of knowing if there is another element to get

---

# Other Languages

At least these languages have an iterator concept

- Java (`java.util.Iterator`)
- Javascript ("iterator protocol")
- Rust (`std::iter::Iterator`)
- Python (`__iter__`/`__next__`)
- Elixir (`Enumerable`)
- Clojure (`ISeq`)

--- 

# Other Languages

| Language | Iterator | Next | Finished? |
| ---------|----------|----------|----------
| Java | `java.util.Iterator` | `next()` | `hasNext() == false`
| Javascript | iterator protocol | `next().value` | `next().done`
| Rust | `std::iter::Iterator` | 10 | £2.50
| Python | `__iter__` / `__next__`| `__next__` | Raises `StopIteration` |

---

# Iterators In Go

"Iterators" in Go are really just special functions

One of three forms

```go
func(func yield() boolean)

func(func yield(k K) boolean) -> iter.Seq[K]

func(func yield(k K, v V) boolean) -> iter.Seq2[K, V]
```

---

# Range Over Functions

The "official" name of the iterator feature

```go
iter0 := NewZeroArgumentIterator()
for range iter0 {}

iter1 := NewOneArgumentIterator()
for i := range iter1 {}

iter2 := NewTwoArgumentIterator()
for k, v := range iter1 {}
```

---

# Yield Function

That's a function which accepts a function, `yield`, which

- accepts 0, 1 or 2 values which are assigned to loop variables
- returns a boolean which communicates whether another value is required

---

# Reimplementing Range Over Slice

Demo

---

# Reimplementing Range Over Slice

```go
func OneArgumentIterator[T any](c []T) iter.Seq[int] {
  i := 0
  return func(yield func(int) boolean) {
    for {
      if i >= len(c) {
        return
      }
      shouldStop := yield(i)
      if shouldStop {
        return
      }
      i++
    }
  }
}
```

Remember, one value `for` over a slice iterates over indices.

---

# Anatomy Of An Iterator

```go
func OneArgumentIterator[T any](c []T) iter.Seq[int] {
  i := 0 // Vars keeping track of what you've already
         // iterated over must be declared 
         // outside the anonymous function & captured
  return func(yield func(int) boolean) {
    for {
      if i >= len(c) {  // - Natural exit condition for finite sequences
        return          // | May be omitted for infinite sequences
      }                 // -
      shouldStop := yield(i) // Yield, passing value(s) to the for-range
      if shouldStop { // - 
        return        // | Exit when the loop requires no more values
      }               // - 
      i++
    }
  }
}
```

---

# Anatomy Of An Iterator

Your iterator functions will generally follow this pattern

Always have

- A call to `yield` with the next value to iterate over
- An `if` -> `return`/`break` on `yield`'s return value

Almost certainly have

- Some "closed over" way to keep track of what's been yielded
- Its own natural exit condition

---

# Who Is Executing?

TODO: diagram passing between for and iterator function

--- 

# Reimplementing Range Over Slice (Two Valued)

Demo

--- 

# Reimplementing Range Over Slice (Two Valued)

```go
func TwoArgumentIterator[T any](c []T) iter.Seq2[int, T] {
  i := 0
  return func(yield func(int, T) boolean) {
    for {
      if i >= len(c) {
        return
      }
      // Only difference is here, we call yield with two values
      shouldStop := yield(i, c[i])
      if shouldStop {
        return
      }
      i++
    }
  }
}
```

--- 

# Reimplementing Range Over Map

Demo

--- 

# Reimplementing Range Over Map

```go
// Special case for string keys
func MapIterator[T any](m [string]T) iter.Seq2[string, T] {
  ks := keys(m)
  i := 0
  return func(yield func(string, T) boolean) {
    for {
      if i >= len(ks) {
        return
      }
      shouldStop := yield(ks[i], m[ks[i]])
      if shouldStop {
        return
      }
      i++
    }
  }  
}
```

---

# Aside: Keys From Map Without Range

```go
func keys[T any](m [string]T) []string {
  vs := reflect.ValueOf(m).MapKeys()
  i := 0
  ks := make([]string, 0)
  for {
    if i >= len(vs) {
      return ks
    }
    ks = append(ks, vs[i].String())
    i++
  }
}
```

Because we need to call a method on the `reflect.Value` depending upon
the type of the key, I haven't figured out how to generalise this

---

# Eager vs Lazy

In what follows I'm going to talk about _eager_ evaluation vs _lazy_ evaluation

Roughly: 

- _eager_ evaluation means stuff is calculated now, even if I don't need it
- _lazy_ evaluation means calculation is defered until I need it to happen


---

# Iterating Over a Custom Type (`Set`)

A set is just a collection of elements that contains no duplicates

Lets consider how we can convert a slice to a set

- eagerly
- lazily

---

# Querying JSON in a Database Without Database Level JSON Support

// TODO

---

# "Infinite" Sequences

Not truly infinite - we have bounded memory, power, time, etc...

Seems esoteric

Different natural way of thinking between functional and imperative programming

Functional - sequence (or pipeline) of operations on data

---

# Generate a Random Hex ID

Recent requirement - 16 digit hex id

Look at Go approach

Look at natural approach in functional programming

---

# Generate a Random Hex ID

Demo

---

# Compared To A Functional Language

<div class="two-columns">

<div> 

Elixir

```elixir
fn -> "ABCDEF0123456789" 
      |> String.graphemes() 
      |> Enum.random() end
|> Stream.repeat()
|> Stream.take(16)
|> Enum.join()
```
</div>

<div> 

Go

```go
chars := []rune("ABCDEFG0123456789")
rc := func() string {
  n := rand.Int31n(16)
  return string(chars[n])
}
r := repeatedly(rc)
t := take(16, r)
id := join(t)
return id
```
</div>

</div>

---

# Compared To A Functional Language

<div class="two-columns">

<div>

Clojure

```clojure
(->> #(rand-nth "ABCDEF0123456789")
     (repeatedly)
     (take 16)
     (clojure.string/join))
```

</div>

<div>

Go

```go
chars := []rune("ABCDEFG0123456789")
rc := func() string {
  n := rand.Int31n(16)
  return string(chars[n])
}
r := repeatedly(rc)
t := take(16, r)
id := join(t)
return id
```

</div>

</div>

---

# Functional Approach

I would argue that the functional approach

- conflates flow control and data transformation less
- reads more like natural language
- therefore can be faster to understand

Your mileage may vary - come debate me after the talk!

---

# Thanks!

For listening

Please give brutally honest feedback and criticism

Do you think this would make a good GopherCon talk? Come tell me (yes or no!)

[uk.linkedin.com/in/scottnasmith](https://uk.linkedin.com/in/scottnasmith)

[github.com/snasphysicist/](https://github.com/snasphysicist/)

All my links, projects and thoughts at my website:

[www.snas.pw](https://www.snas.pw)

(Run as cheaply as possible with no guarantees on uptime)