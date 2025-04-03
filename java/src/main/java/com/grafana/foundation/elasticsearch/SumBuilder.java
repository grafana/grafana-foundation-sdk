// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class SumBuilder implements com.grafana.foundation.cog.Builder<Sum> {
    protected final Sum internal;
    
    public SumBuilder() {
        this.internal = new Sum();
    }
    public SumBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public SumBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public SumBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchSumSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public SumBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public Sum build() {
        return this.internal;
    }
}
