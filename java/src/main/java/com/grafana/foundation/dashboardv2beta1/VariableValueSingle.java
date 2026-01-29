// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = VariableValueSingleDeserializer.class)
public class VariableValueSingle {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("String")
    protected String string;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("Bool")
    protected Boolean bool;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("Float64")
    protected Double float64;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("CustomVariableValue")
    protected CustomVariableValue customVariableValue;
    protected VariableValueSingle() {}
    public static VariableValueSingle createString(String string) {
        VariableValueSingle variableValueSingle = new VariableValueSingle();
        variableValueSingle.string = string;
        return variableValueSingle;
    }
    public static VariableValueSingle createBool(Boolean bool) {
        VariableValueSingle variableValueSingle = new VariableValueSingle();
        variableValueSingle.bool = bool;
        return variableValueSingle;
    }
    public static VariableValueSingle createFloat64(Double float64) {
        VariableValueSingle variableValueSingle = new VariableValueSingle();
        variableValueSingle.float64 = float64;
        return variableValueSingle;
    }
    public static VariableValueSingle createCustomVariableValue(CustomVariableValue customVariableValue) {
        VariableValueSingle variableValueSingle = new VariableValueSingle();
        variableValueSingle.customVariableValue = customVariableValue;
        return variableValueSingle;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
