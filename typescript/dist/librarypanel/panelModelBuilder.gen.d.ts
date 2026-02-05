import * as cog from '../cog';
import * as librarypanel from '../librarypanel';
import * as common from '../common';
import * as dashboard from '../dashboard';
export declare class PanelModelBuilder implements cog.Builder<librarypanel.PanelModel> {
    protected readonly internal: librarypanel.PanelModel;
    constructor();
    /**
     * Builds the object.
     */
    build(): librarypanel.PanelModel;
    type(type: string): this;
    pluginVersion(pluginVersion: string): this;
    targets(targets: cog.Builder<cog.Dataquery>[]): this;
    title(title: string): this;
    description(description: string): this;
    transparent(transparent: boolean): this;
    datasource(datasource: common.DataSourceRef): this;
    links(links: cog.Builder<dashboard.DashboardLink>[]): this;
    repeat(repeat: string): this;
    repeatDirection(repeatDirection: "h" | "v"): this;
    maxPerRow(maxPerRow: number): this;
    maxDataPoints(maxDataPoints: number): this;
    transformations(transformations: dashboard.DataTransformerConfig[]): this;
    interval(interval: string): this;
    timeFrom(timeFrom: string): this;
    timeShift(timeShift: string): this;
    hideTimeOverride(hideTimeOverride: boolean): this;
    cacheTimeout(cacheTimeout: string): this;
    queryCachingTTL(queryCachingTTL: number): this;
    options(options: any): this;
    fieldConfig(fieldConfig: dashboard.FieldConfigSource): this;
}
