// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Result used as replacement with text and color when the value matches
public class ValueMappingResult {
    // Text to display when the value matches
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("text")
    public String text;
    // Text to use when the value matches
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("color")
    public String color;
    // Icon to display when the value matches. Only specific visualizations.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("icon")
    public String icon;
    // Position in the mapping array. Only used internally.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("index")
    public Integer index;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ValueMappingResult> {
        private final ValueMappingResult internal;
        
        public Builder() {
            this.internal = new ValueMappingResult();
        }
    public Builder text(String text) {
    this.internal.text = text;
        return this;
    }
    
    public Builder color(String color) {
    this.internal.color = color;
        return this;
    }
    
    public Builder icon(String icon) {
    this.internal.icon = icon;
        return this;
    }
    
    public Builder index(Integer index) {
    this.internal.index = index;
        return this;
    }
    public ValueMappingResult build() {
            return this.internal;
        }
    }
}
