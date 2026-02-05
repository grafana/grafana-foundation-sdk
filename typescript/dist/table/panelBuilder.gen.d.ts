import * as cog from '../cog';
import * as dashboard from '../dashboard';
import * as common from '../common';
export declare class PanelBuilder implements cog.Builder<dashboard.Panel> {
    protected readonly internal: dashboard.Panel;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboard.Panel;
    id(id: number): this;
    targets(targets: cog.Builder<cog.Dataquery>[]): this;
    withTarget(target: cog.Builder<cog.Dataquery>): this;
    title(title: string): this;
    description(description: string): this;
    transparent(transparent: boolean): this;
    datasource(datasource: common.DataSourceRef): this;
    gridPos(gridPos: dashboard.GridPos): this;
    height(h: number): this;
    span(w: number): this;
    links(links: cog.Builder<dashboard.DashboardLink>[]): this;
    repeat(repeat: string): this;
    repeatDirection(repeatDirection: "h" | "v"): this;
    maxPerRow(maxPerRow: number): this;
    maxDataPoints(maxDataPoints: number): this;
    transformations(transformations: dashboard.DataTransformerConfig[]): this;
    withTransformation(transformation: dashboard.DataTransformerConfig): this;
    interval(interval: string): this;
    timeFrom(timeFrom: string): this;
    timeShift(timeShift: string): this;
    hideTimeOverride(hideTimeOverride: boolean): this;
    libraryPanel(libraryPanel: dashboard.LibraryPanelRef): this;
    cacheTimeout(cacheTimeout: string): this;
    queryCachingTTL(queryCachingTTL: number): this;
    displayName(displayName: string): this;
    unit(unit: string): this;
    decimals(decimals: number): this;
    min(min: number): this;
    max(max: number): this;
    mappings(mappings: dashboard.ValueMapping[]): this;
    thresholds(thresholds: cog.Builder<dashboard.ThresholdsConfig>): this;
    colorScheme(color: cog.Builder<dashboard.FieldColor>): this;
    dataLinks(links: cog.Builder<dashboard.DashboardLink>[]): this;
    noValue(noValue: string): this;
    overrides(overrides: {
        matcher: dashboard.MatcherConfig;
        properties: dashboard.DynamicConfigValue[];
    }[]): this;
    withOverride(override: {
        matcher: dashboard.MatcherConfig;
        properties: dashboard.DynamicConfigValue[];
    }): this;
    overrideByName(name: string, properties: dashboard.DynamicConfigValue[]): this;
    overrideByRegexp(regexp: string, properties: dashboard.DynamicConfigValue[]): this;
    overrideByFieldType(fieldType: string, properties: dashboard.DynamicConfigValue[]): this;
    overrideByQuery(queryRefId: string, properties: dashboard.DynamicConfigValue[]): this;
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
