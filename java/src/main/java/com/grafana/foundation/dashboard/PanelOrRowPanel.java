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
    public static PanelOrRowPanel createPanel(com.grafana.foundation.cog.Builder<Panel> panel) {
        PanelOrRowPanel panelOrRowPanel = new PanelOrRowPanel();
        panelOrRowPanel.panel = panel.build();
        return panelOrRowPanel;
    }
    public static PanelOrRowPanel createRowPanel(com.grafana.foundation.cog.Builder<RowPanel> rowPanel) {
        PanelOrRowPanel panelOrRowPanel = new PanelOrRowPanel();
        panelOrRowPanel.rowPanel = rowPanel.build();
        return panelOrRowPanel;
    }
    
    public String toJSON() throws JsonProcessingException {
        if (panel != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(panel);
        }
        if (rowPanel != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(rowPanel);
        }
        
        return null;
    }

}
