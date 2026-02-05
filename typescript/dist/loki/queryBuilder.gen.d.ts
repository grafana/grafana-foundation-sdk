import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as loki from '../loki';
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
    direction(direction: loki.LokiQueryDirection): this;
}
