// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class QueryEditorArrayExpression {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public QueryEditorArrayExpressionType type;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("expressions")
    public List<QueryEditorExpression> expressions;
    public QueryEditorArrayExpression() {
    }
    
    public QueryEditorArrayExpression(QueryEditorArrayExpressionType type,List<QueryEditorExpression> expressions) {
        this.type = type;
        this.expressions = expressions;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
