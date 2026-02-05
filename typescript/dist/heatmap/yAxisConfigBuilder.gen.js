"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.YAxisConfigBuilder = void 0;
const tslib_1 = require("tslib");
const heatmap = tslib_1.__importStar(require("../heatmap"));
// Configuration options for the yAxis
class YAxisConfigBuilder {
    constructor() {
        this.internal = heatmap.defaultYAxisConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Sets the yAxis unit
    unit(unit) {
        this.internal.unit = unit;
        return this;
    }
    // Reverses the yAxis
    reverse(reverse) {
        this.internal.reverse = reverse;
        return this;
    }
    // Controls the number of decimals for yAxis values
    decimals(decimals) {
        this.internal.decimals = decimals;
        return this;
    }
    // Sets the minimum value for the yAxis
    min(min) {
        this.internal.min = min;
        return this;
    }
    axisPlacement(axisPlacement) {
        this.internal.axisPlacement = axisPlacement;
        return this;
    }
    axisColorMode(axisColorMode) {
        this.internal.axisColorMode = axisColorMode;
        return this;
    }
    axisLabel(axisLabel) {
        this.internal.axisLabel = axisLabel;
        return this;
    }
    axisWidth(axisWidth) {
        this.internal.axisWidth = axisWidth;
        return this;
    }
    axisSoftMin(axisSoftMin) {
        this.internal.axisSoftMin = axisSoftMin;
        return this;
    }
    axisSoftMax(axisSoftMax) {
        this.internal.axisSoftMax = axisSoftMax;
        return this;
    }
    axisGridShow(axisGridShow) {
        this.internal.axisGridShow = axisGridShow;
        return this;
    }
    scaleDistribution(scaleDistribution) {
        const scaleDistributionResource = scaleDistribution.build();
        this.internal.scaleDistribution = scaleDistributionResource;
        return this;
    }
    axisCenteredZero(axisCenteredZero) {
        this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }
    // Sets the maximum value for the yAxis
    max(max) {
        this.internal.max = max;
        return this;
    }
    axisBorderShow(axisBorderShow) {
        this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
}
exports.YAxisConfigBuilder = YAxisConfigBuilder;
//# sourceMappingURL=yAxisConfigBuilder.gen.js.map