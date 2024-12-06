// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class HistogramBuilder implements cog.Builder<elasticsearch.Histogram> {
    protected readonly internal: elasticsearch.Histogram;

    constructor() {
        this.internal = elasticsearch.defaultHistogram();
        this.internal.type = "histogram";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.Histogram {
        return this.internal;
    }

    field(field: string): this {
        this.internal.field = field;
        return this;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    settings(settings: {
	interval?: string;
	min_doc_count?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }
}
