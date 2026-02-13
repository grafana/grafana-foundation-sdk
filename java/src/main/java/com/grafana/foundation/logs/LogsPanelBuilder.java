// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.logs;

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
import com.grafana.foundation.common.LogsSortOrder;
import com.grafana.foundation.common.LogsDedupStrategy;

public class LogsPanelBuilder implements com.grafana.foundation.cog.Builder<Panel> {
    protected final Panel internal;
    
    public LogsPanelBuilder() {
        this.internal = new Panel();
        this.internal.type = "logs";
    }
    public LogsPanelBuilder id(Integer id) {
        this.internal.id = id;
        return this;
    }
    
    public LogsPanelBuilder targets(List<com.grafana.foundation.cog.Builder<Dataquery>> targets) {
        List<Dataquery> targetsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Dataquery> r1 : targets) {
                Dataquery targetsDepth1 = r1.build();
                targetsResources.add(targetsDepth1); 
        }
        this.internal.targets = targetsResources;
        return this;
    }
    
    public LogsPanelBuilder withTarget(com.grafana.foundation.cog.Builder<Dataquery> target) {
		if (this.internal.targets == null) {
			this.internal.targets = new LinkedList<>();
		}
    Dataquery targetResource = target.build();
        this.internal.targets.add(targetResource);
        return this;
    }
    
    public LogsPanelBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public LogsPanelBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public LogsPanelBuilder transparent(Boolean transparent) {
        this.internal.transparent = transparent;
        return this;
    }
    
    public LogsPanelBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public LogsPanelBuilder gridPos(GridPos gridPos) {
        this.internal.gridPos = gridPos;
        return this;
    }
    
    public LogsPanelBuilder height(Integer h) {
        if (!(h > 0)) {
            throw new IllegalArgumentException("h must be > 0");
        }
		if (this.internal.gridPos == null) {
			this.internal.gridPos = new com.grafana.foundation.dashboard.GridPos();
		}
        this.internal.gridPos.h = h;
        return this;
    }
    
    public LogsPanelBuilder span(Integer w) {
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
    
    public LogsPanelBuilder links(List<com.grafana.foundation.cog.Builder<DashboardLink>> links) {
        List<DashboardLink> linksResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<DashboardLink> r1 : links) {
                DashboardLink linksDepth1 = r1.build();
                linksResources.add(linksDepth1); 
        }
        this.internal.links = linksResources;
        return this;
    }
    
    public LogsPanelBuilder repeat(String repeat) {
        this.internal.repeat = repeat;
        return this;
    }
    
    public LogsPanelBuilder repeatDirection(PanelRepeatDirection repeatDirection) {
        this.internal.repeatDirection = repeatDirection;
        return this;
    }
    
    public LogsPanelBuilder maxPerRow(Double maxPerRow) {
        this.internal.maxPerRow = maxPerRow;
        return this;
    }
    
    public LogsPanelBuilder maxDataPoints(Double maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public LogsPanelBuilder transformations(List<DataTransformerConfig> transformations) {
        this.internal.transformations = transformations;
        return this;
    }
    
    public LogsPanelBuilder withTransformation(DataTransformerConfig transformation) {
		if (this.internal.transformations == null) {
			this.internal.transformations = new LinkedList<>();
		}
        this.internal.transformations.add(transformation);
        return this;
    }
    
    public LogsPanelBuilder interval(String interval) {
        this.internal.interval = interval;
        return this;
    }
    
    public LogsPanelBuilder timeFrom(String timeFrom) {
        this.internal.timeFrom = timeFrom;
        return this;
    }
    
    public LogsPanelBuilder timeShift(String timeShift) {
        this.internal.timeShift = timeShift;
        return this;
    }
    
    public LogsPanelBuilder hideTimeOverride(Boolean hideTimeOverride) {
        this.internal.hideTimeOverride = hideTimeOverride;
        return this;
    }
    
    public LogsPanelBuilder libraryPanel(LibraryPanelRef libraryPanel) {
        this.internal.libraryPanel = libraryPanel;
        return this;
    }
    
    public LogsPanelBuilder cacheTimeout(String cacheTimeout) {
        this.internal.cacheTimeout = cacheTimeout;
        return this;
    }
    
    public LogsPanelBuilder queryCachingTTL(Double queryCachingTTL) {
        this.internal.queryCachingTTL = queryCachingTTL;
        return this;
    }
    
    public LogsPanelBuilder displayName(String displayName) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.displayName = displayName;
        return this;
    }
    
    public LogsPanelBuilder unit(String unit) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.unit = unit;
        return this;
    }
    
    public LogsPanelBuilder decimals(Double decimals) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.decimals = decimals;
        return this;
    }
    
    public LogsPanelBuilder min(Double min) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.min = min;
        return this;
    }
    
    public LogsPanelBuilder max(Double max) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.max = max;
        return this;
    }
    
    public LogsPanelBuilder mappings(List<ValueMapping> mappings) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.mappings = mappings;
        return this;
    }
    
    public LogsPanelBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds) {
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
    
    public LogsPanelBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
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
    
    public LogsPanelBuilder dataLinks(List<com.grafana.foundation.cog.Builder<DashboardLink>> links) {
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
    
    public LogsPanelBuilder noValue(String noValue) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.noValue = noValue;
        return this;
    }
    
    public LogsPanelBuilder fieldMinMax(Boolean fieldMinMax) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.fieldMinMax = fieldMinMax;
        return this;
    }
    
    public LogsPanelBuilder overrides(List<com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides>> overrides) {
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
    
    public LogsPanelBuilder withOverride(com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides> override) {
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
    
    public LogsPanelBuilder overrideByName(String name,List<DynamicConfigValue> properties) {
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
    
    public LogsPanelBuilder overrideByRegexp(String regexp,List<DynamicConfigValue> properties) {
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
    
    public LogsPanelBuilder overrideByFieldType(String fieldType,List<DynamicConfigValue> properties) {
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
    
    public LogsPanelBuilder overrideByQuery(String queryRefId,List<DynamicConfigValue> properties) {
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
    
    public LogsPanelBuilder showLabels(Boolean showLabels) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).showLabels = showLabels;
        return this;
    }
    
    public LogsPanelBuilder showCommonLabels(Boolean showCommonLabels) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).showCommonLabels = showCommonLabels;
        return this;
    }
    
    public LogsPanelBuilder showTime(Boolean showTime) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).showTime = showTime;
        return this;
    }
    
    public LogsPanelBuilder showLogContextToggle(Boolean showLogContextToggle) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).showLogContextToggle = showLogContextToggle;
        return this;
    }
    
    public LogsPanelBuilder wrapLogMessage(Boolean wrapLogMessage) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).wrapLogMessage = wrapLogMessage;
        return this;
    }
    
    public LogsPanelBuilder prettifyLogMessage(Boolean prettifyLogMessage) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).prettifyLogMessage = prettifyLogMessage;
        return this;
    }
    
    public LogsPanelBuilder enableLogDetails(Boolean enableLogDetails) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).enableLogDetails = enableLogDetails;
        return this;
    }
    
    public LogsPanelBuilder sortOrder(LogsSortOrder sortOrder) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).sortOrder = sortOrder;
        return this;
    }
    
    public LogsPanelBuilder dedupStrategy(LogsDedupStrategy dedupStrategy) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).dedupStrategy = dedupStrategy;
        return this;
    }
    
    public LogsPanelBuilder enableInfiniteScrolling(Boolean enableInfiniteScrolling) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).enableInfiniteScrolling = enableInfiniteScrolling;
        return this;
    }
    
    public LogsPanelBuilder onClickFilterLabel(Object onClickFilterLabel) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).onClickFilterLabel = onClickFilterLabel;
        return this;
    }
    
    public LogsPanelBuilder onClickFilterOutLabel(Object onClickFilterOutLabel) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).onClickFilterOutLabel = onClickFilterOutLabel;
        return this;
    }
    
    public LogsPanelBuilder isFilterLabelActive(Object isFilterLabelActive) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).isFilterLabelActive = isFilterLabelActive;
        return this;
    }
    
    public LogsPanelBuilder onClickFilterString(Object onClickFilterString) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).onClickFilterString = onClickFilterString;
        return this;
    }
    
    public LogsPanelBuilder onClickFilterOutString(Object onClickFilterOutString) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).onClickFilterOutString = onClickFilterOutString;
        return this;
    }
    
    public LogsPanelBuilder onClickShowField(Object onClickShowField) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).onClickShowField = onClickShowField;
        return this;
    }
    
    public LogsPanelBuilder onClickHideField(Object onClickHideField) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).onClickHideField = onClickHideField;
        return this;
    }
    
    public LogsPanelBuilder logRowMenuIconsBefore(Object logRowMenuIconsBefore) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).logRowMenuIconsBefore = logRowMenuIconsBefore;
        return this;
    }
    
    public LogsPanelBuilder logRowMenuIconsAfter(Object logRowMenuIconsAfter) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).logRowMenuIconsAfter = logRowMenuIconsAfter;
        return this;
    }
    
    public LogsPanelBuilder onNewLogsReceived(Object onNewLogsReceived) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).onNewLogsReceived = onNewLogsReceived;
        return this;
    }
    
    public LogsPanelBuilder displayedFields(List<String> displayedFields) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.logs.Options();
		}
        ((Options) this.internal.options).displayedFields = displayedFields;
        return this;
    }
    public Panel build() {
        return this.internal;
    }
}
