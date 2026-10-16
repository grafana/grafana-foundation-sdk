// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * A cell holds non-panel narrative content (markdown text, code) in a notebook layout.
 * Panel cells are not represented here — they reuse V2PanelKind.
 */
public class CellKind {
    @JsonProperty("kind")
    public String kind;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("spec")
    public CellSpec spec;
    public CellKind() {
        this.kind = "";
        this.spec = new com.grafana.foundation.dashboardv2beta1.CellSpec();
    }
    public CellKind(String kind,CellSpec spec) {
        this.kind = kind;
        this.spec = spec;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
