# Testing Guide

This repository uses Go's standard testing tools for unit, integration, and benchmarking coverage. The recommendations below outline how to validate core simulation behaviors and keep runs reproducible.

## Deterministic and Reproducible Runs

- Seed random number generators at the top of each test file to ensure deterministic behavior (e.g., `rand.Seed(12345)` for `math/rand`).
- Prefer dependency injection for randomness: where possible, pass a `rand.Rand` instance into constructors or functions so tests can use a local RNG seeded per test case.
- Avoid time-based seeds in tests; use fixed constants and document them so failures can be reproduced.
- Keep floating-point comparisons tolerant with small epsilon values (`1e-6`) to account for round-off in neural computations.

## Unit Tests

### Genetic Operators (`internal/ga`)
- **Uniform crossover**: With a fixed seed, assert that child genomes include predictable mixes of parent traits, and that no traits are lost. Validate behavior when parents share keys and when a key exists in only one parent.
- **Gaussian mutation**: Use a seeded RNG and small `Scale` values to ensure perturbations fall within expected ranges. Test immutability by confirming the original genome map is unchanged, and mutation applies per-key adjustments.
- Table-driven tests are recommended to cover varying trait counts and mutation scales without duplicating code.

### Neural Network Forward Passes (`internal/ga`)
- Construct `DenseNetwork` instances with fixed weight matrices and input vectors. Assert exact output vectors for deterministic cases (e.g., identity weights, zero weights, or triangular matrices).
- Include dimension-mismatch scenarios where inputs are shorter than weight rows to confirm graceful handling (ignored extra weights) without panics.
- Use subtests (`t.Run`) to separate activation patterns, making failures easier to trace.

### Quadtree Collision Queries (`internal/physics`)
- Seed blob positions deterministically or handcraft fixtures positioned in and out of the query rectangle.
- Assert that `Query` returns only blobs whose positions fall inside the requested rectangle and that inserts preserve existing blobs.
- Add edge-case coverage for blobs on the boundary and for empty quadtrees to ensure zero-length slices are returned without nils.

## Integration Tests

### Population Cycles (`internal/simulation`)
- Drive the `Engine` with a small set of deterministic blobs that implement `Update` predictably (e.g., stub blobs tracking update counts).
- Run multi-step cycles and assert expected counts of updates, renderer callbacks (`BeforeStep`/`AfterStep`), and error handling when steps are invalid or contexts cancel early.
- Validate that configuration fields (step count, timestep) propagate through the loop and that blobs mutate state across iterations as expected.

## Performance Benchmarks

- Use Go benchmarks (`go test -bench=.`) to profile hotspots. For scalable runs, parameterize benchmark cases with blob counts (e.g., 100, 250, 500, 750) to observe nonlinear growth.
- For quadtree-heavy workloads with 500+ blobs, benchmark both insertion and query patterns: clustered vs. uniformly distributed blobs, and varying query rectangle sizes.
- Benchmark simulation loops by instantiating the `Engine` with prebuilt blob slices and a no-op renderer. Measure step throughput with fixed timesteps to detect regressions.
- When comparing benchmark runs, set `GOMAXPROCS` explicitly and pin seeds so runs are comparable across machines.
