// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class TableColoredBackgroundCellOptionsBuilder implements com.grafana.foundation.cog.Builder<TableColoredBackgroundCellOptions> {
    protected final TableColoredBackgroundCellOptions internal;
    
    public TableColoredBackgroundCellOptionsBuilder() {
        this.internal = new TableColoredBackgroundCellOptions();
    this.internal.type = "color-background";
    }
    public TableColoredBackgroundCellOptionsBuilder mode(TableCellBackgroundDisplayMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public TableColoredBackgroundCellOptionsBuilder applyToRow(Boolean applyToRow) {
    this.internal.applyToRow = applyToRow;
        return this;
    }
    public TableColoredBackgroundCellOptions build() {
        return this.internal;
    }
}
