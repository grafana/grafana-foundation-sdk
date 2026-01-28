// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;
import java.util.Map;
import java.util.HashMap;

public class DashboardBuilder implements com.grafana.foundation.cog.Builder<Dashboard> {
    protected final Dashboard internal;
    
    public DashboardBuilder(String title) {
        this.internal = new Dashboard();
        this.internal.title = title;
    }
    public DashboardBuilder annotations(List<com.grafana.foundation.cog.Builder<AnnotationQueryKind>> annotations) {
        List<AnnotationQueryKind> annotationsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<AnnotationQueryKind> r1 : annotations) {
                AnnotationQueryKind annotationsDepth1 = r1.build();
                annotationsResources.add(annotationsDepth1); 
        }
        this.internal.annotations = annotationsResources;
        return this;
    }
    
    public DashboardBuilder cursorSync(DashboardCursorSync cursorSync) {
        this.internal.cursorSync = cursorSync;
        return this;
    }
    
    public DashboardBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public DashboardBuilder editable(Boolean editable) {
        this.internal.editable = editable;
        return this;
    }
    
    public DashboardBuilder elements(Map<String, Element> elements) {
        this.internal.elements = elements;
        return this;
    }
    
    public DashboardBuilder panel(String key,com.grafana.foundation.cog.Builder<PanelKind> panelKind) {
		if (this.internal.elements == null) {
			this.internal.elements = new HashMap<>();
		}
    PanelKind panelKindResource = panelKind.build();
    Element element = new Element();
        element.panelKind = panelKindResource;
        this.internal.elements.put(key, element);
        return this;
    }
    
    public DashboardBuilder libraryPanel(String key,com.grafana.foundation.cog.Builder<LibraryPanelKind> libraryPanelKind) {
		if (this.internal.elements == null) {
			this.internal.elements = new HashMap<>();
		}
    LibraryPanelKind libraryPanelKindResource = libraryPanelKind.build();
    Element element = new Element();
        element.libraryPanelKind = libraryPanelKindResource;
        this.internal.elements.put(key, element);
        return this;
    }
    
    public DashboardBuilder gridLayout(com.grafana.foundation.cog.Builder<GridLayoutKind> gridLayoutKind) {
    GridLayoutKind gridLayoutKindResource = gridLayoutKind.build();
    GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.gridLayoutKind = gridLayoutKindResource;
        this.internal.layout = gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
        return this;
    }
    
    public DashboardBuilder rowsLayout(com.grafana.foundation.cog.Builder<RowsLayoutKind> rowsLayoutKind) {
    RowsLayoutKind rowsLayoutKindResource = rowsLayoutKind.build();
    GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.rowsLayoutKind = rowsLayoutKindResource;
        this.internal.layout = gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
        return this;
    }
    
    public DashboardBuilder autoGridLayout(com.grafana.foundation.cog.Builder<AutoGridLayoutKind> autoGridLayoutKind) {
    AutoGridLayoutKind autoGridLayoutKindResource = autoGridLayoutKind.build();
    GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.autoGridLayoutKind = autoGridLayoutKindResource;
        this.internal.layout = gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
        return this;
    }
    
    public DashboardBuilder tabsLayout(com.grafana.foundation.cog.Builder<TabsLayoutKind> tabsLayoutKind) {
    TabsLayoutKind tabsLayoutKindResource = tabsLayoutKind.build();
    GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.tabsLayoutKind = tabsLayoutKindResource;
        this.internal.layout = gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
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
    
    public DashboardBuilder liveNow(Boolean liveNow) {
        this.internal.liveNow = liveNow;
        return this;
    }
    
    public DashboardBuilder preload(Boolean preload) {
        this.internal.preload = preload;
        return this;
    }
    
    public DashboardBuilder revision(Short revision) {
        this.internal.revision = revision;
        return this;
    }
    
    public DashboardBuilder tags(List<String> tags) {
        this.internal.tags = tags;
        return this;
    }
    
    public DashboardBuilder timeSettings(com.grafana.foundation.cog.Builder<TimeSettingsSpec> timeSettings) {
    TimeSettingsSpec timeSettingsResource = timeSettings.build();
        this.internal.timeSettings = timeSettingsResource;
        return this;
    }
    
    public DashboardBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public DashboardBuilder variables(List<VariableKind> variables) {
        this.internal.variables = variables;
        return this;
    }
    
    public DashboardBuilder queryVariable(com.grafana.foundation.cog.Builder<QueryVariableKind> queryVariableKind) {
		if (this.internal.variables == null) {
			this.internal.variables = new LinkedList<>();
		}
    QueryVariableKind queryVariableKindResource = queryVariableKind.build();
    VariableKind variableKind = new VariableKind();
        variableKind.queryVariableKind = queryVariableKindResource;
        this.internal.variables.add(variableKind);
        return this;
    }
    
    public DashboardBuilder textVariable(com.grafana.foundation.cog.Builder<TextVariableKind> textVariableKind) {
		if (this.internal.variables == null) {
			this.internal.variables = new LinkedList<>();
		}
    TextVariableKind textVariableKindResource = textVariableKind.build();
    VariableKind variableKind = new VariableKind();
        variableKind.textVariableKind = textVariableKindResource;
        this.internal.variables.add(variableKind);
        return this;
    }
    
    public DashboardBuilder constantVariable(com.grafana.foundation.cog.Builder<ConstantVariableKind> constantVariableKind) {
		if (this.internal.variables == null) {
			this.internal.variables = new LinkedList<>();
		}
    ConstantVariableKind constantVariableKindResource = constantVariableKind.build();
    VariableKind variableKind = new VariableKind();
        variableKind.constantVariableKind = constantVariableKindResource;
        this.internal.variables.add(variableKind);
        return this;
    }
    
    public DashboardBuilder datasourceVariable(com.grafana.foundation.cog.Builder<DatasourceVariableKind> datasourceVariableKind) {
		if (this.internal.variables == null) {
			this.internal.variables = new LinkedList<>();
		}
    DatasourceVariableKind datasourceVariableKindResource = datasourceVariableKind.build();
    VariableKind variableKind = new VariableKind();
        variableKind.datasourceVariableKind = datasourceVariableKindResource;
        this.internal.variables.add(variableKind);
        return this;
    }
    
    public DashboardBuilder intervalVariable(com.grafana.foundation.cog.Builder<IntervalVariableKind> intervalVariableKind) {
		if (this.internal.variables == null) {
			this.internal.variables = new LinkedList<>();
		}
    IntervalVariableKind intervalVariableKindResource = intervalVariableKind.build();
    VariableKind variableKind = new VariableKind();
        variableKind.intervalVariableKind = intervalVariableKindResource;
        this.internal.variables.add(variableKind);
        return this;
    }
    
    public DashboardBuilder customVariable(com.grafana.foundation.cog.Builder<CustomVariableKind> customVariableKind) {
		if (this.internal.variables == null) {
			this.internal.variables = new LinkedList<>();
		}
    CustomVariableKind customVariableKindResource = customVariableKind.build();
    VariableKind variableKind = new VariableKind();
        variableKind.customVariableKind = customVariableKindResource;
        this.internal.variables.add(variableKind);
        return this;
    }
    
    public DashboardBuilder groupByVariable(com.grafana.foundation.cog.Builder<GroupByVariableKind> groupByVariableKind) {
		if (this.internal.variables == null) {
			this.internal.variables = new LinkedList<>();
		}
    GroupByVariableKind groupByVariableKindResource = groupByVariableKind.build();
    VariableKind variableKind = new VariableKind();
        variableKind.groupByVariableKind = groupByVariableKindResource;
        this.internal.variables.add(variableKind);
        return this;
    }
    
    public DashboardBuilder adhocVariable(com.grafana.foundation.cog.Builder<AdhocVariableKind> adhocVariableKind) {
		if (this.internal.variables == null) {
			this.internal.variables = new LinkedList<>();
		}
    AdhocVariableKind adhocVariableKindResource = adhocVariableKind.build();
    VariableKind variableKind = new VariableKind();
        variableKind.adhocVariableKind = adhocVariableKindResource;
        this.internal.variables.add(variableKind);
        return this;
    }
    
    public DashboardBuilder switchVariableKind(com.grafana.foundation.cog.Builder<SwitchVariableKind> switchVariableKind) {
		if (this.internal.variables == null) {
			this.internal.variables = new LinkedList<>();
		}
    SwitchVariableKind switchVariableKindResource = switchVariableKind.build();
    VariableKind variableKind = new VariableKind();
        variableKind.switchVariableKind = switchVariableKindResource;
        this.internal.variables.add(variableKind);
        return this;
    }
    public Dashboard build() {
        return this.internal;
    }
}
