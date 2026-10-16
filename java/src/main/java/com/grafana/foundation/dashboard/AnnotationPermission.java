// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * +k8s:deepcopy-gen=true
 */
public class AnnotationPermission {
    // +k8s:deepcopy-gen=true
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dashboard")
    public AnnotationActions dashboard;
    public AnnotationPermission() {
    }
    public AnnotationPermission(AnnotationActions dashboard) {
        this.dashboard = dashboard;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
