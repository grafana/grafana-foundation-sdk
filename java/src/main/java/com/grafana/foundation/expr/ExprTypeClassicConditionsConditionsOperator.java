// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ExprTypeClassicConditionsConditionsOperator {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public ExprTypeClassicConditionsConditionsOperatorType type;
    public ExprTypeClassicConditionsConditionsOperator() {
        this.type = ExprTypeClassicConditionsConditionsOperatorType.AND;
    }
    public ExprTypeClassicConditionsConditionsOperator(ExprTypeClassicConditionsConditionsOperatorType type) {
        this.type = type;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
