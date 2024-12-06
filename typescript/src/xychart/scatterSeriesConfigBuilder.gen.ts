// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as xychart from '../xychart';
import * as common from '../common';

export class ScatterSeriesConfigBuilder implements cog.Builder<xychart.ScatterSeriesConfig> {
    protected readonly internal: xychart.ScatterSeriesConfig;

    constructor() {
        this.internal = xychart.defaultScatterSeriesConfig();
    }

    /**
     * Builds the object.
     */
    build(): xychart.ScatterSeriesConfig {
        return this.internal;
    }

    x(x: string): this {
        this.internal.x = x;
        return this;
    }

    y(y: string): this {
        this.internal.y = y;
        return this;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    show(show: xychart.ScatterShow): this {
        this.internal.show = show;
        return this;
    }

    pointSize(pointSize: cog.Builder<common.ScaleDimensionConfig>): this {
        const pointSizeResource = pointSize.build();
        this.internal.pointSize = pointSizeResource;
        return this;
    }

    pointColor(pointColor: cog.Builder<common.ColorDimensionConfig>): this {
        const pointColorResource = pointColor.build();
        this.internal.pointColor = pointColorResource;
        return this;
    }

    lineColor(lineColor: cog.Builder<common.ColorDimensionConfig>): this {
        const lineColorResource = lineColor.build();
        this.internal.lineColor = lineColorResource;
        return this;
    }

    lineWidth(lineWidth: number): this {
        if (!(lineWidth >= 0)) {
            throw new Error("lineWidth must be >= 0");
        }
        this.internal.lineWidth = lineWidth;
        return this;
    }

    lineStyle(lineStyle: cog.Builder<common.LineStyle>): this {
        const lineStyleResource = lineStyle.build();
        this.internal.lineStyle = lineStyleResource;
        return this;
    }

    label(label: common.VisibilityMode): this {
        this.internal.label = label;
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

    frame(frame: number): this {
        this.internal.frame = frame;
        return this;
    }

    labelValue(labelValue: cog.Builder<common.TextDimensionConfig>): this {
        const labelValueResource = labelValue.build();
        this.internal.labelValue = labelValueResource;
        return this;
    }

    axisBorderShow(axisBorderShow: boolean): this {
        this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
}
