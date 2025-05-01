// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class BuilderQueryEditorGroupByExpressionBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorGroupByExpression> {
    protected final BuilderQueryEditorGroupByExpression internal;
    
    public BuilderQueryEditorGroupByExpressionBuilder() {
        this.internal = new BuilderQueryEditorGroupByExpression();
    }
    public BuilderQueryEditorGroupByExpressionBuilder property(com.grafana.foundation.cog.Builder<BuilderQueryEditorProperty> property) {
    BuilderQueryEditorProperty propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
    
    public BuilderQueryEditorGroupByExpressionBuilder interval(com.grafana.foundation.cog.Builder<BuilderQueryEditorProperty> interval) {
    BuilderQueryEditorProperty intervalResource = interval.build();
        this.internal.interval = intervalResource;
        return this;
    }
    
    public BuilderQueryEditorGroupByExpressionBuilder focus(Boolean focus) {
        this.internal.focus = focus;
        return this;
    }
    
    public BuilderQueryEditorGroupByExpressionBuilder type(BuilderQueryEditorExpressionType type) {
        this.internal.type = type;
        return this;
    }
    public BuilderQueryEditorGroupByExpression build() {
        return this.internal;
    }
}
