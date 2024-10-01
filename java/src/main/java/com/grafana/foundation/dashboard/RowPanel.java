// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import java.util.LinkedList;

// Row panel
public class RowPanel {
    // The panel type
    @JsonProperty("type")
    public String type;
    // Whether this row should be collapsed or not.
    @JsonProperty("collapsed")
    public Boolean collapsed;
    // Row title
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("title")
    public String title;
    // Name of default datasource for the row
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // Row grid position
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("gridPos")
    public GridPos gridPos;
    // Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
    @JsonProperty("id")
    public Integer id;
    // List of panels in the row
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("panels")
    public List<Panel> panels;
    // Name of template variable to repeat for.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("repeat")
    public String repeat;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<RowPanel> {
        protected final RowPanel internal;
        
        public Builder(String title) {
            this.internal = new RowPanel();
    this.internal.type = "row";
    this.internal.title = title;
        this.collapsed(false);
        }
    public Builder collapsed(Boolean collapsed) {
    this.internal.collapsed = collapsed;
        return this;
    }
    
    public Builder title(String title) {
    this.internal.title = title;
        return this;
    }
    
    public Builder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public Builder gridPos(GridPos gridPos) {
    this.internal.gridPos = gridPos;
        return this;
    }
    
    public Builder id(Integer id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder withPanel(com.grafana.foundation.cog.Builder<Panel> panels) {
		if (this.internal.panels == null) {
			this.internal.panels = new LinkedList<>();
		}
    this.internal.panels.add(panels.build());
        return this;
    }
    
    public Builder repeat(String repeat) {
    this.internal.repeat = repeat;
        return this;
    }
    public RowPanel build() {
            return this.internal;
        }
    }
}
