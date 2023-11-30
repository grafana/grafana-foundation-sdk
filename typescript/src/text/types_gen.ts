// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export enum TextMode {
	HTML = "html",
	Markdown = "markdown",
	Code = "code",
}

export const defaultTextMode = (): TextMode => (TextMode.HTML);

export enum CodeLanguage {
	Json = "json",
	Yaml = "yaml",
	Xml = "xml",
	Typescript = "typescript",
	Sql = "sql",
	Go = "go",
	Markdown = "markdown",
	Html = "html",
	Plaintext = "plaintext",
}

export const defaultCodeLanguage = (): CodeLanguage => (CodeLanguage.Plaintext);

export interface CodeOptions {
	// The language passed to monaco code editor
	language: CodeLanguage;
	showLineNumbers: boolean;
	showMiniMap: boolean;
}

export const defaultCodeOptions = (): CodeOptions => ({
	language: CodeLanguage.Plaintext,
	showLineNumbers: false,
	showMiniMap: false,
});

export interface Options {
	mode: TextMode;
	code?: CodeOptions;
	content: string;
}

export const defaultOptions = (): Options => ({
	mode: TextMode.Markdown,
	content: "# Title\n\nFor markdown syntax help: [commonmark.org/help](https://commonmark.org/help/)",
});

