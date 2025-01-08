// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MetricAggregationTypeBuilder implements com.grafana.foundation.cog.Builder<MetricAggregationType> {
    protected final MetricAggregationType internal;
    
    public MetricAggregationTypeBuilder() {
        this.internal = new MetricAggregationType();
    this.internal.string = "count";
    }
    public MetricAggregationTypeBuilder pipelineMetricAggregationType(PipelineMetricAggregationType pipelineMetricAggregationType) {
    this.internal.pipelineMetricAggregationType = pipelineMetricAggregationType;
        return this;
    }
    public MetricAggregationType build() {
        return this.internal;
    }
}
