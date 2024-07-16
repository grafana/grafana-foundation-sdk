// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class AnnotationActions { 
    @JsonProperty("canAdd")
    public Boolean canAdd; 
    @JsonProperty("canDelete")
    public Boolean canDelete; 
    @JsonProperty("canEdit")
    public Boolean canEdit;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AnnotationActions> {
        private final AnnotationActions internal;
        
        public Builder() {
            this.internal = new AnnotationActions();
        }
    public Builder canAdd(Boolean canAdd) {
    this.internal.canAdd = canAdd;
        return this;
    }
    
    public Builder canDelete(Boolean canDelete) {
    this.internal.canDelete = canDelete;
        return this;
    }
    
    public Builder canEdit(Boolean canEdit) {
    this.internal.canEdit = canEdit;
        return this;
    }
    public AnnotationActions build() {
            return this.internal;
        }
    }
}