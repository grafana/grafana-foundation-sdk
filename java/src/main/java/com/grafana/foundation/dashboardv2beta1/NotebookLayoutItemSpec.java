// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * One ordered item in a notebook layout. `element` references either a CellKind
 * (markdown/code content) or a V2PanelKind in the notebook's elements map. `source`
 * records who authored the cell; `collapsed` hides the body in the UI.
 */
public class NotebookLayoutItemSpec {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("element")
    public ElementReference element;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("source")
    public NotebookLayoutItemSpecSource source;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("collapsed")
    public Boolean collapsed;
    public NotebookLayoutItemSpec() {
        this.element = new com.grafana.foundation.dashboardv2beta1.ElementReference();
        this.source = NotebookLayoutItemSpecSource.ASSISTANT;
    }
    public NotebookLayoutItemSpec(ElementReference element,NotebookLayoutItemSpecSource source,Boolean collapsed) {
        this.element = element;
        this.source = source;
        this.collapsed = collapsed;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
