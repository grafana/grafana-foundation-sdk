// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class RateBuilder implements com.grafana.foundation.cog.Builder<Rate> {
    protected final Rate internal;
    
    public RateBuilder() {
        this.internal = new Rate();
        this.internal.type = "rate";
    }
    public RateBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public RateBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public RateBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchRateSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public RateBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public Rate build() {
        return this.internal;
    }
}
