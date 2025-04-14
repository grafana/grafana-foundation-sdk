// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class BuilderQueryEditorGroupByExpressionArrayBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorGroupByExpressionArray> {
    protected final BuilderQueryEditorGroupByExpressionArray internal;
    
    public BuilderQueryEditorGroupByExpressionArrayBuilder() {
        this.internal = new BuilderQueryEditorGroupByExpressionArray();
    }
    public BuilderQueryEditorGroupByExpressionArrayBuilder expressions(com.grafana.foundation.cog.Builder<List<BuilderQueryEditorGroupByExpression>> expressions) {
        this.internal.expressions = expressions.build();
        return this;
    }
    
    public BuilderQueryEditorGroupByExpressionArrayBuilder type(BuilderQueryEditorExpressionType type) {
        this.internal.type = type;
        return this;
    }
    public BuilderQueryEditorGroupByExpressionArray build() {
        return this.internal;
    }
}
