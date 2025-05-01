// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;
import java.util.LinkedList;

public class PipelineMetricAggregationWithMultipleBucketPathsBuilder implements com.grafana.foundation.cog.Builder<PipelineMetricAggregationWithMultipleBucketPaths> {
    protected final PipelineMetricAggregationWithMultipleBucketPaths internal;
    
    public PipelineMetricAggregationWithMultipleBucketPathsBuilder() {
        this.internal = new PipelineMetricAggregationWithMultipleBucketPaths();
    }
    public PipelineMetricAggregationWithMultipleBucketPathsBuilder pipelineVariables(List<com.grafana.foundation.cog.Builder<PipelineVariable>> pipelineVariables) {
        List<PipelineVariable> pipelineVariablesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<PipelineVariable> r1 : pipelineVariables) {
                PipelineVariable pipelineVariablesDepth1 = r1.build();
                pipelineVariablesResources.add(pipelineVariablesDepth1); 
        }
        this.internal.pipelineVariables = pipelineVariablesResources;
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
