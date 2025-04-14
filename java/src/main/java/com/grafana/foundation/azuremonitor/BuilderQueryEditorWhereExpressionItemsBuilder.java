// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class BuilderQueryEditorWhereExpressionItemsBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpressionItems> {
    protected final BuilderQueryEditorWhereExpressionItems internal;
    
    public BuilderQueryEditorWhereExpressionItemsBuilder() {
        this.internal = new BuilderQueryEditorWhereExpressionItems();
    }
    public BuilderQueryEditorWhereExpressionItemsBuilder property(com.grafana.foundation.cog.Builder<BuilderQueryEditorProperty> property) {
        this.internal.property = property.build();
        return this;
    }
    
    public BuilderQueryEditorWhereExpressionItemsBuilder operator(com.grafana.foundation.cog.Builder<BuilderQueryEditorOperator> operator) {
        this.internal.operator = operator.build();
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
