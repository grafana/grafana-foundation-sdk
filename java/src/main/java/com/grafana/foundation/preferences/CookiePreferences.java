// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferences;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class CookiePreferences {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("analytics")
    public Object analytics;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("performance")
    public Object performance;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("functional")
    public Object functional;
    public CookiePreferences() {
    }
    public CookiePreferences(Object analytics,Object performance,Object functional) {
        this.analytics = analytics;
        this.performance = performance;
        this.functional = functional;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
