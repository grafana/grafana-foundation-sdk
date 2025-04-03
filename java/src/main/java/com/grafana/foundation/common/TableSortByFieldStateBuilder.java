// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class TableSortByFieldStateBuilder implements com.grafana.foundation.cog.Builder<TableSortByFieldState> {
    protected final TableSortByFieldState internal;
    
    public TableSortByFieldStateBuilder() {
        this.internal = new TableSortByFieldState();
    }
    public TableSortByFieldStateBuilder displayName(String displayName) {
        this.internal.displayName = displayName;
        return this;
    }
    
    public TableSortByFieldStateBuilder desc(Boolean desc) {
        this.internal.desc = desc;
        return this;
    }
    public TableSortByFieldState build() {
        return this.internal;
    }
}
