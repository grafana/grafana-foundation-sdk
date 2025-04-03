// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class BasePipelineMetricAggregationBuilder implements com.grafana.foundation.cog.Builder<BasePipelineMetricAggregation> {
    protected final BasePipelineMetricAggregation internal;
    
    public BasePipelineMetricAggregationBuilder() {
        this.internal = new BasePipelineMetricAggregation();
    }
    public BasePipelineMetricAggregationBuilder pipelineAgg(String pipelineAgg) {
        this.internal.pipelineAgg = pipelineAgg;
        return this;
    }
    
    public BasePipelineMetricAggregationBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public BasePipelineMetricAggregationBuilder type(String type) {
        this.internal.type = type;
        return this;
    }
    
    public BasePipelineMetricAggregationBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public BasePipelineMetricAggregationBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public BasePipelineMetricAggregation build() {
        return this.internal;
    }
}
