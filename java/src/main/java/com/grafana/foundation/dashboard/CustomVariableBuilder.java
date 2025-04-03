// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;

public class CustomVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
    protected final VariableModel internal;
    
    public CustomVariableBuilder(String name) {
        this.internal = new VariableModel();
        this.internal.name = name;
        this.internal.type = VariableType.CUSTOM;
    }
    public CustomVariableBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public CustomVariableBuilder label(String label) {
        this.internal.label = label;
        return this;
    }
    
    public CustomVariableBuilder hide(VariableHide hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public CustomVariableBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public CustomVariableBuilder values(StringOrMap query) {
        this.internal.query = query;
        return this;
    }
    
    public CustomVariableBuilder current(VariableOption current) {
        this.internal.current = current;
        return this;
    }
    
    public CustomVariableBuilder multi(Boolean multi) {
        this.internal.multi = multi;
        return this;
    }
    
    public CustomVariableBuilder allowCustomValue(Boolean allowCustomValue) {
        this.internal.allowCustomValue = allowCustomValue;
        return this;
    }
    
    public CustomVariableBuilder options(List<VariableOption> options) {
        this.internal.options = options;
        return this;
    }
    
    public CustomVariableBuilder includeAll(Boolean includeAll) {
        this.internal.includeAll = includeAll;
        return this;
    }
    
    public CustomVariableBuilder allValue(String allValue) {
        this.internal.allValue = allValue;
        return this;
    }
    public VariableModel build() {
        return this.internal;
    }
}
