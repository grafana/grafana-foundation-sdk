"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.CodeOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const text = tslib_1.__importStar(require("../text"));
class CodeOptionsBuilder {
    constructor() {
        this.internal = text.defaultCodeOptions();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // The language passed to monaco code editor
    language(language) {
        this.internal.language = language;
        return this;
    }
    showLineNumbers(showLineNumbers) {
        this.internal.showLineNumbers = showLineNumbers;
        return this;
    }
    showMiniMap(showMiniMap) {
        this.internal.showMiniMap = showMiniMap;
        return this;
    }
}
exports.CodeOptionsBuilder = CodeOptionsBuilder;
//# sourceMappingURL=codeOptionsBuilder.gen.js.map