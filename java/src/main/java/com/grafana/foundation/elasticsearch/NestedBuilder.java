// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class NestedBuilder implements com.grafana.foundation.cog.Builder<Nested> {
    protected final Nested internal;
    
    public NestedBuilder() {
        this.internal = new Nested();
        this.internal.type = "nested";
    }
    public NestedBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public NestedBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public NestedBuilder settings(Object settings) {
        this.internal.settings = settings;
        return this;
    }
    public Nested build() {
        return this.internal;
    }
}
