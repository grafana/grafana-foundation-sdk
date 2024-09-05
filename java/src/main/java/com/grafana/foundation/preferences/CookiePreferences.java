// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferences;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class CookiePreferences {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("analytics")
    public Object analytics;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("performance")
    public Object performance;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("functional")
    public Object functional;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CookiePreferences> {
        private final CookiePreferences internal;
        
        public Builder() {
            this.internal = new CookiePreferences();
        }
    public Builder analytics(Object analytics) {
    this.internal.analytics = analytics;
        return this;
    }
    
    public Builder performance(Object performance) {
    this.internal.performance = performance;
        return this;
    }
    
    public Builder functional(Object functional) {
    this.internal.functional = functional;
        return this;
    }
    public CookiePreferences build() {
            return this.internal;
        }
    }
}
