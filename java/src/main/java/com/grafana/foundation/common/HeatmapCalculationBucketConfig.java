// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class HeatmapCalculationBucketConfig {
    // Sets the bucket calculation mode 
    @JsonProperty("mode")
    public HeatmapCalculationMode mode;
    // The number of buckets to use for the axis in the heatmap 
    @JsonProperty("value")
    public String value;
    // Controls the scale of the buckets 
    @JsonProperty("scale")
    public ScaleDistributionConfig scale;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<HeatmapCalculationBucketConfig> {
        private final HeatmapCalculationBucketConfig internal;
        
        public Builder() {
            this.internal = new HeatmapCalculationBucketConfig();
        }
    public Builder mode(HeatmapCalculationMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder value(String value) {
    this.internal.value = value;
        return this;
    }
    
    public Builder scale(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scale) {
    this.internal.scale = scale.build();
        return this;
    }
    public HeatmapCalculationBucketConfig build() {
            return this.internal;
        }
    }
}
