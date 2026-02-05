"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.FilterValueRangeBuilder = void 0;
const tslib_1 = require("tslib");
const heatmap = tslib_1.__importStar(require("../heatmap"));
// Controls the value filter range
class FilterValueRangeBuilder {
    constructor() {
        this.internal = heatmap.defaultFilterValueRange();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Sets the filter range to values less than or equal to the given value
    le(le) {
        this.internal.le = le;
        return this;
    }
    // Sets the filter range to values greater than or equal to the given value
    ge(ge) {
        this.internal.ge = ge;
        return this;
    }
}
exports.FilterValueRangeBuilder = FilterValueRangeBuilder;
//# sourceMappingURL=filterValueRangeBuilder.gen.js.map