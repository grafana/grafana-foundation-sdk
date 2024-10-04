// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Controls various color options
public class HeatmapColorOptions {
    // Sets the color mode
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mode")
    public HeatmapColorMode mode;
    // Controls the color scheme used
    @JsonProperty("scheme")
    public String scheme;
    // Controls the color fill when in opacity mode
    @JsonProperty("fill")
    public String fill;
    // Controls the color scale
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("scale")
    public HeatmapColorScale scale;
    // Controls the exponent when scale is set to exponential
    @JsonProperty("exponent")
    public Float exponent;
    // Controls the number of color steps
    @JsonProperty("steps")
    public Long steps;
    // Reverses the color scheme
    @JsonProperty("reverse")
    public Boolean reverse;
    // Sets the minimum value for the color scale
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("min")
    public Float min;
    // Sets the maximum value for the color scale
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("max")
    public Float max;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<HeatmapColorOptions> {
        protected final HeatmapColorOptions internal;
        
        public Builder() {
            this.internal = new HeatmapColorOptions();
        }
    public Builder mode(HeatmapColorMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder scheme(String scheme) {
    this.internal.scheme = scheme;
        return this;
    }
    
    public Builder fill(String fill) {
    this.internal.fill = fill;
        return this;
    }
    
    public Builder scale(HeatmapColorScale scale) {
    this.internal.scale = scale;
        return this;
    }
    
    public Builder exponent(Float exponent) {
    this.internal.exponent = exponent;
        return this;
    }
    
    public Builder steps(Long steps) {
        if (!(steps >= 2)) {
            throw new IllegalArgumentException("steps must be >= 2");
        }
        if (!(steps <= 128)) {
            throw new IllegalArgumentException("steps must be <= 128");
        }
    this.internal.steps = steps;
        return this;
    }
    
    public Builder reverse(Boolean reverse) {
    this.internal.reverse = reverse;
        return this;
    }
    
    public Builder min(Float min) {
    this.internal.min = min;
        return this;
    }
    
    public Builder max(Float max) {
    this.internal.max = max;
        return this;
    }
    public HeatmapColorOptions build() {
            return this.internal;
        }
    }
}
