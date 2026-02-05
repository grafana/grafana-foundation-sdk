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
    bucketCount(bucketCount: number): this;
    bucketSize(bucketSize: number): this;
    bucketOffset(bucketOffset: number): this;
    legend(legend: cog.Builder<common.VizLegendOptions>): this;
    tooltip(tooltip: cog.Builder<common.VizTooltipOptions>): this;
    combine(combine: boolean): this;
    lineWidth(lineWidth: number): this;
    fillOpacity(fillOpacity: number): this;
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
    stacking(stacking: cog.Builder<common.StackingConfig>): this;
    gradientMode(gradientMode: common.GraphGradientMode): this;
    axisBorderShow(axisBorderShow: boolean): this;
}
