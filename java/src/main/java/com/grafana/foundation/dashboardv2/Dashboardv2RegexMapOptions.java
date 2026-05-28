// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Dashboardv2RegexMapOptions {
    // Regular expression to match against
    @JsonProperty("pattern")
    public String pattern;
    // Config to apply when the value matches the regex
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("result")
    public ValueMappingResult result;
    public Dashboardv2RegexMapOptions() {
        this.pattern = "";
        this.result = new com.grafana.foundation.dashboardv2.ValueMappingResult();
    }
    public Dashboardv2RegexMapOptions(String pattern,ValueMappingResult result) {
        this.pattern = pattern;
        this.result = result;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
