// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Option to be selected in a variable.
public class VariableOption {
    // Whether the option is selected or not
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("selected")
    public Boolean selected;
    // Text to be displayed for the option
    @JsonProperty("text")
    public StringOrArrayOfString text;
    // Value of the option
    @JsonProperty("value")
    public StringOrArrayOfString value;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
