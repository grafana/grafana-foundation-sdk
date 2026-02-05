// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import java.util.List;
import java.util.LinkedList;

public class SQLExpressionBuilder implements com.grafana.foundation.cog.Builder<SQLExpression> {
    protected final SQLExpression internal;
    
    public SQLExpressionBuilder() {
        this.internal = new SQLExpression();
    }
    public SQLExpressionBuilder columns(List<com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression>> columns) {
        List<QueryEditorFunctionExpression> columnsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression> r1 : columns) {
                QueryEditorFunctionExpression columnsDepth1 = r1.build();
                columnsResources.add(columnsDepth1); 
        }
        this.internal.columns = columnsResources;
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
    
    public SQLExpressionBuilder groupBy(List<com.grafana.foundation.cog.Builder<QueryEditorGroupByExpression>> groupBy) {
        List<QueryEditorGroupByExpression> groupByResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<QueryEditorGroupByExpression> r1 : groupBy) {
                QueryEditorGroupByExpression groupByDepth1 = r1.build();
                groupByResources.add(groupByDepth1); 
        }
        this.internal.groupBy = groupByResources;
        return this;
    }
    
    public SQLExpressionBuilder orderBy(com.grafana.foundation.cog.Builder<QueryEditorPropertyExpression> orderBy) {
    QueryEditorPropertyExpression orderByResource = orderBy.build();
        this.internal.orderBy = orderByResource;
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
