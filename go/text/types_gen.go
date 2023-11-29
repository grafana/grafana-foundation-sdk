// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package text

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
