# lazytest 🧪

![Build Pipeline](https://img.shields.io/github/actions/workflow/status/kampanosg/lazytest/go.yml?style=for-the-badge&logo=go)
![Security Pipeline](https://img.shields.io/github/actions/workflow/status/kampanosg/lazytest/sec.yml?style=for-the-badge&logo=go&label=Security)

LazyTest is a command line tool that helps you view, search and run tests in your project.

<p align="center">
  <img width="1726" alt="Screenshot 2024-03-06 at 16 30 01" src="https://github.com/kampanosg/lazytest/assets/30287348/c744fb01-d8e1-439b-9c12-98076408f88b">
</p>

## Features ⭐

<details>
  <summary>Run a test</summary>
  
  https://github.com/kampanosg/lazytest/assets/30287348/99be1ef1-81b2-47e7-8772-44cd957ad744

</details>

<details>
  <summary>Run all tests</summary>
  
  https://github.com/kampanosg/lazytest/assets/30287348/31c87349-eb59-43e8-8154-bf65f940686c

</details>

<details>
  <summary>Run failed tests</summary>

https://github.com/kampanosg/lazytest/assets/30287348/42a60237-9757-443e-bb7e-197002283249

</details>

<details>
  <summary>Run passed tests</summary>

https://github.com/kampanosg/lazytest/assets/30287348/02a7b69b-a4d8-4ae7-973b-30d8f63a146d

</details>

<details>
  <summary>View Results</summary>

https://github.com/kampanosg/lazytest/assets/30287348/6bd3fef4-8465-428b-b8ea-94e09d4aac5e

</details>

<details>
  <summary>Search for a test</summary>

https://github.com/kampanosg/lazytest/assets/30287348/f0b084b8-822c-4c76-b86b-f575a1e1e4e7

</details>

<details>
  <summary>Viewing and navigating tests</summary>

https://github.com/kampanosg/lazytest/assets/30287348/65bb7520-16dc-473b-b342-1ce79036854a

</details>

## Installation ⬇️

### Go

You can use Go to install LazyTest

```sh
go install github.com/kampanosg/lazytest@latest
```

### Brew

> [!WARNING]
> Brew is not yet available, I am working on it.

### Build from source

> [!WARNING]
> You need **Go v1.22** to build the project.

## Engines 🏎️

"Engines" are a core concept in LazyTest. It allows the tool to be extensible and work with many languages and test frameworks. When LazyTest parses the codebase, it uses the engines to determine if a given file contains tests. Engines also provide instructions on how to run a given test.

You can write a new Engine by implementing the following function

```go
type LazyEngine interface {
  ParseTestSuite(fp string) (*models.LazyTestSuite, error)
}
```

The `ParseTestSuite` expects a `filepath`. If it's a valid test file for the given language, then the engine should parse the file and return the `LazyTestSuite` which contains all the tests and instructions on how to run them. As an example, have a look at the Go engine:

https://github.com/kampanosg/lazytest/blob/c4e9a5800f76c01d780e798e0511b95288de0057/pkg/engines/golang/engine.go#L26-L54

### Available Engines

LazyTests comes packed with the following engines:

* Golang

### Upcoming Engines

The following engines are being built:

* Rust
* Python
* Zig
* Jest

## Usage ⚙️

### Use it with ToggleTerm

## Inspiration & Similar Projects 💬

LazyTest is heavily inspired by the following projects. Go check them out - their work is excellent!

* [LazyGit](https://github.com/jesseduffield/lazygit) - A simple terminal UI for git commands
* [NeoTest](https://github.com/nvim-neotest/neotest) - A framework for interacting with tests within NeoVim.

## Mandatory GIF 🖼️

<p align="center">
  <img src="https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExMHk3eTA0Z3ZkdjV2dmh2anJjaW85N2F1OGl4d2F2bXFyeDl1OGpuYyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/SV34NE5jEh4cM/giphy.gif" />
  <br />
  <i>The average developer when being asked to write tests</i>
</p>
