// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class AnnotationPermission {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("dashboard")
    public AnnotationActions dashboard;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("organization")
    public AnnotationActions organization;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AnnotationPermission> {
        private final AnnotationPermission internal;
        
        public Builder() {
            this.internal = new AnnotationPermission();
        }
    public Builder dashboard(com.grafana.foundation.cog.Builder<AnnotationActions> dashboard) {
    this.internal.dashboard = dashboard.build();
        return this;
    }
    
    public Builder organization(com.grafana.foundation.cog.Builder<AnnotationActions> organization) {
    this.internal.organization = organization.build();
        return this;
    }
    public AnnotationPermission build() {
            return this.internal;
        }
    }
}
