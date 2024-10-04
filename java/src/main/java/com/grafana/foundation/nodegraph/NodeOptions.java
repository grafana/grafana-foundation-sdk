// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.nodegraph;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class NodeOptions {
    // Unit for the main stat to override what ever is set in the data frame.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mainStatUnit")
    public String mainStatUnit;
    // Unit for the secondary stat to override what ever is set in the data frame.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("secondaryStatUnit")
    public String secondaryStatUnit;
    // Define which fields are shown as part of the node arc (colored circle around the node).
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("arcs")
    public List<ArcOption> arcs;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<NodeOptions> {
        protected final NodeOptions internal;
        
        public Builder() {
            this.internal = new NodeOptions();
        }
    public Builder mainStatUnit(String mainStatUnit) {
    this.internal.mainStatUnit = mainStatUnit;
        return this;
    }
    
    public Builder secondaryStatUnit(String secondaryStatUnit) {
    this.internal.secondaryStatUnit = secondaryStatUnit;
        return this;
    }
    
    public Builder arcs(com.grafana.foundation.cog.Builder<List<ArcOption>> arcs) {
    this.internal.arcs = arcs.build();
        return this;
    }
    public NodeOptions build() {
            return this.internal;
        }
    }
}
