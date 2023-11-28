// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class DateHistogramSettingsBuilder implements cog.OptionsBuilder<elasticsearch.DateHistogramSettings> {
    private readonly internal: elasticsearch.DateHistogramSettings;

    constructor() {
        this.internal = elasticsearch.defaultDateHistogramSettings();
    }

    build(): elasticsearch.DateHistogramSettings {
        return this.internal;
    }

    interval(interval: string): this {
        this.internal.interval = interval;
        return this;
    }

    min_doc_count(min_doc_count: string): this {
        this.internal.min_doc_count = min_doc_count;
        return this;
    }

    trimEdges(trimEdges: string): this {
        this.internal.trimEdges = trimEdges;
        return this;
    }

    offset(offset: string): this {
        this.internal.offset = offset;
        return this;
    }

    timeZone(timeZone: string): this {
        this.internal.timeZone = timeZone;
        return this;
    }
}
