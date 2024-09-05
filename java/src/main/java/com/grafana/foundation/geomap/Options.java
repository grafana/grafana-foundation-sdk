// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.MapLayerOptions;
import java.util.List;

public class Options {
    @JsonProperty("view")
    public MapViewConfig view;
    @JsonProperty("controls")
    public ControlsOptions controls;
    @JsonProperty("basemap")
    public MapLayerOptions basemap;
    @JsonProperty("layers")
    public List<MapLayerOptions> layers;
    @JsonProperty("tooltip")
    public TooltipOptions tooltip;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
