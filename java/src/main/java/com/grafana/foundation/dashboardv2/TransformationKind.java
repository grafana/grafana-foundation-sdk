// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class TransformationKind {
    @JsonProperty("kind")
    public String kind;
    // The group is the transformation ID
    @JsonProperty("group")
    public String group;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("spec")
    public TransformationSpec spec;
    public TransformationKind() {
        this.kind = "";
        this.group = "";
        this.spec = new com.grafana.foundation.dashboardv2.TransformationSpec();
    }
    public TransformationKind(String kind,String group,TransformationSpec spec) {
        this.kind = kind;
        this.group = group;
        this.spec = spec;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
