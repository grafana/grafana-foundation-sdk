// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class PipelineVariableBuilder implements com.grafana.foundation.cog.Builder<PipelineVariable> {
    protected final PipelineVariable internal;
    
    public PipelineVariableBuilder() {
        this.internal = new PipelineVariable();
    }
    public PipelineVariableBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public PipelineVariableBuilder pipelineAgg(String pipelineAgg) {
    this.internal.pipelineAgg = pipelineAgg;
        return this;
    }
    public PipelineVariable build() {
        return this.internal;
    }
}
