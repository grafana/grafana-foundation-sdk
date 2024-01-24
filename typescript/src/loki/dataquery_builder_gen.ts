// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as loki from '../loki';

export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    private readonly internal: loki.dataquery;

    constructor() {
        this.internal = loki.defaultDataquery();
    }

    build(): loki.dataquery {
        return this.internal;
    }

    expr(expr: string): this {
        this.internal.expr = expr;
        return this;
    }

    legendFormat(legendFormat: string): this {
        this.internal.legendFormat = legendFormat;
        return this;
    }

    maxLines(maxLines: number): this {
        this.internal.maxLines = maxLines;
        return this;
    }

    resolution(resolution: number): this {
        this.internal.resolution = resolution;
        return this;
    }

    editorMode(editorMode: loki.QueryEditorMode): this {
        this.internal.editorMode = editorMode;
        return this;
    }

    range(range: boolean): this {
        this.internal.range = range;
        return this;
    }

    instant(instant: boolean): this {
        this.internal.instant = instant;
        return this;
    }

    step(step: string): this {
        this.internal.step = step;
        return this;
    }

    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    datasource(datasource: any): this {
        this.internal.datasource = datasource;
        return this;
    }
}
