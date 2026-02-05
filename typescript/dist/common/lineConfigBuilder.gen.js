"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.LineConfigBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// TODO docs
class LineConfigBuilder {
    constructor() {
        this.internal = common.defaultLineConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    lineColor(lineColor) {
        this.internal.lineColor = lineColor;
        return this;
    }
    lineWidth(lineWidth) {
        this.internal.lineWidth = lineWidth;
        return this;
    }
    lineInterpolation(lineInterpolation) {
        this.internal.lineInterpolation = lineInterpolation;
        return this;
    }
    lineStyle(lineStyle) {
        const lineStyleResource = lineStyle.build();
        this.internal.lineStyle = lineStyleResource;
        return this;
    }
    // Indicate if null values should be treated as gaps or connected.
    // When the value is a number, it represents the maximum delta in the
    // X axis that should be considered connected.  For timeseries, this is milliseconds
    spanNulls(spanNulls) {
        this.internal.spanNulls = spanNulls;
        return this;
    }
}
exports.LineConfigBuilder = LineConfigBuilder;
//# sourceMappingURL=lineConfigBuilder.gen.js.map