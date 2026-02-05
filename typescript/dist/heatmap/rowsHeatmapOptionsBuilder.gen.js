"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.RowsHeatmapOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const heatmap = tslib_1.__importStar(require("../heatmap"));
// Controls frame rows options
class RowsHeatmapOptionsBuilder {
    constructor() {
        this.internal = heatmap.defaultRowsHeatmapOptions();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Sets the name of the cell when not calculating from data
    value(value) {
        this.internal.value = value;
        return this;
    }
    // Controls tick alignment when not calculating from data
    layout(layout) {
        this.internal.layout = layout;
        return this;
    }
}
exports.RowsHeatmapOptionsBuilder = RowsHeatmapOptionsBuilder;
//# sourceMappingURL=rowsHeatmapOptionsBuilder.gen.js.map