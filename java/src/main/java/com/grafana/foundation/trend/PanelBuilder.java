// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.trend;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
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
import com.grafana.foundation.common.GraphDrawStyle;
import com.grafana.foundation.common.GraphGradientMode;
import com.grafana.foundation.common.GraphThresholdsStyleConfig;
import com.grafana.foundation.common.LineInterpolation;
import com.grafana.foundation.common.LineStyle;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.BarAlignment;
import com.grafana.foundation.common.StackingConfig;
import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.GraphTransform;
import com.grafana.foundation.common.BoolOrFloat64;
import com.grafana.foundation.common.BoolOrUint32;

public class PanelBuilder implements com.grafana.foundation.cog.Builder<Panel> {
    private Panel internal;

    public PanelBuilder() {
        this.internal = new Panel();
    this.internal.type = "trend";
        this.transparent(false);
        this.height(9);
        this.span(12);
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
    public PanelBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.trend.Options();
		}
        com.grafana.foundation.trend.Options optionsResource = (com.grafana.foundation.trend.Options) this.internal.options;
        optionsResource.legend = legend.build();
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.trend.Options();
		}
        com.grafana.foundation.trend.Options optionsResource = (com.grafana.foundation.trend.Options) this.internal.options;
        optionsResource.tooltip = tooltip.build();
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder xField(String xField) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.trend.Options();
		}
        com.grafana.foundation.trend.Options optionsResource = (com.grafana.foundation.trend.Options) this.internal.options;
        optionsResource.xField = xField;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder drawStyle(GraphDrawStyle drawStyle) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.drawStyle = drawStyle;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder gradientMode(GraphGradientMode gradientMode) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.gradientMode = gradientMode;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder thresholdsStyle(com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> thresholdsStyle) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.thresholdsStyle = thresholdsStyle.build();
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder lineColor(String lineColor) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.lineColor = lineColor;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder lineWidth(Double lineWidth) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.lineWidth = lineWidth;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder lineInterpolation(LineInterpolation lineInterpolation) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.lineInterpolation = lineInterpolation;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.lineStyle = lineStyle.build();
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder fillColor(String fillColor) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.fillColor = fillColor;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder fillOpacity(Double fillOpacity) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.fillOpacity = fillOpacity;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder showPoints(VisibilityMode showPoints) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.showPoints = showPoints;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder pointSize(Double pointSize) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.pointSize = pointSize;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder pointColor(String pointColor) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.pointColor = pointColor;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.scaleDistribution = scaleDistribution.build();
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.axisCenteredZero = axisCenteredZero;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder barAlignment(BarAlignment barAlignment) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.barAlignment = barAlignment;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder barWidthFactor(Double barWidthFactor) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.barWidthFactor = barWidthFactor;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.stacking = stacking.build();
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.hideFrom = hideFrom.build();
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder transform(GraphTransform transform) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.transform = transform;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder spanNulls(BoolOrFloat64 spanNulls) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.spanNulls = spanNulls;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder fillBelowTo(String fillBelowTo) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.fillBelowTo = fillBelowTo;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder pointSymbol(String pointSymbol) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.pointSymbol = pointSymbol;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.axisBorderShow = axisBorderShow;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder barMaxWidth(Double barMaxWidth) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.barMaxWidth = barMaxWidth;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public PanelBuilder insertNulls(BoolOrUint32 insertNulls) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
		if (this.internal.fieldConfig.defaults.custom == null) {
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.trend.FieldConfig();
		}
        com.grafana.foundation.trend.FieldConfig fieldConfigResource = (com.grafana.foundation.trend.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.insertNulls = insertNulls;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    
    public Panel build() {
        return this.internal;
    }
}
