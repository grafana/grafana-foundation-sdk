// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;


public class QueryEditorOperatorExpressionBuilder implements com.grafana.foundation.cog.Builder<QueryEditorOperatorExpression> {
    protected final QueryEditorOperatorExpression internal;
    
    public QueryEditorOperatorExpressionBuilder() {
        this.internal = new QueryEditorOperatorExpression();
    }
    public QueryEditorOperatorExpressionBuilder property(com.grafana.foundation.cog.Builder<QueryEditorProperty> property) {
        this.internal.property = property.build();
        return this;
    }
    
    public QueryEditorOperatorExpressionBuilder operator(com.grafana.foundation.cog.Builder<QueryEditorOperator> operator) {
        this.internal.operator = operator.build();
        return this;
    }
    public QueryEditorOperatorExpression build() {
        return this.internal;
    }
}
