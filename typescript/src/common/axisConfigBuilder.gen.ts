// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class AxisConfigBuilder implements cog.Builder<common.AxisConfig> {
    protected readonly internal: common.AxisConfig;

    constructor() {
        this.internal = common.defaultAxisConfig();
    }

    build(): common.AxisConfig {
        return this.internal;
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

    axisBorderShow(axisBorderShow: boolean): this {
        this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
}
