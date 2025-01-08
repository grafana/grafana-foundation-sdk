// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class HeatmapCalculationOptions {
    // The number of buckets to use for the xAxis in the heatmap
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("xBuckets")
    public HeatmapCalculationBucketConfig xBuckets;
    // The number of buckets to use for the yAxis in the heatmap
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("yBuckets")
    public HeatmapCalculationBucketConfig yBuckets;
    public HeatmapCalculationOptions() {
    }
    
    public HeatmapCalculationOptions(HeatmapCalculationBucketConfig xBuckets,HeatmapCalculationBucketConfig yBuckets) {
        this.xBuckets = xBuckets;
        this.yBuckets = yBuckets;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
