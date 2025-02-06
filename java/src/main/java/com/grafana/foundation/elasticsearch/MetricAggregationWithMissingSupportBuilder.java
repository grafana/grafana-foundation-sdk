// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MetricAggregationWithMissingSupportBuilder implements com.grafana.foundation.cog.Builder<MetricAggregationWithMissingSupport> {
    protected final MetricAggregationWithMissingSupport internal;
    
    public MetricAggregationWithMissingSupportBuilder() {
        this.internal = new MetricAggregationWithMissingSupport();
    }
    public MetricAggregationWithMissingSupportBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchMetricAggregationWithMissingSupportSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public MetricAggregationWithMissingSupportBuilder type(MetricAggregationType type) {
        this.internal.type = type;
        return this;
    }
    
    public MetricAggregationWithMissingSupportBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public MetricAggregationWithMissingSupportBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public MetricAggregationWithMissingSupport build() {
        return this.internal;
    }
}
