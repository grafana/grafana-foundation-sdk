// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Support for legacy graph panel.
// @deprecated this a deprecated panel type
public class GraphPanel {
    @JsonProperty("type")
    public String type;
    // @deprecated this is part of deprecated graph panel
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("legend")
    public DashboardGraphPanelLegend legend;
    public GraphPanel() {
    }
    
    public GraphPanel(String type,DashboardGraphPanelLegend legend) {
        this.type = type;
        this.legend = legend;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
