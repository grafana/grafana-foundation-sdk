// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class SQLExpression {
    // SELECT part of the SQL expression
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("select")
    public QueryEditorFunctionExpression select;
    // FROM part of the SQL expression
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("from")
    public QueryEditorPropertyExpressionOrQueryEditorFunctionExpression from;
    // WHERE part of the SQL expression
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("where")
    public QueryEditorArrayExpression where;
    // GROUP BY part of the SQL expression
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("groupBy")
    public QueryEditorArrayExpression groupBy;
    // ORDER BY part of the SQL expression
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("orderBy")
    public QueryEditorFunctionExpression orderBy;
    // The sort order of the SQL expression, `ASC` or `DESC`
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("orderByDirection")
    public String orderByDirection;
    // LIMIT part of the SQL expression
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("limit")
    public Long limit;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<SQLExpression> {
        protected final SQLExpression internal;
        
        public Builder() {
            this.internal = new SQLExpression();
        }
    public Builder select(com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression> select) {
    this.internal.select = select.build();
        return this;
    }
    
    public Builder from(QueryEditorPropertyExpressionOrQueryEditorFunctionExpression from) {
    this.internal.from = from;
        return this;
    }
    
    public Builder where(com.grafana.foundation.cog.Builder<QueryEditorArrayExpression> where) {
    this.internal.where = where.build();
        return this;
    }
    
    public Builder groupBy(com.grafana.foundation.cog.Builder<QueryEditorArrayExpression> groupBy) {
    this.internal.groupBy = groupBy.build();
        return this;
    }
    
    public Builder orderBy(com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression> orderBy) {
    this.internal.orderBy = orderBy.build();
        return this;
    }
    
    public Builder orderByDirection(String orderByDirection) {
    this.internal.orderByDirection = orderByDirection;
        return this;
    }
    
    public Builder limit(Long limit) {
    this.internal.limit = limit;
        return this;
    }
    public SQLExpression build() {
            return this.internal;
        }
    }
}
