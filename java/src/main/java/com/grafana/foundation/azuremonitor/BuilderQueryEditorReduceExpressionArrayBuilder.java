// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import java.util.LinkedList;

public class BuilderQueryEditorReduceExpressionArrayBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorReduceExpressionArray> {
    protected final BuilderQueryEditorReduceExpressionArray internal;
    
    public BuilderQueryEditorReduceExpressionArrayBuilder() {
        this.internal = new BuilderQueryEditorReduceExpressionArray();
    }
    public BuilderQueryEditorReduceExpressionArrayBuilder expressions(List<com.grafana.foundation.cog.Builder<BuilderQueryEditorReduceExpression>> expressions) {
        List<BuilderQueryEditorReduceExpression> expressionsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<BuilderQueryEditorReduceExpression> r1 : expressions) {
                BuilderQueryEditorReduceExpression expressionsDepth1 = r1.build();
                expressionsResources.add(expressionsDepth1); 
        }
        this.internal.expressions = expressionsResources;
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
