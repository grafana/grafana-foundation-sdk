// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class TableImageCellOptionsBuilder implements com.grafana.foundation.cog.Builder<TableImageCellOptions> {
    protected final TableImageCellOptions internal;
    
    public TableImageCellOptionsBuilder() {
        this.internal = new TableImageCellOptions();
    this.internal.type = "image";
    }
    public TableImageCellOptionsBuilder alt(String alt) {
    this.internal.alt = alt;
        return this;
    }
    
    public TableImageCellOptionsBuilder title(String title) {
    this.internal.title = title;
        return this;
    }
    public TableImageCellOptions build() {
        return this.internal;
    }
}
