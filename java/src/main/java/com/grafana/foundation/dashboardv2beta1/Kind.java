// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// --- Common types ---
public class Kind {
    @JsonProperty("kind")
    public String kind;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("spec")
    public Object spec;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("metadata")
    public Object metadata;
    public Kind() {
        this.kind = "";
        this.spec = new Object();
    }
    public Kind(String kind,Object spec,Object metadata) {
        this.kind = kind;
        this.spec = spec;
        this.metadata = metadata;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
