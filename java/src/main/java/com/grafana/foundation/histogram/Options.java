// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.histogram;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;

public class Options {
    // Size of each bucket
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("bucketSize")
    public Integer bucketSize;
    // Offset buckets by this amount
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("bucketOffset")
    public Integer bucketOffset;
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    // Combines multiple series into a single histogram
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("combine")
    public Boolean combine;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
