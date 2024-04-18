export declare enum TextMode {
    HTML = "html",
    Markdown = "markdown",
    Code = "code"
}
export declare const defaultTextMode: () => TextMode;
export declare enum CodeLanguage {
    Json = "json",
    Yaml = "yaml",
    Xml = "xml",
    Typescript = "typescript",
    Sql = "sql",
    Go = "go",
    Markdown = "markdown",
    Html = "html",
    Plaintext = "plaintext"
}
export declare const defaultCodeLanguage: () => CodeLanguage;
export interface CodeOptions {
    language: CodeLanguage;
    showLineNumbers: boolean;
    showMiniMap: boolean;
}
export declare const defaultCodeOptions: () => CodeOptions;
export interface Options {
    mode: TextMode;
    code?: CodeOptions;
    content: string;
}
export declare const defaultOptions: () => Options;
