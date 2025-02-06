// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class CountBuilder implements com.grafana.foundation.cog.Builder<Count> {
    protected final Count internal;
    
    public CountBuilder() {
        this.internal = new Count();
        this.internal.type = "count";
    }
    public CountBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public CountBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public Count build() {
        return this.internal;
    }
}
