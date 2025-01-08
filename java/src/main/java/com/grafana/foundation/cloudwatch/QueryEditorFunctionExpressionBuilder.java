// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import java.util.List;

public class QueryEditorFunctionExpressionBuilder implements com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression> {
    protected final QueryEditorFunctionExpression internal;
    
    public QueryEditorFunctionExpressionBuilder() {
        this.internal = new QueryEditorFunctionExpression();
    this.internal.type = "function";
    }
    public QueryEditorFunctionExpressionBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public QueryEditorFunctionExpressionBuilder parameters(com.grafana.foundation.cog.Builder<List<QueryEditorFunctionParameterExpression>> parameters) {
    this.internal.parameters = parameters.build();
        return this;
    }
    public QueryEditorFunctionExpression build() {
        return this.internal;
    }
}
