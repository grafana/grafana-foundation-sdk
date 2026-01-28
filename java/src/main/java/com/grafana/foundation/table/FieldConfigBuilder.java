// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.table;

import com.grafana.foundation.common.FieldTextAlignment;
import com.grafana.foundation.common.TableCellDisplayMode;
import com.grafana.foundation.common.TableCellOptions;

public class FieldConfigBuilder implements com.grafana.foundation.cog.Builder<FieldConfig> {
    protected final FieldConfig internal;
    
    public FieldConfigBuilder() {
        this.internal = new FieldConfig();
    }
    public FieldConfigBuilder width(Double width) {
        this.internal.width = width;
        return this;
    }
    
    public FieldConfigBuilder minWidth(Double minWidth) {
        this.internal.minWidth = minWidth;
        return this;
    }
    
    public FieldConfigBuilder align(FieldTextAlignment align) {
        this.internal.align = align;
        return this;
    }
    
    public FieldConfigBuilder displayMode(TableCellDisplayMode displayMode) {
        this.internal.displayMode = displayMode;
        return this;
    }
    
    public FieldConfigBuilder cellOptions(TableCellOptions cellOptions) {
        this.internal.cellOptions = cellOptions;
        return this;
    }
    
    public FieldConfigBuilder hidden(Boolean hidden) {
        this.internal.hidden = hidden;
        return this;
    }
    
    public FieldConfigBuilder inspect(Boolean inspect) {
        this.internal.inspect = inspect;
        return this;
    }
    
    public FieldConfigBuilder filterable(Boolean filterable) {
        this.internal.filterable = filterable;
        return this;
    }
    
    public FieldConfigBuilder hideHeader(Boolean hideHeader) {
        this.internal.hideHeader = hideHeader;
        return this;
    }
    public FieldConfig build() {
        return this.internal;
    }
}
