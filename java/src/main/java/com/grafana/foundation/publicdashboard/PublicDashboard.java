// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.publicdashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class PublicDashboard {
    // Unique public dashboard identifier
    @JsonProperty("uid")
    public String uid;
    // Dashboard unique identifier referenced by this public dashboard
    @JsonProperty("dashboardUid")
    public String dashboardUid;
    // Unique public access token
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("accessToken")
    public String accessToken;
    // Flag that indicates if the public dashboard is enabled
    @JsonProperty("isEnabled")
    public Boolean isEnabled;
    // Flag that indicates if annotations are enabled
    @JsonProperty("annotationsEnabled")
    public Boolean annotationsEnabled;
    // Flag that indicates if the time range picker is enabled
    @JsonProperty("timeSelectionEnabled")
    public Boolean timeSelectionEnabled;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<PublicDashboard> {
        protected final PublicDashboard internal;
        
        public Builder() {
            this.internal = new PublicDashboard();
        }
    public Builder uid(String uid) {
    this.internal.uid = uid;
        return this;
    }
    
    public Builder dashboardUid(String dashboardUid) {
    this.internal.dashboardUid = dashboardUid;
        return this;
    }
    
    public Builder accessToken(String accessToken) {
    this.internal.accessToken = accessToken;
        return this;
    }
    
    public Builder isEnabled(Boolean isEnabled) {
    this.internal.isEnabled = isEnabled;
        return this;
    }
    
    public Builder annotationsEnabled(Boolean annotationsEnabled) {
    this.internal.annotationsEnabled = annotationsEnabled;
        return this;
    }
    
    public Builder timeSelectionEnabled(Boolean timeSelectionEnabled) {
    this.internal.timeSelectionEnabled = timeSelectionEnabled;
        return this;
    }
    public PublicDashboard build() {
            return this.internal;
        }
    }
}
