// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;


public class QueryEditorFunctionParameterExpressionBuilder implements com.grafana.foundation.cog.Builder<QueryEditorFunctionParameterExpression> {
    protected final QueryEditorFunctionParameterExpression internal;
    
    public QueryEditorFunctionParameterExpressionBuilder() {
        this.internal = new QueryEditorFunctionParameterExpression();
        this.internal.type = "functionParameter";
    }
    public QueryEditorFunctionParameterExpressionBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    public QueryEditorFunctionParameterExpression build() {
        return this.internal;
    }
}
