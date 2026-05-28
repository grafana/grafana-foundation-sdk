// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.timeseries;

import com.grafana.foundation.dashboardv2beta1.VizConfigKind;
import java.util.List;
import com.grafana.foundation.dashboardv2beta1.ValueMapping;
import com.grafana.foundation.dashboardv2beta1.ThresholdsConfig;
import com.grafana.foundation.dashboardv2beta1.FieldColor;
import com.grafana.foundation.dashboardv2beta1.Action;
import java.util.LinkedList;
import com.grafana.foundation.dashboardv2beta1.NullValueMode;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides;
import com.grafana.foundation.dashboardv2beta1.MatcherConfig;
import com.grafana.foundation.dashboardv2beta1.DynamicConfigValue;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import com.grafana.foundation.common.VizOrientation;
import com.grafana.foundation.common.GraphDrawStyle;
import com.grafana.foundation.common.GraphGradientMode;
import com.grafana.foundation.common.GraphThresholdsStyleConfig;
import com.grafana.foundation.common.GraphTransform;
import com.grafana.foundation.common.LineInterpolation;
import com.grafana.foundation.common.LineStyle;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.BarAlignment;
import com.grafana.foundation.common.StackingConfig;
import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.BoolOrFloat64;

public class TimeseriesVisualizationBuilder implements com.grafana.foundation.cog.Builder<VizConfigKind> {
    protected final VizConfigKind internal;
    
