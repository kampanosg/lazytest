<p align="center">
  <b>LazyTest 🥱🧪</b>
  <br>
  <i>LazyTest is a command line tool that helps you view, search and run tests in your project.</i>
  <br />
  <br />
  <img alt="the lazytest tui" src="https://github.com/kampanosg/lazytest/assets/30287348/e8843e0e-c847-40a7-83d2-6e09231b121e">
</p>

<p align="center">
  <img src="https://img.shields.io/github/actions/workflow/status/kampanosg/lazytest/go.yml?style=for-the-badge&logo=go" />
  <img src="https://img.shields.io/github/actions/workflow/status/kampanosg/lazytest/sec.yml?style=for-the-badge&logo=go&label=Security" />
</p>

## Features ⭐

<details>
  <summary>Run a test</summary>

  <a href="https://asciinema.org/a/W4MXbeoTxGzyLjCCAgXjHDD2U?t=4" target="_blank"><img src="https://asciinema.org/a/W4MXbeoTxGzyLjCCAgXjHDD2U.svg" /></a>

</details>

<details>
  <summary>Run a test suite</summary>

  <a href="https://asciinema.org/a/yrAjoTbk6xNea0EP1imokWLA8?t=3" target="_blank"><img src="https://asciinema.org/a/yrAjoTbk6xNea0EP1imokWLA8.svg" /></a>

</details>

<details>
  <summary>Run all tests</summary>

  <a href="https://asciinema.org/a/xrvObScQMKAGbrwq1WO93znrl?t=2" target="_blank"><img src="https://asciinema.org/a/xrvObScQMKAGbrwq1WO93znrl.svg" /></a>

</details>

<details>
  <summary>Run failed tests</summary>

<a href="https://asciinema.org/a/uRx39aO9qbhwWQV2cYLCsZQYT?t=2" target="_blank"><img src="https://asciinema.org/a/uRx39aO9qbhwWQV2cYLCsZQYT.svg" /></a>

</details>

<details>
  <summary>Run passed tests</summary>

  <a href="https://asciinema.org/a/qkCh5I8DtiRpE9Trb5QQjFKkr?t=3.1" target="_blank"><img src="https://asciinema.org/a/qkCh5I8DtiRpE9Trb5QQjFKkr.svg" /></a>

</details>

<details>
  <summary>Search for a test</summary>

<a href="https://asciinema.org/a/wODn3nhYWeeqxvpUcUTH1KqgO?t=2" target="_blank"><img src="https://asciinema.org/a/wODn3nhYWeeqxvpUcUTH1KqgO.svg" /></a>

</details>

<details>
  <summary>Resize the panes</summary>

  Sometimes test names may exceed the size of the pane. Or you may need more space for the output text. Unfortunately, [tview](https://github.com/rivo/tview), the TUI library that LazyTest uses, does not support horizontal scrolling (and it's _probably_ [not](https://github.com/rivo/tview/issues/707#issuecomment-1991260955) going to be implemented any time soon).

  As an alternative, LazyTest panes can be resized with the `+` and `-` keys.

<a href="https://asciinema.org/a/Pj9sSFz9I2doITi3sQQkvgpgO?t=3.2" target="_blank"><img src="https://asciinema.org/a/Pj9sSFz9I2doITi3sQQkvgpgO.svg" /></a>
  
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

"Engines" are a core concept in LazyTest. It allows the tool to be extensible and work with many languages and test frameworks. LazyTest uses engines to parse the codebase and determine which files contain tests. Engines also provide instructions on how to run a given test.

You can write a new Engine by implementing the following interface:

```go
type LazyEngine interface {
	Load(dir string) (*models.LazyTree, error)
	GetIcon() string
}
```

The `Load` function parses the contents of `dir` and constructs a `LazyTree` which includes all the tests. The tree is then passed to the TUI that renders it on the terminal

https://github.com/kampanosg/lazytest/blob/c0ce2bfdbc501a91c26662c20bbe154c907dd500/pkg/engines/golang/engine.go#L37-L71

### Available Engines

LazyTests comes packed with the following engines:

* Go (requires [`go`](https://go.dev/))
* Rust (requires [`cargo`](https://www.rust-lang.org/))
* Bashunit (requires [`bashunit`](https://bashunit.typeddevs.com/))

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
| `3`           | Focus on the History Panel           |              |
| `4`           | Focus on the Timings Panel           |              |
| `r`           | Run a Test or Test Suite             | Tests Panel  |
| `a`           | Run all Tests                        | Tests Panel  |
| `p`           | Run Passed Tests (from previous run) | Tests Panel  |
| `f`           | Run Failed Tests (from previous run) | Tests Panel  |
| `/`           | Enter Search Mode                    | Tests Panel  |
| `Esc`         | Exit Search Mode                     | Search Panel |
| `Enter`       | Go to Search Results                 | Search Panel |
| `C`           | Clear Search Results                 | Tests Panel  |
| `+`           | Increase the Tests Panel size        |              |
| `-`           | Decrease the Tests Panel size        |              |
| `0`           | Reset the layout                     |              |
| `y`           | Copy the test name/test suite path   | Tests Panel  |
| `Y`           | Copy the (current) output            | Tests Panel  |
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

<a href="https://asciinema.org/a/tlLvhwKDe7ruyHPyLsc8aZzzg?t=3" target="_blank"><img src="https://asciinema.org/a/tlLvhwKDe7ruyHPyLsc8aZzzg.svg" /></a>

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
