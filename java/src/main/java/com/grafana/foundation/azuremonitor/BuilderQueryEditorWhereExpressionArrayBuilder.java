// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import java.util.LinkedList;

public class BuilderQueryEditorWhereExpressionArrayBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpressionArray> {
    protected final BuilderQueryEditorWhereExpressionArray internal;
    
    public BuilderQueryEditorWhereExpressionArrayBuilder() {
        this.internal = new BuilderQueryEditorWhereExpressionArray();
    }
    public BuilderQueryEditorWhereExpressionArrayBuilder expressions(List<com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpression>> expressions) {
        List<BuilderQueryEditorWhereExpression> expressionsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpression> r1 : expressions) {
                BuilderQueryEditorWhereExpression expressionsDepth1 = r1.build();
                expressionsResources.add(expressionsDepth1); 
        }
        this.internal.expressions = expressionsResources;
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
