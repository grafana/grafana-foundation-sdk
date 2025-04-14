// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class BuilderQueryEditorReduceExpressionBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorReduceExpression> {
    protected final BuilderQueryEditorReduceExpression internal;
    
    public BuilderQueryEditorReduceExpressionBuilder() {
        this.internal = new BuilderQueryEditorReduceExpression();
    }
    public BuilderQueryEditorReduceExpressionBuilder property(com.grafana.foundation.cog.Builder<BuilderQueryEditorProperty> property) {
        this.internal.property = property.build();
        return this;
    }
    
    public BuilderQueryEditorReduceExpressionBuilder reduce(com.grafana.foundation.cog.Builder<BuilderQueryEditorProperty> reduce) {
        this.internal.reduce = reduce.build();
        return this;
    }
    
    public BuilderQueryEditorReduceExpressionBuilder parameters(com.grafana.foundation.cog.Builder<List<BuilderQueryEditorFunctionParameterExpression>> parameters) {
        this.internal.parameters = parameters.build();
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
