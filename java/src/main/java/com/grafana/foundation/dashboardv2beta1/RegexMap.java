// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Maps regular expressions to replacement text and a color.
// For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
public class RegexMap {
    @JsonProperty("type")
    public MappingType type;
    // Regular expression to match against and the result to apply when the value matches the regex
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("options")
    public Dashboardv2beta1RegexMapOptions options;
    public RegexMap() {
        this.type = MappingType.REGEX;
        this.options = new com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1RegexMapOptions();
    }
    public RegexMap(Dashboardv2beta1RegexMapOptions options) {
        this.type = MappingType.REGEX;
        this.options = options;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
