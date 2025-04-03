// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.tempo;


public class TraceqlFilterBuilder implements com.grafana.foundation.cog.Builder<TraceqlFilter> {
    protected final TraceqlFilter internal;
    
    public TraceqlFilterBuilder() {
        this.internal = new TraceqlFilter();
    }
    public TraceqlFilterBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public TraceqlFilterBuilder tag(String tag) {
        this.internal.tag = tag;
        return this;
    }
    
    public TraceqlFilterBuilder operator(String operator) {
        this.internal.operator = operator;
        return this;
    }
    
    public TraceqlFilterBuilder value(StringOrArrayOfString value) {
        this.internal.value = value;
        return this;
    }
    
    public TraceqlFilterBuilder valueType(String valueType) {
        this.internal.valueType = valueType;
        return this;
    }
    
    public TraceqlFilterBuilder scope(TraceqlSearchScope scope) {
        this.internal.scope = scope;
        return this;
    }
    public TraceqlFilter build() {
        return this.internal;
    }
}
