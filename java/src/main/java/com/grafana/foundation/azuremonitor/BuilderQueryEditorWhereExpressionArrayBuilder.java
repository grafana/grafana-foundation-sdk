// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class BuilderQueryEditorWhereExpressionArrayBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpressionArray> {
    protected final BuilderQueryEditorWhereExpressionArray internal;
    
    public BuilderQueryEditorWhereExpressionArrayBuilder() {
        this.internal = new BuilderQueryEditorWhereExpressionArray();
    }
    public BuilderQueryEditorWhereExpressionArrayBuilder expressions(com.grafana.foundation.cog.Builder<List<BuilderQueryEditorWhereExpression>> expressions) {
        this.internal.expressions = expressions.build();
        return this;
    }
    
    public BuilderQueryEditorWhereExpressionArrayBuilder type(BuilderQueryEditorExpressionType type) {
        this.internal.type = type;
        return this;
    }
    public BuilderQueryEditorWhereExpressionArray build() {
        return this.internal;
    }
}
