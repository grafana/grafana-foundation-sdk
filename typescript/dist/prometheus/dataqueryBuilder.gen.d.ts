import * as cog from '../cog';
import * as prometheus from '../prometheus';
import * as common from '../common';
export declare class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: prometheus.dataquery;
    constructor();
    /**
     * Builds the object.
     */
    build(): prometheus.dataquery;
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
    datasource(datasource: common.DataSourceRef): this;
    interval(interval: string): this;
    rangeAndInstant(): this;
}
