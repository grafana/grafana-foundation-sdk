// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * The notebook's own panel chain. It is a copy of the dashboard one down to the transformation,
 * which follows the dashboard v2 shape rather than the v2beta1 shape in this package. The chain has
 * to be forked rather than shared because PanelKind reaches TransformationKind through
 * QueryGroupKind, and those three are what Dashboard v2beta1 serves. Everything the chain does not
 * change (DataLink, VizConfigKind, PanelQueryKind, QueryOptionsSpec) stays shared.
 */
public class V2PanelKind {
    @JsonProperty("kind")
    public String kind;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("spec")
    public V2PanelSpec spec;
    public V2PanelKind() {
        this.kind = "";
        this.spec = new com.grafana.foundation.dashboardv2beta1.V2PanelSpec();
    }
    public V2PanelKind(String kind,V2PanelSpec spec) {
        this.kind = kind;
        this.spec = spec;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
