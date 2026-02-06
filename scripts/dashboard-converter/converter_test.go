package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockHTTPClient is a mock implementation of HTTP client
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

// TestNewConverter tests the creation of a new Converter
func TestNewConverter(t *testing.T) {
	config := Config{
		Format:    "go",
		InputPath: "test.json",
		Output:    "output.go",
	}

	converter := NewConverter(config)

	assert.NotNil(t, converter)
	assert.Equal(t, config, converter.config)
}

// TestReadInputFromFile tests reading input from a file
func TestReadInputFromFile(t *testing.T) {
	// Create a temporary test file
	tempFile, err := os.CreateTemp("", "dashboard-*.json")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Write test data to the file
	testData := map[string]interface{}{
		"title": "Test Dashboard",
		"panels": []map[string]interface{}{
			{"title": "Panel 1", "type": "graph"},
		},
	}
	jsonData, err := json.Marshal(testData)
	require.NoError(t, err)
	_, err = tempFile.Write(jsonData)
	require.NoError(t, err)
	require.NoError(t, tempFile.Close())

	// Create converter with the temp file as input
	config := Config{
		Format:    "go",
		InputPath: tempFile.Name(),
		Output:    "-",
	}
	converter := NewConverter(config)

	// Test reading input
	ctx := context.Background()
	data, err := converter.readInput(ctx)
	require.NoError(t, err)

	// Verify the data
	dataMap, ok := data.(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, "Test Dashboard", dataMap["title"])
}

// TestReadInputFromStdin tests reading input from stdin
func TestReadInputFromStdin(t *testing.T) {
	// Save and restore the original stdin
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	// Create a pipe to simulate stdin
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdin = r

	// Write test data to the pipe
	testData := map[string]interface{}{
		"title": "Stdin Dashboard",
	}
	jsonData, err := json.Marshal(testData)
	require.NoError(t, err)

	go func() {
		defer w.Close()
		_, err := w.Write(jsonData)
		require.NoError(t, err)
	}()

	// Create converter with stdin as input
	config := Config{
		Format:    "go",
		InputPath: "-",
		Output:    "-",
	}
	converter := NewConverter(config)

	// Test reading input
	ctx := context.Background()
	data, err := converter.readInput(ctx)
	require.NoError(t, err)

	// Verify the data
	dataMap, ok := data.(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, "Stdin Dashboard", dataMap["title"])
}

// TestWriteOutput tests writing output to file
func TestWriteOutput(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "converter-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	testContent := []byte("package main\n\nfunc main() {}\n")

	tests := []struct {
		name        string
		config      Config
		checkOutput func(t *testing.T)
	}{
		{
			name: "write to stdout",
			config: Config{
				Output: "-",
			},
			checkOutput: func(t *testing.T) {
				// This is tested separately with a buffer
			},
		},
		{
			name: "write to file",
			config: Config{
				Output: filepath.Join(tempDir, "output.go"),
			},
			checkOutput: func(t *testing.T) {
				outputPath := filepath.Join(tempDir, "output.go")
				content, err := os.ReadFile(outputPath)
				require.NoError(t, err)
				assert.Equal(t, testContent, content)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter := NewConverter(tt.config)

			if tt.config.Output == "-" {
				// Capture stdout
				oldStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w

				err := converter.writeOutput(testContent)
				require.NoError(t, err)

				w.Close()
				os.Stdout = oldStdout

				var buf bytes.Buffer
				_, err = io.Copy(&buf, r)
				require.NoError(t, err)
				assert.Equal(t, string(testContent), buf.String())
			} else {
				err := converter.writeOutput(testContent)
				require.NoError(t, err)
				tt.checkOutput(t)
			}
		})
	}
}

// TestConvert tests the full conversion process
func TestConvert(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "converter-full-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a test input file
	inputPath := filepath.Join(tempDir, "input.json")
	testDashboard := map[string]interface{}{
		"title": "Full Test Dashboard",
		"uid":   "full-test",
		"panels": []interface{}{
			map[string]interface{}{
				"title": "Test Panel",
				"type":  "graph",
			},
		},
	}
	dashboardJSON, err := json.MarshalIndent(testDashboard, "", "  ")
	require.NoError(t, err)
	err = os.WriteFile(inputPath, dashboardJSON, 0644)
	require.NoError(t, err)

	// Set up output path
	outputPath := filepath.Join(tempDir, "output.go")

	// Create converter with test configuration
	config := Config{
		Format:    "go",
		InputPath: inputPath,
		Output:    outputPath,
		Cleanup:   true,
		Verbose:   true,
	}
	converter := NewConverter(config)

	// Run the conversion
	err = converter.Convert(context.Background())
	require.NoError(t, err)

	// Verify the output file exists
	_, err = os.Stat(outputPath)
	require.NoError(t, err)

	// Read the file and verify it contains expected content
	content, err := os.ReadFile(outputPath)
	require.NoError(t, err)
	contentStr := string(content)

	assert.Contains(t, contentStr, "package main")
	assert.Contains(t, contentStr, "Full Test Dashboard")
	assert.Contains(t, contentStr, "full-test")
}
