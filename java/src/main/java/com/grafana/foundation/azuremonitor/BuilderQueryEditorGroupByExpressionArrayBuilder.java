// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import java.util.LinkedList;

public class BuilderQueryEditorGroupByExpressionArrayBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorGroupByExpressionArray> {
    protected final BuilderQueryEditorGroupByExpressionArray internal;
    
    public BuilderQueryEditorGroupByExpressionArrayBuilder() {
        this.internal = new BuilderQueryEditorGroupByExpressionArray();
    }
    public BuilderQueryEditorGroupByExpressionArrayBuilder expressions(List<com.grafana.foundation.cog.Builder<BuilderQueryEditorGroupByExpression>> expressions) {
        List<BuilderQueryEditorGroupByExpression> expressionsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<BuilderQueryEditorGroupByExpression> r1 : expressions) {
                BuilderQueryEditorGroupByExpression expressionsDepth1 = r1.build();
                expressionsResources.add(expressionsDepth1); 
        }
        this.internal.expressions = expressionsResources;
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
