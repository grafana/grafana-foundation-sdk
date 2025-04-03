// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bargauge;

import com.grafana.foundation.dashboard.Panel;
import java.util.List;
import com.grafana.foundation.cog.variants.Dataquery;
import java.util.LinkedList;
import com.grafana.foundation.dashboard.DataSourceRef;
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
import com.grafana.foundation.common.BarGaugeDisplayMode;
import com.grafana.foundation.common.BarGaugeValueMode;
import com.grafana.foundation.common.BarGaugeNamePlacement;
import com.grafana.foundation.common.BarGaugeSizing;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;

public class PanelBuilder implements com.grafana.foundation.cog.Builder<Panel> {
    protected final Panel internal;
    
    public PanelBuilder() {
        this.internal = new Panel();
        this.internal.type = "bargauge";
    }
    public PanelBuilder id(Integer id) {
        this.internal.id = id;
        return this;
    }
    
    public PanelBuilder targets(com.grafana.foundation.cog.Builder<List<Dataquery>> targets) {
        this.internal.targets = targets.build();
        return this;
    }
    
    public PanelBuilder withTarget(com.grafana.foundation.cog.Builder<Dataquery> target) {
		if (this.internal.targets == null) {
			this.internal.targets = new LinkedList<>();
		}
        this.internal.targets.add(target.build());
        return this;
    }
    
    public PanelBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public PanelBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public PanelBuilder transparent(Boolean transparent) {
        this.internal.transparent = transparent;
        return this;
    }
    
    public PanelBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public PanelBuilder gridPos(GridPos gridPos) {
        this.internal.gridPos = gridPos;
        return this;
    }
    
    public PanelBuilder height(Integer h) {
        if (!(h > 0)) {
            throw new IllegalArgumentException("h must be > 0");
        }
		if (this.internal.gridPos == null) {
			this.internal.gridPos = new com.grafana.foundation.dashboard.GridPos();
		}
        this.internal.gridPos.h = h;
        return this;
    }
    
    public PanelBuilder span(Integer w) {
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
    
    public PanelBuilder links(com.grafana.foundation.cog.Builder<List<DashboardLink>> links) {
        this.internal.links = links.build();
        return this;
    }
    
    public PanelBuilder repeat(String repeat) {
        this.internal.repeat = repeat;
        return this;
    }
    
    public PanelBuilder repeatDirection(PanelRepeatDirection repeatDirection) {
        this.internal.repeatDirection = repeatDirection;
        return this;
    }
    
    public PanelBuilder maxPerRow(Double maxPerRow) {
        this.internal.maxPerRow = maxPerRow;
        return this;
    }
    
    public PanelBuilder maxDataPoints(Double maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public PanelBuilder transformations(List<DataTransformerConfig> transformations) {
        this.internal.transformations = transformations;
        return this;
    }
    
    public PanelBuilder withTransformation(DataTransformerConfig transformation) {
		if (this.internal.transformations == null) {
			this.internal.transformations = new LinkedList<>();
		}
        this.internal.transformations.add(transformation);
        return this;
    }
    
    public PanelBuilder interval(String interval) {
        this.internal.interval = interval;
        return this;
    }
    
    public PanelBuilder timeFrom(String timeFrom) {
        this.internal.timeFrom = timeFrom;
        return this;
    }
    
    public PanelBuilder timeShift(String timeShift) {
        this.internal.timeShift = timeShift;
        return this;
    }
    
    public PanelBuilder hideTimeOverride(Boolean hideTimeOverride) {
        this.internal.hideTimeOverride = hideTimeOverride;
        return this;
    }
    
    public PanelBuilder libraryPanel(LibraryPanelRef libraryPanel) {
        this.internal.libraryPanel = libraryPanel;
        return this;
    }
    
    public PanelBuilder cacheTimeout(String cacheTimeout) {
        this.internal.cacheTimeout = cacheTimeout;
        return this;
    }
    
    public PanelBuilder queryCachingTTL(Double queryCachingTTL) {
        this.internal.queryCachingTTL = queryCachingTTL;
        return this;
    }
    
    public PanelBuilder displayName(String displayName) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.displayName = displayName;
        return this;
    }
    
    public PanelBuilder unit(String unit) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.unit = unit;
        return this;
    }
    
    public PanelBuilder decimals(Double decimals) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.decimals = decimals;
        return this;
    }
    
    public PanelBuilder min(Double min) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.min = min;
        return this;
    }
    
    public PanelBuilder max(Double max) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.max = max;
        return this;
    }
    
    public PanelBuilder mappings(List<ValueMapping> mappings) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.mappings = mappings;
        return this;
    }
    
    public PanelBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.thresholds = thresholds.build();
        return this;
    }
    
    public PanelBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.color = color.build();
        return this;
    }
    
    public PanelBuilder dataLinks(com.grafana.foundation.cog.Builder<List<DashboardLink>> links) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.links = links.build();
        return this;
    }
    
    public PanelBuilder noValue(String noValue) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.noValue = noValue;
        return this;
    }
    
    public PanelBuilder overrides(com.grafana.foundation.cog.Builder<List<DashboardFieldConfigSourceOverrides>> overrides) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
        this.internal.fieldConfig.overrides = overrides.build();
        return this;
    }
    
    public PanelBuilder withOverride(com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides> override) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.overrides == null) {
			this.internal.fieldConfig.overrides = new LinkedList<>();
		}
        this.internal.fieldConfig.overrides.add(override.build());
        return this;
    }
    
    public PanelBuilder overrideByName(String name,List<DynamicConfigValue> properties) {
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
    
    public PanelBuilder overrideByRegexp(String regexp,List<DynamicConfigValue> properties) {
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
    
    public PanelBuilder overrideByFieldType(String fieldType,List<DynamicConfigValue> properties) {
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
    
    public PanelBuilder overrideByQuery(String queryRefId,List<DynamicConfigValue> properties) {
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
    
    public PanelBuilder displayMode(BarGaugeDisplayMode displayMode) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.options).displayMode = displayMode;
        return this;
    }
    
    public PanelBuilder valueMode(BarGaugeValueMode valueMode) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.options).valueMode = valueMode;
        return this;
    }
    
    public PanelBuilder namePlacement(BarGaugeNamePlacement namePlacement) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.options).namePlacement = namePlacement;
        return this;
    }
    
    public PanelBuilder showUnfilled(Boolean showUnfilled) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.options).showUnfilled = showUnfilled;
        return this;
    }
    
    public PanelBuilder sizing(BarGaugeSizing sizing) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.options).sizing = sizing;
        return this;
    }
    
    public PanelBuilder minVizWidth(Integer minVizWidth) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.options).minVizWidth = minVizWidth;
        return this;
    }
    
    public PanelBuilder minVizHeight(Integer minVizHeight) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.options).minVizHeight = minVizHeight;
        return this;
    }
    
    public PanelBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.options).legend = legend.build();
        return this;
    }
    
    public PanelBuilder reduceOptions(com.grafana.foundation.cog.Builder<ReduceDataOptions> reduceOptions) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.options).reduceOptions = reduceOptions.build();
        return this;
    }
    
    public PanelBuilder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.options).text = text.build();
        return this;
    }
    
    public PanelBuilder maxVizHeight(Integer maxVizHeight) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.options).maxVizHeight = maxVizHeight;
        return this;
    }
    
    public PanelBuilder orientation(VizOrientation orientation) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.options).orientation = orientation;
        return this;
    }
    public Panel build() {
        return this.internal;
    }
}
