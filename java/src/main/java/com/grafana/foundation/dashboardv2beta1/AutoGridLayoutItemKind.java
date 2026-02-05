// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class AutoGridLayoutItemKind {
    @JsonProperty("kind")
    public String kind;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("spec")
    public AutoGridLayoutItemSpec spec;
    public AutoGridLayoutItemKind() {
        this.kind = "";
        this.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutItemSpec();
    }
    public AutoGridLayoutItemKind(String kind,AutoGridLayoutItemSpec spec) {
        this.kind = kind;
        this.spec = spec;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
