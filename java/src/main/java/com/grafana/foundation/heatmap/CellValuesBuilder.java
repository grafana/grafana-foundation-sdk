// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;


public class CellValuesBuilder implements com.grafana.foundation.cog.Builder<CellValues> {
    protected final CellValues internal;
    
    public CellValuesBuilder() {
        this.internal = new CellValues();
    }
    public CellValuesBuilder unit(String unit) {
        this.internal.unit = unit;
        return this;
    }
    
    public CellValuesBuilder decimals(Float decimals) {
        this.internal.decimals = decimals;
        return this;
    }
    public CellValues build() {
        return this.internal;
    }
}
