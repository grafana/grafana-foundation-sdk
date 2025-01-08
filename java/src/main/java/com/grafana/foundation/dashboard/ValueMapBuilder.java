// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.Map;

public class ValueMapBuilder implements com.grafana.foundation.cog.Builder<ValueMap> {
    protected final ValueMap internal;
    
    public ValueMapBuilder() {
        this.internal = new ValueMap();
    this.internal.type = "value";
    }
    public ValueMapBuilder options(Map<String, ValueMappingResult> options) {
    this.internal.options = options;
        return this;
    }
    public ValueMap build() {
        return this.internal;
    }
}
