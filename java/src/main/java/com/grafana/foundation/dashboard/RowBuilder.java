// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.LinkedList;

public class RowBuilder implements com.grafana.foundation.cog.Builder<RowPanel> {
    protected final RowPanel internal;
    
    public RowBuilder(String title) {
        this.internal = new RowPanel();
    this.internal.type = "row";
    this.internal.title = title;
    }
    public RowBuilder collapsed(Boolean collapsed) {
    this.internal.collapsed = collapsed;
        return this;
    }
    
    public RowBuilder title(String title) {
    this.internal.title = title;
        return this;
    }
    
    public RowBuilder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public RowBuilder gridPos(GridPos gridPos) {
    this.internal.gridPos = gridPos;
        return this;
    }
    
    public RowBuilder id(Integer id) {
    this.internal.id = id;
        return this;
    }
    
    public RowBuilder withPanel(com.grafana.foundation.cog.Builder<Panel> panel) {
		if (this.internal.panels == null) {
			this.internal.panels = new LinkedList<>();
		}
    this.internal.panels.add(panel.build());
        return this;
    }
    
    public RowBuilder repeat(String repeat) {
    this.internal.repeat = repeat;
        return this;
    }
    public RowPanel build() {
        return this.internal;
    }
}
