// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// Sparkline cell options
export class TableSparklineCellOptionsBuilder implements cog.Builder<common.TableSparklineCellOptions> {
    protected readonly internal: common.TableSparklineCellOptions;

    constructor() {
        this.internal = common.defaultTableSparklineCellOptions();
        this.internal.type = "sparkline";
    }

    /**
     * Builds the object.
     */
    build(): common.TableSparklineCellOptions {
        return this.internal;
    }

    drawStyle(drawStyle: common.GraphDrawStyle): this {
        this.internal.drawStyle = drawStyle;
        return this;
    }

    gradientMode(gradientMode: common.GraphGradientMode): this {
        this.internal.gradientMode = gradientMode;
        return this;
    }

    thresholdsStyle(thresholdsStyle: cog.Builder<common.GraphThresholdsStyleConfig>): this {
        const thresholdsStyleResource = thresholdsStyle.build();
        this.internal.thresholdsStyle = thresholdsStyleResource;
        return this;
    }

    lineColor(lineColor: string): this {
        this.internal.lineColor = lineColor;
        return this;
    }

    lineWidth(lineWidth: number): this {
        this.internal.lineWidth = lineWidth;
        return this;
    }

    lineInterpolation(lineInterpolation: common.LineInterpolation): this {
        this.internal.lineInterpolation = lineInterpolation;
        return this;
    }

    lineStyle(lineStyle: cog.Builder<common.LineStyle>): this {
        const lineStyleResource = lineStyle.build();
        this.internal.lineStyle = lineStyleResource;
        return this;
    }

    fillColor(fillColor: string): this {
        this.internal.fillColor = fillColor;
        return this;
    }

    fillOpacity(fillOpacity: number): this {
        this.internal.fillOpacity = fillOpacity;
        return this;
    }

    showPoints(showPoints: common.VisibilityMode): this {
        this.internal.showPoints = showPoints;
        return this;
    }

    pointSize(pointSize: number): this {
        this.internal.pointSize = pointSize;
        return this;
    }

    pointColor(pointColor: string): this {
        this.internal.pointColor = pointColor;
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

    barAlignment(barAlignment: common.BarAlignment): this {
        this.internal.barAlignment = barAlignment;
        return this;
    }

    barWidthFactor(barWidthFactor: number): this {
        this.internal.barWidthFactor = barWidthFactor;
        return this;
    }

    stacking(stacking: cog.Builder<common.StackingConfig>): this {
        const stackingResource = stacking.build();
        this.internal.stacking = stackingResource;
        return this;
    }

    hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>): this {
        const hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
        return this;
    }

    hideValue(hideValue: boolean): this {
        this.internal.hideValue = hideValue;
        return this;
    }

    transform(transform: common.GraphTransform): this {
        this.internal.transform = transform;
        return this;
    }

    // Indicate if null values should be treated as gaps or connected.
    // When the value is a number, it represents the maximum delta in the
    // X axis that should be considered connected.  For timeseries, this is milliseconds
    spanNulls(spanNulls: boolean | number): this {
        this.internal.spanNulls = spanNulls;
        return this;
    }

    fillBelowTo(fillBelowTo: string): this {
        this.internal.fillBelowTo = fillBelowTo;
        return this;
    }

    pointSymbol(pointSymbol: string): this {
        this.internal.pointSymbol = pointSymbol;
        return this;
    }

    axisBorderShow(axisBorderShow: boolean): this {
        this.internal.axisBorderShow = axisBorderShow;
        return this;
    }

    barMaxWidth(barMaxWidth: number): this {
        this.internal.barMaxWidth = barMaxWidth;
        return this;
    }
}
