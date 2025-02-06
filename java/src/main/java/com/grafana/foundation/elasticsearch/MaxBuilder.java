// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MaxBuilder implements com.grafana.foundation.cog.Builder<Max> {
    protected final Max internal;
    
    public MaxBuilder() {
        this.internal = new Max();
        this.internal.type = "max";
    }
    public MaxBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public MaxBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public MaxBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchMaxSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public MaxBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public Max build() {
        return this.internal;
    }
}
