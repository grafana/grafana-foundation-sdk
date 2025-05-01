// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MetricAggregationWithInlineScriptBuilder implements com.grafana.foundation.cog.Builder<MetricAggregationWithInlineScript> {
    protected final MetricAggregationWithInlineScript internal;
    
    public MetricAggregationWithInlineScriptBuilder() {
        this.internal = new MetricAggregationWithInlineScript();
    }
    public MetricAggregationWithInlineScriptBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchMetricAggregationWithInlineScriptSettings> settings) {
    ElasticsearchMetricAggregationWithInlineScriptSettings settingsResource = settings.build();
        this.internal.settings = settingsResource;
        return this;
    }
    
    public MetricAggregationWithInlineScriptBuilder type(MetricAggregationType type) {
        this.internal.type = type;
        return this;
    }
    
    public MetricAggregationWithInlineScriptBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public MetricAggregationWithInlineScriptBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public MetricAggregationWithInlineScript build() {
        return this.internal;
    }
}
