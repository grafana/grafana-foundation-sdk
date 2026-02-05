"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AxisConfigBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// TODO docs
class AxisConfigBuilder {
    constructor() {
        this.internal = common.defaultAxisConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
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
    axisBorderShow(axisBorderShow) {
        this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
}
exports.AxisConfigBuilder = AxisConfigBuilder;
//# sourceMappingURL=axisConfigBuilder.gen.js.map