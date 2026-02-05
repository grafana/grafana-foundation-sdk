// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.table;

import com.grafana.foundation.dashboard.Panel;
import java.util.List;
import com.grafana.foundation.cog.variants.Dataquery;
import java.util.LinkedList;
import com.grafana.foundation.common.DataSourceRef;
import com.grafana.foundation.dashboard.GridPos;
import com.grafana.foundation.dashboard.DashboardLink;
import com.grafana.foundation.dashboard.PanelRepeatDirection;
import com.grafana.foundation.dashboard.DataTransformerConfig;
import com.grafana.foundation.dashboard.LibraryPanelRef;
import com.grafana.foundation.dashboard.ValueMapping;
import com.grafana.foundation.dashboard.ThresholdsConfig;
import com.grafana.foundation.dashboard.FieldColor;
import com.grafana.foundation.dashboard.DashboardFieldConfigSourceOverrides;
import com.grafana.foundation.dashboard.DynamicConfigValue;
import com.grafana.foundation.dashboard.MatcherConfig;
import com.grafana.foundation.common.TableSortByFieldState;
import com.grafana.foundation.common.TableFooterOptions;
import com.grafana.foundation.common.TableCellHeight;
import com.grafana.foundation.common.FieldTextAlignment;
import com.grafana.foundation.common.TableCellDisplayMode;
import com.grafana.foundation.common.TableCellOptions;

public class TablePanelBuilder implements com.grafana.foundation.cog.Builder<Panel> {
    protected final Panel internal;
    
    public TablePanelBuilder() {
        this.internal = new Panel();
        this.internal.type = "table";
    }
    public TablePanelBuilder id(Integer id) {
        this.internal.id = id;
        return this;
    }
    
