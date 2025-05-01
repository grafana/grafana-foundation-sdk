// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class BuilderQueryEditorOperatorTypeBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorOperatorType> {
    protected final BuilderQueryEditorOperatorType internal;
    
    public BuilderQueryEditorOperatorTypeBuilder() {
        this.internal = new BuilderQueryEditorOperatorType();
    }
    public BuilderQueryEditorOperatorTypeBuilder string(String string) {
        this.internal.string = string;
        return this;
    }
    
    public BuilderQueryEditorOperatorTypeBuilder bool(Boolean bool) {
        this.internal.bool = bool;
        return this;
    }
    
    public BuilderQueryEditorOperatorTypeBuilder float64(Double float64) {
        this.internal.float64 = float64;
        return this;
    }
    
    public BuilderQueryEditorOperatorTypeBuilder selectableValue(com.grafana.foundation.cog.Builder<SelectableValue> selectableValue) {
    SelectableValue selectableValueResource = selectableValue.build();
        this.internal.selectableValue = selectableValueResource;
        return this;
    }
    public BuilderQueryEditorOperatorType build() {
        return this.internal;
    }
}
