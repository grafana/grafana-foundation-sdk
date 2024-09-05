// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.ResourceDimensionConfig;

public class BackgroundConfig {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("color")
    public ColorDimensionConfig color;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("image")
    public ResourceDimensionConfig image;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("size")
    public BackgroundImageSize size;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
