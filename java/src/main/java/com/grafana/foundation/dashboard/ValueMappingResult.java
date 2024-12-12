// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Result used as replacement with text and color when the value matches
public class ValueMappingResult {
    // Text to display when the value matches
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public String text;
    // Text to use when the value matches
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("color")
    public String color;
    // Icon to display when the value matches. Only specific visualizations.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("icon")
    public String icon;
    // Position in the mapping array. Only used internally.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("index")
    public Integer index;
    public ValueMappingResult() {}
    
    public ValueMappingResult(String text,String color,String icon,Integer index) {
        this.text = text;
        this.color = color;
        this.icon = icon;
        this.index = index;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
