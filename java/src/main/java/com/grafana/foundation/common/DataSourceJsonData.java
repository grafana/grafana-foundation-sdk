// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class DataSourceJsonData {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("authType")
    public String authType;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("defaultRegion")
    public String defaultRegion;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("profile")
    public String profile;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("manageAlerts")
    public Boolean manageAlerts;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alertmanagerUid")
    public String alertmanagerUid;
    public DataSourceJsonData() {
    }
    public DataSourceJsonData(String authType,String defaultRegion,String profile,Boolean manageAlerts,String alertmanagerUid) {
        this.authType = authType;
        this.defaultRegion = defaultRegion;
        this.profile = profile;
        this.manageAlerts = manageAlerts;
        this.alertmanagerUid = alertmanagerUid;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
