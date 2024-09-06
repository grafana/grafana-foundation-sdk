// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Options {
    // Enable inline editing
    @JsonProperty("inlineEditing")
    public Boolean inlineEditing;
    // Show all available element types
    @JsonProperty("showAdvancedTypes")
    public Boolean showAdvancedTypes;
    // Enable pan and zoom
    @JsonProperty("panZoom")
    public Boolean panZoom;
    // The root element of canvas (frame), where all canvas elements are nested
    // TODO: Figure out how to define a default value for this
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("root")
    public CanvasOptionsRoot root;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
