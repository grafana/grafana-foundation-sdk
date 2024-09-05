// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.ScaleDimensionConfig;
import java.util.List;

public class CanvasConnection {
    @JsonProperty("source")
    public ConnectionCoordinates source;
    @JsonProperty("target")
    public ConnectionCoordinates target;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("targetName")
    public String targetName;
    @JsonProperty("path")
    public ConnectionPath path;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("color")
    public ColorDimensionConfig color;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("size")
    public ScaleDimensionConfig size;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("vertices")
    public List<ConnectionCoordinates> vertices;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("sourceOriginal")
    public ConnectionCoordinates sourceOriginal;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("targetOriginal")
    public ConnectionCoordinates targetOriginal;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
