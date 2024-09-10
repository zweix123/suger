# functional

+ [Map](./functional.go)

  > Inspired by [samber/lo](https://github.com/samber/lo)
  > + By default, concurrent mode is used because loops can be easily implemented without the need for concurrency. Therefore, map is recommended for use in concurrent scenarios.
  > + The provided parameter function has only one parameter because I believe that processing each element should be symmetrical and not distinguished by index.
  >
  > 灵感来自[samber/lo](https://github.com/samber/lo)
  > + 默认使用并发模式, 因为假如不需要并发使用循环可以简单的实现, 所以map推荐在并发的场景使用.
  > + 提供的参数函数只有一个参数, 因为我认为处理每个元素是对称的, 而不应该按照索引区分.
