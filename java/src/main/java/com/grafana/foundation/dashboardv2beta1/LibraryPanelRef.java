// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// A library panel is a reusable panel that you can use in any dashboard.
// When you make a change to a library panel, that change propagates to all instances of where the panel is used.
// Library panels streamline reuse of panels across multiple dashboards.
public class LibraryPanelRef {
    // Library panel name
    @JsonProperty("name")
    public String name;
    // Library panel uid
    @JsonProperty("uid")
    public String uid;
    public LibraryPanelRef() {
        this.name = "";
        this.uid = "";
    }
    public LibraryPanelRef(String name,String uid) {
        this.name = name;
        this.uid = uid;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
