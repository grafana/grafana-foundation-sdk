// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class BuilderQueryEditorWhereExpressionBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpression> {
    protected final BuilderQueryEditorWhereExpression internal;
    
    public BuilderQueryEditorWhereExpressionBuilder() {
        this.internal = new BuilderQueryEditorWhereExpression();
    }
    public BuilderQueryEditorWhereExpressionBuilder type(BuilderQueryEditorExpressionType type) {
        this.internal.type = type;
        return this;
    }
    
    public BuilderQueryEditorWhereExpressionBuilder expressions(com.grafana.foundation.cog.Builder<List<BuilderQueryEditorWhereExpressionItems>> expressions) {
        this.internal.expressions = expressions.build();
        return this;
    }
    public BuilderQueryEditorWhereExpression build() {
        return this.internal;
    }
}
