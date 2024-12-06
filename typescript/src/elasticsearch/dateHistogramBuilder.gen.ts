// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class DateHistogramBuilder implements cog.Builder<elasticsearch.DateHistogram> {
    protected readonly internal: elasticsearch.DateHistogram;

    constructor() {
        this.internal = elasticsearch.defaultDateHistogram();
        this.internal.type = "date_histogram";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.DateHistogram {
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
	trimEdges?: string;
	offset?: string;
	timeZone?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }
}
