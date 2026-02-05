import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as prometheus from '../prometheus';
export declare class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind;
    version(version: string): this;
    datasource(ref: {
        name?: string;
    }): this;
    expr(expr: string): this;
    instant(): this;
    range(): this;
    exemplar(exemplar: boolean): this;
    editorMode(editorMode: prometheus.QueryEditorMode): this;
    format(format: prometheus.PromQueryFormat): this;
    legendFormat(legendFormat: string): this;
    intervalFactor(intervalFactor: number): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    interval(interval: string): this;
    rangeAndInstant(): this;
}
