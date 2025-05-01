// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class BuilderQueryEditorWhereExpressionItemsBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpressionItems> {
    protected final BuilderQueryEditorWhereExpressionItems internal;
    
    public BuilderQueryEditorWhereExpressionItemsBuilder() {
        this.internal = new BuilderQueryEditorWhereExpressionItems();
    }
    public BuilderQueryEditorWhereExpressionItemsBuilder property(com.grafana.foundation.cog.Builder<BuilderQueryEditorProperty> property) {
    BuilderQueryEditorProperty propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
    
    public BuilderQueryEditorWhereExpressionItemsBuilder operator(com.grafana.foundation.cog.Builder<BuilderQueryEditorOperator> operator) {
    BuilderQueryEditorOperator operatorResource = operator.build();
        this.internal.operator = operatorResource;
        return this;
    }
    
    public BuilderQueryEditorWhereExpressionItemsBuilder type(BuilderQueryEditorExpressionType type) {
        this.internal.type = type;
        return this;
    }
    public BuilderQueryEditorWhereExpressionItems build() {
        return this.internal;
    }
}
