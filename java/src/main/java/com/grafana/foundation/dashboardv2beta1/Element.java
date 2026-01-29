// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

// Supported dashboard elements
// |* more element types in the future
@JsonDeserialize(using = ElementDeserializer.class)
public class Element {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected PanelKind panelKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected LibraryPanelKind libraryPanelKind;
    protected Element() {}
    public static Element createPanelKind(PanelKind panelKind) {
        Element element = new Element();
        element.panelKind = panelKind;
        return element;
    }
    public static Element createLibraryPanelKind(LibraryPanelKind libraryPanelKind) {
        Element element = new Element();
        element.libraryPanelKind = libraryPanelKind;
        return element;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
