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

public class PanelBuilder implements com.grafana.foundation.cog.Builder<Panel> {
    protected final Panel internal;
    
    public PanelBuilder() {
        this.internal = new Panel();
    this.internal.type = "barchart";
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
    
    public PanelBuilder xField(String xField) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.xField = xField;
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder colorByField(String colorByField) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.colorByField = colorByField;
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder orientation(VizOrientation orientation) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.orientation = orientation;
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder barRadius(Double barRadius) {
        if (!(barRadius >= 0)) {
            throw new IllegalArgumentException("barRadius must be >= 0");
        }
        if (!(barRadius <= 0.5)) {
            throw new IllegalArgumentException("barRadius must be <= 0.5");
        }
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.barRadius = barRadius;
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder xTickLabelRotation(Integer xTickLabelRotation) {
        if (!(xTickLabelRotation >= -90)) {
            throw new IllegalArgumentException("xTickLabelRotation must be >= -90");
        }
        if (!(xTickLabelRotation <= 90)) {
            throw new IllegalArgumentException("xTickLabelRotation must be <= 90");
        }
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.xTickLabelRotation = xTickLabelRotation;
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder xTickLabelMaxLength(Integer xTickLabelMaxLength) {
        if (!(xTickLabelMaxLength >= 0)) {
            throw new IllegalArgumentException("xTickLabelMaxLength must be >= 0");
        }
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.xTickLabelMaxLength = xTickLabelMaxLength;
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder xTickLabelSpacing(Integer xTickLabelSpacing) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.xTickLabelSpacing = xTickLabelSpacing;
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder stacking(StackingMode stacking) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.stacking = stacking;
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder showValue(VisibilityMode showValue) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.showValue = showValue;
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder barWidth(Double barWidth) {
        if (!(barWidth >= 0)) {
            throw new IllegalArgumentException("barWidth must be >= 0");
        }
        if (!(barWidth <= 1)) {
            throw new IllegalArgumentException("barWidth must be <= 1");
        }
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.barWidth = barWidth;
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder groupWidth(Double groupWidth) {
        if (!(groupWidth >= 0)) {
            throw new IllegalArgumentException("groupWidth must be >= 0");
        }
        if (!(groupWidth <= 1)) {
            throw new IllegalArgumentException("groupWidth must be <= 1");
        }
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.groupWidth = groupWidth;
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.legend = legend.build();
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.tooltip = tooltip.build();
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.text = text.build();
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder fullHighlight(Boolean fullHighlight) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.barchart.Options();
		}
        com.grafana.foundation.barchart.Options optionsResource = (com.grafana.foundation.barchart.Options) this.internal.options;
        optionsResource.fullHighlight = fullHighlight;
    this.internal.options = optionsResource;
        return this;
    }
    
    public PanelBuilder lineWidth(Integer lineWidth) {
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
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.lineWidth = lineWidth;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    
    public PanelBuilder fillOpacity(Integer fillOpacity) {
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
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.fillOpacity = fillOpacity;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.gradientMode = gradientMode;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.scaleDistribution = scaleDistribution.build();
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.hideFrom = hideFrom.build();
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.thresholdsStyle = thresholdsStyle.build();
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
			this.internal.fieldConfig.defaults.custom = new com.grafana.foundation.barchart.FieldConfig();
		}
        com.grafana.foundation.barchart.FieldConfig fieldConfigResource = (com.grafana.foundation.barchart.FieldConfig) this.internal.fieldConfig.defaults.custom;
        fieldConfigResource.axisCenteredZero = axisCenteredZero;
    this.internal.fieldConfig.defaults.custom = fieldConfigResource;
        return this;
    }
    public Panel build() {
        return this.internal;
    }
}
