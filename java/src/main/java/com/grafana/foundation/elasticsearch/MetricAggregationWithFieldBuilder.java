// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MetricAggregationWithFieldBuilder implements com.grafana.foundation.cog.Builder<MetricAggregationWithField> {
    protected final MetricAggregationWithField internal;
    
    public MetricAggregationWithFieldBuilder() {
        this.internal = new MetricAggregationWithField();
    }
    public MetricAggregationWithFieldBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public MetricAggregationWithFieldBuilder type(MetricAggregationType type) {
        this.internal.type = type;
        return this;
    }
    
    public MetricAggregationWithFieldBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public MetricAggregationWithFieldBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public MetricAggregationWithField build() {
        return this.internal;
    }
}
