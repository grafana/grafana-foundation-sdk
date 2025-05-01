// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import java.util.LinkedList;

public class BuilderQueryEditorReduceExpressionBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorReduceExpression> {
    protected final BuilderQueryEditorReduceExpression internal;
    
    public BuilderQueryEditorReduceExpressionBuilder() {
        this.internal = new BuilderQueryEditorReduceExpression();
    }
    public BuilderQueryEditorReduceExpressionBuilder property(com.grafana.foundation.cog.Builder<BuilderQueryEditorProperty> property) {
    BuilderQueryEditorProperty propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
    
    public BuilderQueryEditorReduceExpressionBuilder reduce(com.grafana.foundation.cog.Builder<BuilderQueryEditorProperty> reduce) {
    BuilderQueryEditorProperty reduceResource = reduce.build();
        this.internal.reduce = reduceResource;
        return this;
    }
    
    public BuilderQueryEditorReduceExpressionBuilder parameters(List<com.grafana.foundation.cog.Builder<BuilderQueryEditorFunctionParameterExpression>> parameters) {
        List<BuilderQueryEditorFunctionParameterExpression> parametersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<BuilderQueryEditorFunctionParameterExpression> r1 : parameters) {
                BuilderQueryEditorFunctionParameterExpression parametersDepth1 = r1.build();
                parametersResources.add(parametersDepth1); 
        }
        this.internal.parameters = parametersResources;
        return this;
    }
    
    public BuilderQueryEditorReduceExpressionBuilder focus(Boolean focus) {
        this.internal.focus = focus;
        return this;
    }
    public BuilderQueryEditorReduceExpression build() {
        return this.internal;
    }
}
