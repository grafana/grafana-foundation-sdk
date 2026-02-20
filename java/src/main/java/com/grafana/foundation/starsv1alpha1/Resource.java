// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.starsv1alpha1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class Resource {
    @JsonProperty("group")
    public String group;
    @JsonProperty("kind")
    public String kind;
    // The set of resources
    // +listType=set
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("names")
    public List<String> names;
    public Resource() {
        this.group = "";
        this.kind = "";
        this.names = new LinkedList<>();
    }
    public Resource(String group,String kind,List<String> names) {
        this.group = group;
        this.kind = kind;
        this.names = names;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
