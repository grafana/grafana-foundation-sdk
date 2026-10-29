// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

/**
 * #SearchResultResource is the full identity of a hit.
 */
public class SearchResultResource {
    @JsonProperty("group")
    public String group;
    @JsonProperty("resource")
    public String resource;
    @JsonProperty("kind")
    public String kind;
    @JsonProperty("name")
    public String name;
    public SearchResultResource() {
        this.group = "";
        this.resource = "";
        this.kind = "";
        this.name = "";
    }
    public SearchResultResource(String group,String resource,String kind,String name) {
        this.group = group;
        this.resource = resource;
        this.kind = kind;
        this.name = name;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
