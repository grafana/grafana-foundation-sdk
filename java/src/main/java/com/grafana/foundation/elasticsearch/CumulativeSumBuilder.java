// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class CumulativeSumBuilder implements com.grafana.foundation.cog.Builder<CumulativeSum> {
    protected final CumulativeSum internal;
    
    public CumulativeSumBuilder() {
        this.internal = new CumulativeSum();
        this.internal.type = "cumulative_sum";
    }
    public CumulativeSumBuilder pipelineAgg(String pipelineAgg) {
        this.internal.pipelineAgg = pipelineAgg;
        return this;
    }
    
    public CumulativeSumBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public CumulativeSumBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public CumulativeSumBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchCumulativeSumSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public CumulativeSumBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public CumulativeSum build() {
        return this.internal;
    }
}
