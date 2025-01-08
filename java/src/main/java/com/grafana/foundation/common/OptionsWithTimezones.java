// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// TODO docs
public class OptionsWithTimezones {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timezone")
    public List<String> timezone;
    public OptionsWithTimezones() {
    }
    
    public OptionsWithTimezones(List<String> timezone) {
        this.timezone = timezone;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
