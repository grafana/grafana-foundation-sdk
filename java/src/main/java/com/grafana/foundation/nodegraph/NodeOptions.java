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
    public NodeOptions() {
    }
    
    public NodeOptions(String mainStatUnit,String secondaryStatUnit,List<ArcOption> arcs) {
        this.mainStatUnit = mainStatUnit;
        this.secondaryStatUnit = secondaryStatUnit;
        this.arcs = arcs;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
