"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.VizTooltipOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// TODO docs
class VizTooltipOptionsBuilder {
    constructor() {
        this.internal = common.defaultVizTooltipOptions();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    mode(mode) {
        this.internal.mode = mode;
        return this;
    }
    sort(sort) {
        this.internal.sort = sort;
        return this;
    }
    maxWidth(maxWidth) {
        this.internal.maxWidth = maxWidth;
        return this;
    }
    maxHeight(maxHeight) {
        this.internal.maxHeight = maxHeight;
        return this;
    }
    hideZeros(hideZeros) {
        this.internal.hideZeros = hideZeros;
        return this;
    }
}
exports.VizTooltipOptionsBuilder = VizTooltipOptionsBuilder;
//# sourceMappingURL=vizTooltipOptionsBuilder.gen.js.map