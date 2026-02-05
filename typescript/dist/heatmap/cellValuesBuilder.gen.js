"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.CellValuesBuilder = void 0;
const tslib_1 = require("tslib");
const heatmap = tslib_1.__importStar(require("../heatmap"));
// Controls cell value options
class CellValuesBuilder {
    constructor() {
        this.internal = heatmap.defaultCellValues();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Controls the cell value unit
    unit(unit) {
        this.internal.unit = unit;
        return this;
    }
    // Controls the number of decimals for cell values
    decimals(decimals) {
        this.internal.decimals = decimals;
        return this;
    }
}
exports.CellValuesBuilder = CellValuesBuilder;
//# sourceMappingURL=cellValuesBuilder.gen.js.map