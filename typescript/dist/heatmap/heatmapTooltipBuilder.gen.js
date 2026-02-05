"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.HeatmapTooltipBuilder = void 0;
const tslib_1 = require("tslib");
const heatmap = tslib_1.__importStar(require("../heatmap"));
// Controls tooltip options
class HeatmapTooltipBuilder {
    constructor() {
        this.internal = heatmap.defaultHeatmapTooltip();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Controls how the tooltip is shown
    mode(mode) {
        this.internal.mode = mode;
        return this;
    }
    maxHeight(maxHeight) {
        this.internal.maxHeight = maxHeight;
        return this;
    }
    maxWidth(maxWidth) {
        this.internal.maxWidth = maxWidth;
        return this;
    }
    // Controls if the tooltip shows a histogram of the y-axis values
    yHistogram(yHistogram) {
        this.internal.yHistogram = yHistogram;
        return this;
    }
    // Controls if the tooltip shows a color scale in header
    showColorScale(showColorScale) {
        this.internal.showColorScale = showColorScale;
        return this;
    }
}
exports.HeatmapTooltipBuilder = HeatmapTooltipBuilder;
//# sourceMappingURL=heatmapTooltipBuilder.gen.js.map