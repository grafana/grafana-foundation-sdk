// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package text

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
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

type Options struct {
	Mode    TextMode     `json:"mode"`
	Code    *CodeOptions `json:"code,omitempty"`
	Content string       `json:"content"`
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "text",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
