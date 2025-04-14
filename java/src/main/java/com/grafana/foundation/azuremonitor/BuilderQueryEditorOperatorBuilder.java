// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class BuilderQueryEditorOperatorBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorOperator> {
    protected final BuilderQueryEditorOperator internal;
    
    public BuilderQueryEditorOperatorBuilder() {
        this.internal = new BuilderQueryEditorOperator();
    }
    public BuilderQueryEditorOperatorBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public BuilderQueryEditorOperatorBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    
    public BuilderQueryEditorOperatorBuilder labelValue(String labelValue) {
        this.internal.labelValue = labelValue;
        return this;
    }
    public BuilderQueryEditorOperator build() {
        return this.internal;
    }
}
