// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class AverageBuilder implements com.grafana.foundation.cog.Builder<Average> {
    protected final Average internal;
    
    public AverageBuilder() {
        this.internal = new Average();
    this.internal.type = "avg";
    }
    public AverageBuilder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public AverageBuilder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public AverageBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchAverageSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    
    public AverageBuilder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    public Average build() {
        return this.internal;
    }
}
