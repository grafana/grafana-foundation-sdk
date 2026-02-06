# dashboard-converter

A command-line tool for converting Grafana dashboard JSON files into Go code using the Grafana Foundation SDK.

This enables you to:

- Version control your dashboards as code
- Programmatically modify dashboard content
- Generate dashboards within your Go applications
- Integrate with existing Go-based tooling

## Installation

```bash
make build
```

## Usage

```bash
# From file
./build/dashboard-converter my-dashboard.json -o dashboard.go
```

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-f, --format` | `go` | Output format |
| `-o, --output` | `-` | Output path ('-' for stdout) |
| `--cleanup` | `true` | Remove empty strings and null values |
| `--verbose` | `false` | Verbose output |

## Tips

```bash
# Use stdin and stdout
./build/dashboard-converter < my-dashboard.json > /path/to/my-ascode-dashboard/dashboard.go

# Convert a dashboard file to Go
./build/dashboard-converter my-dashboard.json -o /path/to/my-ascode-dashboard/dashboard.go
```

```
cd /path/to/my-ascode-dashboard
go mod init my-ascode-dashboard
go get github.com/grafana/grafana-foundation-sdk/go@v11.6.x+cog-v0.0.x
go mod tidy

go run dashboard.go > dashboard.json
```

### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format`, | `-f` | `go` | Output format (only 'go' is currently supported) |
| `--output`, | `-o` | `-` | Output path ('-' for stdout, or file.go file) |
| `--cleanup` | | `true` | Remove empty strings and null values from JSON |
| `--verbose` | | `false` | Enable verbose output for additional process information |

## Cleanup Recommendations

The generated Go code works out-of-the-box, but you may want to clean it up for better readability:

* Remove `Version(<number>)` line - version is often managed externally
* Remove `Id()` calls - IDs are typically assigned by Grafana
* Remove empty `Annotations()` stanzas if no annotations are defined
* Remove `Editable()` calls - dashboards as code are typically not editable in the UI
* Extract repeated query expressions into constants/variables for better maintainability

## Testing

```bash
make test
```
