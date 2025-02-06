// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class AdHocVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
    protected final VariableModel internal;
    
    public AdHocVariableBuilder(String name) {
        this.internal = new VariableModel();
        this.internal.name = name;
        this.internal.type = VariableType.ADHOC;
    }
    public AdHocVariableBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public AdHocVariableBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public AdHocVariableBuilder label(String label) {
        this.internal.label = label;
        return this;
    }
    
    public AdHocVariableBuilder hide(VariableHide hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public AdHocVariableBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public AdHocVariableBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public AdHocVariableBuilder allFormat(String allFormat) {
        this.internal.allFormat = allFormat;
        return this;
    }
    public VariableModel build() {
        return this.internal;
    }
}
