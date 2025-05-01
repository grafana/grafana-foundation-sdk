// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

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
import com.grafana.foundation.common.ScaleDimensionConfig;
import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.LineStyle;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.TextDimensionConfig;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;

public class PanelBuilder implements com.grafana.foundation.cog.Builder<Panel> {
    protected final Panel internal;
    
    public PanelBuilder() {
        this.internal = new Panel();
        this.internal.type = "xychart";
    }
    public PanelBuilder id(Integer id) {
        this.internal.id = id;
        return this;
    }
    
    public PanelBuilder targets(List<com.grafana.foundation.cog.Builder<Dataquery>> targets) {
        List<Dataquery> targetsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Dataquery> r1 : targets) {
                Dataquery targetsDepth1 = r1.build();
                targetsResources.add(targetsDepth1); 
        }
        this.internal.targets = targetsResources;
        return this;
    }
    
    public PanelBuilder withTarget(com.grafana.foundation.cog.Builder<Dataquery> target) {
		if (this.internal.targets == null) {
			this.internal.targets = new LinkedList<>();
		}
    Dataquery targetResource = target.build();
        this.internal.targets.add(targetResource);
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
    
    public PanelBuilder links(List<com.grafana.foundation.cog.Builder<DashboardLink>> links) {
        List<DashboardLink> linksResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<DashboardLink> r1 : links) {
                DashboardLink linksDepth1 = r1.build();
                linksResources.add(linksDepth1); 
        }
        this.internal.links = linksResources;
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
    ThresholdsConfig thresholdsResource = thresholds.build();
        this.internal.fieldConfig.defaults.thresholds = thresholdsResource;
        return this;
    }
    
    public PanelBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
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
    
    public PanelBuilder dataLinks(List<com.grafana.foundation.cog.Builder<DashboardLink>> links) {
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
    
    public PanelBuilder overrides(List<com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides>> overrides) {
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
    
    public PanelBuilder withOverride(com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides> override) {
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
    
    public PanelBuilder show(ScatterShow show) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).show = show;
        return this;
    }
    
    public PanelBuilder pointSize(com.grafana.foundation.cog.Builder<ScaleDimensionConfig> pointSize) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
    ScaleDimensionConfig pointSizeResource = pointSize.build();
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).pointSize = pointSizeResource;
        return this;
    }
    
    public PanelBuilder pointColor(com.grafana.foundation.cog.Builder<ColorDimensionConfig> pointColor) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
    ColorDimensionConfig pointColorResource = pointColor.build();
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).pointColor = pointColorResource;
        return this;
    }
    
    public PanelBuilder lineColor(com.grafana.foundation.cog.Builder<ColorDimensionConfig> lineColor) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
    ColorDimensionConfig lineColorResource = lineColor.build();
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).lineColor = lineColorResource;
        return this;
    }
    
    public PanelBuilder lineWidth(Integer lineWidth) {
        if (!(lineWidth >= 0)) {
            throw new IllegalArgumentException("lineWidth must be >= 0");
        }
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).lineWidth = lineWidth;
        return this;
    }
    
    public PanelBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
    LineStyle lineStyleResource = lineStyle.build();
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).lineStyle = lineStyleResource;
        return this;
    }
    
    public PanelBuilder label(VisibilityMode label) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).label = label;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
    HideSeriesConfig hideFromResource = hideFrom.build();
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).hideFrom = hideFromResource;
        return this;
    }
    
    public PanelBuilder axisPlacement(AxisPlacement axisPlacement) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisPlacement = axisPlacement;
        return this;
    }
    
    public PanelBuilder axisColorMode(AxisColorMode axisColorMode) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisColorMode = axisColorMode;
        return this;
    }
    
    public PanelBuilder axisLabel(String axisLabel) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisLabel = axisLabel;
        return this;
    }
    
    public PanelBuilder axisWidth(Double axisWidth) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisWidth = axisWidth;
        return this;
    }
    
    public PanelBuilder axisSoftMin(Double axisSoftMin) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisSoftMin = axisSoftMin;
        return this;
    }
    
    public PanelBuilder axisSoftMax(Double axisSoftMax) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisSoftMax = axisSoftMax;
        return this;
    }
    
    public PanelBuilder axisGridShow(Boolean axisGridShow) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisGridShow = axisGridShow;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
    ScaleDistributionConfig scaleDistributionResource = scaleDistribution.build();
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).scaleDistribution = scaleDistributionResource;
        return this;
    }
    
    public PanelBuilder axisCenteredZero(Boolean axisCenteredZero) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisCenteredZero = axisCenteredZero;
        return this;
    }
    
    public PanelBuilder labelValue(com.grafana.foundation.cog.Builder<TextDimensionConfig> labelValue) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
    TextDimensionConfig labelValueResource = labelValue.build();
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).labelValue = labelValueResource;
        return this;
    }
    
    public PanelBuilder axisBorderShow(Boolean axisBorderShow) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.xychart.FieldConfig();
		}
        ((FieldConfig) this.internal.fieldConfig.defaults.custom).axisBorderShow = axisBorderShow;
        return this;
    }
    
    public PanelBuilder seriesMapping(SeriesMapping seriesMapping) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.xychart.Options();
		}
        ((Options) this.internal.options).seriesMapping = seriesMapping;
        return this;
    }
    
    public PanelBuilder dims(com.grafana.foundation.cog.Builder<XYDimensionConfig> dims) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.xychart.Options();
		}
    XYDimensionConfig dimsResource = dims.build();
        ((Options) this.internal.options).dims = dimsResource;
        return this;
    }
    
    public PanelBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.xychart.Options();
		}
    VizLegendOptions legendResource = legend.build();
        ((Options) this.internal.options).legend = legendResource;
        return this;
    }
    
    public PanelBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.xychart.Options();
		}
    VizTooltipOptions tooltipResource = tooltip.build();
        ((Options) this.internal.options).tooltip = tooltipResource;
        return this;
    }
    
    public PanelBuilder series(List<com.grafana.foundation.cog.Builder<ScatterSeriesConfig>> series) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.xychart.Options();
		}
        List<ScatterSeriesConfig> seriesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<ScatterSeriesConfig> r1 : series) {
                ScatterSeriesConfig seriesDepth1 = r1.build();
                seriesResources.add(seriesDepth1); 
        }
        ((Options) this.internal.options).series = seriesResources;
        return this;
    }
    public Panel build() {
        return this.internal;
    }
}
