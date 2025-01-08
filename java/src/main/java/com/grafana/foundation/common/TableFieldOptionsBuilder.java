// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class TableFieldOptionsBuilder implements com.grafana.foundation.cog.Builder<TableFieldOptions> {
    protected final TableFieldOptions internal;
    
    public TableFieldOptionsBuilder() {
        this.internal = new TableFieldOptions();
    }
    public TableFieldOptionsBuilder width(Double width) {
    this.internal.width = width;
        return this;
    }
    
    public TableFieldOptionsBuilder minWidth(Double minWidth) {
    this.internal.minWidth = minWidth;
        return this;
    }
    
    public TableFieldOptionsBuilder align(FieldTextAlignment align) {
    this.internal.align = align;
        return this;
    }
    
    public TableFieldOptionsBuilder displayMode(TableCellDisplayMode displayMode) {
    this.internal.displayMode = displayMode;
        return this;
    }
    
    public TableFieldOptionsBuilder cellOptions(TableCellOptions cellOptions) {
    this.internal.cellOptions = cellOptions;
        return this;
    }
    
    public TableFieldOptionsBuilder hidden(Boolean hidden) {
    this.internal.hidden = hidden;
        return this;
    }
    
    public TableFieldOptionsBuilder inspect(Boolean inspect) {
    this.internal.inspect = inspect;
        return this;
    }
    
    public TableFieldOptionsBuilder filterable(Boolean filterable) {
    this.internal.filterable = filterable;
        return this;
    }
    
    public TableFieldOptionsBuilder hideHeader(Boolean hideHeader) {
    this.internal.hideHeader = hideHeader;
        return this;
    }
    public TableFieldOptions build() {
        return this.internal;
    }
}
