// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as prometheus from '../prometheus';

export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    private readonly internal: prometheus.dataquery;

    constructor() {
        this.internal = prometheus.defaultDataquery();
    }

    build(): prometheus.dataquery {
        return this.internal;
    }

    expr(expr: string): this {
        this.internal.expr = expr;
        return this;
    }

    instant(instant: boolean): this {
        this.internal.instant = instant;
        return this;
    }

    range(range: boolean): this {
        this.internal.range = range;
        return this;
    }

    exemplar(exemplar: boolean): this {
        this.internal.exemplar = exemplar;
        return this;
    }

    editorMode(editorMode: prometheus.QueryEditorMode): this {
        this.internal.editorMode = editorMode;
        return this;
    }

    format(format: prometheus.PromQueryFormat): this {
        this.internal.format = format;
        return this;
    }

    legendFormat(legendFormat: string): this {
        this.internal.legendFormat = legendFormat;
        return this;
    }

    intervalFactor(intervalFactor: number): this {
        this.internal.intervalFactor = intervalFactor;
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

    // An additional lower limit for the step parameter of the Prometheus query and for the
    // `$__interval` and `$__rate_interval` variables.
    interval(interval: string): this {
        this.internal.interval = interval;
        return this;
    }
}
