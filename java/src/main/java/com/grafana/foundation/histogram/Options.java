// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.histogram;

import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Options {
    // Bucket count (approx) 
    @JsonProperty("bucketCount")
    public Integer bucketCount;
    // Size of each bucket 
    @JsonProperty("bucketSize")
    public Integer bucketSize;
    // Offset buckets by this amount 
    @JsonProperty("bucketOffset")
    public Float bucketOffset; 
    @JsonProperty("legend")
    public VizLegendOptions legend; 
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    // Combines multiple series into a single histogram 
    @JsonProperty("combine")
    public Boolean combine;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
