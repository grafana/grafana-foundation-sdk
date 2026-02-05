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
    frameIndex(frameIndex: number): this;
    showHeader(showHeader: boolean): this;
    showTypeIcons(showTypeIcons: boolean): this;
    sortBy(sortBy: cog.Builder<common.TableSortByFieldState>[]): this;
    footer(footer: cog.Builder<common.TableFooterOptions>): this;
    cellHeight(cellHeight: common.TableCellHeight): this;
    width(width: number): this;
    minWidth(minWidth: number): this;
    align(align: common.FieldTextAlignment): this;
    displayMode(displayMode: common.TableCellDisplayMode): this;
    cellOptions(cellOptions: common.TableCellOptions): this;
    hidden(hidden: boolean): this;
    inspect(inspect: boolean): this;
    filterable(filterable: boolean): this;
    hideHeader(hideHeader: boolean): this;
}
