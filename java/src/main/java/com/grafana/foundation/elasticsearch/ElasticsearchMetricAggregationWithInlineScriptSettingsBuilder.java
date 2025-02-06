// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchMetricAggregationWithInlineScriptSettings> {
    protected final ElasticsearchMetricAggregationWithInlineScriptSettings internal;
    
    public ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder() {
        this.internal = new ElasticsearchMetricAggregationWithInlineScriptSettings();
    }
    public ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder script(InlineScript script) {
        this.internal.script = script;
        return this;
    }
    public ElasticsearchMetricAggregationWithInlineScriptSettings build() {
        return this.internal;
    }
}
