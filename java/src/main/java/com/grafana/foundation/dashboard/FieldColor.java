// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Map a field to a color.
public class FieldColor {
    // The main color scheme mode. 
    @JsonProperty("mode")
    public FieldColorModeId mode;
    // The fixed color value for fixed or shades color modes. 
    @JsonProperty("fixedColor")
    public String fixedColor;
    // Some visualizations need to know how to assign a series color from by value color schemes. 
    @JsonProperty("seriesBy")
    public FieldColorSeriesByMode seriesBy;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<FieldColor> {
        private final FieldColor internal;
        
        public Builder() {
            this.internal = new FieldColor();
        }
    public Builder mode(FieldColorModeId mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder fixedColor(String fixedColor) {
    this.internal.fixedColor = fixedColor;
        return this;
    }
    
    public Builder seriesBy(FieldColorSeriesByMode seriesBy) {
    this.internal.seriesBy = seriesBy;
        return this;
    }
    public FieldColor build() {
            return this.internal;
        }
    }
}
