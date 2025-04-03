// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class QueryEditorGroupByExpression {
    @JsonProperty("type")
    public QueryEditorExpressionType type;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("property")
    public QueryEditorProperty property;
    public QueryEditorGroupByExpression() {
        this.type = QueryEditorExpressionType.GROUP_BY;
    }
    public QueryEditorGroupByExpression(QueryEditorProperty property) {
        this.type = QueryEditorExpressionType.GROUP_BY;
        this.property = property;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
