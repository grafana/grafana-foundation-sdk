// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MinBuilder implements com.grafana.foundation.cog.Builder<Min> {
    protected final Min internal;
    
    public MinBuilder() {
        this.internal = new Min();
        this.internal.type = "min";
    }
    public MinBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public MinBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public MinBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchMinSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public MinBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public Min build() {
        return this.internal;
    }
}
