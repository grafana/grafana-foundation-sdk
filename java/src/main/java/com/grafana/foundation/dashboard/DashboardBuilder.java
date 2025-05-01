// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;
import java.util.LinkedList;

public class DashboardBuilder implements com.grafana.foundation.cog.Builder<Dashboard> {
    protected final Dashboard internal;
    private Integer currentY;
    private Integer currentX;
    private Integer lastPanelHeight;
    
    public DashboardBuilder(String title) {
        this.internal = new Dashboard();
        this.internal.title = title;
    this.currentY = 0;
    this.currentX = 0;
    this.lastPanelHeight = 0;
    }
    public DashboardBuilder id(Long id) {
        this.internal.id = id;
        return this;
    }
    
    public DashboardBuilder uid(String uid) {
        this.internal.uid = uid;
        return this;
    }
    
    public DashboardBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public DashboardBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public DashboardBuilder revision(Long revision) {
        this.internal.revision = revision;
        return this;
    }
    
    public DashboardBuilder gnetId(String gnetId) {
        this.internal.gnetId = gnetId;
        return this;
    }
    
    public DashboardBuilder tags(List<String> tags) {
        this.internal.tags = tags;
        return this;
    }
    
    public DashboardBuilder timezone(String timezone) {
        this.internal.timezone = timezone;
        return this;
    }
    
    public DashboardBuilder editable() {
        this.internal.editable = true;
        return this;
    }
    
    public DashboardBuilder readonly() {
        this.internal.editable = false;
        return this;
    }
    
    public DashboardBuilder tooltip(DashboardCursorSync graphTooltip) {
        this.internal.graphTooltip = graphTooltip;
        return this;
    }
    
    public DashboardBuilder time(com.grafana.foundation.cog.Builder<DashboardDashboardTime> time) {
    DashboardDashboardTime timeResource = time.build();
        this.internal.time = timeResource;
        return this;
    }
    
    public DashboardBuilder timepicker(com.grafana.foundation.cog.Builder<TimePickerConfig> timepicker) {
    TimePickerConfig timepickerResource = timepicker.build();
        this.internal.timepicker = timepickerResource;
        return this;
    }
    
    public DashboardBuilder fiscalYearStartMonth(Integer fiscalYearStartMonth) {
        if (!(fiscalYearStartMonth < 12)) {
            throw new IllegalArgumentException("fiscalYearStartMonth must be < 12");
        }
        this.internal.fiscalYearStartMonth = fiscalYearStartMonth;
        return this;
    }
    
    public DashboardBuilder liveNow(Boolean liveNow) {
        this.internal.liveNow = liveNow;
        return this;
    }
    
    public DashboardBuilder weekStart(String weekStart) {
        this.internal.weekStart = weekStart;
        return this;
    }
    
    public DashboardBuilder refresh(String refresh) {
        this.internal.refresh = refresh;
        return this;
    }
    
    public DashboardBuilder version(Integer version) {
        this.internal.version = version;
        return this;
    }
    
    public DashboardBuilder withPanel(com.grafana.foundation.cog.Builder<Panel> panel) {
		if (this.internal.panels == null) {
			this.internal.panels = new LinkedList<>();
		}
    Panel panelResource = panel.build();
    PanelOrRowPanel panelOrRowPanel = new PanelOrRowPanel();
        panelOrRowPanel.panel = panelResource;

    if (panelOrRowPanel.panel.gridPos == null) {
        panelOrRowPanel.panel.gridPos = new GridPos();
    }
    if (panelOrRowPanel.panel.gridPos.x == null) {
        panelOrRowPanel.panel.gridPos.x = 0;
    }
    if (panelOrRowPanel.panel.gridPos.y == null) {
        panelOrRowPanel.panel.gridPos.y = 0;
    }
    if (panelOrRowPanel.panel.gridPos.w == null) {
        panelOrRowPanel.panel.gridPos.w = 0;
    }
    if (panelOrRowPanel.panel.gridPos.h == null) {
        panelOrRowPanel.panel.gridPos.h = 0;
    }
    // The panel either has no position set, or it is the first panel of the dashboard.
    // In that case, we position it on the grid
    if (panelOrRowPanel.panel.gridPos.x == 0 && panelOrRowPanel.panel.gridPos.y == 0) {
        panelOrRowPanel.panel.gridPos.x = this.currentX;
        panelOrRowPanel.panel.gridPos.y = this.currentY;
    }
        this.internal.panels.add(panelOrRowPanel);

	// Prepare the coordinates for the next panel
	this.currentX += panelOrRowPanel.panel.gridPos.w;
	this.lastPanelHeight = java.lang.Math.max(this.lastPanelHeight, panelOrRowPanel.panel.gridPos.h);

	// Check for grid width overflow?
	if (this.currentX >= 24) {
		this.currentX = 0;
		this.currentY += this.lastPanelHeight;
		this.lastPanelHeight = 0;
	}
        return this;
    }
    
