// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class SQLExpression {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("columns")
    public List<QueryEditorFunctionExpression> columns;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("from")
    public String from;
    // whereJsonTree?: _
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("whereString")
    public String whereString;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("groupBy")
    public List<QueryEditorGroupByExpression> groupBy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("orderBy")
    public QueryEditorPropertyExpression orderBy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("orderByDirection")
    public OrderByDirection orderByDirection;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("limit")
    public Long limit;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("offset")
    public Long offset;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<SQLExpression> {
        protected final SQLExpression internal;
        
        public Builder() {
            this.internal = new SQLExpression();
        }
    public Builder columns(com.grafana.foundation.cog.Builder<List<QueryEditorFunctionExpression>> columns) {
    this.internal.columns = columns.build();
        return this;
    }
    
    public Builder from(String from) {
    this.internal.from = from;
        return this;
    }
    
    public Builder whereString(String whereString) {
    this.internal.whereString = whereString;
        return this;
    }
    
    public Builder groupBy(com.grafana.foundation.cog.Builder<List<QueryEditorGroupByExpression>> groupBy) {
    this.internal.groupBy = groupBy.build();
        return this;
    }
    
    public Builder orderBy(com.grafana.foundation.cog.Builder<QueryEditorPropertyExpression> orderBy) {
    this.internal.orderBy = orderBy.build();
        return this;
    }
    
    public Builder orderByDirection(OrderByDirection orderByDirection) {
    this.internal.orderByDirection = orderByDirection;
        return this;
    }
    
    public Builder limit(Long limit) {
    this.internal.limit = limit;
        return this;
    }
    
    public Builder offset(Long offset) {
    this.internal.offset = offset;
        return this;
    }
    public SQLExpression build() {
            return this.internal;
        }
    }
}
