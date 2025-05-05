// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.resource;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Manifest {
    @JsonProperty("apiVersion")
    public String apiVersion;
    @JsonProperty("kind")
    public String kind;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("metadata")
    public Metadata metadata;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("spec")
    public Object spec;
    public Manifest() {
        this.apiVersion = "";
        this.kind = "";
        this.metadata = new com.grafana.foundation.resource.Metadata();
        this.spec = new Object();
    }
    public Manifest(String apiVersion,String kind,Metadata metadata,Object spec) {
        this.apiVersion = apiVersion;
        this.kind = kind;
        this.metadata = metadata;
        this.spec = spec;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
