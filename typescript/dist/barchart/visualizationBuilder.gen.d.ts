import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
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
    xField(xField: string): this;
    colorByField(colorByField: string): this;
    orientation(orientation: common.VizOrientation): this;
    barRadius(barRadius: number): this;
    xTickLabelRotation(xTickLabelRotation: number): this;
    xTickLabelMaxLength(xTickLabelMaxLength: number): this;
    xTickLabelSpacing(xTickLabelSpacing: number): this;
    stacking(stacking: common.StackingMode): this;
    showValue(showValue: common.VisibilityMode): this;
    barWidth(barWidth: number): this;
    groupWidth(groupWidth: number): this;
    legend(legend: cog.Builder<common.VizLegendOptions>): this;
    tooltip(tooltip: cog.Builder<common.VizTooltipOptions>): this;
    text(text: cog.Builder<common.VizTextDisplayOptions>): this;
    fullHighlight(fullHighlight: boolean): this;
    lineWidth(lineWidth: number): this;
    fillOpacity(fillOpacity: number): this;
    gradientMode(gradientMode: common.GraphGradientMode): this;
    axisPlacement(axisPlacement: common.AxisPlacement): this;
    axisColorMode(axisColorMode: common.AxisColorMode): this;
    axisLabel(axisLabel: string): this;
    axisWidth(axisWidth: number): this;
    axisSoftMin(axisSoftMin: number): this;
    axisSoftMax(axisSoftMax: number): this;
    axisGridShow(axisGridShow: boolean): this;
    scaleDistribution(scaleDistribution: cog.Builder<common.ScaleDistributionConfig>): this;
    axisCenteredZero(axisCenteredZero: boolean): this;
    hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>): this;
    thresholdsStyle(thresholdsStyle: cog.Builder<common.GraphThresholdsStyleConfig>): this;
    axisBorderShow(axisBorderShow: boolean): this;
}
