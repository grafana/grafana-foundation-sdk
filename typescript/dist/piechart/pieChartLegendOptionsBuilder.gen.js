"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.PieChartLegendOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const piechart = tslib_1.__importStar(require("../piechart"));
class PieChartLegendOptionsBuilder {
    constructor() {
        this.internal = piechart.defaultPieChartLegendOptions();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    values(values) {
        this.internal.values = values;
        return this;
    }
    displayMode(displayMode) {
        this.internal.displayMode = displayMode;
        return this;
    }
    placement(placement) {
        this.internal.placement = placement;
        return this;
    }
    showLegend(showLegend) {
        this.internal.showLegend = showLegend;
        return this;
    }
    asTable(asTable) {
        this.internal.asTable = asTable;
        return this;
    }
    isVisible(isVisible) {
        this.internal.isVisible = isVisible;
        return this;
    }
    sortBy(sortBy) {
        this.internal.sortBy = sortBy;
        return this;
    }
    sortDesc(sortDesc) {
        this.internal.sortDesc = sortDesc;
        return this;
    }
    width(width) {
        this.internal.width = width;
        return this;
    }
    calcs(calcs) {
        this.internal.calcs = calcs;
        return this;
    }
}
exports.PieChartLegendOptionsBuilder = PieChartLegendOptionsBuilder;
//# sourceMappingURL=pieChartLegendOptionsBuilder.gen.js.map