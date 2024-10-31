// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as loki from '../loki';
import * as dashboard from '../dashboard';

export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: loki.dataquery;

    constructor() {
        this.internal = loki.defaultDataquery();
    }

    /**
     * Builds the object.
     */
    build(): loki.dataquery {
        return this.internal;
    }

    // The LogQL query.
    expr(expr: string): this {
        this.internal.expr = expr;
        return this;
    }

    // Used to override the name of the series.
    legendFormat(legendFormat: string): this {
        this.internal.legendFormat = legendFormat;
        return this;
    }

    // Used to limit the number of log rows returned.
    maxLines(maxLines: number): this {
        this.internal.maxLines = maxLines;
        return this;
    }

    // @deprecated, now use step.
    resolution(resolution: number): this {
        this.internal.resolution = resolution;
        return this;
    }

    editorMode(editorMode: loki.QueryEditorMode): this {
        this.internal.editorMode = editorMode;
        return this;
    }

    // @deprecated, now use queryType.
    range(range: boolean): this {
        this.internal.range = range;
        return this;
    }

    // @deprecated, now use queryType.
    instant(instant: boolean): this {
        this.internal.instant = instant;
        return this;
    }

    // Used to set step value for range queries.
    step(step: string): this {
        this.internal.step = step;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }
}
