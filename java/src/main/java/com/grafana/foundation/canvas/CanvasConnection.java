// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.ScaleDimensionConfig;

public class CanvasConnection {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("source")
    public ConnectionCoordinates source;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("target")
    public ConnectionCoordinates target;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("targetName")
    public String targetName;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("path")
    public ConnectionPath path;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("color")
    public ColorDimensionConfig color;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("size")
    public ScaleDimensionConfig size;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
