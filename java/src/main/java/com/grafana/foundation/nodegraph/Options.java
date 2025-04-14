// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.nodegraph;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Options {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("nodes")
    public NodeOptions nodes;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("edges")
    public EdgeOptions edges;
    // How to handle zoom/scroll events in the node graph
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("zoomMode")
    public ZoomMode zoomMode;
    // How to layout the nodes in the node graph
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("layoutAlgorithm")
    public LayoutAlgorithm layoutAlgorithm;
    public Options() {
    }
    public Options(NodeOptions nodes,EdgeOptions edges,ZoomMode zoomMode,LayoutAlgorithm layoutAlgorithm) {
        this.nodes = nodes;
        this.edges = edges;
        this.zoomMode = zoomMode;
        this.layoutAlgorithm = layoutAlgorithm;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
