// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as statetimeline from '../statetimeline';
import * as common from '../common';

export class FieldConfigBuilder implements cog.Builder<statetimeline.FieldConfig> {
    protected readonly internal: statetimeline.FieldConfig;

    constructor() {
        this.internal = statetimeline.defaultFieldConfig();
    }

    /**
     * Builds the object.
     */
    build(): statetimeline.FieldConfig {
        return this.internal;
    }

    lineWidth(lineWidth: number): this {
        if (!(lineWidth <= 10)) {
            throw new Error("lineWidth must be <= 10");
        }
        this.internal.lineWidth = lineWidth;
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

    hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>): this {
        const hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
        return this;
    }

    fillOpacity(fillOpacity: number): this {
        if (!(fillOpacity <= 100)) {
            throw new Error("fillOpacity must be <= 100");
        }
        this.internal.fillOpacity = fillOpacity;
        return this;
    }

    axisBorderShow(axisBorderShow: boolean): this {
        this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
}

