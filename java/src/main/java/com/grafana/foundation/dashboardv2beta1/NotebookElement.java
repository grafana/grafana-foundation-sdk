// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * A notebook element is a narrative cell, a panel, or a library panel. Unlike the dashboard
 * Element union, this one includes CellKind — and it is referenced ONLY by NotebookSpec.
 * CellKind is listed first so it is the generated default (a notebook is narrative-first).
 */
@JsonDeserialize(using = NotebookElementDeserializer.class)
public class NotebookElement {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected CellKind cellKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected V2PanelKind v2PanelKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected LibraryPanelKind libraryPanelKind;
    protected NotebookElement() {}
    public static NotebookElement createCellKind(CellKind cellKind) {
        NotebookElement notebookElement = new NotebookElement();
        notebookElement.cellKind = cellKind;
        return notebookElement;
    }
    public static NotebookElement createV2PanelKind(V2PanelKind v2PanelKind) {
        NotebookElement notebookElement = new NotebookElement();
        notebookElement.v2PanelKind = v2PanelKind;
        return notebookElement;
    }
    public static NotebookElement createLibraryPanelKind(LibraryPanelKind libraryPanelKind) {
        NotebookElement notebookElement = new NotebookElement();
        notebookElement.libraryPanelKind = libraryPanelKind;
        return notebookElement;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
