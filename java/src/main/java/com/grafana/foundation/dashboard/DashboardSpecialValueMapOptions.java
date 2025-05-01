// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class DashboardSpecialValueMapOptions {
    // Special value to match against
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("match")
    public SpecialValueMatch match;
    // Config to apply when the value matches the special value
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("result")
    public ValueMappingResult result;
    public DashboardSpecialValueMapOptions() {
        this.match = SpecialValueMatch.TRUE;
        this.result = new com.grafana.foundation.dashboard.ValueMappingResult();
    }
    public DashboardSpecialValueMapOptions(SpecialValueMatch match,ValueMappingResult result) {
        this.match = match;
        this.result = result;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
