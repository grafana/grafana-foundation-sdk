// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class PercentilesBuilder implements com.grafana.foundation.cog.Builder<Percentiles> {
    protected final Percentiles internal;
    
    public PercentilesBuilder() {
        this.internal = new Percentiles();
    }
    public PercentilesBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public PercentilesBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public PercentilesBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchPercentilesSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public PercentilesBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public Percentiles build() {
        return this.internal;
    }
}
