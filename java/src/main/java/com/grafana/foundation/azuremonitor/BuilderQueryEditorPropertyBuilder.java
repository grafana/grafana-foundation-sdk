// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class BuilderQueryEditorPropertyBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorProperty> {
    protected final BuilderQueryEditorProperty internal;
    
    public BuilderQueryEditorPropertyBuilder() {
        this.internal = new BuilderQueryEditorProperty();
    }
    public BuilderQueryEditorPropertyBuilder type(BuilderQueryEditorPropertyType type) {
        this.internal.type = type;
        return this;
    }
    
    public BuilderQueryEditorPropertyBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    public BuilderQueryEditorProperty build() {
        return this.internal;
    }
}
