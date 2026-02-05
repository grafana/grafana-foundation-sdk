// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.grafana.foundation.common.DataSourceRef;
import java.util.List;

public class QueryVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
    protected final VariableModel internal;
    
    public QueryVariableBuilder(String name) {
        this.internal = new VariableModel();
        this.internal.name = name;
        this.internal.type = VariableType.QUERY;
    }
    public QueryVariableBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public QueryVariableBuilder label(String label) {
        this.internal.label = label;
        return this;
    }
    
    public QueryVariableBuilder hide(VariableHide hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public QueryVariableBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public QueryVariableBuilder query(StringOrMap query) {
        this.internal.query = query;
        return this;
    }
    
    public QueryVariableBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public QueryVariableBuilder current(VariableOption current) {
        this.internal.current = current;
        return this;
    }
    
    public QueryVariableBuilder multi(Boolean multi) {
        this.internal.multi = multi;
        return this;
    }
    
    public QueryVariableBuilder allowCustomValue(Boolean allowCustomValue) {
        this.internal.allowCustomValue = allowCustomValue;
        return this;
    }
    
    public QueryVariableBuilder options(List<VariableOption> options) {
        this.internal.options = options;
        return this;
    }
    
    public QueryVariableBuilder refresh(VariableRefresh refresh) {
        this.internal.refresh = refresh;
        return this;
    }
    
    public QueryVariableBuilder sort(VariableSort sort) {
        this.internal.sort = sort;
        return this;
    }
    
    public QueryVariableBuilder includeAll(Boolean includeAll) {
        this.internal.includeAll = includeAll;
        return this;
    }
    
    public QueryVariableBuilder allValue(String allValue) {
        this.internal.allValue = allValue;
        return this;
    }
    
    public QueryVariableBuilder regex(String regex) {
        this.internal.regex = regex;
        return this;
    }
    public VariableModel build() {
        return this.internal;
    }
}
