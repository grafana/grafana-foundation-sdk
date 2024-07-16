// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Support for legacy graph panel.
// @deprecated this a deprecated panel type
public class GraphPanel { 
    @JsonProperty("type")
    public String type;
    // @deprecated this is part of deprecated graph panel 
    @JsonProperty("legend")
    public DashboardGraphPanelLegend legend;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<GraphPanel> {
        private final GraphPanel internal;
        
        public Builder() {
            this.internal = new GraphPanel();
    this.internal.type = "graph";
        }
    public Builder legend(com.grafana.foundation.cog.Builder<DashboardGraphPanelLegend> legend) {
    this.internal.legend = legend.build();
        return this;
    }
    public GraphPanel build() {
            return this.internal;
        }
    }
}
