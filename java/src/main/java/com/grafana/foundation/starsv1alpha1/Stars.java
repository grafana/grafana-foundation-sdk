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

public class Stars {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("resource")
    public List<Resource> resource;
    public Stars() {
        this.resource = new LinkedList<>();
    }
    public Stars(List<Resource> resource) {
        this.resource = resource;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
