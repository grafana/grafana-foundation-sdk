// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.statushistory;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.HideSeriesConfig;

public class FieldConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lineWidth")
    public Integer lineWidth;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hideFrom")
    public HideSeriesConfig hideFrom;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillOpacity")
    public Integer fillOpacity;
    public FieldConfig() {
        this.lineWidth = 1;
        this.fillOpacity = 70;
    }
    
    public FieldConfig(Integer lineWidth,HideSeriesConfig hideFrom,Integer fillOpacity) {
        this.lineWidth = lineWidth;
        this.hideFrom = hideFrom;
        this.fillOpacity = fillOpacity;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
