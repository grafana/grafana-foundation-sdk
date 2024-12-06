// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class HistogramSettingsBuilder implements cog.Builder<elasticsearch.HistogramSettings> {
    protected readonly internal: elasticsearch.HistogramSettings;

    constructor() {
        this.internal = elasticsearch.defaultHistogramSettings();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.HistogramSettings {
        return this.internal;
    }

    interval(interval: string): this {
        this.internal.interval = interval;
        return this;
    }

    minDocCount(minDocCount: string): this {
        this.internal.min_doc_count = minDocCount;
        return this;
    }
}
