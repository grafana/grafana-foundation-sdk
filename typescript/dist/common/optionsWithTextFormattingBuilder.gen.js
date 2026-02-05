"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.OptionsWithTextFormattingBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// TODO docs
class OptionsWithTextFormattingBuilder {
    constructor() {
        this.internal = common.defaultOptionsWithTextFormatting();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    text(text) {
        const textResource = text.build();
        this.internal.text = textResource;
        return this;
    }
}
exports.OptionsWithTextFormattingBuilder = OptionsWithTextFormattingBuilder;
//# sourceMappingURL=optionsWithTextFormattingBuilder.gen.js.map