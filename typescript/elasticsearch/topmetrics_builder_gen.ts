// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class TopMetricsBuilder implements cog.OptionsBuilder<elasticsearch.TopMetrics> {
    private readonly internal: elasticsearch.TopMetrics;

    constructor() {
        this.internal = elasticsearch.defaultTopMetrics();
        this.internal.type = "top_metrics";
    }

    build(): elasticsearch.TopMetrics {
        return this.internal;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    settings(settings: {
	order?: string;
	orderBy?: string;
	metrics?: string[];
}): this {
        this.internal.settings = settings;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
