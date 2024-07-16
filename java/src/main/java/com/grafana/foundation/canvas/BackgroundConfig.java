// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.ResourceDimensionConfig;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class BackgroundConfig { 
    @JsonProperty("color")
    public ColorDimensionConfig color; 
    @JsonProperty("image")
    public ResourceDimensionConfig image; 
    @JsonProperty("size")
    public BackgroundImageSize size;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
