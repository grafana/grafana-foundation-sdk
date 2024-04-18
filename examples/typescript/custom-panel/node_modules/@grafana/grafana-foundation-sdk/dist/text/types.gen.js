"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultOptions = exports.defaultCodeOptions = exports.defaultCodeLanguage = exports.CodeLanguage = exports.defaultTextMode = exports.TextMode = void 0;
var TextMode;
(function (TextMode) {
    TextMode["HTML"] = "html";
    TextMode["Markdown"] = "markdown";
    TextMode["Code"] = "code";
})(TextMode || (exports.TextMode = TextMode = {}));
const defaultTextMode = () => (TextMode.HTML);
exports.defaultTextMode = defaultTextMode;
var CodeLanguage;
(function (CodeLanguage) {
    CodeLanguage["Json"] = "json";
    CodeLanguage["Yaml"] = "yaml";
    CodeLanguage["Xml"] = "xml";
    CodeLanguage["Typescript"] = "typescript";
    CodeLanguage["Sql"] = "sql";
    CodeLanguage["Go"] = "go";
    CodeLanguage["Markdown"] = "markdown";
    CodeLanguage["Html"] = "html";
    CodeLanguage["Plaintext"] = "plaintext";
})(CodeLanguage || (exports.CodeLanguage = CodeLanguage = {}));
const defaultCodeLanguage = () => (CodeLanguage.Plaintext);
exports.defaultCodeLanguage = defaultCodeLanguage;
const defaultCodeOptions = () => ({
    language: CodeLanguage.Plaintext,
    showLineNumbers: false,
    showMiniMap: false,
});
exports.defaultCodeOptions = defaultCodeOptions;
const defaultOptions = () => ({
    mode: TextMode.Markdown,
    content: "# Title\n\nFor markdown syntax help: [commonmark.org/help](https://commonmark.org/help/)",
});
exports.defaultOptions = defaultOptions;
//# sourceMappingURL=types.gen.js.map