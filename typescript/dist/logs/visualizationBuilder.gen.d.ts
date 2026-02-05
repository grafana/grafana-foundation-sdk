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
    showLabels(showLabels: boolean): this;
    showCommonLabels(showCommonLabels: boolean): this;
    showTime(showTime: boolean): this;
    showLogContextToggle(showLogContextToggle: boolean): this;
    wrapLogMessage(wrapLogMessage: boolean): this;
    prettifyLogMessage(prettifyLogMessage: boolean): this;
    enableLogDetails(enableLogDetails: boolean): this;
    sortOrder(sortOrder: common.LogsSortOrder): this;
    dedupStrategy(dedupStrategy: common.LogsDedupStrategy): this;
    enableInfiniteScrolling(enableInfiniteScrolling: boolean): this;
    onClickFilterLabel(onClickFilterLabel: any): this;
    onClickFilterOutLabel(onClickFilterOutLabel: any): this;
    isFilterLabelActive(isFilterLabelActive: any): this;
    onClickFilterString(onClickFilterString: any): this;
    onClickFilterOutString(onClickFilterOutString: any): this;
    onClickShowField(onClickShowField: any): this;
    onClickHideField(onClickHideField: any): this;
    logRowMenuIconsBefore(logRowMenuIconsBefore: any): this;
    logRowMenuIconsAfter(logRowMenuIconsAfter: any): this;
    onNewLogsReceived(onNewLogsReceived: any): this;
    displayedFields(displayedFields: string[]): this;
}
