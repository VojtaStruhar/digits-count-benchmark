# Counting even / odd digits in a number

## Task

Write a function that returns true if the number of even digits is greater than the number of odd digits in the input number.

My friend had this task on his programming exam in the first semester of uni. I wanted to benchmark different approaches to solving this problem.

## Benchmark

1. Convert the number to a string, iterate over the now letters (digits), convert them back into integers and do the calculation.
2. Convert the number to a string and iterate over the now letters, but instead of converting them back to integers, just look at the byte value. The evenness of ASCII is the same as of the numbers themselves.
3. Numeric solution. Check if the number is even and then divide it by 10, throwing away the last digit.

| Approach                     | Iterations | speed [ns/op] |
| ---------------------------- | ---------: | ------------: |
| Text solution                |   43787802 |         27.20 |
| Text solution + optimization |   76066716 |         15.93 |
| Numeric solution             |  521218287 |         2.302 |

>  This is imperfect code altogether. The text solutions don't even work on negative numbers, but still - the results are interesting.

**Benchmark command**

```shell
go test -bench=.
```

**Benchmark output**

```
goos: darwin
goarch: arm64
pkg: evenOddDigits
BenchmarkEvenOddCount/Text_solution-8                           43787802                27.20 ns/op
BenchmarkEvenOddCount/Text_solution_-_shortcut-8                76066716                15.93 ns/op
BenchmarkEvenOddCount/Numeric_solution-8                        521218287                2.302 ns/op
PASS
ok      evenOddDigits   4.796s
```

Benchmarked on base model M2 Macbook Air
