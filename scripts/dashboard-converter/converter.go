package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	_ "embed"

	"github.com/grafana/grafana-foundation-sdk/go/cog/plugins"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"golang.org/x/tools/imports"
)

// NOTE: ./main.go.tmpl contains `main.go` content of the generated code,
// with a single %s placeholder for the generated Foundation SDK code from
// calling dashboard.DashboardConverter(...).

//go:embed main.go.tmpl
var main_go_template string

// Converter handles dashboard conversion operations
type Converter struct {
	config Config
}

// NewConverter creates a new Converter instance
func NewConverter(config Config) *Converter {
	return &Converter{
		config: config,
	}
}

// Convert processes the dashboard conversion (Go output only).
func (c *Converter) Convert(ctx context.Context) error {
	plugins.RegisterDefaultPlugins()

	data, err := c.readInput(ctx)
	if err != nil {
		return err
	}

	if c.config.Cleanup {
		data = cleanJSON(data)
		if data == nil {
			return fmt.Errorf("dashboard is empty after cleanup")
		}
	}

	dashJSON, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal cleaned dashboard: %w", err)
	}

	var dash dashboard.Dashboard
	if err = json.Unmarshal(dashJSON, &dash); err != nil {
		return fmt.Errorf("failed to parse dashboard JSON: %w", err)
	}

	formattedCode, err := c.generateCode(dash)
	if err != nil {
		return err
	}

	return c.writeOutput(formattedCode)
}

// readInput handles different input sources
func (c *Converter) readInput(ctx context.Context) (any, error) {
	// Add verbose logging
	if c.config.Verbose {
		switch {
		case c.config.InputPath == "-":
			fmt.Println("Reading input from stdin")
		default:
			fmt.Println("Reading input from file:", c.config.InputPath)
		}
	}

	var rawData []byte
	var err error

	switch {
	case c.config.InputPath == "-":
		rawData, err = io.ReadAll(os.Stdin)
	default:
		rawData, err = os.ReadFile(c.config.InputPath)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	var data any
	if err = json.Unmarshal(rawData, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w\nInput: %s", err, string(rawData))
	}

	return data, nil
}

// generateCode creates the formatted Go code
func (c *Converter) generateCode(dash dashboard.Dashboard) ([]byte, error) {
	code := fmt.Sprintf(main_go_template, dashboard.DashboardConverter(dash))

	return imports.Process("", []byte(code), &imports.Options{
		TabWidth:  8,
		TabIndent: true,
		Comments:  true,
	})
}

// writeOutput handles the output of the converted code
func (c *Converter) writeOutput(content []byte) error {
	if c.config.Output == "-" {
		_, err := os.Stdout.Write(content)
		return err
	}

	outputPath := c.config.Output

	if c.config.Verbose {
		fmt.Println("Writing output to file:", outputPath)
	}

	if err := os.WriteFile(outputPath, content, 0o644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}
