// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Dashboard specific preferences (applied per dashboard = all users using the dashboard)
public class Preferences {
    // default layout template to be used when new containers are created
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("layout")
    public AutoGridLayoutKindOrGridLayoutKind layout;
    public Preferences() {
    }
    public Preferences(AutoGridLayoutKindOrGridLayoutKind layout) {
        this.layout = layout;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
