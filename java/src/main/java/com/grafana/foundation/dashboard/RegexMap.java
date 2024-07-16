// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Maps regular expressions to replacement text and a color.
// For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
public class RegexMap { 
    @JsonProperty("type")
    public String type;
    // Regular expression to match against and the result to apply when the value matches the regex 
    @JsonProperty("options")
    public DashboardRegexMapOptions options;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<RegexMap> {
        private final RegexMap internal;
        
        public Builder() {
            this.internal = new RegexMap();
    this.internal.type = "regex";
        }
    public Builder options(com.grafana.foundation.cog.Builder<DashboardRegexMapOptions> options) {
    this.internal.options = options.build();
        return this;
    }
    public RegexMap build() {
            return this.internal;
        }
    }
}
