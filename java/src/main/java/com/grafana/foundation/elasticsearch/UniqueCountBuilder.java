// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class UniqueCountBuilder implements com.grafana.foundation.cog.Builder<UniqueCount> {
    protected final UniqueCount internal;
    
    public UniqueCountBuilder() {
        this.internal = new UniqueCount();
    this.internal.type = "cardinality";
    }
    public UniqueCountBuilder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public UniqueCountBuilder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public UniqueCountBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchUniqueCountSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    
    public UniqueCountBuilder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    public UniqueCount build() {
        return this.internal;
    }
}
