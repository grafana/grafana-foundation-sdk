// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class PanelRef {
    @JsonProperty("dashboardUID")
    public String dashboardUID;
    @JsonProperty("panelID")
    public Long panelID;
    public PanelRef() {
        this.dashboardUID = "";
        this.panelID = 0L;
    }
    public PanelRef(String dashboardUID,Long panelID) {
        this.dashboardUID = dashboardUID;
        this.panelID = panelID;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
