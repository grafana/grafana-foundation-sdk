// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Counterpart for TypeScript's TimeOption type.
public class TimeOption {
    @JsonProperty("display")
    public String display;
    @JsonProperty("from")
    public String from;
    @JsonProperty("to")
    public String to;
    public TimeOption() {
    }
    public TimeOption(String display,String from,String to) {
        this.display = display;
        this.from = from;
        this.to = to;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
