// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class AnnotationPermission {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dashboard")
    public AnnotationActions dashboard;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("organization")
    public AnnotationActions organization;
    public AnnotationPermission() {
    }
    public AnnotationPermission(AnnotationActions dashboard,AnnotationActions organization) {
        this.dashboard = dashboard;
        this.organization = organization;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