    public TablePanelBuilder targets(List<com.grafana.foundation.cog.Builder<Dataquery>> targets) {
        List<Dataquery> targetsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Dataquery> r1 : targets) {
                Dataquery targetsDepth1 = r1.build();
                targetsResources.add(targetsDepth1); 
        }
        this.internal.targets = targetsResources;
        return this;
    }
    
    public TablePanelBuilder withTarget(com.grafana.foundation.cog.Builder<Dataquery> target) {
		if (this.internal.targets == null) {
			this.internal.targets = new LinkedList<>();
		}
    Dataquery targetResource = target.build();
        this.internal.targets.add(targetResource);
        return this;
    }
    
    public TablePanelBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public TablePanelBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public TablePanelBuilder transparent(Boolean transparent) {
        this.internal.transparent = transparent;
        return this;
    }
    
    public TablePanelBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public TablePanelBuilder gridPos(GridPos gridPos) {
        this.internal.gridPos = gridPos;
        return this;
    }
    
    public TablePanelBuilder height(Integer h) {
        if (!(h > 0)) {
            throw new IllegalArgumentException("h must be > 0");
        }
		if (this.internal.gridPos == null) {
			this.internal.gridPos = new com.grafana.foundation.dashboard.GridPos();
		}
        this.internal.gridPos.h = h;
        return this;
    }
    
    public TablePanelBuilder span(Integer w) {
        if (!(w > 0)) {
            throw new IllegalArgumentException("w must be > 0");
        }
        if (!(w <= 24)) {
            throw new IllegalArgumentException("w must be <= 24");
        }
		if (this.internal.gridPos == null) {
			this.internal.gridPos = new com.grafana.foundation.dashboard.GridPos();
		}
        this.internal.gridPos.w = w;
        return this;
    }
    
    public TablePanelBuilder links(List<com.grafana.foundation.cog.Builder<DashboardLink>> links) {
        List<DashboardLink> linksResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<DashboardLink> r1 : links) {
                DashboardLink linksDepth1 = r1.build();
                linksResources.add(linksDepth1); 
        }
        this.internal.links = linksResources;
        return this;
    }
    
    public TablePanelBuilder repeat(String repeat) {
        this.internal.repeat = repeat;
        return this;
    }
    
    public TablePanelBuilder repeatDirection(PanelRepeatDirection repeatDirection) {
        this.internal.repeatDirection = repeatDirection;
        return this;
    }
    
    public TablePanelBuilder maxPerRow(Double maxPerRow) {
        this.internal.maxPerRow = maxPerRow;
        return this;
    }
    
    public TablePanelBuilder maxDataPoints(Double maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public TablePanelBuilder transformations(List<DataTransformerConfig> transformations) {
        this.internal.transformations = transformations;
        return this;
    }
    
    public TablePanelBuilder withTransformation(DataTransformerConfig transformation) {
		if (this.internal.transformations == null) {
			this.internal.transformations = new LinkedList<>();
		}
        this.internal.transformations.add(transformation);
        return this;
    }
    
    public TablePanelBuilder interval(String interval) {
        this.internal.interval = interval;
        return this;
    }
    
    public TablePanelBuilder timeFrom(String timeFrom) {
        this.internal.timeFrom = timeFrom;
        return this;
    }
    
    public TablePanelBuilder timeShift(String timeShift) {
        this.internal.timeShift = timeShift;
        return this;
    }
    
    public TablePanelBuilder hideTimeOverride(Boolean hideTimeOverride) {
        this.internal.hideTimeOverride = hideTimeOverride;
        return this;
    }
    
    public TablePanelBuilder libraryPanel(LibraryPanelRef libraryPanel) {
        this.internal.libraryPanel = libraryPanel;
        return this;
    }
    
    public TablePanelBuilder cacheTimeout(String cacheTimeout) {
        this.internal.cacheTimeout = cacheTimeout;
        return this;
    }
    
    public TablePanelBuilder queryCachingTTL(Double queryCachingTTL) {
        this.internal.queryCachingTTL = queryCachingTTL;
        return this;
    }
    
    public TablePanelBuilder displayName(String displayName) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.displayName = displayName;
        return this;
    }
    
    public TablePanelBuilder unit(String unit) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.unit = unit;
        return this;
    }
    
    public TablePanelBuilder decimals(Double decimals) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.decimals = decimals;
        return this;
    }
    
    public TablePanelBuilder min(Double min) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.min = min;
        return this;
    }
    
    public TablePanelBuilder max(Double max) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.max = max;
        return this;
    }
    
    public TablePanelBuilder mappings(List<ValueMapping> mappings) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.mappings = mappings;
        return this;
    }
    
    public TablePanelBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    ThresholdsConfig thresholdsResource = thresholds.build();
        this.internal.fieldConfig.defaults.thresholds = thresholdsResource;
        return this;
    }
    
    public TablePanelBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    FieldColor colorResource = color.build();
        this.internal.fieldConfig.defaults.color = colorResource;
        return this;
    }
    
    public TablePanelBuilder dataLinks(List<com.grafana.foundation.cog.Builder<DashboardLink>> links) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        List<DashboardLink> linksResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<DashboardLink> r1 : links) {
                DashboardLink linksDepth1 = r1.build();
                linksResources.add(linksDepth1); 
        }
        this.internal.fieldConfig.defaults.links = linksResources;
        return this;
    }
    
    public TablePanelBuilder noValue(String noValue) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.noValue = noValue;
        return this;
    }
    
    public TablePanelBuilder overrides(List<com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides>> overrides) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
        List<DashboardFieldConfigSourceOverrides> overridesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides> r1 : overrides) {
                DashboardFieldConfigSourceOverrides overridesDepth1 = r1.build();
                overridesResources.add(overridesDepth1); 
        }
        this.internal.fieldConfig.overrides = overridesResources;
        return this;
    }
    
    public TablePanelBuilder withOverride(com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides> override) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.overrides == null) {
			this.internal.fieldConfig.overrides = new LinkedList<>();
		}
    DashboardFieldConfigSourceOverrides overrideResource = override.build();
        this.internal.fieldConfig.overrides.add(overrideResource);
        return this;
    }
    
    public TablePanelBuilder overrideByName(String name,List<DynamicConfigValue> properties) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.overrides == null) {
			this.internal.fieldConfig.overrides = new LinkedList<>();
		}
    MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byName";
        matcherConfig.options = name;
    DashboardFieldConfigSourceOverrides dashboardFieldConfigSourceOverrides = new DashboardFieldConfigSourceOverrides();
        dashboardFieldConfigSourceOverrides.matcher = matcherConfig;
        dashboardFieldConfigSourceOverrides.properties = properties;
        this.internal.fieldConfig.overrides.add(dashboardFieldConfigSourceOverrides);
        return this;
    }
    
    public TablePanelBuilder overrideByRegexp(String regexp,List<DynamicConfigValue> properties) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.overrides == null) {
			this.internal.fieldConfig.overrides = new LinkedList<>();
		}
    MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byRegexp";
        matcherConfig.options = regexp;
    DashboardFieldConfigSourceOverrides dashboardFieldConfigSourceOverrides = new DashboardFieldConfigSourceOverrides();
        dashboardFieldConfigSourceOverrides.matcher = matcherConfig;
        dashboardFieldConfigSourceOverrides.properties = properties;
        this.internal.fieldConfig.overrides.add(dashboardFieldConfigSourceOverrides);
        return this;
    }
    
    public TablePanelBuilder overrideByFieldType(String fieldType,List<DynamicConfigValue> properties) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.overrides == null) {
			this.internal.fieldConfig.overrides = new LinkedList<>();
		}
    MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byType";
        matcherConfig.options = fieldType;
    DashboardFieldConfigSourceOverrides dashboardFieldConfigSourceOverrides = new DashboardFieldConfigSourceOverrides();
        dashboardFieldConfigSourceOverrides.matcher = matcherConfig;
        dashboardFieldConfigSourceOverrides.properties = properties;
        this.internal.fieldConfig.overrides.add(dashboardFieldConfigSourceOverrides);
        return this;
    }
    
    public TablePanelBuilder overrideByQuery(String queryRefId,List<DynamicConfigValue> properties) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.overrides == null) {
			this.internal.fieldConfig.overrides = new LinkedList<>();
		}
    MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byFrameRefID";
        matcherConfig.options = queryRefId;
    DashboardFieldConfigSourceOverrides dashboardFieldConfigSourceOverrides = new DashboardFieldConfigSourceOverrides();
        dashboardFieldConfigSourceOverrides.matcher = matcherConfig;
        dashboardFieldConfigSourceOverrides.properties = properties;
        this.internal.fieldConfig.overrides.add(dashboardFieldConfigSourceOverrides);
        return this;
    }
    
    public TablePanelBuilder frameIndex(Double frameIndex) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.table.Options();
		}
        ((Options) this.internal.options).frameIndex = frameIndex;
        return this;
    }
    
    public TablePanelBuilder showHeader(Boolean showHeader) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.table.Options();
		}
        ((Options) this.internal.options).showHeader = showHeader;
        return this;
    }
    
    public TablePanelBuilder showTypeIcons(Boolean showTypeIcons) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.table.Options();
		}
        ((Options) this.internal.options).showTypeIcons = showTypeIcons;
        return this;
    }
    
    public TablePanelBuilder sortBy(List<com.grafana.foundation.cog.Builder<TableSortByFieldState>> sortBy) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.table.Options();
		}
        List<TableSortByFieldState> sortByResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TableSortByFieldState> r1 : sortBy) {
                TableSortByFieldState sortByDepth1 = r1.build();
                sortByResources.add(sortByDepth1); 
        }
        ((Options) this.internal.options).sortBy = sortByResources;
        return this;
    }
    
    public TablePanelBuilder footer(com.grafana.foundation.cog.Builder<TableFooterOptions> footer) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.table.Options();
		}
    TableFooterOptions footerResource = footer.build();
        ((Options) this.internal.options).footer = footerResource;
        return this;
    }
    
    public TablePanelBuilder cellHeight(TableCellHeight cellHeight) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.table.Options();
		}
        ((Options) this.internal.options).cellHeight = cellHeight;
        return this;
    }
    
    public TablePanelBuilder width(Double width) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.table.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).width = width;
        return this;
    }
    
    public TablePanelBuilder minWidth(Double minWidth) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.table.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).minWidth = minWidth;
        return this;
    }
    
    public TablePanelBuilder align(FieldTextAlignment align) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.table.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).align = align;
        return this;
    }
    
    public TablePanelBuilder displayMode(TableCellDisplayMode displayMode) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.table.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).displayMode = displayMode;
        return this;
    }
    
    public TablePanelBuilder cellOptions(TableCellOptions cellOptions) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.table.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).cellOptions = cellOptions;
        return this;
    }
    
    public TablePanelBuilder hidden(Boolean hidden) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.table.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).hidden = hidden;
        return this;
    }
    
    public TablePanelBuilder inspect(Boolean inspect) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.table.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).inspect = inspect;
        return this;
    }
    
    public TablePanelBuilder filterable(Boolean filterable) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.table.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).filterable = filterable;
        return this;
    }
    
    public TablePanelBuilder hideHeader(Boolean hideHeader) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.table.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).hideHeader = hideHeader;
        return this;
    }
    public Panel build() {
        return this.internal;
    }
}
