// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class BuilderQueryExpressionBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryExpression> {
    protected final BuilderQueryExpression internal;
    
    public BuilderQueryExpressionBuilder() {
        this.internal = new BuilderQueryExpression();
    }
    public BuilderQueryExpressionBuilder from(com.grafana.foundation.cog.Builder<BuilderQueryEditorPropertyExpression> from) {
    BuilderQueryEditorPropertyExpression fromResource = from.build();
        this.internal.from = fromResource;
        return this;
    }
    
    public BuilderQueryExpressionBuilder columns(com.grafana.foundation.cog.Builder<BuilderQueryEditorColumnsExpression> columns) {
    BuilderQueryEditorColumnsExpression columnsResource = columns.build();
        this.internal.columns = columnsResource;
        return this;
    }
    
    public BuilderQueryExpressionBuilder where(com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpressionArray> where) {
    BuilderQueryEditorWhereExpressionArray whereResource = where.build();
        this.internal.where = whereResource;
        return this;
    }
    
    public BuilderQueryExpressionBuilder reduce(com.grafana.foundation.cog.Builder<BuilderQueryEditorReduceExpressionArray> reduce) {
    BuilderQueryEditorReduceExpressionArray reduceResource = reduce.build();
        this.internal.reduce = reduceResource;
        return this;
    }
    
    public BuilderQueryExpressionBuilder groupBy(com.grafana.foundation.cog.Builder<BuilderQueryEditorGroupByExpressionArray> groupBy) {
    BuilderQueryEditorGroupByExpressionArray groupByResource = groupBy.build();
        this.internal.groupBy = groupByResource;
        return this;
    }
    
    public BuilderQueryExpressionBuilder limit(Long limit) {
        this.internal.limit = limit;
        return this;
    }
    
    public BuilderQueryExpressionBuilder orderBy(com.grafana.foundation.cog.Builder<BuilderQueryEditorOrderByExpressionArray> orderBy) {
    BuilderQueryEditorOrderByExpressionArray orderByResource = orderBy.build();
        this.internal.orderBy = orderByResource;
        return this;
    }
    
    public BuilderQueryExpressionBuilder fuzzySearch(com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpressionArray> fuzzySearch) {
    BuilderQueryEditorWhereExpressionArray fuzzySearchResource = fuzzySearch.build();
        this.internal.fuzzySearch = fuzzySearchResource;
        return this;
    }
    
    public BuilderQueryExpressionBuilder timeFilter(com.grafana.foundation.cog.Builder<BuilderQueryEditorWhereExpressionArray> timeFilter) {
    BuilderQueryEditorWhereExpressionArray timeFilterResource = timeFilter.build();
        this.internal.timeFilter = timeFilterResource;
        return this;
    }
    public BuilderQueryExpression build() {
        return this.internal;
    }
}
