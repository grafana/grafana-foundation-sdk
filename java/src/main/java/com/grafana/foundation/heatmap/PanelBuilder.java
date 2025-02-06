// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

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
import com.grafana.foundation.common.HeatmapCalculationOptions;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.HideSeriesConfig;

public class PanelBuilder implements com.grafana.foundation.cog.Builder<Panel> {
    protected final Panel internal;
    
    public PanelBuilder() {
        this.internal = new Panel();
        this.internal.type = "heatmap";
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
    
    public PanelBuilder libraryPanel(LibraryPanelRef libraryPanel) {
        this.internal.libraryPanel = libraryPanel;
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
    
    public PanelBuilder calculate(Boolean calculate) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.options).calculate = calculate;
        return this;
    }
    
    public PanelBuilder calculation(com.grafana.foundation.cog.Builder<HeatmapCalculationOptions> calculation) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.options).calculation = calculation.build();
        return this;
    }
    
    public PanelBuilder color(com.grafana.foundation.cog.Builder<HeatmapColorOptions> color) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.options).color = color.build();
        return this;
    }
    
    public PanelBuilder filterValues(com.grafana.foundation.cog.Builder<FilterValueRange> filterValues) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.options).filterValues = filterValues.build();
        return this;
    }
    
    public PanelBuilder rowsFrame(com.grafana.foundation.cog.Builder<RowsHeatmapOptions> rowsFrame) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.options).rowsFrame = rowsFrame.build();
        return this;
    }
    
    public PanelBuilder showValue(VisibilityMode showValue) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.options).showValue = showValue;
        return this;
    }
    
    public PanelBuilder cellGap(Integer cellGap) {
        if (!(cellGap <= 25)) {
            throw new IllegalArgumentException("cellGap must be <= 25");
        }
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.options).cellGap = cellGap;
        return this;
    }
    
    public PanelBuilder cellRadius(Float cellRadius) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.options).cellRadius = cellRadius;
        return this;
    }
    
    public PanelBuilder cellValues(com.grafana.foundation.cog.Builder<CellValues> cellValues) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.options).cellValues = cellValues.build();
        return this;
    }
    
    public PanelBuilder yAxis(com.grafana.foundation.cog.Builder<YAxisConfig> yAxis) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.options).yAxis = yAxis.build();
        return this;
    }
    
    public PanelBuilder showLegend() {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
		if (((Options) this.internal.options).legend == null) {
			((Options) this.internal.options).legend = new com.grafana.foundation.heatmap.HeatmapLegendBuilder().build();
		}
        ((Options) this.internal.options).legend.show = true;
        return this;
    }
    
    public PanelBuilder hideLegend() {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
		if (((Options) this.internal.options).legend == null) {
			((Options) this.internal.options).legend = new com.grafana.foundation.heatmap.HeatmapLegendBuilder().build();
		}
        ((Options) this.internal.options).legend.show = false;
        return this;
    }
    
    public PanelBuilder showTooltip() {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
		if (((Options) this.internal.options).tooltip == null) {
			((Options) this.internal.options).tooltip = new com.grafana.foundation.heatmap.HeatmapTooltipBuilder().build();
		}
        ((Options) this.internal.options).tooltip.show = true;
        return this;
    }
    
    public PanelBuilder hideTooltip() {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
		if (((Options) this.internal.options).tooltip == null) {
			((Options) this.internal.options).tooltip = new com.grafana.foundation.heatmap.HeatmapTooltipBuilder().build();
		}
        ((Options) this.internal.options).tooltip.show = false;
        return this;
    }
    
    public PanelBuilder showYHistogram() {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
		if (((Options) this.internal.options).tooltip == null) {
			((Options) this.internal.options).tooltip = new com.grafana.foundation.heatmap.HeatmapTooltipBuilder().build();
		}
        ((Options) this.internal.options).tooltip.yHistogram = true;
        return this;
    }
    
    public PanelBuilder hideYHistogram() {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
		if (((Options) this.internal.options).tooltip == null) {
			((Options) this.internal.options).tooltip = new com.grafana.foundation.heatmap.HeatmapTooltipBuilder().build();
		}
        ((Options) this.internal.options).tooltip.yHistogram = false;
        return this;
    }
    
    public PanelBuilder exemplarsColor(String color) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.heatmap.Options();
		}
		if (((Options) this.internal.options).exemplars == null) {
			((Options) this.internal.options).exemplars = new com.grafana.foundation.heatmap.ExemplarConfigBuilder().build();
		}
        ((Options) this.internal.options).exemplars.color = color;
        return this;
    }
    
    public PanelBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.heatmap.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).scaleDistribution = scaleDistribution.build();
        return this;
    }
    
    public PanelBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.heatmap.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).hideFrom = hideFrom.build();
        return this;
    }
    public Panel build() {
        return this.internal;
    }
}
