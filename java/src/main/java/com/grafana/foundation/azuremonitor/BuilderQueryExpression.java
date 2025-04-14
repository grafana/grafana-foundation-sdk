// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BuilderQueryExpression {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("from")
    public BuilderQueryEditorPropertyExpression from;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("columns")
    public BuilderQueryEditorColumnsExpression columns;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("where")
    public BuilderQueryEditorWhereExpressionArray where;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("reduce")
    public BuilderQueryEditorReduceExpressionArray reduce;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("groupBy")
    public BuilderQueryEditorGroupByExpressionArray groupBy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("limit")
    public Long limit;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("orderBy")
    public BuilderQueryEditorOrderByExpressionArray orderBy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fuzzySearch")
    public BuilderQueryEditorWhereExpressionArray fuzzySearch;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeFilter")
    public BuilderQueryEditorWhereExpressionArray timeFilter;
    public BuilderQueryExpression() {
    }
    public BuilderQueryExpression(BuilderQueryEditorPropertyExpression from,BuilderQueryEditorColumnsExpression columns,BuilderQueryEditorWhereExpressionArray where,BuilderQueryEditorReduceExpressionArray reduce,BuilderQueryEditorGroupByExpressionArray groupBy,Long limit,BuilderQueryEditorOrderByExpressionArray orderBy,BuilderQueryEditorWhereExpressionArray fuzzySearch,BuilderQueryEditorWhereExpressionArray timeFilter) {
        this.from = from;
        this.columns = columns;
        this.where = where;
        this.reduce = reduce;
        this.groupBy = groupBy;
        this.limit = limit;
        this.orderBy = orderBy;
        this.fuzzySearch = fuzzySearch;
        this.timeFilter = timeFilter;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
