// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;


public class QueryEditorOperatorExpressionBuilder implements com.grafana.foundation.cog.Builder<QueryEditorOperatorExpression> {
    protected final QueryEditorOperatorExpression internal;
    
    public QueryEditorOperatorExpressionBuilder() {
        this.internal = new QueryEditorOperatorExpression();
    }
    public QueryEditorOperatorExpressionBuilder property(com.grafana.foundation.cog.Builder<QueryEditorProperty> property) {
    QueryEditorProperty propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
    
    public QueryEditorOperatorExpressionBuilder operator(com.grafana.foundation.cog.Builder<QueryEditorOperator> operator) {
    QueryEditorOperator operatorResource = operator.build();
        this.internal.operator = operatorResource;
        return this;
    }
    public QueryEditorOperatorExpression build() {
        return this.internal;
    }
}
