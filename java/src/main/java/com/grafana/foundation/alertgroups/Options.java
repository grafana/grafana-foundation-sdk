// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertgroups;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Options {
    // Comma-separated list of values used to filter alert results 
    @JsonProperty("labels")
    public String labels;
    // Name of the alertmanager used as a source for alerts 
    @JsonProperty("alertmanager")
    public String alertmanager;
    // Expand all alert groups by default 
    @JsonProperty("expandAll")
    public Boolean expandAll;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
