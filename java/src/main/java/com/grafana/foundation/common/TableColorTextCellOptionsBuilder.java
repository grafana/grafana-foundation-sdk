// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class TableColorTextCellOptionsBuilder implements com.grafana.foundation.cog.Builder<TableColorTextCellOptions> {
    protected final TableColorTextCellOptions internal;
    
    public TableColorTextCellOptionsBuilder() {
        this.internal = new TableColorTextCellOptions();
    this.internal.type = "color-text";
    }
    public TableColorTextCellOptionsBuilder wrapText(Boolean wrapText) {
    this.internal.wrapText = wrapText;
        return this;
    }
    public TableColorTextCellOptions build() {
        return this.internal;
    }
}
