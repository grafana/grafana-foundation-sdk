// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;

public class TextBoxVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
    protected final VariableModel internal;
    
    public TextBoxVariableBuilder(String name) {
        this.internal = new VariableModel();
        this.internal.name = name;
        this.internal.type = VariableType.TEXTBOX;
    }
    public TextBoxVariableBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public TextBoxVariableBuilder label(String label) {
        this.internal.label = label;
        return this;
    }
    
    public TextBoxVariableBuilder hide(VariableHide hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public TextBoxVariableBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public TextBoxVariableBuilder defaultValue(StringOrMap query) {
        this.internal.query = query;
        return this;
    }
    
    public TextBoxVariableBuilder current(VariableOption current) {
        this.internal.current = current;
        return this;
    }
    
    public TextBoxVariableBuilder allowCustomValue(Boolean allowCustomValue) {
        this.internal.allowCustomValue = allowCustomValue;
        return this;
    }
    
    public TextBoxVariableBuilder options(List<VariableOption> options) {
        this.internal.options = options;
        return this;
    }
    
    public TextBoxVariableBuilder definition(String definition) {
        this.internal.definition = definition;
        return this;
    }
    public VariableModel build() {
        return this.internal;
    }
}
