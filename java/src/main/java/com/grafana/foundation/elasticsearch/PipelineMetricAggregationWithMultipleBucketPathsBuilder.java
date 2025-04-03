// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;

public class PipelineMetricAggregationWithMultipleBucketPathsBuilder implements com.grafana.foundation.cog.Builder<PipelineMetricAggregationWithMultipleBucketPaths> {
    protected final PipelineMetricAggregationWithMultipleBucketPaths internal;
    
    public PipelineMetricAggregationWithMultipleBucketPathsBuilder() {
        this.internal = new PipelineMetricAggregationWithMultipleBucketPaths();
    }
    public PipelineMetricAggregationWithMultipleBucketPathsBuilder pipelineVariables(com.grafana.foundation.cog.Builder<List<PipelineVariable>> pipelineVariables) {
        this.internal.pipelineVariables = pipelineVariables.build();
        return this;
    }
    
    public PipelineMetricAggregationWithMultipleBucketPathsBuilder type(MetricAggregationType type) {
        this.internal.type = type;
        return this;
    }
    
    public PipelineMetricAggregationWithMultipleBucketPathsBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public PipelineMetricAggregationWithMultipleBucketPathsBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public PipelineMetricAggregationWithMultipleBucketPaths build() {
        return this.internal;
    }
}
