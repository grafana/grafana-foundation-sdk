import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as candlestick from '../candlestick';
import * as common from '../common';
export declare class VisualizationBuilder implements cog.Builder<dashboardv2beta1.VizConfigKind> {
    protected readonly internal: dashboardv2beta1.VizConfigKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.VizConfigKind;
    displayName(displayName: string): this;
    displayNameFromDS(displayNameFromDS: string): this;
    description(description: string): this;
    path(path: string): this;
    unit(unit: string): this;
    decimals(decimals: number): this;
    min(min: number): this;
    max(max: number): this;
    mappings(mappings: dashboardv2beta1.ValueMapping[]): this;
    thresholds(thresholds: cog.Builder<dashboardv2beta1.ThresholdsConfig>): this;
    colorScheme(color: cog.Builder<dashboardv2beta1.FieldColor>): this;
    dataLinks(links: any[]): this;
    actions(actions: cog.Builder<dashboardv2beta1.Action>[]): this;
    noValue(noValue: string): this;
    overrides(overrides: {
        matcher: dashboardv2beta1.MatcherConfig;
        properties: dashboardv2beta1.DynamicConfigValue[];
    }[]): this;
    override(override: {
        matcher: dashboardv2beta1.MatcherConfig;
        properties: dashboardv2beta1.DynamicConfigValue[];
    }): this;
    overrideByName(name: string, properties: dashboardv2beta1.DynamicConfigValue[]): this;
    overrideByRegexp(regexp: string, properties: dashboardv2beta1.DynamicConfigValue[]): this;
    overrideByFieldType(fieldType: string, properties: dashboardv2beta1.DynamicConfigValue[]): this;
    overrideByQuery(queryRefId: string, properties: dashboardv2beta1.DynamicConfigValue[]): this;
    mode(mode: candlestick.VizDisplayMode): this;
    candleStyle(candleStyle: candlestick.CandleStyle): this;
    colorStrategy(colorStrategy: candlestick.ColorStrategy): this;
    fields(fields: cog.Builder<candlestick.CandlestickFieldMap>): this;
    colors(colors: cog.Builder<candlestick.CandlestickColors>): this;
    legend(legend: cog.Builder<common.VizLegendOptions>): this;
    tooltip(tooltip: cog.Builder<common.VizTooltipOptions>): this;
    includeAllFields(includeAllFields: boolean): this;
    drawStyle(drawStyle: common.GraphDrawStyle): this;
    gradientMode(gradientMode: common.GraphGradientMode): this;
    thresholdsStyle(thresholdsStyle: cog.Builder<common.GraphThresholdsStyleConfig>): this;
    transform(transform: common.GraphTransform): this;
    lineColor(lineColor: string): this;
    lineWidth(lineWidth: number): this;
    lineInterpolation(lineInterpolation: common.LineInterpolation): this;
    lineStyle(lineStyle: cog.Builder<common.LineStyle>): this;
    fillColor(fillColor: string): this;
    fillOpacity(fillOpacity: number): this;
    showPoints(showPoints: common.VisibilityMode): this;
    pointSize(pointSize: number): this;
    pointColor(pointColor: string): this;
    axisPlacement(axisPlacement: common.AxisPlacement): this;
    axisColorMode(axisColorMode: common.AxisColorMode): this;
    axisLabel(axisLabel: string): this;
    axisWidth(axisWidth: number): this;
    axisSoftMin(axisSoftMin: number): this;
    axisSoftMax(axisSoftMax: number): this;
    axisGridShow(axisGridShow: boolean): this;
    scaleDistribution(scaleDistribution: cog.Builder<common.ScaleDistributionConfig>): this;
    axisCenteredZero(axisCenteredZero: boolean): this;
    barAlignment(barAlignment: common.BarAlignment): this;
    barWidthFactor(barWidthFactor: number): this;
    stacking(stacking: cog.Builder<common.StackingConfig>): this;
    hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>): this;
    insertNulls(insertNulls: boolean | number): this;
    spanNulls(spanNulls: boolean | number): this;
    fillBelowTo(fillBelowTo: string): this;
    pointSymbol(pointSymbol: string): this;
    axisBorderShow(axisBorderShow: boolean): this;
    barMaxWidth(barMaxWidth: number): this;
}
