// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import java.util.List;

public class SQLExpressionBuilder implements com.grafana.foundation.cog.Builder<SQLExpression> {
    protected final SQLExpression internal;
    
    public SQLExpressionBuilder() {
        this.internal = new SQLExpression();
    }
    public SQLExpressionBuilder columns(com.grafana.foundation.cog.Builder<List<QueryEditorFunctionExpression>> columns) {
    this.internal.columns = columns.build();
        return this;
    }
    
    public SQLExpressionBuilder from(String from) {
    this.internal.from = from;
        return this;
    }
    
    public SQLExpressionBuilder whereString(String whereString) {
    this.internal.whereString = whereString;
        return this;
    }
    
    public SQLExpressionBuilder groupBy(com.grafana.foundation.cog.Builder<List<QueryEditorGroupByExpression>> groupBy) {
    this.internal.groupBy = groupBy.build();
        return this;
    }
    
    public SQLExpressionBuilder orderBy(com.grafana.foundation.cog.Builder<QueryEditorPropertyExpression> orderBy) {
    this.internal.orderBy = orderBy.build();
        return this;
    }
    
    public SQLExpressionBuilder orderByDirection(OrderByDirection orderByDirection) {
    this.internal.orderByDirection = orderByDirection;
        return this;
    }
    
    public SQLExpressionBuilder limit(Long limit) {
    this.internal.limit = limit;
        return this;
    }
    
    public SQLExpressionBuilder offset(Long offset) {
    this.internal.offset = offset;
        return this;
    }
    public SQLExpression build() {
        return this.internal;
    }
}
