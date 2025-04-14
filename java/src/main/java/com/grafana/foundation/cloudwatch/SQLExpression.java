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
    public SQLExpression() {
    }
    public SQLExpression(QueryEditorFunctionExpression select,QueryEditorPropertyExpressionOrQueryEditorFunctionExpression from,QueryEditorArrayExpression where,QueryEditorArrayExpression groupBy,QueryEditorFunctionExpression orderBy,String orderByDirection,Long limit) {
        this.select = select;
        this.from = from;
        this.where = where;
        this.groupBy = groupBy;
        this.orderBy = orderBy;
        this.orderByDirection = orderByDirection;
        this.limit = limit;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
