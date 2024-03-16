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

https://github.com/kampanosg/lazytest/assets/30287348/9da52725-06d5-4cb1-ac19-43c14d4f1d9f

</details>

<details>
  <summary>Run failed tests</summary>

https://github.com/kampanosg/lazytest/assets/30287348/42a60237-9757-443e-bb7e-197002283249

</details>

<details>
  <summary>Run passed tests</summary>

https://github.com/kampanosg/lazytest/assets/30287348/3c81c5de-9c7b-4eb5-9e6a-9daab8cc1f80

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

<details>
  <summary>Select text</summary>

  You can select output text with `Shift+Drag` and copy it with your OS's default buffer (e.g. `CMD+c` on the Mac)

https://github.com/kampanosg/lazytest/assets/30287348/285ea274-1ec4-4794-8a29-4778238cbb41

</details>

<details>
  <summary>Resize the panes</sumary>

  Sometimes test names may exceed the size of the pane. Similarly, you may need more space for the output. Unfortunately, [tview](), the TUI library that LazyTest uses does not support horizontal scrolling (and it's _probably_ [not]() going to be implemented any time soon).

  As an alternative, LazyTest panes can be resized with the `<` and `>` keys

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

> [!IMPORTANT]
> You need **Go v1.22** to build the project.

#### Makefile

You can build the project with `make`

```sh
make
```

An executable called `lazytest` will be built. You can move it to your `$PATH`

#### Bakery

If you're a [Bakery](https://github.com/kampanosg/bakery) fan (well, thank you!), LazyTest comes with a `Bakefile` that you can use with `bake`. You can just run

```sh
bake
```

An executable called `lazytest` will be built. You can move it to your `$PATH`

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

The following engines are being assembled:

* Rust
* Python
* Zig
* Jest

## Usage ⚙️

LazyTest aims to be as intuitive and easy to use as possible. If you're familiar with other terminal applications, LazyTest should (hopefully) feel natural. The table below lists all the available keys and modes.

| **Key**       | **Description**                      | **Where**    |
|---------------|--------------------------------------|--------------|
| `?`           | Opens the help dialog                |              |
| `j/k`         | Navigate up/down                     | Tests Panel  |
| `↑/↓`         | Navigate up down                     | Tests Panel  |
| `Enter/Space` | Fold/Unfold a node                   | Tests Panel  |
| `1`           | Focus on the Tests Panel             |              |
| `2`           | Focus on the Output Panel            |              |
| `r`           | Run a Test or Test Suite             | Tests Panel  |
| `a`           | Run all Tests                        | Tests Panel  |
| `p`           | Run Passed Tests (from previous run) | Tests Panel  |
| `f`           | Run Failed Tests (from previous run) | Tests Panel  |
| `/`           | Enter Search Mode                    | Tests Panel  |
| `Esc`         | Exit Search Mode                     | Search Panel |
| `Enter`       | Go to Search Results                 | Search Panel |
| `C`           | Clear Search Results                 | Tests Panel  |
| `>`           | Resize Right (make tests pane bigger)|              |
| `<`           | Resize Left (make output pane bigger)|              |
| `q`           | Quit the app                         |              |

### Use it with ToggleTerm

[ToggleTerm](https://github.com/akinsho/toggleterm.nvim) is a popular NeoVim plugin that lets you "persist and toggle multiple terminals during an editing session". It can be used with LazyTest, to avoid context-switching when you have to run your tests.

You can add the following to your NeoVim config:

```lua
local Terminal  = require('toggleterm.terminal').Terminal

local lazytest = Terminal:new({
  cmd = "lazytest",
  dir = ".",
  direction = "float",
  float_opts = {
    border = "curved",
  },
  -- function to run on opening the terminal
  on_open = function(term)
    vim.cmd("startinsert!")
    vim.api.nvim_buf_set_keymap(term.bufnr, "n", "q", "<cmd>close<CR>", {noremap = true, silent = true})
  end,
  -- function to run on closing the terminal
  on_close = function(term)
    vim.cmd("startinsert!")
  end,
})

function _lazytest_toggle()
  lazytest:toggle()
end

vim.api.nvim_set_keymap("n", "<C-t>", "<cmd>lua _lazytest_toggle()<CR>", {noremap = true, silent = true})
```

The above binds `<C-t>` to bring up a new floating terminal and executes the `./lazytest` command. You can quit the terminal by pressing `q`. Make sure you include this config to your `init.lua`. 

> [!NOTE]
> You can view and example of my NeoVim x LazyTest configuration [here](https://github.com/kampanosg/.dotfiles/commit/328dea4fe9f1b5f2cec13a188e7330bd11a2c0ed)

https://github.com/kampanosg/lazytest/assets/30287348/41052e0a-18ed-4908-8857-4271fd33abbd

### Flags

LazyTest supports the following flags that can be passed with the executable

| **Flag**  | **Description**                            | **Example**                           |
|-----------|--------------------------------------------|---------------------------------------|
| `dir`     | the directory to start searching for tests | `lazytest -dir /Code/example/project` |
| `excl`    | engines to exclude                         | `lazytest -excl=rust,zig`             |
| `version` | the current version of LazyTest            |                                       |

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
