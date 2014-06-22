Benchmarks
====

### Type switch

If the passed-in value is not a `reflect.Value` use a type switch. 
If no match is found, create and check `reflect.Value` value.

```
BenchmarkIsZeroCheck	    20000000	       146 ns/op	      12 B/op	       0 allocs/op
BenchmarkIsZeroCheckReflect	50000000	        67.9 ns/op	       6 B/op	       0 allocs/op
```


### Removal of type assertions

```
BenchmarkIsZeroCheck	    20000000	       140 ns/op	      12 B/op	       0 allocs/op
BenchmarkIsZeroCheckReflect	50000000	        67.3 ns/op	       6 B/op	       0 allocs/op
```

Performance was slightly but consistently higher without the type assertions. Memory usage stays the same. 


### Removal of explicit switch cases  

```
BenchmarkIsZeroCheck	    10000000	       176 ns/op	      13 B/op	       0 allocs/op
BenchmarkIsZeroCheckReflect	20000000	        90.8 ns/op	       6 B/op	       0 allocs/op
```

Performance turned significantly worse. Reversed change.


### Added Cache

A `map[reflect.Type]interface{}` and a mutex were used to cache results of calls to `reflect.Zero'.

```
BenchmarkIsZeroCheck	    10000000	       150 ns/op	       7 B/op	       0 allocs/op
BenchmarkIsZeroCheckReflect	20000000	        71.8 ns/op	       3 B/op	       0 allocs/op
```

Memory usage per operation was reduced significantly. But performance dropped too much. Removing cache. 