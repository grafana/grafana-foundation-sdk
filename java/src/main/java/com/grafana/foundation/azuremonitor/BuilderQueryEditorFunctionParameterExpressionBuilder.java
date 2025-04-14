// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class BuilderQueryEditorFunctionParameterExpressionBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorFunctionParameterExpression> {
    protected final BuilderQueryEditorFunctionParameterExpression internal;
    
    public BuilderQueryEditorFunctionParameterExpressionBuilder() {
        this.internal = new BuilderQueryEditorFunctionParameterExpression();
    }
    public BuilderQueryEditorFunctionParameterExpressionBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    
    public BuilderQueryEditorFunctionParameterExpressionBuilder fieldType(BuilderQueryEditorPropertyType fieldType) {
        this.internal.fieldType = fieldType;
        return this;
    }
    
    public BuilderQueryEditorFunctionParameterExpressionBuilder type(BuilderQueryEditorExpressionType type) {
        this.internal.type = type;
        return this;
    }
    public BuilderQueryEditorFunctionParameterExpression build() {
        return this.internal;
    }
}
