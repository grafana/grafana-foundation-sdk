// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.MapLayerOptions;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class Options {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("view")
    public MapViewConfig view;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("controls")
    public ControlsOptions controls;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("basemap")
    public MapLayerOptions basemap;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("layers")
    public List<MapLayerOptions> layers;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public TooltipOptions tooltip;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
