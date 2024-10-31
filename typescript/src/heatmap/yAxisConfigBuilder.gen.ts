// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as heatmap from '../heatmap';
import * as common from '../common';

// Configuration options for the yAxis
export class YAxisConfigBuilder implements cog.Builder<heatmap.YAxisConfig> {
    protected readonly internal: heatmap.YAxisConfig;

    constructor() {
        this.internal = heatmap.defaultYAxisConfig();
    }

    /**
     * Builds the object.
     */
    build(): heatmap.YAxisConfig {
        return this.internal;
    }

    // Sets the yAxis unit
    unit(unit: string): this {
        this.internal.unit = unit;
        return this;
    }

    // Reverses the yAxis
    reverse(reverse: boolean): this {
        this.internal.reverse = reverse;
        return this;
    }

    // Controls the number of decimals for yAxis values
    decimals(decimals: number): this {
        this.internal.decimals = decimals;
        return this;
    }

    // Sets the minimum value for the yAxis
    min(min: number): this {
        this.internal.min = min;
        return this;
    }

    axisPlacement(axisPlacement: common.AxisPlacement): this {
        this.internal.axisPlacement = axisPlacement;
        return this;
    }

    axisColorMode(axisColorMode: common.AxisColorMode): this {
        this.internal.axisColorMode = axisColorMode;
        return this;
    }

    axisLabel(axisLabel: string): this {
        this.internal.axisLabel = axisLabel;
        return this;
    }

    axisWidth(axisWidth: number): this {
        this.internal.axisWidth = axisWidth;
        return this;
    }

    axisSoftMin(axisSoftMin: number): this {
        this.internal.axisSoftMin = axisSoftMin;
        return this;
    }

    axisSoftMax(axisSoftMax: number): this {
        this.internal.axisSoftMax = axisSoftMax;
        return this;
    }

    axisGridShow(axisGridShow: boolean): this {
        this.internal.axisGridShow = axisGridShow;
        return this;
    }

    scaleDistribution(scaleDistribution: cog.Builder<common.ScaleDistributionConfig>): this {
        const scaleDistributionResource = scaleDistribution.build();
        this.internal.scaleDistribution = scaleDistributionResource;
        return this;
    }

    axisCenteredZero(axisCenteredZero: boolean): this {
        this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }

    // Sets the maximum value for the yAxis
    max(max: number): this {
        this.internal.max = max;
        return this;
    }

    axisBorderShow(axisBorderShow: boolean): this {
        this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
}
