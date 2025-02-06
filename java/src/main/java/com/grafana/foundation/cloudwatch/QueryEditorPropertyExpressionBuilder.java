// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;


public class QueryEditorPropertyExpressionBuilder implements com.grafana.foundation.cog.Builder<QueryEditorPropertyExpression> {
    protected final QueryEditorPropertyExpression internal;
    
    public QueryEditorPropertyExpressionBuilder() {
        this.internal = new QueryEditorPropertyExpression();
        this.internal.type = "property";
    }
    public QueryEditorPropertyExpressionBuilder property(com.grafana.foundation.cog.Builder<QueryEditorProperty> property) {
        this.internal.property = property.build();
        return this;
    }
    public QueryEditorPropertyExpression build() {
        return this.internal;
    }
}
