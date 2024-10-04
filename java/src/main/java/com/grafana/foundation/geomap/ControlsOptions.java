// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ControlsOptions {
    // Zoom (upper left)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showZoom")
    public Boolean showZoom;
    // let the mouse wheel zoom
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mouseWheelZoom")
    public Boolean mouseWheelZoom;
    // Lower right
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showAttribution")
    public Boolean showAttribution;
    // Scale options
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showScale")
    public Boolean showScale;
    // Show debug
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showDebug")
    public Boolean showDebug;
    // Show measure
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showMeasure")
    public Boolean showMeasure;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ControlsOptions> {
        protected final ControlsOptions internal;
        
        public Builder() {
            this.internal = new ControlsOptions();
        }
    public Builder showZoom(Boolean showZoom) {
    this.internal.showZoom = showZoom;
        return this;
    }
    
    public Builder mouseWheelZoom(Boolean mouseWheelZoom) {
    this.internal.mouseWheelZoom = mouseWheelZoom;
        return this;
    }
    
    public Builder showAttribution(Boolean showAttribution) {
    this.internal.showAttribution = showAttribution;
        return this;
    }
    
    public Builder showScale(Boolean showScale) {
    this.internal.showScale = showScale;
        return this;
    }
    
    public Builder showDebug(Boolean showDebug) {
    this.internal.showDebug = showDebug;
        return this;
    }
    
    public Builder showMeasure(Boolean showMeasure) {
    this.internal.showMeasure = showMeasure;
        return this;
    }
    public ControlsOptions build() {
            return this.internal;
        }
    }
}
