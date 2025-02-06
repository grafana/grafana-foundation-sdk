// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class ConstantVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
    protected final VariableModel internal;
    
    public ConstantVariableBuilder(String name) {
        this.internal = new VariableModel();
        this.internal.name = name;
        this.internal.type = VariableType.CONSTANT;
        this.internal.hide = VariableHide.HIDE_VARIABLE;
    }
    public ConstantVariableBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public ConstantVariableBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public ConstantVariableBuilder label(String label) {
        this.internal.label = label;
        return this;
    }
    
    public ConstantVariableBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public ConstantVariableBuilder value(StringOrMap query) {
        this.internal.query = query;
        return this;
    }
    
    public ConstantVariableBuilder allFormat(String allFormat) {
        this.internal.allFormat = allFormat;
        return this;
    }
    public VariableModel build() {
        return this.internal;
    }
}
