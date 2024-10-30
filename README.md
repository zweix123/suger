# suger

What is my motivation for creating this library?

我作该库的发心是什么?

1. Go places data structures such as dynamic arrays and hash tables at the language level, rather than at the standard library level like C++, resulting in a lack of official encapsulation for commonly used operations on these data structures.

   Go将诸如动态数组、哈希表这样的数据结构放在语言层面, 而非像C++一样放在标准库层面, 致使对这些数据结构的常用操作没有官方封装.

2. The exceptions in Go are propagated in the form of return values, resulting in most functions having more than one return value, which is inconvenient for nested use of functions. I have no intention of overturning this design, but only want to encapsulate it in appropriate scenarios.

   One solution to this prob lem is in a functional style, while another solution is to use monad's Options, Results, Either, and other types.
   
   Go的异常通过返回值的形式拓传, 导致大部分函数的返回值个数都大于一个, 不方便函数嵌套使用, 我无意颠覆该设计, 只想在合适的场景作封装.

   这个问题的一个解法是函数式的风格, 另一个解法是monad的Option、Result、Either等类型.

What principles will I follow?

我会遵循什么原则呢?

1. Go is not C++, Java, Python, it is Go; Simply rewriting code in other languages is not good, and Go's specifications and best practices should be followed.

   Go不是C++、Java、Python, 它是Go; 单纯的重写在其他语言的代码并不好, 应该遵循Go的规范和最佳实践.

2. As abstract and universal as possible
  
   尽可能的抽象和通用

## Install

```bash
go get -u github.com/zweix123/suger
```

## Usage

+ [async](./async)
  + `Future`
+ [common](./common)
  + `Assert`
  + `HandlePanic`
+ [dict](./dict)
  + `Contains`
  + `Keys`
  + `Values`
+ [functional](./functional)
  + `MapSerial`
  + `MapParallel`
  + `MapParallelWithLimit`
  + `Reduce`
  + `Filter`
+ [slice](./slice)
  + `Contains`
  + `Equal`
  + `Reverse`
  + `Uniq`

  + `Chunk`
  + `Flatten`
