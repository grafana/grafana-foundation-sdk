// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class BuilderQueryEditorReduceExpressionArrayBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorReduceExpressionArray> {
    protected final BuilderQueryEditorReduceExpressionArray internal;
    
    public BuilderQueryEditorReduceExpressionArrayBuilder() {
        this.internal = new BuilderQueryEditorReduceExpressionArray();
    }
    public BuilderQueryEditorReduceExpressionArrayBuilder expressions(com.grafana.foundation.cog.Builder<List<BuilderQueryEditorReduceExpression>> expressions) {
        this.internal.expressions = expressions.build();
        return this;
    }
    
    public BuilderQueryEditorReduceExpressionArrayBuilder type(BuilderQueryEditorExpressionType type) {
        this.internal.type = type;
        return this;
    }
    public BuilderQueryEditorReduceExpressionArray build() {
        return this.internal;
    }
}
