// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;


public class SQLExpressionBuilder implements com.grafana.foundation.cog.Builder<SQLExpression> {
    protected final SQLExpression internal;
    
    public SQLExpressionBuilder() {
        this.internal = new SQLExpression();
    }
    public SQLExpressionBuilder select(com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression> select) {
    QueryEditorFunctionExpression selectResource = select.build();
        this.internal.select = selectResource;
        return this;
    }
    
    public SQLExpressionBuilder from(QueryEditorPropertyExpressionOrQueryEditorFunctionExpression from) {
        this.internal.from = from;
        return this;
    }
    
    public SQLExpressionBuilder where(com.grafana.foundation.cog.Builder<QueryEditorArrayExpression> where) {
    QueryEditorArrayExpression whereResource = where.build();
        this.internal.where = whereResource;
        return this;
    }
    
    public SQLExpressionBuilder groupBy(com.grafana.foundation.cog.Builder<QueryEditorArrayExpression> groupBy) {
    QueryEditorArrayExpression groupByResource = groupBy.build();
        this.internal.groupBy = groupByResource;
        return this;
    }
    
    public SQLExpressionBuilder orderBy(com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression> orderBy) {
    QueryEditorFunctionExpression orderByResource = orderBy.build();
        this.internal.orderBy = orderByResource;
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
