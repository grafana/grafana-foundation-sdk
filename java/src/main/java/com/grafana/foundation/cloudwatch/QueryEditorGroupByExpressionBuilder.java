// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;


public class QueryEditorGroupByExpressionBuilder implements com.grafana.foundation.cog.Builder<QueryEditorGroupByExpression> {
    protected final QueryEditorGroupByExpression internal;
    
    public QueryEditorGroupByExpressionBuilder() {
        this.internal = new QueryEditorGroupByExpression();
        this.internal.type = "groupBy";
    }
    public QueryEditorGroupByExpressionBuilder property(com.grafana.foundation.cog.Builder<QueryEditorProperty> property) {
        this.internal.property = property.build();
        return this;
    }
    public QueryEditorGroupByExpression build() {
        return this.internal;
    }
}
