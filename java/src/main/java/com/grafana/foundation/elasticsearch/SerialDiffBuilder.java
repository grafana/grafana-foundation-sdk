// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class SerialDiffBuilder implements com.grafana.foundation.cog.Builder<SerialDiff> {
    protected final SerialDiff internal;
    
    public SerialDiffBuilder() {
        this.internal = new SerialDiff();
    this.internal.type = "serial_diff";
    }
    public SerialDiffBuilder pipelineAgg(String pipelineAgg) {
    this.internal.pipelineAgg = pipelineAgg;
        return this;
    }
    
    public SerialDiffBuilder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public SerialDiffBuilder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public SerialDiffBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchSerialDiffSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    
    public SerialDiffBuilder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    public SerialDiff build() {
        return this.internal;
    }
}
