// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.ScaleDimensionConfig;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class CanvasConnection { 
    @JsonProperty("source")
    public ConnectionCoordinates source; 
    @JsonProperty("target")
    public ConnectionCoordinates target; 
    @JsonProperty("targetName")
    public String targetName; 
    @JsonProperty("path")
    public ConnectionPath path; 
    @JsonProperty("color")
    public ColorDimensionConfig color; 
    @JsonProperty("size")
    public ScaleDimensionConfig size; 
    @JsonProperty("vertices")
    public List<ConnectionCoordinates> vertices; 
    @JsonProperty("sourceOriginal")
    public ConnectionCoordinates sourceOriginal; 
    @JsonProperty("targetOriginal")
    public ConnectionCoordinates targetOriginal;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
