// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Custom formatter variable
public class CustomFormatterVariable {
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public VariableType type;
    @JsonProperty("multi")
    public Boolean multi;
    @JsonProperty("includeAll")
    public Boolean includeAll;
    public CustomFormatterVariable() {
        this.name = "";
        this.type = VariableType.QUERY;
        this.multi = false;
        this.includeAll = false;
    }
    public CustomFormatterVariable(String name,VariableType type,Boolean multi,Boolean includeAll) {
        this.name = name;
        this.type = type;
        this.multi = multi;
        this.includeAll = includeAll;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
