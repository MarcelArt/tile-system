# Repository Guidelines

## Project Structure & Module Organization

This Go/raylib application renders a tile grid and simulates heat transfer. `main.go` creates the game and activates the world scene.

- `internal/data/`: tile models, block definitions, and material properties.
- `internal/entities/`: tile generation, rendering, coordinate conversion, and thermal simulation.
- `internal/scenes/`: world scene, simulation timing, and tooltip UI.
- `pkg/engine/`: game loop, scene/entity interfaces, and save support.
- `pkg/array2d/`, `pkg/rng/`, and `pkg/no/`: reusable utilities.

There is currently no dedicated asset directory or test suite. Place tests beside the code they exercise. `tmp/` contains ignored temporary output.

## Build, Test, and Development Commands

Use the Go toolchain version declared in `go.mod` (currently `1.27.1`). Running the application requires a graphical environment and raylib runtime dependencies.

- `make go` or `go run .`: launch the application locally.
- `make dev`: launch Air for development; install `air` separately and configure it as needed.
- `make build`: compile the executable to `builds/tile-system`.
- `go test ./...`: run all package tests.
- `go vet ./...`: check for common Go mistakes.
- `gofmt -w <changed-files>`: format changed Go files.

## Coding Style & Naming Conventions

Use standard Go formatting with tabs and `gofmt`. Keep package names lowercase, exported identifiers in PascalCase, and unexported identifiers in camelCase. Preserve existing role suffixes such as `tile.data.go`, `tile_grid.entity.go`, and `world.scene.go`. Raylib imports use the alias `rl`.

Keep application behavior in `internal/` and reusable infrastructure in `pkg/`. Follow the existing separation between drawing, frame updates, and simulation updates.

## Testing Guidelines

Use Go's standard `testing` package, `*_test.go` files, and `TestName` functions. No coverage threshold is configured. Prefer deterministic, table-driven tests for grid boundaries, coordinate conversion, and heat-transfer behavior. For rendering or tooltip changes, also run the application and check cursor behavior near window edges.

## Commit & Pull Request Guidelines

History uses short, lowercase subjects with prefixes such as `feat: tooltip` and `wip: heat transfer`. Keep commits focused and describe the concrete change.

PRs should explain the behavior change, link relevant issues, and report validation commands and results. Include screenshots for visible UI changes and describe simulation assumptions when modifying thermal behavior.
