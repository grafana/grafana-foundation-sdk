// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.barchart;

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
import com.grafana.foundation.common.VizOrientation;
import com.grafana.foundation.common.StackingMode;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.GraphGradientMode;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.GraphThresholdsStyleConfig;

public class BarchartPanelBuilder implements com.grafana.foundation.cog.Builder<Panel> {
    protected final Panel internal;
    
    public BarchartPanelBuilder() {
        this.internal = new Panel();
        this.internal.type = "barchart";
    }
    public BarchartPanelBuilder id(Integer id) {
        this.internal.id = id;
        return this;
    }
    
    public BarchartPanelBuilder targets(List<com.grafana.foundation.cog.Builder<Dataquery>> targets) {
        List<Dataquery> targetsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Dataquery> r1 : targets) {
                Dataquery targetsDepth1 = r1.build();
                targetsResources.add(targetsDepth1); 
        }
        this.internal.targets = targetsResources;
        return this;
    }
    
    public BarchartPanelBuilder withTarget(com.grafana.foundation.cog.Builder<Dataquery> target) {
		if (this.internal.targets == null) {
			this.internal.targets = new LinkedList<>();
		}
    Dataquery targetResource = target.build();
        this.internal.targets.add(targetResource);
        return this;
    }
    
    public BarchartPanelBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public BarchartPanelBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public BarchartPanelBuilder transparent(Boolean transparent) {
        this.internal.transparent = transparent;
        return this;
    }
    
    public BarchartPanelBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public BarchartPanelBuilder gridPos(GridPos gridPos) {
        this.internal.gridPos = gridPos;
        return this;
    }
    
    public BarchartPanelBuilder height(Integer h) {
        if (!(h > 0)) {
            throw new IllegalArgumentException("h must be > 0");
        }
		if (this.internal.gridPos == null) {
			this.internal.gridPos = new com.grafana.foundation.dashboard.GridPos();
		}
        this.internal.gridPos.h = h;
        return this;
    }
    
    public BarchartPanelBuilder span(Integer w) {
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
    
    public BarchartPanelBuilder links(List<com.grafana.foundation.cog.Builder<DashboardLink>> links) {
        List<DashboardLink> linksResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<DashboardLink> r1 : links) {
                DashboardLink linksDepth1 = r1.build();
                linksResources.add(linksDepth1); 
        }
        this.internal.links = linksResources;
        return this;
    }
    
    public BarchartPanelBuilder repeat(String repeat) {
        this.internal.repeat = repeat;
        return this;
    }
    
    public BarchartPanelBuilder repeatDirection(PanelRepeatDirection repeatDirection) {
        this.internal.repeatDirection = repeatDirection;
        return this;
    }
    
    public BarchartPanelBuilder maxPerRow(Double maxPerRow) {
        this.internal.maxPerRow = maxPerRow;
        return this;
    }
    
    public BarchartPanelBuilder maxDataPoints(Double maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public BarchartPanelBuilder transformations(List<DataTransformerConfig> transformations) {
        this.internal.transformations = transformations;
        return this;
    }
    
    public BarchartPanelBuilder withTransformation(DataTransformerConfig transformation) {
		if (this.internal.transformations == null) {
			this.internal.transformations = new LinkedList<>();
		}
        this.internal.transformations.add(transformation);
        return this;
    }
    
    public BarchartPanelBuilder interval(String interval) {
        this.internal.interval = interval;
        return this;
    }
    
    public BarchartPanelBuilder timeFrom(String timeFrom) {
        this.internal.timeFrom = timeFrom;
        return this;
    }
    
    public BarchartPanelBuilder timeShift(String timeShift) {
        this.internal.timeShift = timeShift;
        return this;
    }
    
    public BarchartPanelBuilder hideTimeOverride(Boolean hideTimeOverride) {
        this.internal.hideTimeOverride = hideTimeOverride;
        return this;
    }
    
    public BarchartPanelBuilder libraryPanel(LibraryPanelRef libraryPanel) {
        this.internal.libraryPanel = libraryPanel;
        return this;
    }
    
    public BarchartPanelBuilder displayName(String displayName) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.displayName = displayName;
        return this;
    }
    
    public BarchartPanelBuilder unit(String unit) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.unit = unit;
        return this;
    }
    
    public BarchartPanelBuilder decimals(Double decimals) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.decimals = decimals;
        return this;
    }
    
    public BarchartPanelBuilder min(Double min) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.min = min;
        return this;
    }
    
    public BarchartPanelBuilder max(Double max) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.max = max;
        return this;
    }
    
    public BarchartPanelBuilder mappings(List<ValueMapping> mappings) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.mappings = mappings;
        return this;
    }
    
    public BarchartPanelBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds) {
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
    
    public BarchartPanelBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
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
    
    public BarchartPanelBuilder dataLinks(List<com.grafana.foundation.cog.Builder<DashboardLink>> links) {
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
    
    public BarchartPanelBuilder noValue(String noValue) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
        this.internal.fieldConfig.defaults.noValue = noValue;
        return this;
    }
    
    public BarchartPanelBuilder overrides(List<com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides>> overrides) {
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
    
    public BarchartPanelBuilder withOverride(com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides> override) {
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
    
    public BarchartPanelBuilder overrideByName(String name,List<DynamicConfigValue> properties) {
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
    
    public BarchartPanelBuilder overrideByRegexp(String regexp,List<DynamicConfigValue> properties) {
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
    
    public BarchartPanelBuilder overrideByFieldType(String fieldType,List<DynamicConfigValue> properties) {
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
    
    public BarchartPanelBuilder overrideByQuery(String queryRefId,List<DynamicConfigValue> properties) {
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
    
    public BarchartPanelBuilder xField(String xField) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        ((Options) this.internal.options).xField = xField;
        return this;
    }
    
    public BarchartPanelBuilder colorByField(String colorByField) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        ((Options) this.internal.options).colorByField = colorByField;
        return this;
    }
    
    public BarchartPanelBuilder orientation(VizOrientation orientation) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        ((Options) this.internal.options).orientation = orientation;
        return this;
    }
    
    public BarchartPanelBuilder barRadius(Double barRadius) {
        if (!(barRadius >= 0)) {
            throw new IllegalArgumentException("barRadius must be >= 0");
        }
        if (!(barRadius <= 0.5)) {
            throw new IllegalArgumentException("barRadius must be <= 0.5");
        }
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        ((Options) this.internal.options).barRadius = barRadius;
        return this;
    }
    
    public BarchartPanelBuilder xTickLabelRotation(Integer xTickLabelRotation) {
        if (!(xTickLabelRotation >= -90)) {
            throw new IllegalArgumentException("xTickLabelRotation must be >= -90");
        }
        if (!(xTickLabelRotation <= 90)) {
            throw new IllegalArgumentException("xTickLabelRotation must be <= 90");
        }
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        ((Options) this.internal.options).xTickLabelRotation = xTickLabelRotation;
        return this;
    }
    
    public BarchartPanelBuilder xTickLabelMaxLength(Integer xTickLabelMaxLength) {
        if (!(xTickLabelMaxLength >= 0)) {
            throw new IllegalArgumentException("xTickLabelMaxLength must be >= 0");
        }
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        ((Options) this.internal.options).xTickLabelMaxLength = xTickLabelMaxLength;
        return this;
    }
    
    public BarchartPanelBuilder xTickLabelSpacing(Integer xTickLabelSpacing) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        ((Options) this.internal.options).xTickLabelSpacing = xTickLabelSpacing;
        return this;
    }
    
    public BarchartPanelBuilder stacking(StackingMode stacking) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        ((Options) this.internal.options).stacking = stacking;
        return this;
    }
    
    public BarchartPanelBuilder showValue(VisibilityMode showValue) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        ((Options) this.internal.options).showValue = showValue;
        return this;
    }
    
    public BarchartPanelBuilder barWidth(Double barWidth) {
        if (!(barWidth >= 0)) {
            throw new IllegalArgumentException("barWidth must be >= 0");
        }
        if (!(barWidth <= 1)) {
            throw new IllegalArgumentException("barWidth must be <= 1");
        }
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        ((Options) this.internal.options).barWidth = barWidth;
        return this;
    }
    
    public BarchartPanelBuilder groupWidth(Double groupWidth) {
        if (!(groupWidth >= 0)) {
            throw new IllegalArgumentException("groupWidth must be >= 0");
        }
        if (!(groupWidth <= 1)) {
            throw new IllegalArgumentException("groupWidth must be <= 1");
        }
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        ((Options) this.internal.options).groupWidth = groupWidth;
        return this;
    }
    
    public BarchartPanelBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
    VizLegendOptions legendResource = legend.build();
        ((Options) this.internal.options).legend = legendResource;
        return this;
    }
    
    public BarchartPanelBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
    VizTooltipOptions tooltipResource = tooltip.build();
        ((Options) this.internal.options).tooltip = tooltipResource;
        return this;
    }
    
    public BarchartPanelBuilder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
    VizTextDisplayOptions textResource = text.build();
        ((Options) this.internal.options).text = textResource;
        return this;
    }
    
    public BarchartPanelBuilder fullHighlight(Boolean fullHighlight) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        ((Options) this.internal.options).fullHighlight = fullHighlight;
        return this;
    }
    
    public BarchartPanelBuilder lineWidth(Integer lineWidth) {
        if (!(lineWidth <= 10)) {
            throw new IllegalArgumentException("lineWidth must be <= 10");
        }
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).lineWidth = lineWidth;
        return this;
    }
    
    public BarchartPanelBuilder fillOpacity(Integer fillOpacity) {
        if (!(fillOpacity <= 100)) {
            throw new IllegalArgumentException("fillOpacity must be <= 100");
        }
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).fillOpacity = fillOpacity;
        return this;
    }
    
    public BarchartPanelBuilder gradientMode(GraphGradientMode gradientMode) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).gradientMode = gradientMode;
        return this;
    }
    
    public BarchartPanelBuilder axisPlacement(AxisPlacement axisPlacement) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisPlacement = axisPlacement;
        return this;
    }
    
    public BarchartPanelBuilder axisColorMode(AxisColorMode axisColorMode) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisColorMode = axisColorMode;
        return this;
    }
    
    public BarchartPanelBuilder axisLabel(String axisLabel) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisLabel = axisLabel;
        return this;
    }
    
    public BarchartPanelBuilder axisWidth(Double axisWidth) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisWidth = axisWidth;
        return this;
    }
    
    public BarchartPanelBuilder axisSoftMin(Double axisSoftMin) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisSoftMin = axisSoftMin;
        return this;
    }
    
    public BarchartPanelBuilder axisSoftMax(Double axisSoftMax) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisSoftMax = axisSoftMax;
        return this;
    }
    
    public BarchartPanelBuilder axisGridShow(Boolean axisGridShow) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisGridShow = axisGridShow;
        return this;
    }
    
    public BarchartPanelBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
    ScaleDistributionConfig scaleDistributionResource = scaleDistribution.build();
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).scaleDistribution = scaleDistributionResource;
        return this;
    }
    
    public BarchartPanelBuilder axisCenteredZero(Boolean axisCenteredZero) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisCenteredZero = axisCenteredZero;
        return this;
    }
    
    public BarchartPanelBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
    HideSeriesConfig hideFromResource = hideFrom.build();
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).hideFrom = hideFromResource;
        return this;
    }
    
    public BarchartPanelBuilder thresholdsStyle(com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> thresholdsStyle) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
    GraphThresholdsStyleConfig thresholdsStyleResource = thresholdsStyle.build();
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).thresholdsStyle = thresholdsStyleResource;
        return this;
    }
    
    public BarchartPanelBuilder axisBorderShow(Boolean axisBorderShow) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisBorderShow = axisBorderShow;
        return this;
    }
    public Panel build() {
        return this.internal;
    }
}
