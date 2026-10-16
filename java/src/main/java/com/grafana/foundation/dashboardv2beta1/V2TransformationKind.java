// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * Dashboard v2 shape: the transformation ID moved from `kind` to `group`.
 */
public class V2TransformationKind {
    @JsonProperty("kind")
    public String kind;
    // The group is the transformation ID
    @JsonProperty("group")
    public String group;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("spec")
    public V2TransformationSpec spec;
    public V2TransformationKind() {
        this.kind = "";
        this.group = "";
        this.spec = new com.grafana.foundation.dashboardv2beta1.V2TransformationSpec();
    }
    public V2TransformationKind(String kind,String group,V2TransformationSpec spec) {
        this.kind = kind;
        this.group = group;
        this.spec = spec;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
