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
    public SQLExpression() {
    }
    
    public SQLExpression(List<QueryEditorFunctionExpression> columns,String from,String whereString,List<QueryEditorGroupByExpression> groupBy,QueryEditorPropertyExpression orderBy,OrderByDirection orderByDirection,Long limit,Long offset) {
        this.columns = columns;
        this.from = from;
        this.whereString = whereString;
        this.groupBy = groupBy;
        this.orderBy = orderBy;
        this.orderByDirection = orderByDirection;
        this.limit = limit;
        this.offset = offset;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
