// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Maps regular expressions to replacement text and a color.
// For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
public class RegexMap {
    @JsonProperty("type")
    public String type;
    // Regular expression to match against and the result to apply when the value matches the regex
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("options")
    public DashboardRegexMapOptions options;
    public RegexMap() {
    }
    
    public RegexMap(String type,DashboardRegexMapOptions options) {
        this.type = type;
        this.options = options;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
