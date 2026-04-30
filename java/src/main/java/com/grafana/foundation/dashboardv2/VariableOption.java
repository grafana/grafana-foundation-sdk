// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;

// Variable option specification
public class VariableOption {
    // Whether the option is selected or not
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("selected")
    public Boolean selected;
    // Text to be displayed for the option
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("text")
    public StringOrArrayOfString text;
    // Value of the option
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("value")
    public StringOrArrayOfString value;
    // Additional properties for multi-props variables
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("properties")
    public Map<String, String> properties;
    public VariableOption() {
        this.text = new com.grafana.foundation.dashboardv2.StringOrArrayOfString();
        this.value = new com.grafana.foundation.dashboardv2.StringOrArrayOfString();
    }
    public VariableOption(Boolean selected,StringOrArrayOfString text,StringOrArrayOfString value,Map<String, String> properties) {
        this.selected = selected;
        this.text = text;
        this.value = value;
        this.properties = properties;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
