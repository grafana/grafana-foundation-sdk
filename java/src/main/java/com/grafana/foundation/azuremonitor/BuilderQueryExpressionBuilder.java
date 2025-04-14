// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class BuilderQueryExpressionBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryExpression> {
    protected final BuilderQueryExpression internal;
    
    public BuilderQueryExpressionBuilder() {
        this.internal = new BuilderQueryExpression();
    }
    public BuilderQueryExpressionBuilder from(com.grafana.foundation.cog.Builder<BuilderQueryEditorPropertyExpression> from) {
        this.internal.from = from.build();
        return this;
    }
    
    public BuilderQueryExpressionBuilder columns(com.grafana.foundation.cog.Builder<BuilderQueryEditorColumnsExpression> columns) {
        this.internal.columns = columns.build();
        return this;
    }
    
    public BuilderQueryExpressionBuilder where(com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpressionArray> where) {
        this.internal.where = where.build();
        return this;
    }
    
    public BuilderQueryExpressionBuilder reduce(com.grafana.foundation.cog.Builder<BuilderQueryEditorReduceExpressionArray> reduce) {
        this.internal.reduce = reduce.build();
        return this;
    }
    
    public BuilderQueryExpressionBuilder groupBy(com.grafana.foundation.cog.Builder<BuilderQueryEditorGroupByExpressionArray> groupBy) {
        this.internal.groupBy = groupBy.build();
        return this;
    }
    
    public BuilderQueryExpressionBuilder limit(Long limit) {
        this.internal.limit = limit;
        return this;
    }
    
    public BuilderQueryExpressionBuilder orderBy(com.grafana.foundation.cog.Builder<BuilderQueryEditorOrderByExpressionArray> orderBy) {
        this.internal.orderBy = orderBy.build();
        return this;
    }
    
    public BuilderQueryExpressionBuilder fuzzySearch(com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpressionArray> fuzzySearch) {
        this.internal.fuzzySearch = fuzzySearch.build();
        return this;
    }
    
    public BuilderQueryExpressionBuilder timeFilter(com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpressionArray> timeFilter) {
        this.internal.timeFilter = timeFilter.build();
        return this;
    }
    public BuilderQueryExpression build() {
        return this.internal;
    }
}
