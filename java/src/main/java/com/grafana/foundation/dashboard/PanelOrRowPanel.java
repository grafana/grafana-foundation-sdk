// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = PanelOrRowPanelDeserializer.class)
public class PanelOrRowPanel { 
    @JsonUnwrapped
    public Panel panel; 
    @JsonUnwrapped
    public RowPanel rowPanel;
    
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
