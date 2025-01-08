// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class DerivativeBuilder implements com.grafana.foundation.cog.Builder<Derivative> {
    protected final Derivative internal;
    
    public DerivativeBuilder() {
        this.internal = new Derivative();
    this.internal.type = "derivative";
    }
    public DerivativeBuilder pipelineAgg(String pipelineAgg) {
    this.internal.pipelineAgg = pipelineAgg;
        return this;
    }
    
    public DerivativeBuilder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public DerivativeBuilder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public DerivativeBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchDerivativeSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    
    public DerivativeBuilder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    public Derivative build() {
        return this.internal;
    }
}
