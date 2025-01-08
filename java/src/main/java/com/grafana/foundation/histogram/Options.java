// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.histogram;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;

public class Options {
    // Bucket count (approx)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("bucketCount")
    public Integer bucketCount;
    // Size of each bucket
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("bucketSize")
    public Integer bucketSize;
    // Offset buckets by this amount
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("bucketOffset")
    public Float bucketOffset;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    // Combines multiple series into a single histogram
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("combine")
    public Boolean combine;
    public Options() {
        this.bucketCount = 30;
        this.bucketOffset = 0.0f;
    }
    
    public Options(Integer bucketCount,Integer bucketSize,Float bucketOffset,VizLegendOptions legend,VizTooltipOptions tooltip,Boolean combine) {
        this.bucketCount = bucketCount;
        this.bucketSize = bucketSize;
        this.bucketOffset = bucketOffset;
        this.legend = legend;
        this.tooltip = tooltip;
        this.combine = combine;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
