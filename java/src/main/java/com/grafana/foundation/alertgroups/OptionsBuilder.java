// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertgroups;


public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder labels(String labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public OptionsBuilder alertmanager(String alertmanager) {
        this.internal.alertmanager = alertmanager;
        return this;
    }
    
    public OptionsBuilder expandAll(Boolean expandAll) {
        this.internal.expandAll = expandAll;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
