// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as xychart from '../xychart';
import * as common from '../common';

export class FieldConfigBuilder implements cog.Builder<xychart.FieldConfig> {
    protected readonly internal: xychart.FieldConfig;

    constructor() {
        this.internal = xychart.defaultFieldConfig();
    }

    /**
     * Builds the object.
     */
    build(): xychart.FieldConfig {
        return this.internal;
    }

    show(show: xychart.XYShowMode): this {
        this.internal.show = show;
        return this;
    }

    pointSize(pointSize: {
	fixed?: number;
	min?: number;
	max?: number;
}): this {
        this.internal.pointSize = pointSize;
        return this;
    }

    pointShape(pointShape: xychart.PointShape): this {
        this.internal.pointShape = pointShape;
        return this;
    }

    pointStrokeWidth(pointStrokeWidth: number): this {
        if (!(pointStrokeWidth >= 0)) {
            throw new Error("pointStrokeWidth must be >= 0");
        }
        this.internal.pointStrokeWidth = pointStrokeWidth;
        return this;
    }

    fillOpacity(fillOpacity: number): this {
        if (!(fillOpacity <= 100)) {
            throw new Error("fillOpacity must be <= 100");
        }
        this.internal.fillOpacity = fillOpacity;
        return this;
    }

    lineWidth(lineWidth: number): this {
        if (!(lineWidth >= 0)) {
            throw new Error("lineWidth must be >= 0");
        }
        this.internal.lineWidth = lineWidth;
        return this;
    }

    hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>): this {
        const hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
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

    lineStyle(lineStyle: cog.Builder<common.LineStyle>): this {
        const lineStyleResource = lineStyle.build();
        this.internal.lineStyle = lineStyleResource;
        return this;
    }

    axisBorderShow(axisBorderShow: boolean): this {
        this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
}

