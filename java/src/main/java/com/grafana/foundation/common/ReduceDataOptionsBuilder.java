// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import java.util.List;

public class ReduceDataOptionsBuilder implements com.grafana.foundation.cog.Builder<ReduceDataOptions> {
    protected final ReduceDataOptions internal;
    
    public ReduceDataOptionsBuilder() {
        this.internal = new ReduceDataOptions();
    }
    public ReduceDataOptionsBuilder values(Boolean values) {
        this.internal.values = values;
        return this;
    }
    
    public ReduceDataOptionsBuilder limit(Double limit) {
        this.internal.limit = limit;
        return this;
    }
    
    public ReduceDataOptionsBuilder calcs(List<String> calcs) {
        this.internal.calcs = calcs;
        return this;
    }
    
    public ReduceDataOptionsBuilder fields(String fields) {
        this.internal.fields = fields;
        return this;
    }
    public ReduceDataOptions build() {
        return this.internal;
    }
}
