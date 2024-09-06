// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class MapViewConfig {
    @JsonProperty("id")
    public String id;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lat")
    public Long lat;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lon")
    public Long lon;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("zoom")
    public Long zoom;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("minZoom")
    public Long minZoom;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxZoom")
    public Long maxZoom;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("padding")
    public Long padding;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("allLayers")
    public Boolean allLayers;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lastOnly")
    public Boolean lastOnly;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("layer")
    public String layer;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("shared")
    public Boolean shared;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
