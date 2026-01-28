// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class PanelQuerySpec {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("query")
    public DataQueryKind query;
    @JsonProperty("refId")
    public String refId;
    @JsonProperty("hidden")
    public Boolean hidden;
    public PanelQuerySpec() {
        this.query = new com.grafana.foundation.dashboardv2beta1.DataQueryKind();
        this.refId = "A";
        this.hidden = false;
    }
    public PanelQuerySpec(DataQueryKind query,String refId,Boolean hidden) {
        this.query = query;
        this.refId = refId;
        this.hidden = hidden;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
