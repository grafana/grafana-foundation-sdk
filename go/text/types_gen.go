// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package text

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type TextMode string

const (
	TextModeHTML     TextMode = "html"
	TextModeMarkdown TextMode = "markdown"
	TextModeCode     TextMode = "code"
)

type CodeLanguage string

const (
	CodeLanguageJson       CodeLanguage = "json"
	CodeLanguageYaml       CodeLanguage = "yaml"
	CodeLanguageXml        CodeLanguage = "xml"
	CodeLanguageTypescript CodeLanguage = "typescript"
	CodeLanguageSql        CodeLanguage = "sql"
	CodeLanguageGo         CodeLanguage = "go"
	CodeLanguageMarkdown   CodeLanguage = "markdown"
	CodeLanguageHtml       CodeLanguage = "html"
	CodeLanguagePlaintext  CodeLanguage = "plaintext"
)

type CodeOptions struct {
	// The language passed to monaco code editor
	Language        CodeLanguage `json:"language"`
	ShowLineNumbers bool         `json:"showLineNumbers"`
	ShowMiniMap     bool         `json:"showMiniMap"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CodeOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CodeOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "language"
	if fields["language"] != nil {
		if string(fields["language"]) != "null" {
			if err := json.Unmarshal(fields["language"], &resource.Language); err != nil {
				errs = append(errs, cog.MakeBuildErrors("language", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("language", errors.New("required field is null"))...)

		}
		delete(fields, "language")
	} else {
		errs = append(errs, cog.MakeBuildErrors("language", errors.New("required field is missing from input"))...)
	}
	// Field "showLineNumbers"
	if fields["showLineNumbers"] != nil {
		if string(fields["showLineNumbers"]) != "null" {
			if err := json.Unmarshal(fields["showLineNumbers"], &resource.ShowLineNumbers); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showLineNumbers", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showLineNumbers", errors.New("required field is null"))...)

		}
		delete(fields, "showLineNumbers")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showLineNumbers", errors.New("required field is missing from input"))...)
	}
	// Field "showMiniMap"
	if fields["showMiniMap"] != nil {
		if string(fields["showMiniMap"]) != "null" {
			if err := json.Unmarshal(fields["showMiniMap"], &resource.ShowMiniMap); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showMiniMap", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showMiniMap", errors.New("required field is null"))...)

		}
		delete(fields, "showMiniMap")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showMiniMap", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CodeOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CodeOptions` objects.
func (resource CodeOptions) Equals(other CodeOptions) bool {
	if resource.Language != other.Language {
		return false
	}
	if resource.ShowLineNumbers != other.ShowLineNumbers {
		return false
	}
	if resource.ShowMiniMap != other.ShowMiniMap {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CodeOptions` fields for violations and returns them.
func (resource CodeOptions) Validate() error {
	return nil
}

type Options struct {
	Mode    TextMode     `json:"mode"`
	Code    *CodeOptions `json:"code,omitempty"`
	Content string       `json:"content"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Options) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is null"))...)

		}
		delete(fields, "mode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is missing from input"))...)
	}
	// Field "code"
	if fields["code"] != nil {
		if string(fields["code"]) != "null" {

			resource.Code = &CodeOptions{}
			if err := resource.Code.UnmarshalJSONStrict(fields["code"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("code", err)...)
			}

		}
		delete(fields, "code")

	}
	// Field "content"
	if fields["content"] != nil {
		if string(fields["content"]) != "null" {
			if err := json.Unmarshal(fields["content"], &resource.Content); err != nil {
				errs = append(errs, cog.MakeBuildErrors("content", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("content", errors.New("required field is null"))...)

		}
		delete(fields, "content")
	} else {
		errs = append(errs, cog.MakeBuildErrors("content", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Options", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Options` objects.
func (resource Options) Equals(other Options) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.Code == nil && other.Code != nil || resource.Code != nil && other.Code == nil {
		return false
	}

	if resource.Code != nil {
		if !resource.Code.Equals(*other.Code) {
			return false
		}
	}
	if resource.Content != other.Content {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if resource.Code != nil {
		if err := resource.Code.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("code", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// VariantConfig returns the configuration related to text panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "text",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := json.Unmarshal(raw, options); err != nil {
				return nil, err
			}

			return options, nil
		},
		StrictOptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := options.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return options, nil
		},
		GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}

			return PanelConverter(inputPanel.(dashboard.Panel))
		},
	}
}
