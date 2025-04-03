// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;


public class SQLExpressionBuilder implements com.grafana.foundation.cog.Builder<SQLExpression> {
    protected final SQLExpression internal;
    
    public SQLExpressionBuilder() {
        this.internal = new SQLExpression();
    }
    public SQLExpressionBuilder select(com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression> select) {
        this.internal.select = select.build();
        return this;
    }
    
    public SQLExpressionBuilder from(QueryEditorPropertyExpressionOrQueryEditorFunctionExpression from) {
        this.internal.from = from;
        return this;
    }
    
    public SQLExpressionBuilder where(com.grafana.foundation.cog.Builder<QueryEditorArrayExpression> where) {
        this.internal.where = where.build();
        return this;
    }
    
    public SQLExpressionBuilder groupBy(com.grafana.foundation.cog.Builder<QueryEditorArrayExpression> groupBy) {
        this.internal.groupBy = groupBy.build();
        return this;
    }
    
    public SQLExpressionBuilder orderBy(com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression> orderBy) {
        this.internal.orderBy = orderBy.build();
        return this;
    }
    
    public SQLExpressionBuilder orderByDirection(String orderByDirection) {
        this.internal.orderByDirection = orderByDirection;
        return this;
    }
    
    public SQLExpressionBuilder limit(Long limit) {
        this.internal.limit = limit;
        return this;
    }
    public SQLExpression build() {
        return this.internal;
    }
}
