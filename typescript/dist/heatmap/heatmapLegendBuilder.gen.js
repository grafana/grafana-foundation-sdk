"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.HeatmapLegendBuilder = void 0;
const tslib_1 = require("tslib");
const heatmap = tslib_1.__importStar(require("../heatmap"));
// Controls legend options
class HeatmapLegendBuilder {
    constructor() {
        this.internal = heatmap.defaultHeatmapLegend();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Controls if the legend is shown
    show(show) {
        this.internal.show = show;
        return this;
    }
}
exports.HeatmapLegendBuilder = HeatmapLegendBuilder;
//# sourceMappingURL=heatmapLegendBuilder.gen.js.map