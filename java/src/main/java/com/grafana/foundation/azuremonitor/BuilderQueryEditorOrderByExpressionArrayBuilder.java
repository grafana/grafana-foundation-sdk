// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import java.util.LinkedList;

public class BuilderQueryEditorOrderByExpressionArrayBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorOrderByExpressionArray> {
    protected final BuilderQueryEditorOrderByExpressionArray internal;
    
    public BuilderQueryEditorOrderByExpressionArrayBuilder() {
        this.internal = new BuilderQueryEditorOrderByExpressionArray();
    }
    public BuilderQueryEditorOrderByExpressionArrayBuilder expressions(List<com.grafana.foundation.cog.Builder<BuilderQueryEditorOrderByExpression>> expressions) {
        List<BuilderQueryEditorOrderByExpression> expressionsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<BuilderQueryEditorOrderByExpression> r1 : expressions) {
                BuilderQueryEditorOrderByExpression expressionsDepth1 = r1.build();
                expressionsResources.add(expressionsDepth1); 
        }
        this.internal.expressions = expressionsResources;
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
