// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = PanelOrRowPanelDeserializer.class)
public class PanelOrRowPanel {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Panel panel;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected RowPanel rowPanel;
    protected PanelOrRowPanel() {}
    public static PanelOrRowPanel createPanel(Panel panel) {
        PanelOrRowPanel panelOrRowPanel = new PanelOrRowPanel();
        panelOrRowPanel.panel = panel;
        return panelOrRowPanel;
    }
    public static PanelOrRowPanel createRowPanel(RowPanel rowPanel) {
        PanelOrRowPanel panelOrRowPanel = new PanelOrRowPanel();
        panelOrRowPanel.rowPanel = rowPanel;
        return panelOrRowPanel;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
