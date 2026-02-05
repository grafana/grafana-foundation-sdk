// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// FIXME: should we introduce this? --- Variable value option
public class VariableValueOption {
    @JsonProperty("label")
    public String label;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("value")
    public VariableValueSingle value;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group")
    public String group;
    public VariableValueOption() {
        this.label = "";
        this.value = new com.grafana.foundation.dashboardv2beta1.VariableValueSingle();
    }
    public VariableValueOption(String label,VariableValueSingle value,String group) {
        this.label = label;
        this.value = value;
        this.group = group;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
