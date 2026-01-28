// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.datagrid;


public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder selectedSeries(Integer selectedSeries) {
        if (!(selectedSeries >= 0)) {
            throw new IllegalArgumentException("selectedSeries must be >= 0");
        }
        this.internal.selectedSeries = selectedSeries;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
