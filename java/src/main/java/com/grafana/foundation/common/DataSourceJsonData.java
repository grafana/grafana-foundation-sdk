// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// TODO docs
public class DataSourceJsonData { 
    @JsonProperty("authType")
    public String authType; 
    @JsonProperty("defaultRegion")
    public String defaultRegion; 
    @JsonProperty("profile")
    public String profile; 
    @JsonProperty("manageAlerts")
    public Boolean manageAlerts; 
    @JsonProperty("alertmanagerUid")
    public String alertmanagerUid;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<DataSourceJsonData> {
        private final DataSourceJsonData internal;
        
        public Builder() {
            this.internal = new DataSourceJsonData();
        }
    public Builder authType(String authType) {
    this.internal.authType = authType;
        return this;
    }
    
    public Builder defaultRegion(String defaultRegion) {
    this.internal.defaultRegion = defaultRegion;
        return this;
    }
    
    public Builder profile(String profile) {
    this.internal.profile = profile;
        return this;
    }
    
    public Builder manageAlerts(Boolean manageAlerts) {
    this.internal.manageAlerts = manageAlerts;
        return this;
    }
    
    public Builder alertmanagerUid(String alertmanagerUid) {
    this.internal.alertmanagerUid = alertmanagerUid;
        return this;
    }
    public DataSourceJsonData build() {
            return this.internal;
        }
    }
}
