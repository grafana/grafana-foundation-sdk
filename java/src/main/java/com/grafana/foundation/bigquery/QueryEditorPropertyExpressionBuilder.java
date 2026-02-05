// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;


public class QueryEditorPropertyExpressionBuilder implements com.grafana.foundation.cog.Builder<QueryEditorPropertyExpression> {
    protected final QueryEditorPropertyExpression internal;
    
    public QueryEditorPropertyExpressionBuilder() {
        this.internal = new QueryEditorPropertyExpression();
    }
    public QueryEditorPropertyExpressionBuilder property(com.grafana.foundation.cog.Builder<QueryEditorProperty> property) {
    QueryEditorProperty propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
    public QueryEditorPropertyExpression build() {
        return this.internal;
    }
}
