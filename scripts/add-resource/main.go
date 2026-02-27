package main

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/url"
	"os"
	"path"
	"reflect"
	"regexp"
	"strings"

	"github.com/charmbracelet/huh"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

type schemaDescriptor struct {
	ResourceType          string
	ResourceIdentifier    string
	PackageName           string
	SchemaType            string
	SchemaUrl             string
	SchemaSubpath         string
	SchemaEnvelope        string
	AppPlatformApiVersion string
	AppPlatformKind       string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s resources_dir\n", os.Args[0])
		os.Exit(1)
	}

	resourcesDir := os.Args[1]
	descriptor := schemaDescriptor{}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Resource type").
				Options(
					huh.NewOption("Resource", "resource"),
					huh.NewOption("App Platform Resource", "app_platform_resource"),
					huh.NewOption("Panel", "panel"),
					huh.NewOption("Dataquery", "dataquery"),
				).
				Value(&descriptor.ResourceType),

			huh.NewInput().
				Title("Package").
				Description("Name of the package in which code for this resource will be generated.\nExample: folderv1beta1").
				Validate(func(str string) error {
					matched, err := regexp.Match(`^[a-z0-9]+$`, []byte(str))
					if err != nil {
						return err
					}

					if !matched {
						return errors.New("Packages can only contain lowercased alpha-numeric characters.")
					}

					return nil
				}).
				Value(&descriptor.PackageName),

			huh.NewSelect[string]().
				Title("Schema type").
				Options(
					huh.NewOption("CUE", "cue"),
					huh.NewOption("JSON Schema", "jsonschema"),
					huh.NewOption("OpenAPI", "openapi"),
				).
				Value(&descriptor.SchemaType),

			huh.NewInput().
				Title("Schema URL").
				Description("Example: https://raw.githubusercontent.com/grafana/grafana/v12.3.2/apps/folder/kinds/folder.cue").
				Validate(func(str string) error {
					_, err := url.ParseRequestURI(str)
					return err
				}).
				Value(&descriptor.SchemaUrl),
		).Title("Schema"),

		huh.NewGroup(
			huh.NewInput().
				Title("Resource identifier").
				Description("Example: 'grafana-athena-datasource' for the Athena datasource, 'timeseries' for the Timeseries panel.").
				Validate(func(str string) error {
					matched, err := regexp.Match(`^[a-z0-9\-]+$`, []byte(str))
					if err != nil {
						return err
					}

					if !matched {
						return errors.New("Identifiers can only contain lowercased alpha-numeric characters and dashes.")
					}

					return nil
				}).
				Value(&descriptor.ResourceIdentifier),
		).Title("Composable resource options").WithHideFunc(func() bool {
			return descriptor.ResourceType != "panel" && descriptor.ResourceType != "dataquery"
		}),

		huh.NewGroup(
			huh.NewInput().
				Title("Schema path").
				Description("If the actual schema isn't at the root of the file, this path describes where to find it.\nExample: folder.versions.v1beta1.schema.spec\nOptional.").
				Value(&descriptor.SchemaSubpath),

			huh.NewInput().
				Title("Schema envelope").
				Description("If the schema is defined under the `spec` field, it will be used as resource name unless told otherwise.\nExample: Folder\nOptional.").
				Value(&descriptor.SchemaEnvelope),
		).Title("CUE schema options").WithHideFunc(func() bool {
			return descriptor.SchemaType != "cue"
		}),

		huh.NewGroup(
			huh.NewInput().
				Title("Resource API Version").
				Description("Example: folder.grafana.app/v1beta1").
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("Required.")
					}
					return nil
				}).
				Value(&descriptor.AppPlatformApiVersion),

			huh.NewInput().
				Title("Resource API Kind").
				Description("Example: Folder").
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("Required.")
					}
					return nil
				}).
				Value(&descriptor.AppPlatformKind),
		).Title("App Platform resource options").WithHideFunc(func() bool {
			return descriptor.ResourceType != "app_platform_resource"
		}),
	)

	if err := form.Run(); err != nil {
		panic(err)
	}

	resourceDir := path.Join(resourcesDir, descriptor.PackageName)
	if err := writeCodegenConfig(descriptor, resourceDir); err != nil {
		panic(err)
	}

	fmt.Printf("✅ Resource configuration written in %s\n", resourceDir)
	fmt.Println("Run the `make preview` command to preview your changes.")
	fmt.Println("This command will generate an updated version of the SDK in `./workspace/foundation-sdk`. Explore it to verify the changes or run `git diff main..release-preview` to see the diff with main.")
}

func writeCodegenConfig(descriptor schemaDescriptor, resourceDir string) error {
	tmpl := template.New("").Option("missingkey=error").Funcs(template.FuncMap{
		"default": func(d any, given ...any) any {
			if empty(given) || empty(given[0]) {
				return d
			}
			return given[0]
		},
	})
	if err := parseFS(tmpl, templateFS, "templates"); err != nil {
		return err
	}

	if err := os.Mkdir(resourceDir, 0744); err != nil {
		return fmt.Errorf("could not create resource directory: %w", err)
	}

	// Codegen unit file
	buf := bytes.Buffer{}
	err := tmpl.ExecuteTemplate(&buf, "config.yaml.tmpl", map[string]any{
		"Descriptor": descriptor,
	})
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(resourceDir, "config.yaml"), buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("could not write codegen unit file: %w", err)
	}

	// Schema transforms file
	buf = bytes.Buffer{}
	err = tmpl.ExecuteTemplate(&buf, "schema.transforms.yaml.tmpl", map[string]any{
		"Descriptor": descriptor,
	})
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(resourceDir, "schema.transforms.yaml"), buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("could not write schema.transforms.yaml file: %w", err)
	}

	// Builder transforms file
	buf = bytes.Buffer{}
	err = tmpl.ExecuteTemplate(&buf, "common.builder.transforms.yaml.tmpl", map[string]any{
		"Descriptor": descriptor,
	})
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(resourceDir, "common.builder.transforms.yaml"), buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("could not write common.builder.transforms.yaml file: %w", err)
	}

	return nil
}

func parseFS(tmpl *template.Template, vfs fs.FS, rootDir string) error {
	if vfs == nil {
		return nil
	}

	err := fs.WalkDir(vfs, rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		fileHandle, err := vfs.Open(path)
		if err != nil {
			return err
		}

		contents, err := io.ReadAll(fileHandle)
		if err != nil {
			return err
		}

		templateName := strings.TrimPrefix(strings.TrimPrefix(path, rootDir), "/")
		t := tmpl.New(templateName)
		_, err = t.Parse(string(contents))

		return err
	})

	return err
}

// empty returns true if the given value has the zero value for its type.
// see https://github.com/Masterminds/sprig/blob/e708470d529a10ac1a3f02ab6fdd339b65958372/defaults.go#L35
func empty(given any) bool {
	g := reflect.ValueOf(given)
	if !g.IsValid() {
		return true
	}

	// Basically adapted from text/template.isTrue
	switch g.Kind() {
	default:
		return g.IsNil()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return g.Len() == 0
	case reflect.Bool:
		return !g.Bool()
	case reflect.Complex64, reflect.Complex128:
		return g.Complex() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return g.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return g.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return g.Float() == 0
	case reflect.Struct:
		return false
	}
}
