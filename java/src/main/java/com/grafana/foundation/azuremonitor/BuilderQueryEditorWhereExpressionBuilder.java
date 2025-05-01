// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import java.util.LinkedList;

public class BuilderQueryEditorWhereExpressionBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpression> {
    protected final BuilderQueryEditorWhereExpression internal;
    
    public BuilderQueryEditorWhereExpressionBuilder() {
        this.internal = new BuilderQueryEditorWhereExpression();
    }
    public BuilderQueryEditorWhereExpressionBuilder type(BuilderQueryEditorExpressionType type) {
        this.internal.type = type;
        return this;
    }
    
    public BuilderQueryEditorWhereExpressionBuilder expressions(List<com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpressionItems>> expressions) {
        List<BuilderQueryEditorWhereExpressionItems> expressionsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpressionItems> r1 : expressions) {
                BuilderQueryEditorWhereExpressionItems expressionsDepth1 = r1.build();
                expressionsResources.add(expressionsDepth1); 
        }
        this.internal.expressions = expressionsResources;
        return this;
    }
    public BuilderQueryEditorWhereExpression build() {
        return this.internal;
    }
}
