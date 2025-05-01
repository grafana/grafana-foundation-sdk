// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class BuilderQueryEditorPropertyExpressionBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorPropertyExpression> {
    protected final BuilderQueryEditorPropertyExpression internal;
    
    public BuilderQueryEditorPropertyExpressionBuilder() {
        this.internal = new BuilderQueryEditorPropertyExpression();
    }
    public BuilderQueryEditorPropertyExpressionBuilder property(com.grafana.foundation.cog.Builder<BuilderQueryEditorProperty> property) {
    BuilderQueryEditorProperty propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
    
    public BuilderQueryEditorPropertyExpressionBuilder type(BuilderQueryEditorExpressionType type) {
        this.internal.type = type;
        return this;
    }
    public BuilderQueryEditorPropertyExpression build() {
        return this.internal;
    }
}
