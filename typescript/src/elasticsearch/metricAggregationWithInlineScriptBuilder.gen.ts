// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class MetricAggregationWithInlineScriptBuilder implements cog.Builder<elasticsearch.MetricAggregationWithInlineScript> {
    protected readonly internal: elasticsearch.MetricAggregationWithInlineScript;

    constructor() {
        this.internal = elasticsearch.defaultMetricAggregationWithInlineScript();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.MetricAggregationWithInlineScript {
        return this.internal;
    }

    settings(settings: {
	script?: elasticsearch.InlineScript;
}): this {
        this.internal.settings = settings;
        return this;
    }

    type(type: elasticsearch.MetricAggregationType): this {
        this.internal.type = type;
        return this;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
