// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class TableAutoCellOptionsBuilder implements com.grafana.foundation.cog.Builder<TableAutoCellOptions> {
    protected final TableAutoCellOptions internal;
    
    public TableAutoCellOptionsBuilder() {
        this.internal = new TableAutoCellOptions();
    this.internal.type = "auto";
    }
    public TableAutoCellOptionsBuilder wrapText(Boolean wrapText) {
    this.internal.wrapText = wrapText;
        return this;
    }
    public TableAutoCellOptions build() {
        return this.internal;
    }
}
