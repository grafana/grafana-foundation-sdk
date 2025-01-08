// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class HeatmapCalculationBucketConfig {
    // Sets the bucket calculation mode
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mode")
    public HeatmapCalculationMode mode;
    // The number of buckets to use for the axis in the heatmap
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("value")
    public String value;
    // Controls the scale of the buckets
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("scale")
    public ScaleDistributionConfig scale;
    public HeatmapCalculationBucketConfig() {
    }
    
    public HeatmapCalculationBucketConfig(HeatmapCalculationMode mode,String value,ScaleDistributionConfig scale) {
        this.mode = mode;
        this.value = value;
        this.scale = scale;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
