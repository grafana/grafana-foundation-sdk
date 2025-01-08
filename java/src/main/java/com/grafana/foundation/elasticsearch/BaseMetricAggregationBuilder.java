// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class BaseMetricAggregationBuilder implements com.grafana.foundation.cog.Builder<BaseMetricAggregation> {
    protected final BaseMetricAggregation internal;
    
    public BaseMetricAggregationBuilder() {
        this.internal = new BaseMetricAggregation();
    }
    public BaseMetricAggregationBuilder type(com.grafana.foundation.cog.Builder<MetricAggregationType> type) {
    this.internal.type = type.build();
        return this;
    }
    
    public BaseMetricAggregationBuilder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public BaseMetricAggregationBuilder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    public BaseMetricAggregation build() {
        return this.internal;
    }
}
