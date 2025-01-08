// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class TableBarGaugeCellOptionsBuilder implements com.grafana.foundation.cog.Builder<TableBarGaugeCellOptions> {
    protected final TableBarGaugeCellOptions internal;
    
    public TableBarGaugeCellOptionsBuilder() {
        this.internal = new TableBarGaugeCellOptions();
    this.internal.type = "gauge";
    }
    public TableBarGaugeCellOptionsBuilder mode(BarGaugeDisplayMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public TableBarGaugeCellOptionsBuilder valueDisplayMode(BarGaugeValueMode valueDisplayMode) {
    this.internal.valueDisplayMode = valueDisplayMode;
        return this;
    }
    public TableBarGaugeCellOptions build() {
        return this.internal;
    }
}
