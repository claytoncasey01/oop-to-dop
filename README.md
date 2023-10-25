# Object Oriented Programming to Data Oriented Programming

## Description

After seeing many talks on Data Oriented Programming but few examples on how to actually get started with it I decided to make this repo. The goal is to provide easy to digest examples in OOP as well as DOP.

### Running

#### Go

>        cd go
>
> > ##### Tests
> >
> >        go test ./...
> >
> > ##### Benchmarks
> >
> >       go test ./... -bench=. -benchmem

> #### Rust
>
>      cd rust
>
> > #### Tests
> >
> >      cargo test
> >
> > #### Benchmarks
> >
> >      cargo bench

## Benchmarks

## Rust

| Function Name       | OOP Time (µs or ns) | DOP Time (µs or ns) | Speed Difference     |
| ------------------- | ------------------- | ------------------- | -------------------- |
| find_by_id          | 4.2601 µs           | 35.571 ns           | DOP is 8.35x faster  |
| find_by_title       | 91.119 ns           | 77.583 ns           | DOP is 1.17x faster  |
| find_by_author_name | 38.961 µs           | 40.111 µs           | OOP is 1.03x faster  |
| update              | 62.089 ns           | 77.203 ns           | OOP is 1.24x faster  |
| publish             | 40.402 ns           | 22.747 ns           | DOP is 1.78x faster  |
| delete              | 726.47 µs           | 20.470 ns           | DOP is 35.49x faster |
| add                 | 163.68 ns           | 465.10 ns           | OOP is 2.84x faster  |

## Go

| Function Name        | OOP Time (µs or ns) | DOP Time (µs or ns) | Speed Difference      |
| -------------------- | ------------------- | ------------------- | --------------------- |
| FindPostById         | 34.672 µs           | 8.427 µs            | DOP is 4.11x faster   |
| FindPostByTitle      | 17.51 ns            | 5.517 ns            | DOP is 3.17x faster   |
| FindPostByAuthorName | 61.521 µs           | 282.1 ns            | DOP is 4.59x faster   |
| PublishPost          | 420.491 µs          | 7.466 ns            | DOP is 56.32x faster  |
| UpdatePost           | 7726.18 µs          | 42.84 ns            | DOP is 180.35x faster |
| AddPost              | 848.7 ns            | 639.6 ns            | DOP is 1.33x faster   |
| DeletePost           | 27815.649 µs        | 4830.309 µs         | DOP is 5.76x faster   |

#### TODO

- ~~Go Examples~~
- ~~Rust Examples~~
- Zig Examples
- C++ Examples
- Others (Kotlin, Typescript, Java, C#)?
- Real-World Examples in each language (Web Server, CLI, etc possibly)
