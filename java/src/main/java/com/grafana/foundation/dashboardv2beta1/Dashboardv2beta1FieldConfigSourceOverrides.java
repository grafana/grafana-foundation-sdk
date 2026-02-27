// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class Dashboardv2beta1FieldConfigSourceOverrides {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("__systemRef")
    public String systemRef;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("matcher")
    public MatcherConfig matcher;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("properties")
    public List<DynamicConfigValue> properties;
    public Dashboardv2beta1FieldConfigSourceOverrides() {
        this.matcher = new com.grafana.foundation.dashboardv2beta1.MatcherConfig();
        this.properties = new LinkedList<>();
    }
    public Dashboardv2beta1FieldConfigSourceOverrides(String systemRef,MatcherConfig matcher,List<DynamicConfigValue> properties) {
        this.systemRef = systemRef;
        this.matcher = matcher;
        this.properties = properties;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
