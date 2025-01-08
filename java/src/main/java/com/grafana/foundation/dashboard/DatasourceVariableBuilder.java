// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class DatasourceVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
    protected final VariableModel internal;
    
    public DatasourceVariableBuilder(String name) {
        this.internal = new VariableModel();
    this.internal.name = name;
    this.internal.type = VariableType.DATASOURCE;
    }
    public DatasourceVariableBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public DatasourceVariableBuilder label(String label) {
    this.internal.label = label;
        return this;
    }
    
    public DatasourceVariableBuilder hide(VariableHide hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public DatasourceVariableBuilder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public DatasourceVariableBuilder type(String string) {
		if (this.internal.query == null) {
			this.internal.query = new com.grafana.foundation.dashboard.StringOrMap();
		}
    this.internal.query.string = string;
        return this;
    }
    
    public DatasourceVariableBuilder current(VariableOption current) {
    this.internal.current = current;
        return this;
    }
    
    public DatasourceVariableBuilder multi(Boolean multi) {
    this.internal.multi = multi;
        return this;
    }
    
    public DatasourceVariableBuilder includeAll(Boolean includeAll) {
    this.internal.includeAll = includeAll;
        return this;
    }
    
    public DatasourceVariableBuilder allValue(String allValue) {
    this.internal.allValue = allValue;
        return this;
    }
    
    public DatasourceVariableBuilder regex(String regex) {
    this.internal.regex = regex;
        return this;
    }
    public VariableModel build() {
        return this.internal;
    }
}
