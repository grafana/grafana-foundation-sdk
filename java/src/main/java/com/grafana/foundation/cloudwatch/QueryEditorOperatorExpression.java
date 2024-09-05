// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class QueryEditorOperatorExpression {
    @JsonProperty("type")
    public String type;
    @JsonProperty("property")
    public QueryEditorProperty property;
    // TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
    @JsonProperty("operator")
    public QueryEditorOperator operator;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<QueryEditorOperatorExpression> {
        private final QueryEditorOperatorExpression internal;
        
        public Builder() {
            this.internal = new QueryEditorOperatorExpression();
    this.internal.type = "operator";
        }
    public Builder property(com.grafana.foundation.cog.Builder<QueryEditorProperty> property) {
    this.internal.property = property.build();
        return this;
    }
    
    public Builder operator(com.grafana.foundation.cog.Builder<QueryEditorOperator> operator) {
    this.internal.operator = operator.build();
        return this;
    }
    public QueryEditorOperatorExpression build() {
            return this.internal;
        }
    }
}
