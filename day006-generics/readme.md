# Getting started with generics

## non-generic functions

- non-generic functions are functions that work with a specific types

```go
func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}

// SumFloats adds together the values of m
func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}
```

## add a generic function to handle multiple types

we can create a generic function that can handle multiple types

```go
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}
```

- `comparable` is a constraint that tells the compiler that the type `K` must be comparable, it is a built-in interface in G
- this function accepts a map with keys of type `K` and values of type `V, where `V`can be either`int64`or`float64`

## Declare a type constraint

- we can declare a type constraint using the `type` keyword

```go
type Number interface {
  int64 | float64
}
```

then repalce the generic function with the type constraint

```go
func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
  var sum V
  for _, v := range m {
    sum += v
  }
  return sum
}
```
