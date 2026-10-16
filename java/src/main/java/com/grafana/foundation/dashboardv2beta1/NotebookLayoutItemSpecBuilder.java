// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


/**
 * One ordered item in a notebook layout. `element` references either a CellKind
 * (markdown/code content) or a V2PanelKind in the notebook's elements map. `source`
 * records who authored the cell; `collapsed` hides the body in the UI.
 */
public class NotebookLayoutItemSpecBuilder implements com.grafana.foundation.cog.Builder<NotebookLayoutItemSpec> {
    protected final NotebookLayoutItemSpec internal;
    
    public NotebookLayoutItemSpecBuilder() {
        this.internal = new NotebookLayoutItemSpec();
    }
    public NotebookLayoutItemSpecBuilder element(com.grafana.foundation.cog.Builder<ElementReference> element) {
    ElementReference elementResource = element.build();
        this.internal.element = elementResource;
        return this;
    }
    
    public NotebookLayoutItemSpecBuilder source(NotebookLayoutItemSpecSource source) {
        this.internal.source = source;
        return this;
    }
    
    public NotebookLayoutItemSpecBuilder collapsed(Boolean collapsed) {
        this.internal.collapsed = collapsed;
        return this;
    }
    public NotebookLayoutItemSpec build() {
        return this.internal;
    }
}
