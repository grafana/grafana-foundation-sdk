// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class QueryEditorPropertyExpression {
    @JsonProperty("type")
    public QueryEditorExpressionType type;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("property")
    public QueryEditorProperty property;
    public QueryEditorPropertyExpression() {
        this.type = QueryEditorExpressionType.PROPERTY;
        this.property = new com.grafana.foundation.bigquery.QueryEditorPropertyBuilder().build();
    }
    public QueryEditorPropertyExpression(QueryEditorProperty property) {
        this.type = QueryEditorExpressionType.PROPERTY;
        this.property = property;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