    public DashboardBuilder withRow(com.grafana.foundation.cog.Builder<RowPanel> rowPanel) {
		if (this.internal.panels == null) {
			this.internal.panels = new LinkedList<>();
		}
    RowPanel rowPanelResource = rowPanel.build();
    PanelOrRowPanel panelOrRowPanel = new PanelOrRowPanel();
        panelOrRowPanel.rowPanel = rowPanelResource;

    // Position the row on the grid
    if (panelOrRowPanel.rowPanel.gridPos == null || (panelOrRowPanel.rowPanel.gridPos.x == 0 && panelOrRowPanel.rowPanel.gridPos.y == 0)) {
        GridPos gridPos = new GridPos();
        gridPos.x = 0; // beginning of the line
        gridPos.y = this.currentY;
        gridPos.h = 1;
        gridPos.w = 24; // full width
        panelOrRowPanel.rowPanel.gridPos = gridPos;
    }
        this.internal.panels.add(panelOrRowPanel);

    // Reset the state for the next row
	this.currentX = 0;
	this.currentY += panelOrRowPanel.rowPanel.gridPos.h;
	this.lastPanelHeight = 0;
        return this;
    }
    
    public DashboardBuilder variables(List<com.grafana.foundation.cog.Builder<VariableModel>> variables) {
		if (this.internal.templating == null) {
			this.internal.templating = new com.grafana.foundation.dashboard.DashboardDashboardTemplatingBuilder().build();
		}
        List<VariableModel> variablesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<VariableModel> r1 : variables) {
                VariableModel variablesDepth1 = r1.build();
                variablesResources.add(variablesDepth1); 
        }
        this.internal.templating.list = variablesResources;
        return this;
    }
    
    public DashboardBuilder withVariable(com.grafana.foundation.cog.Builder<VariableModel> variable) {
		if (this.internal.templating == null) {
			this.internal.templating = new com.grafana.foundation.dashboard.DashboardDashboardTemplatingBuilder().build();
		}
		if (this.internal.templating.list == null) {
			this.internal.templating.list = new LinkedList<>();
		}
    VariableModel variableResource = variable.build();
        this.internal.templating.list.add(variableResource);
        return this;
    }
    
    public DashboardBuilder annotations(List<com.grafana.foundation.cog.Builder<AnnotationQuery>> annotations) {
		if (this.internal.annotations == null) {
			this.internal.annotations = new com.grafana.foundation.dashboard.AnnotationContainer();
		}
        List<AnnotationQuery> annotationsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<AnnotationQuery> r1 : annotations) {
                AnnotationQuery annotationsDepth1 = r1.build();
                annotationsResources.add(annotationsDepth1); 
        }
        this.internal.annotations.list = annotationsResources;
        return this;
    }
    
    public DashboardBuilder annotation(com.grafana.foundation.cog.Builder<AnnotationQuery> annotation) {
		if (this.internal.annotations == null) {
			this.internal.annotations = new com.grafana.foundation.dashboard.AnnotationContainer();
		}
		if (this.internal.annotations.list == null) {
			this.internal.annotations.list = new LinkedList<>();
		}
    AnnotationQuery annotationResource = annotation.build();
        this.internal.annotations.list.add(annotationResource);
        return this;
    }
    
    public DashboardBuilder links(List<com.grafana.foundation.cog.Builder<DashboardLink>> links) {
        List<DashboardLink> linksResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<DashboardLink> r1 : links) {
                DashboardLink linksDepth1 = r1.build();
                linksResources.add(linksDepth1); 
        }
        this.internal.links = linksResources;
        return this;
    }
    
    public DashboardBuilder link(com.grafana.foundation.cog.Builder<DashboardLink> link) {
		if (this.internal.links == null) {
			this.internal.links = new LinkedList<>();
		}
    DashboardLink linkResource = link.build();
        this.internal.links.add(linkResource);
        return this;
    }
    
    public DashboardBuilder snapshot(com.grafana.foundation.cog.Builder<Snapshot> snapshot) {
    Snapshot snapshotResource = snapshot.build();
        this.internal.snapshot = snapshotResource;
        return this;
    }
    
    public DashboardBuilder preload(Boolean preload) {
        this.internal.preload = preload;
        return this;
    }
    public Dashboard build() {
        return this.internal;
    }
}
