# suger

What is my motivation for creating this library?

我作该库的发心是什么?

1. Go places data structures such as dynamic arrays and hash tables at the language level, rather than at the standard library level like C++, resulting in a lack of official encapsulation for commonly used operations on these data structures.

   Go将诸如动态数组、哈希表这样的数据结构放在语言层面, 而非像C++一样放在标准库层面, 致使对这些数据结构的常用操作没有官方封装.

2. The exceptions in Go are propagated in the form of return values, resulting in most functions having more than one return value, which is inconvenient for nested use of functions. I have no intention of overturning this design, but only want to encapsulate it in appropriate scenarios.

   Go的异常通过返回值的形式拓传, 导致大部分函数的返回值个数都大于一个, 不方便函数嵌套使用, 我无意颠覆该设计, 只想在合适的场景作封装.

What principles will I follow?

我会遵循什么原则呢?

Go is not C++, Java, Python, it is Go; Simply rewriting code in other languages is not good, and Go's specifications and best practices should be followed.

Go不是C++、Java、Python, 它是Go; 单纯的重写在其他语言的代码并不好, 应该遵循Go的规范和最佳实践.

## Install

```bash
go get github.com/zweix123/suger
```

## Usage

+ [common](./common)
+ [slice](./slice)
+ [map](./map)
+ [functional](./functional)
+ [snippet](./snippet)

  This is a special directory with some typical and repetitive scenarios. The functions encapsulated for it do not conform to the design philosophy of Go, but they are indeed very sweet. Therefore, they are placed here and used with caution. This part of the code is often not robust.

  这是一个特别的目录, 有一些典型的、重复的场景，为其封装的函数并不符合Go的设计哲学, 但是确实很甜, 于是放在这里, 使用请谨慎, 这部分代码往往并不鲁棒.

  + [json](./json)

  + [future](./snippet/future.go): Implementing Promise and Future as Concurrent Primitives in Golang.

    [future](./snippet/future.go): 在Golang中实现promise和future这两个并发原语.