    public TimeseriesVisualizationBuilder() {
        this.internal = new VizConfigKind();
        this.internal.kind = "VizConfig";
        this.internal.group = "timeseries";
    }
    public TimeseriesVisualizationBuilder displayName(String displayName) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.displayName = displayName;
        return this;
    }
    
    public TimeseriesVisualizationBuilder displayNameFromDS(String displayNameFromDS) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.displayNameFromDS = displayNameFromDS;
        return this;
    }
    
    public TimeseriesVisualizationBuilder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.description = description;
        return this;
    }
    
    public TimeseriesVisualizationBuilder path(String path) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.path = path;
        return this;
    }
    
    public TimeseriesVisualizationBuilder unit(String unit) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.unit = unit;
        return this;
    }
    
    public TimeseriesVisualizationBuilder decimals(Double decimals) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.decimals = decimals;
        return this;
    }
    
    public TimeseriesVisualizationBuilder min(Double min) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.min = min;
        return this;
    }
    
    public TimeseriesVisualizationBuilder max(Double max) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.max = max;
        return this;
    }
    
    public TimeseriesVisualizationBuilder mappings(List<ValueMapping> mappings) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.mappings = mappings;
        return this;
    }
    
    public TimeseriesVisualizationBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
    ThresholdsConfig thresholdsResource = thresholds.build();
        this.internal.spec.fieldConfig.defaults.thresholds = thresholdsResource;
        return this;
    }
    
    public TimeseriesVisualizationBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
    FieldColor colorResource = color.build();
        this.internal.spec.fieldConfig.defaults.color = colorResource;
        return this;
    }
    
    public TimeseriesVisualizationBuilder dataLinks(List<Object> links) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.links = links;
        return this;
    }
    
    public TimeseriesVisualizationBuilder actions(List<com.grafana.foundation.cog.Builder<Action>> actions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        List<Action> actionsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Action> r1 : actions) {
                Action actionsDepth1 = r1.build();
                actionsResources.add(actionsDepth1); 
        }
        this.internal.spec.fieldConfig.defaults.actions = actionsResources;
        return this;
    }
    
    public TimeseriesVisualizationBuilder noValue(String noValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.noValue = noValue;
        return this;
    }
    
    public TimeseriesVisualizationBuilder fieldMinMax(Boolean fieldMinMax) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.fieldMinMax = fieldMinMax;
        return this;
    }
    
    public TimeseriesVisualizationBuilder nullValueMode(NullValueMode nullValueMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.nullValueMode = nullValueMode;
        return this;
    }
    
    public TimeseriesVisualizationBuilder overrides(List<com.grafana.foundation.cog.Builder<Dashboardv2beta1FieldConfigSourceOverrides>> overrides) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
        List<Dashboardv2beta1FieldConfigSourceOverrides> overridesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Dashboardv2beta1FieldConfigSourceOverrides> r1 : overrides) {
                Dashboardv2beta1FieldConfigSourceOverrides overridesDepth1 = r1.build();
                overridesResources.add(overridesDepth1); 
        }
        this.internal.spec.fieldConfig.overrides = overridesResources;
        return this;
    }
    
    public TimeseriesVisualizationBuilder override(MatcherConfig matcher,List<DynamicConfigValue> properties) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.overrides == null) {
			this.internal.spec.fieldConfig.overrides = new LinkedList<>();
		}
    Dashboardv2beta1FieldConfigSourceOverrides dashboardv2beta1FieldConfigSourceOverrides = new Dashboardv2beta1FieldConfigSourceOverrides();
        dashboardv2beta1FieldConfigSourceOverrides.matcher = matcher;
        dashboardv2beta1FieldConfigSourceOverrides.properties = properties;
        this.internal.spec.fieldConfig.overrides.add(dashboardv2beta1FieldConfigSourceOverrides);
        return this;
    }
    
    public TimeseriesVisualizationBuilder timezone(List<String> timezone) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.timeseries.Options();
		}
        ((Options) this.internal.spec.options).timezone = timezone;
        return this;
    }
    
    public TimeseriesVisualizationBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.timeseries.Options();
		}
    VizLegendOptions legendResource = legend.build();
        ((Options) this.internal.spec.options).legend = legendResource;
        return this;
    }
    
    public TimeseriesVisualizationBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.timeseries.Options();
		}
    VizTooltipOptions tooltipResource = tooltip.build();
        ((Options) this.internal.spec.options).tooltip = tooltipResource;
        return this;
    }
    
    public TimeseriesVisualizationBuilder orientation(VizOrientation orientation) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.timeseries.Options();
		}
        ((Options) this.internal.spec.options).orientation = orientation;
        return this;
    }
    
    public TimeseriesVisualizationBuilder drawStyle(GraphDrawStyle drawStyle) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).drawStyle = drawStyle;
        return this;
    }
    
    public TimeseriesVisualizationBuilder gradientMode(GraphGradientMode gradientMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).gradientMode = gradientMode;
        return this;
    }
    
    public TimeseriesVisualizationBuilder thresholdsStyle(com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> thresholdsStyle) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
    GraphThresholdsStyleConfig thresholdsStyleResource = thresholdsStyle.build();
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).thresholdsStyle = thresholdsStyleResource;
        return this;
    }
    
    public TimeseriesVisualizationBuilder transform(GraphTransform transform) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).transform = transform;
        return this;
    }
    
    public TimeseriesVisualizationBuilder lineColor(String lineColor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).lineColor = lineColor;
        return this;
    }
    
    public TimeseriesVisualizationBuilder lineWidth(Double lineWidth) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).lineWidth = lineWidth;
        return this;
    }
    
    public TimeseriesVisualizationBuilder lineInterpolation(LineInterpolation lineInterpolation) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).lineInterpolation = lineInterpolation;
        return this;
    }
    
    public TimeseriesVisualizationBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
    LineStyle lineStyleResource = lineStyle.build();
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).lineStyle = lineStyleResource;
        return this;
    }
    
    public TimeseriesVisualizationBuilder fillColor(String fillColor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).fillColor = fillColor;
        return this;
    }
    
    public TimeseriesVisualizationBuilder fillOpacity(Double fillOpacity) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).fillOpacity = fillOpacity;
        return this;
    }
    
    public TimeseriesVisualizationBuilder showPoints(VisibilityMode showPoints) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).showPoints = showPoints;
        return this;
    }
    
    public TimeseriesVisualizationBuilder pointSize(Double pointSize) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).pointSize = pointSize;
        return this;
    }
    
    public TimeseriesVisualizationBuilder pointColor(String pointColor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).pointColor = pointColor;
        return this;
    }
    
    public TimeseriesVisualizationBuilder axisPlacement(AxisPlacement axisPlacement) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisPlacement = axisPlacement;
        return this;
    }
    
    public TimeseriesVisualizationBuilder axisColorMode(AxisColorMode axisColorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisColorMode = axisColorMode;
        return this;
    }
    
    public TimeseriesVisualizationBuilder axisLabel(String axisLabel) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisLabel = axisLabel;
        return this;
    }
    
    public TimeseriesVisualizationBuilder axisWidth(Double axisWidth) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisWidth = axisWidth;
        return this;
    }
    
    public TimeseriesVisualizationBuilder axisSoftMin(Double axisSoftMin) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisSoftMin = axisSoftMin;
        return this;
    }
    
    public TimeseriesVisualizationBuilder axisSoftMax(Double axisSoftMax) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisSoftMax = axisSoftMax;
        return this;
    }
    
    public TimeseriesVisualizationBuilder axisGridShow(Boolean axisGridShow) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisGridShow = axisGridShow;
        return this;
    }
    
    public TimeseriesVisualizationBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
    ScaleDistributionConfig scaleDistributionResource = scaleDistribution.build();
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).scaleDistribution = scaleDistributionResource;
        return this;
    }
    
    public TimeseriesVisualizationBuilder axisCenteredZero(Boolean axisCenteredZero) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisCenteredZero = axisCenteredZero;
        return this;
    }
    
    public TimeseriesVisualizationBuilder barAlignment(BarAlignment barAlignment) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).barAlignment = barAlignment;
        return this;
    }
    
    public TimeseriesVisualizationBuilder barWidthFactor(Double barWidthFactor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).barWidthFactor = barWidthFactor;
        return this;
    }
    
    public TimeseriesVisualizationBuilder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
    StackingConfig stackingResource = stacking.build();
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).stacking = stackingResource;
        return this;
    }
    
    public TimeseriesVisualizationBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
    HideSeriesConfig hideFromResource = hideFrom.build();
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).hideFrom = hideFromResource;
        return this;
    }
    
    public TimeseriesVisualizationBuilder insertNulls(BoolOrFloat64 insertNulls) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).insertNulls = insertNulls;
        return this;
    }
    
    public TimeseriesVisualizationBuilder spanNulls(BoolOrFloat64 spanNulls) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).spanNulls = spanNulls;
        return this;
    }
    
    public TimeseriesVisualizationBuilder fillBelowTo(String fillBelowTo) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).fillBelowTo = fillBelowTo;
        return this;
    }
    
    public TimeseriesVisualizationBuilder pointSymbol(String pointSymbol) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).pointSymbol = pointSymbol;
        return this;
    }
    
    public TimeseriesVisualizationBuilder axisBorderShow(Boolean axisBorderShow) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisBorderShow = axisBorderShow;
        return this;
    }
    
    public TimeseriesVisualizationBuilder barMaxWidth(Double barMaxWidth) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.timeseries.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).barMaxWidth = barMaxWidth;
        return this;
    }
    
    public TimeseriesVisualizationBuilder overrideByName(String name,List<DynamicConfigValue> properties) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.overrides == null) {
			this.internal.spec.fieldConfig.overrides = new LinkedList<>();
		}
    MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byName";
        matcherConfig.options = name;
    Dashboardv2beta1FieldConfigSourceOverrides dashboardv2beta1FieldConfigSourceOverrides = new Dashboardv2beta1FieldConfigSourceOverrides();
        dashboardv2beta1FieldConfigSourceOverrides.matcher = matcherConfig;
        dashboardv2beta1FieldConfigSourceOverrides.properties = properties;
        this.internal.spec.fieldConfig.overrides.add(dashboardv2beta1FieldConfigSourceOverrides);
        return this;
    }
    
    public TimeseriesVisualizationBuilder overrideByRegexp(String regexp,List<DynamicConfigValue> properties) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.overrides == null) {
			this.internal.spec.fieldConfig.overrides = new LinkedList<>();
		}
    MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byRegexp";
        matcherConfig.options = regexp;
    Dashboardv2beta1FieldConfigSourceOverrides dashboardv2beta1FieldConfigSourceOverrides = new Dashboardv2beta1FieldConfigSourceOverrides();
        dashboardv2beta1FieldConfigSourceOverrides.matcher = matcherConfig;
        dashboardv2beta1FieldConfigSourceOverrides.properties = properties;
        this.internal.spec.fieldConfig.overrides.add(dashboardv2beta1FieldConfigSourceOverrides);
        return this;
    }
    
    public TimeseriesVisualizationBuilder overrideByFieldType(String fieldType,List<DynamicConfigValue> properties) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.overrides == null) {
			this.internal.spec.fieldConfig.overrides = new LinkedList<>();
		}
    MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byType";
        matcherConfig.options = fieldType;
    Dashboardv2beta1FieldConfigSourceOverrides dashboardv2beta1FieldConfigSourceOverrides = new Dashboardv2beta1FieldConfigSourceOverrides();
        dashboardv2beta1FieldConfigSourceOverrides.matcher = matcherConfig;
        dashboardv2beta1FieldConfigSourceOverrides.properties = properties;
        this.internal.spec.fieldConfig.overrides.add(dashboardv2beta1FieldConfigSourceOverrides);
        return this;
    }
    
    public TimeseriesVisualizationBuilder overrideByQuery(String queryRefId,List<DynamicConfigValue> properties) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.overrides == null) {
			this.internal.spec.fieldConfig.overrides = new LinkedList<>();
		}
    MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byFrameRefID";
        matcherConfig.options = queryRefId;
    Dashboardv2beta1FieldConfigSourceOverrides dashboardv2beta1FieldConfigSourceOverrides = new Dashboardv2beta1FieldConfigSourceOverrides();
        dashboardv2beta1FieldConfigSourceOverrides.matcher = matcherConfig;
        dashboardv2beta1FieldConfigSourceOverrides.properties = properties;
        this.internal.spec.fieldConfig.overrides.add(dashboardv2beta1FieldConfigSourceOverrides);
        return this;
    }
    public VizConfigKind build() {
        return this.internal;
    }
}
