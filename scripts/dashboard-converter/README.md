# dashboard-converter

A command-line tool for converting Grafana dashboard JSON files into Go code using the Grafana Foundation SDK.

This enables you to:

- Version control your dashboards as code
- Programmatically modify dashboard content
- Generate dashboards within your Go applications
- Integrate with existing Go-based tooling

## Usage

### With Docker

```bash
# From file
docker run -v $(pwd):/workspace grafana/dashboard-converter \
  /workspace/my-dashboard.json -o /workspace/dashboard.go

# From stdin
cat my-dashboard.json | docker run -i grafana/dashboard-converter > dashboard.go
```

### With local binary

```bash
# From file
./build/dashboard-converter my-dashboard.json -o dashboard.go

# From stdin
cat my-dashboard.json | ./build/dashboard-converter > dashboard.go
```

## Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--output` | `-o` | `-` | Output path ('-' for stdout) |
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

## How to create a release for the dashboard-converter

Releases are triggered by pushing a tag with the format `dashboard-converter-vX.Y.Z`. The [Dashboard Converter workflow](https://github.com/grafana/grafana-foundation-sdk/actions/workflows/dashboard-converter.yaml) runs automatically on tag push.

```console
git tag dashboard-converter-v1.0.0
git push origin dashboard-converter-v1.0.0
```
