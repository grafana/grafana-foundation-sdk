// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MovingFunctionBuilder implements com.grafana.foundation.cog.Builder<MovingFunction> {
    protected final MovingFunction internal;
    
    public MovingFunctionBuilder() {
        this.internal = new MovingFunction();
    }
    public MovingFunctionBuilder pipelineAgg(String pipelineAgg) {
        this.internal.pipelineAgg = pipelineAgg;
        return this;
    }
    
    public MovingFunctionBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public MovingFunctionBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public MovingFunctionBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchMovingFunctionSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public MovingFunctionBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public MovingFunction build() {
        return this.internal;
    }
}
