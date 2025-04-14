// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// +k8s:deepcopy-gen=true
public class AnnotationActions {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("canAdd")
    public Boolean canAdd;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("canDelete")
    public Boolean canDelete;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("canEdit")
    public Boolean canEdit;
    public AnnotationActions() {
    }
    public AnnotationActions(Boolean canAdd,Boolean canDelete,Boolean canEdit) {
        this.canAdd = canAdd;
        this.canDelete = canDelete;
        this.canEdit = canEdit;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
