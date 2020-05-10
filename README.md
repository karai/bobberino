![bobberino](https://user-images.githubusercontent.com/34389545/81494289-59913300-926d-11ea-871a-9a827650655a.png)

[![Discord](https://img.shields.io/discord/388915017187328002?label=Join%20Discord)](http://chat.turtlecoin.lol) [![GitHub issues](https://img.shields.io/github/issues/karai/bobberino?label=Issues)](https://github.com/karai/bobberino/issues) ![GitHub stars](https://img.shields.io/github/stars/karai/bobberino?label=Github%20Stars) ![Build](https://github.com/karai/bobberino/workflows/Build/badge.svg) ![GitHub](https://img.shields.io/github/license/karai/bobberino) ![GitHub issues by-label](https://img.shields.io/github/issues/karai/bobberino/Todo)

## Usage

**Launch Bobberino**

```bash
./bobberino
```

## Dependencies

-   Golang 1.13+ [[Download]](https://golang.org)

## Operating System

-   Linux
-   MacOS (Need testers)
-   BSD (Need testers)
-   Windows
    -   Note: Windows requires Git BASH for proper color rendering [Download](https://gitforwindows.org/)

## Building

```bash
go build main.go
```

**Optional:** Compile with all errors displayed, then run binary. Avoids "too many errors" from hiding error info.

```bash
go build -gcflags="-e" && ./bobberino
```

## Contributing

-   MIT License
-   `gofmt` is used on all files.
-   go modules are used to manage dependencies.
