// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class BuilderQueryEditorOrderByExpressionArrayBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorOrderByExpressionArray> {
    protected final BuilderQueryEditorOrderByExpressionArray internal;
    
    public BuilderQueryEditorOrderByExpressionArrayBuilder() {
        this.internal = new BuilderQueryEditorOrderByExpressionArray();
    }
    public BuilderQueryEditorOrderByExpressionArrayBuilder expressions(com.grafana.foundation.cog.Builder<List<BuilderQueryEditorOrderByExpression>> expressions) {
        this.internal.expressions = expressions.build();
        return this;
    }
    
    public BuilderQueryEditorOrderByExpressionArrayBuilder type(BuilderQueryEditorExpressionType type) {
        this.internal.type = type;
        return this;
    }
    public BuilderQueryEditorOrderByExpressionArray build() {
        return this.internal;
    }
}
