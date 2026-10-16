// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class CellSpec {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("content")
    public CellContentKind content;
    public CellSpec() {
        this.content = new com.grafana.foundation.dashboardv2beta1.CellContentKind();
    }
    public CellSpec(CellContentKind content) {
        this.content = content;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
