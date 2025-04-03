// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ExtendedStatBuilder implements com.grafana.foundation.cog.Builder<ExtendedStat> {
    protected final ExtendedStat internal;
    
    public ExtendedStatBuilder() {
        this.internal = new ExtendedStat();
    }
    public ExtendedStatBuilder label(String label) {
        this.internal.label = label;
        return this;
    }
    
    public ExtendedStatBuilder value(ExtendedStatMetaType value) {
        this.internal.value = value;
        return this;
    }
    public ExtendedStat build() {
        return this.internal;
    }
}
