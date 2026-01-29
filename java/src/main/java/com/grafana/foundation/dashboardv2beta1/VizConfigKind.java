// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = VizConfigKindDeserializer.class)
public class VizConfigKind {
    @JsonProperty("kind")
    public String kind;
    // The group is the plugin ID
    @JsonProperty("group")
    public String group;
    @JsonProperty("version")
    public String version;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("spec")
    public VizConfigSpec spec;
    public VizConfigKind() {
        this.kind = "";
        this.group = "";
        this.version = "";
        this.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
    }
    public VizConfigKind(String kind,String group,String version,VizConfigSpec spec) {
        this.kind = kind;
        this.group = group;
        this.version = version;
        this.spec = spec;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
