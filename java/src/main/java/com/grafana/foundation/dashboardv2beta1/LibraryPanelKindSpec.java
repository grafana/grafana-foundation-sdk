// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class LibraryPanelKindSpec {
    // Panel ID for the library panel in the dashboard
    @JsonProperty("id")
    public Double id;
    // Title for the library panel in the dashboard
    @JsonProperty("title")
    public String title;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("libraryPanel")
    public LibraryPanelRef libraryPanel;
    public LibraryPanelKindSpec() {
        this.id = 0.0;
        this.title = "";
        this.libraryPanel = new com.grafana.foundation.dashboardv2beta1.LibraryPanelRef();
    }
    public LibraryPanelKindSpec(Double id,String title,LibraryPanelRef libraryPanel) {
        this.id = id;
        this.title = title;
        this.libraryPanel = libraryPanel;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
