package main

import (
	"context"
	"flag"
	"fmt"
	"os"
)

// Config holds the application configuration
type Config struct {
	Format    string
	Output    string
	Cleanup   bool
	InputPath string
	Verbose   bool
}

// cleanJSON recursively removes empty strings and null values, aiming to let
// Grafana use defaults instead, while also working around some schema type
// mismatches, e.g. "" vs [] which will raise JSON unmarshal errors.

// Common fixups to apply while cleaning JSON, from recent schema changes.
func fixups(path []string, key string, val any) bool {
	// Create the full path including current key
	fullPath := append(append([]string{}, path...), key)

	// Skip root "refresh": false, as it doesn't match recent schemas where refresh is a string
	if len(fullPath) == 1 && fullPath[0] == "refresh" {
		if boolVal, ok := val.(bool); ok && !boolVal {
			return true
		}
	}
	return false // Don't skip by default
}

func cleanJSON(data any) any {
	return cleanJSONWithPath(data, []string{})
}

func cleanJSONWithPath(data any, path []string) any {
	switch v := data.(type) {
	case map[string]any:
		result := make(map[string]any)
		for key, val := range v {
			// Apply fixups before cleaning
			if fixups(path, key, val) {
				continue
			}

			// Create new path for this level
			newPath := append(append([]string{}, path...), key)

			if cleaned := cleanJSONWithPath(val, newPath); cleaned != nil {
				if str, ok := cleaned.(string); !ok || str != "" {
					result[key] = cleaned
				}
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	case []any:
		var result []any
		for i, val := range v {
			// Create new path with array index
			arrayIndexPath := append(append([]string{}, path...), fmt.Sprintf("[%d]", i))

			if cleaned := cleanJSONWithPath(val, arrayIndexPath); cleaned != nil {
				if str, ok := cleaned.(string); !ok || str != "" {
					result = append(result, cleaned)
				}
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	case string:
		if v == "" {
			return nil
		}
		return v
	default:
		return v
	}
}

// parseArgs parses command line arguments into a Config
func parseArgs() (Config, error) {
	var config Config

	flag.StringVar(&config.Format, "format", "go", "Output format (only 'go' is currently supported)")
	flag.StringVar(&config.Format, "f", "go", "Output format (short for --format)")
	flag.StringVar(&config.Output, "output", "-", "Output path ('-' for stdout)")
	flag.StringVar(&config.Output, "o", "-", "Output path (short for --output)")
	flag.BoolVar(&config.Cleanup, "cleanup", true, "Remove empty strings and null values from JSON")
	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose output for additional process information")

	flag.Parse()

	// Handle input path
	args := flag.Args()
	if len(args) == 0 {
		if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
			config.InputPath = "-"
		} else {
			return config, fmt.Errorf("input path is required")
		}
	} else {
		config.InputPath = args[0]
	}

	return config, nil
}

func main() {
	config, err := parseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintln(os.Stderr, "Usage: dashboard-converter [options] [input]")
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	ctx := context.Background()
	if err := NewConverter(config).Convert(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
