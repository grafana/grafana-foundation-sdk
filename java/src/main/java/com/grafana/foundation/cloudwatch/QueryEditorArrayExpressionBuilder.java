// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import java.util.List;

public class QueryEditorArrayExpressionBuilder implements com.grafana.foundation.cog.Builder<QueryEditorArrayExpression> {
    protected final QueryEditorArrayExpression internal;
    
    public QueryEditorArrayExpressionBuilder() {
        this.internal = new QueryEditorArrayExpression();
    }
    public QueryEditorArrayExpressionBuilder type(QueryEditorArrayExpressionType type) {
        this.internal.type = type;
        return this;
    }
    
    public QueryEditorArrayExpressionBuilder expressions(List<QueryEditorExpression> expressions) {
        this.internal.expressions = expressions;
        return this;
    }
    public QueryEditorArrayExpression build() {
        return this.internal;
    }
}
