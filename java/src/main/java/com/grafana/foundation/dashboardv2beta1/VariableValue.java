// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// Variable types
@JsonDeserialize(using = VariableValueDeserializer.class)
public class VariableValue {
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
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("ArrayOfVariableValueSingle")
    protected List<VariableValueSingle> arrayOfVariableValueSingle;
    protected VariableValue() {}
    public static VariableValue createString(String string) {
        VariableValue variableValue = new VariableValue();
        variableValue.string = string;
        return variableValue;
    }
    public static VariableValue createBool(Boolean bool) {
        VariableValue variableValue = new VariableValue();
        variableValue.bool = bool;
        return variableValue;
    }
    public static VariableValue createFloat64(Double float64) {
        VariableValue variableValue = new VariableValue();
        variableValue.float64 = float64;
        return variableValue;
    }
    public static VariableValue createCustomVariableValue(CustomVariableValue customVariableValue) {
        VariableValue variableValue = new VariableValue();
        variableValue.customVariableValue = customVariableValue;
        return variableValue;
    }
    public static VariableValue createArrayOfVariableValueSingle(List<VariableValueSingle> arrayOfVariableValueSingle) {
        VariableValue variableValue = new VariableValue();
        variableValue.arrayOfVariableValueSingle = arrayOfVariableValueSingle;
        return variableValue;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
