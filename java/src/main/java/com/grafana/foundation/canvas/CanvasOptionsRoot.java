// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class CanvasOptionsRoot {
    // Name of the root element
    @JsonProperty("name")
    public String name;
    // Type of root element (frame)
    @JsonProperty("type")
    public String type;
    // The list of canvas elements attached to the root element
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("elements")
    public List<CanvasElementOptions> elements;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CanvasOptionsRoot> {
        protected final CanvasOptionsRoot internal;
        
        public Builder() {
            this.internal = new CanvasOptionsRoot();
    this.internal.type = "frame";
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder elements(com.grafana.foundation.cog.Builder<List<CanvasElementOptions>> elements) {
    this.internal.elements = elements.build();
        return this;
    }
    public CanvasOptionsRoot build() {
            return this.internal;
        }
    }
}
