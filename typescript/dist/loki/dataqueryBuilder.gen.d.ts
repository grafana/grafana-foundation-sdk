import * as cog from '../cog';
import * as loki from '../loki';
import * as common from '../common';
export declare class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: loki.dataquery;
    constructor();
    /**
     * Builds the object.
     */
    build(): loki.dataquery;
    expr(expr: string): this;
    legendFormat(legendFormat: string): this;
    maxLines(maxLines: number): this;
    resolution(resolution: number): this;
    editorMode(editorMode: loki.QueryEditorMode): this;
    range(range: boolean): this;
    instant(instant: boolean): this;
    step(step: string): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    datasource(datasource: common.DataSourceRef): this;
    direction(direction: loki.LokiQueryDirection): this;
}
