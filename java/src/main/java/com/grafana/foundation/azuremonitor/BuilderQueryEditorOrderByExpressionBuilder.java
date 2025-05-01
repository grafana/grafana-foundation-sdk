// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class BuilderQueryEditorOrderByExpressionBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorOrderByExpression> {
    protected final BuilderQueryEditorOrderByExpression internal;
    
    public BuilderQueryEditorOrderByExpressionBuilder() {
        this.internal = new BuilderQueryEditorOrderByExpression();
    }
    public BuilderQueryEditorOrderByExpressionBuilder property(com.grafana.foundation.cog.Builder<BuilderQueryEditorProperty> property) {
    BuilderQueryEditorProperty propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
    
    public BuilderQueryEditorOrderByExpressionBuilder order(BuilderQueryEditorOrderByOptions order) {
        this.internal.order = order;
        return this;
    }
    
    public BuilderQueryEditorOrderByExpressionBuilder type(BuilderQueryEditorExpressionType type) {
        this.internal.type = type;
        return this;
    }
    public BuilderQueryEditorOrderByExpression build() {
        return this.internal;
    }
}
