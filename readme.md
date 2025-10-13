## GoLang Learn Playground

Welcome to **GoLang Learn**, a hands-on playground for learning the Go programming language through bite-sized examples. This repo powers the **Coder's Gyan Go learning path**, and each numbered folder contains a focused program that demonstrates a single core concept, making it easy to run, tweak, and internalize Go fundamentals.

## Getting Started
- **Prerequisites:** Install Go 1.21 or later from [go.dev/dl](https://go.dev/dl/).
- **Clone:** `git clone https://github.com/toufikforyou/GoLangLearn.git`
- **Run an example:** Navigate into a lesson folder and run `go run <file>.go` (e.g. `cd 5_loops; go run loops.go`).
- **Experiment:** Modify the code and re-run to see how the behavior changes.

## Lesson Map
- `1_hello_world/` – First Go program and the `fmt` package.
- `2_simple_values/` – Working with numbers, strings, and booleans.
- `3_variables/` – Variable declarations, short assignment, and mutability.
- `4_constants/` – Defining constants and iota patterns.
- `5_loops/` – Loop constructs including `for` and `while`-style loops.
- `6_if_else/` – Conditional branching with `if`, `else if`, and `else`.
- `7_switch/` – Multi-way branching with `switch` statements.
- `8_arrays/` – Fixed-length collections and iteration.
- `9_slices/` – Dynamic slices, append, and slicing semantics.
- `10_maps/` – Key-value storage, lookups, and deletes.
- `11_range/` – Iterating over collections with `range`.
- `12_functions/` – Function declarations, return values, and scope.
- `13_variadic_functions/` – Variadic parameters and spread syntax.
- `14_closures/` – Capturing state with closures and anonymous functions.
- `15_pointers/` – Pointer basics, dereferencing, and mutation.
- `16_struct/` – Defining structs and associated data models.
- `17_interface/` – Interfaces, method sets, and polymorphism.
- `18_enums/` – Enum-like patterns with constants and custom types.
- `19_generics/` – Parameterized types and constraints in Go 1.18+.
- `20_goroutines/` – Concurrency fundamentals with goroutines.
- `21_channels/` – Channel communication patterns and synchronization.
- `22_mutex/` – Coordinating access to shared state with mutexes.
- `23_files/` – Reading from and writing to files.
- `24_packages/` – Structuring code with modules, packages, and imports.

## How To Run Everything
- **One-off execution:** `go run ./...` from the repo root runs all top-level examples.
- **Specific lesson:** `cd <lesson_dir>; go run .` executes the current folder entry point.
- **Format and vet:** `go fmt ./...` and `go vet ./...` keep code idiomatic and catch pitfalls.
- **Reset modules:** Run `go clean -modcache` if you experience dependency cache issues.

## Suggested Study Flow
1. Start with `1_hello_world/` and progress numerically—each concept builds on prior lessons.
2. After running an example, tweak inputs or add `fmt.Printf` statements to observe state.
3. Revisit earlier topics when exploring later sections (e.g., functions + slices) to reinforce understanding.

## Additional Tips
- Use `go fmt` to keep your code idiomatic: `go fmt ./...`.
- Pair these examples with the [Go Tour](https://go.dev/tour/) or [Effective Go](https://go.dev/doc/effective_go) for deeper dives.
- When ready for testing, explore Go’s built-in `testing` package to validate behavior.

## Contributing
- Open an issue or PR with new lesson ideas, fixes, or documentation improvements.
- Keep contributions focused and include a brief explanation in the PR description.
- Run `go fmt ./...` before submitting to ensure consistent formatting.

Happy coding! Feel free to submit issues or pull requests with improvements or new learning modules.
