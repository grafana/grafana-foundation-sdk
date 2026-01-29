// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ConditionalRenderingVariableSpec {
    @JsonProperty("variable")
    public String variable;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("operator")
    public ConditionalRenderingVariableSpecOperator operator;
    @JsonProperty("value")
    public String value;
    public ConditionalRenderingVariableSpec() {
        this.variable = "";
        this.operator = ConditionalRenderingVariableSpecOperator.EQUALS;
        this.value = "";
    }
    public ConditionalRenderingVariableSpec(String variable,ConditionalRenderingVariableSpecOperator operator,String value) {
        this.variable = variable;
        this.operator = operator;
        this.value = value;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
