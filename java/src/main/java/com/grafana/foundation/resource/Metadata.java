// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.resource;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;

public class Metadata {
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("namespace")
    public String namespace;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("labels")
    public Map<String, String> labels;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("annotations")
    public Map<String, String> annotations;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("uid")
    public String uid;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resourceVersion")
    public String resourceVersion;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("generation")
    public Long generation;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("creationTimestamp")
    public String creationTimestamp;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("updateTimestamp")
    public String updateTimestamp;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("deletionTimestamp")
    public String deletionTimestamp;
    public Metadata() {
        this.name = "";
    }
    public Metadata(String name,String namespace,Map<String, String> labels,Map<String, String> annotations,String uid,String resourceVersion,Long generation,String creationTimestamp,String updateTimestamp,String deletionTimestamp) {
        this.name = name;
        this.namespace = namespace;
        this.labels = labels;
        this.annotations = annotations;
        this.uid = uid;
        this.resourceVersion = resourceVersion;
        this.generation = generation;
        this.creationTimestamp = creationTimestamp;
        this.updateTimestamp = updateTimestamp;
        this.deletionTimestamp = deletionTimestamp;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
