// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class QueryEditorOperatorExpression {
    @JsonProperty("type")
    public QueryEditorExpressionType type;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("property")
    public QueryEditorProperty property;
    // TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("operator")
    public QueryEditorOperator operator;
    public QueryEditorOperatorExpression() {
        this.type = QueryEditorExpressionType.OPERATOR;
        this.property = new com.grafana.foundation.cloudwatch.QueryEditorProperty();
        this.operator = new com.grafana.foundation.cloudwatch.QueryEditorOperator();
    }
    public QueryEditorOperatorExpression(QueryEditorProperty property,QueryEditorOperator operator) {
        this.type = QueryEditorExpressionType.OPERATOR;
        this.property = property;
        this.operator = operator;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
