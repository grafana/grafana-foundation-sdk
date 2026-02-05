import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as heatmap from '../heatmap';
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
    calculate(calculate: boolean): this;
    calculation(calculation: cog.Builder<common.HeatmapCalculationOptions>): this;
    color(color: cog.Builder<heatmap.HeatmapColorOptions>): this;
    filterValues(filterValues: cog.Builder<heatmap.FilterValueRange>): this;
    rowsFrame(rowsFrame: cog.Builder<heatmap.RowsHeatmapOptions>): this;
    showValue(showValue: common.VisibilityMode): this;
    cellGap(cellGap: number): this;
    cellRadius(cellRadius: number): this;
    cellValues(cellValues: cog.Builder<heatmap.CellValues>): this;
    yAxis(yAxis: cog.Builder<heatmap.YAxisConfig>): this;
    legend(legend: cog.Builder<heatmap.HeatmapLegend>): this;
    tooltip(tooltip: cog.Builder<heatmap.HeatmapTooltip>): this;
    exemplars(exemplars: cog.Builder<heatmap.ExemplarConfig>): this;
    selectionMode(selectionMode: heatmap.HeatmapSelectionMode): this;
    scaleDistribution(scaleDistribution: cog.Builder<common.ScaleDistributionConfig>): this;
    hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>): this;
}
