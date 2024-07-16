// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.statetimeline;

import com.grafana.foundation.common.HideSeriesConfig;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class FieldConfig { 
    @JsonProperty("lineWidth")
    public Integer lineWidth; 
    @JsonProperty("hideFrom")
    public HideSeriesConfig hideFrom; 
    @JsonProperty("fillOpacity")
    public Integer fillOpacity;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
