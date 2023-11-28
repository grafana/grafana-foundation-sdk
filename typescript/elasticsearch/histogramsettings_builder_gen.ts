// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class HistogramSettingsBuilder implements cog.Builder<elasticsearch.HistogramSettings> {
    private readonly internal: elasticsearch.HistogramSettings;

    constructor() {
        this.internal = elasticsearch.defaultHistogramSettings();
    }

    build(): elasticsearch.HistogramSettings {
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
}
