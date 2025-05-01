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
    public PublicDashboard() {
        this.uid = "";
        this.dashboardUid = "";
        this.isEnabled = false;
        this.annotationsEnabled = false;
        this.timeSelectionEnabled = false;
    }
    public PublicDashboard(String uid,String dashboardUid,String accessToken,Boolean isEnabled,Boolean annotationsEnabled,Boolean timeSelectionEnabled) {
        this.uid = uid;
        this.dashboardUid = dashboardUid;
        this.accessToken = accessToken;
        this.isEnabled = isEnabled;
        this.annotationsEnabled = annotationsEnabled;
        this.timeSelectionEnabled = timeSelectionEnabled;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
