// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import java.util.List;

public class TableFooterOptionsBuilder implements com.grafana.foundation.cog.Builder<TableFooterOptions> {
    protected final TableFooterOptions internal;
    
    public TableFooterOptionsBuilder() {
        this.internal = new TableFooterOptions();
    }
    public TableFooterOptionsBuilder show(Boolean show) {
    this.internal.show = show;
        return this;
    }
    
    public TableFooterOptionsBuilder reducer(List<String> reducer) {
    this.internal.reducer = reducer;
        return this;
    }
    
    public TableFooterOptionsBuilder fields(List<String> fields) {
    this.internal.fields = fields;
        return this;
    }
    
    public TableFooterOptionsBuilder enablePagination(Boolean enablePagination) {
    this.internal.enablePagination = enablePagination;
        return this;
    }
    
    public TableFooterOptionsBuilder countRows(Boolean countRows) {
    this.internal.countRows = countRows;
        return this;
    }
    public TableFooterOptions build() {
        return this.internal;
    }
}
