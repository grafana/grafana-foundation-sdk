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
    private Panel internal;

    public PanelBuilder() {
        this.internal = new Panel();
    this.internal.type = "xychart";
        this.transparent(false);
        this.height(9);
        this.span(12);
        this.show(ScatterShow.POINTS);
        this.label(VisibilityMode.AUTO);
    }
    public PanelBuilder id(Integer id) {
    this.internal.id = id;
        return this;
    }
    public PanelBuilder targets(com.grafana.foundation.cog.Builder<List<Dataquery>> targets) {
    this.internal.targets = targets.build();
        return this;
    }
    public PanelBuilder withTarget(com.grafana.foundation.cog.Builder<Dataquery> targets) {
		if (this.internal.targets == null) {
			this.internal.targets = new LinkedList<>();
		}
    this.internal.targets.add(targets.build());
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
    public PanelBuilder withTransformation(DataTransformerConfig transformations) {
		if (this.internal.transformations == null) {
			this.internal.transformations = new LinkedList<>();
		}
    this.internal.transformations.add(transformations);
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
    public PanelBuilder withOverride(com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides> overrides) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.overrides == null) {
			this.internal.fieldConfig.overrides = new LinkedList<>();
		}
    this.internal.fieldConfig.overrides.add(overrides.build());
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.show = show;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.pointSize = pointSize.build();
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.pointColor = pointColor.build();
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.lineColor = lineColor.build();
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.lineWidth = lineWidth;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.lineStyle = lineStyle.build();
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.label = label;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.hideFrom = hideFrom.build();
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.axisPlacement = axisPlacement;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.axisColorMode = axisColorMode;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.axisLabel = axisLabel;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.axisWidth = axisWidth;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.axisSoftMin = axisSoftMin;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.axisSoftMax = axisSoftMax;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.axisGridShow = axisGridShow;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.scaleDistribution = scaleDistribution.build();
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.labelValue = labelValue.build();
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
        com.grafana.foundation.xychart.FieldConfig fieldConfigResource = (com.grafana.foundation.xychart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.axisCenteredZero = axisCenteredZero;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder seriesMapping(SeriesMapping seriesMapping) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.xychart.Options();
		}
        com.grafana.foundation.xychart.Options optionsResource = (com.grafana.foundation.xychart.Options) this.internal.options;
        optionsResource.seriesMapping = seriesMapping;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder dims(XYDimensionConfig dims) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.xychart.Options();
		}
        com.grafana.foundation.xychart.Options optionsResource = (com.grafana.foundation.xychart.Options) this.internal.options;
        optionsResource.dims = dims;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.xychart.Options();
		}
        com.grafana.foundation.xychart.Options optionsResource = (com.grafana.foundation.xychart.Options) this.internal.options;
        optionsResource.legend = legend.build();
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.xychart.Options();
		}
        com.grafana.foundation.xychart.Options optionsResource = (com.grafana.foundation.xychart.Options) this.internal.options;
        optionsResource.tooltip = tooltip.build();
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder series(List<ScatterSeriesConfig> series) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.xychart.Options();
		}
        com.grafana.foundation.xychart.Options optionsResource = (com.grafana.foundation.xychart.Options) this.internal.options;
        optionsResource.series = series;
    this.internal.options = optionsResource;
        return this;
    }
    
    public Panel build() {
        return this.internal;
    }
}
