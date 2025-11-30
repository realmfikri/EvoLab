# EvoLab

EvoLab is a Go-based artificial life simulation that models evolving agents competing for food. The simulation demonstrates how basic evolutionary pressures like mutation, resource scarcity, and tradeoffs between movement speed and body size influence emergent behaviors over time.

## Prerequisites
- Go 1.22+ (module-aware build).
- Standard Go toolchain (no external dependencies beyond the Go ecosystem).

## Building and Running
1. **Clone and enter the repository**
   ```bash
   git clone <repo-url>
   cd EvoLab
   ```
2. **Build**
   ```bash
   go build ./...
   ```
3. **Run** the simulation with optional flags to tune behavior:
   ```bash
   go run ./cmd/evolab \
     --mutation-rate 0.05 \
     --food-scarcity 0.6 \
     --speed-size-cost 1.2
   ```

### Configuration Options
You can tune simulation behavior via command-line flags or a configuration file.

- **Mutation rate** (`--mutation-rate`, float): Probability of trait mutation per generation. Higher values increase diversity but may destabilize populations.
- **Food scarcity** (`--food-scarcity`, float 0-1): Fraction representing how scarce food is; larger values mean less food available.
- **Speed-size cost** (`--speed-size-cost`, float): Multiplier balancing movement speed against body size; higher values penalize speed for larger agents.

#### Sample command-line usage
```bash
go run ./cmd/evolab --mutation-rate 0.08 --food-scarcity 0.45 --speed-size-cost 1.5
```

#### Sample configuration file (`config.yaml`)
```yaml
simulation:
  mutation_rate: 0.05
  food_scarcity: 0.6
  speed_size_cost: 1.2
```
Then run:
```bash
go run ./cmd/evolab --config config.yaml
```

## Roadmap
- **Neural view**: Visualize agent neural activity and decision pathways during simulation runs.
- **Species graph**: Track lineage divergence over time with a graph showing species relationships and dominance shifts.

