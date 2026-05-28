// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.candlestick;

import com.grafana.foundation.dashboardv2.VizConfigKind;
import java.util.List;
import com.grafana.foundation.dashboardv2.ValueMapping;
import com.grafana.foundation.dashboardv2.ThresholdsConfig;
import com.grafana.foundation.dashboardv2.FieldColor;
import com.grafana.foundation.dashboardv2.Action;
import java.util.LinkedList;
import com.grafana.foundation.dashboardv2.NullValueMode;
import com.grafana.foundation.dashboardv2.Dashboardv2FieldConfigSourceOverrides;
import com.grafana.foundation.dashboardv2.MatcherConfig;
import com.grafana.foundation.dashboardv2.DynamicConfigValue;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
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

public class CandlestickVisualizationV2Builder implements com.grafana.foundation.cog.Builder<VizConfigKind> {
    protected final VizConfigKind internal;
    
    public CandlestickVisualizationV2Builder() {
        this.internal = new VizConfigKind();
        this.internal.kind = "VizConfig";
        this.internal.group = "candlestick";
    }
    public CandlestickVisualizationV2Builder displayName(String displayName) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.displayName = displayName;
        return this;
    }
    
    public CandlestickVisualizationV2Builder displayNameFromDS(String displayNameFromDS) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.displayNameFromDS = displayNameFromDS;
        return this;
    }
    
    public CandlestickVisualizationV2Builder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.description = description;
        return this;
    }
    
    public CandlestickVisualizationV2Builder path(String path) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.path = path;
        return this;
    }
    
    public CandlestickVisualizationV2Builder unit(String unit) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.unit = unit;
        return this;
    }
    
    public CandlestickVisualizationV2Builder decimals(Double decimals) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.decimals = decimals;
        return this;
    }
    
    public CandlestickVisualizationV2Builder min(Double min) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.min = min;
        return this;
    }
    
    public CandlestickVisualizationV2Builder max(Double max) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.max = max;
        return this;
    }
    
    public CandlestickVisualizationV2Builder mappings(List<ValueMapping> mappings) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.mappings = mappings;
        return this;
    }
    
    public CandlestickVisualizationV2Builder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
    ThresholdsConfig thresholdsResource = thresholds.build();
        this.internal.spec.fieldConfig.defaults.thresholds = thresholdsResource;
        return this;
    }
    
    public CandlestickVisualizationV2Builder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
    FieldColor colorResource = color.build();
        this.internal.spec.fieldConfig.defaults.color = colorResource;
        return this;
    }
    
    public CandlestickVisualizationV2Builder dataLinks(List<Object> links) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.links = links;
        return this;
    }
    
    public CandlestickVisualizationV2Builder actions(List<com.grafana.foundation.cog.Builder<Action>> actions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        List<Action> actionsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Action> r1 : actions) {
                Action actionsDepth1 = r1.build();
                actionsResources.add(actionsDepth1); 
        }
        this.internal.spec.fieldConfig.defaults.actions = actionsResources;
        return this;
    }
    
    public CandlestickVisualizationV2Builder noValue(String noValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.noValue = noValue;
        return this;
    }
    
    public CandlestickVisualizationV2Builder fieldMinMax(Boolean fieldMinMax) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.fieldMinMax = fieldMinMax;
        return this;
    }
    
    public CandlestickVisualizationV2Builder nullValueMode(NullValueMode nullValueMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
        this.internal.spec.fieldConfig.defaults.nullValueMode = nullValueMode;
        return this;
    }
    
    public CandlestickVisualizationV2Builder overrides(List<com.grafana.foundation.cog.Builder<Dashboardv2FieldConfigSourceOverrides>> overrides) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
        List<Dashboardv2FieldConfigSourceOverrides> overridesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Dashboardv2FieldConfigSourceOverrides> r1 : overrides) {
                Dashboardv2FieldConfigSourceOverrides overridesDepth1 = r1.build();
                overridesResources.add(overridesDepth1); 
        }
        this.internal.spec.fieldConfig.overrides = overridesResources;
        return this;
    }
    
    public CandlestickVisualizationV2Builder override(MatcherConfig matcher,List<DynamicConfigValue> properties) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.overrides == null) {
			this.internal.spec.fieldConfig.overrides = new LinkedList<>();
		}
    Dashboardv2FieldConfigSourceOverrides dashboardv2FieldConfigSourceOverrides = new Dashboardv2FieldConfigSourceOverrides();
        dashboardv2FieldConfigSourceOverrides.matcher = matcher;
        dashboardv2FieldConfigSourceOverrides.properties = properties;
        this.internal.spec.fieldConfig.overrides.add(dashboardv2FieldConfigSourceOverrides);
        return this;
    }
    
    public CandlestickVisualizationV2Builder mode(VizDisplayMode mode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.candlestick.Options();
		}
        ((Options) this.internal.spec.options).mode = mode;
        return this;
    }
    
    public CandlestickVisualizationV2Builder candleStyle(CandleStyle candleStyle) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.candlestick.Options();
		}
        ((Options) this.internal.spec.options).candleStyle = candleStyle;
        return this;
    }
    
    public CandlestickVisualizationV2Builder colorStrategy(ColorStrategy colorStrategy) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.candlestick.Options();
		}
        ((Options) this.internal.spec.options).colorStrategy = colorStrategy;
        return this;
    }
    
    public CandlestickVisualizationV2Builder fields(com.grafana.foundation.cog.Builder<CandlestickFieldMap> fields) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.candlestick.Options();
		}
    CandlestickFieldMap fieldsResource = fields.build();
        ((Options) this.internal.spec.options).fields = fieldsResource;
        return this;
    }
    
    public CandlestickVisualizationV2Builder colors(com.grafana.foundation.cog.Builder<CandlestickColors> colors) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.candlestick.Options();
		}
    CandlestickColors colorsResource = colors.build();
        ((Options) this.internal.spec.options).colors = colorsResource;
        return this;
    }
    
    public CandlestickVisualizationV2Builder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.candlestick.Options();
		}
    VizLegendOptions legendResource = legend.build();
        ((Options) this.internal.spec.options).legend = legendResource;
        return this;
    }
    
    public CandlestickVisualizationV2Builder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.candlestick.Options();
		}
    VizTooltipOptions tooltipResource = tooltip.build();
        ((Options) this.internal.spec.options).tooltip = tooltipResource;
        return this;
    }
    
    public CandlestickVisualizationV2Builder includeAllFields(Boolean includeAllFields) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.candlestick.Options();
		}
        ((Options) this.internal.spec.options).includeAllFields = includeAllFields;
        return this;
    }
    
    public CandlestickVisualizationV2Builder drawStyle(GraphDrawStyle drawStyle) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).drawStyle = drawStyle;
        return this;
    }
    
    public CandlestickVisualizationV2Builder gradientMode(GraphGradientMode gradientMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).gradientMode = gradientMode;
        return this;
    }
    
    public CandlestickVisualizationV2Builder thresholdsStyle(com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> thresholdsStyle) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
    GraphThresholdsStyleConfig thresholdsStyleResource = thresholdsStyle.build();
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).thresholdsStyle = thresholdsStyleResource;
        return this;
    }
    
    public CandlestickVisualizationV2Builder transform(GraphTransform transform) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).transform = transform;
        return this;
    }
    
    public CandlestickVisualizationV2Builder lineColor(String lineColor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).lineColor = lineColor;
        return this;
    }
    
    public CandlestickVisualizationV2Builder lineWidth(Double lineWidth) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).lineWidth = lineWidth;
        return this;
    }
    
    public CandlestickVisualizationV2Builder lineInterpolation(LineInterpolation lineInterpolation) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).lineInterpolation = lineInterpolation;
        return this;
    }
    
    public CandlestickVisualizationV2Builder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
    LineStyle lineStyleResource = lineStyle.build();
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).lineStyle = lineStyleResource;
        return this;
    }
    
    public CandlestickVisualizationV2Builder fillColor(String fillColor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).fillColor = fillColor;
        return this;
    }
    
    public CandlestickVisualizationV2Builder fillOpacity(Double fillOpacity) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).fillOpacity = fillOpacity;
        return this;
    }
    
    public CandlestickVisualizationV2Builder showPoints(VisibilityMode showPoints) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).showPoints = showPoints;
        return this;
    }
    
    public CandlestickVisualizationV2Builder pointSize(Double pointSize) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).pointSize = pointSize;
        return this;
    }
    
    public CandlestickVisualizationV2Builder pointColor(String pointColor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).pointColor = pointColor;
        return this;
    }
    
    public CandlestickVisualizationV2Builder axisPlacement(AxisPlacement axisPlacement) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisPlacement = axisPlacement;
        return this;
    }
    
    public CandlestickVisualizationV2Builder axisColorMode(AxisColorMode axisColorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisColorMode = axisColorMode;
        return this;
    }
    
    public CandlestickVisualizationV2Builder axisLabel(String axisLabel) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisLabel = axisLabel;
        return this;
    }
    
    public CandlestickVisualizationV2Builder axisWidth(Double axisWidth) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisWidth = axisWidth;
        return this;
    }
    
    public CandlestickVisualizationV2Builder axisSoftMin(Double axisSoftMin) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisSoftMin = axisSoftMin;
        return this;
    }
    
    public CandlestickVisualizationV2Builder axisSoftMax(Double axisSoftMax) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisSoftMax = axisSoftMax;
        return this;
    }
    
    public CandlestickVisualizationV2Builder axisGridShow(Boolean axisGridShow) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisGridShow = axisGridShow;
        return this;
    }
    
    public CandlestickVisualizationV2Builder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
    ScaleDistributionConfig scaleDistributionResource = scaleDistribution.build();
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).scaleDistribution = scaleDistributionResource;
        return this;
    }
    
    public CandlestickVisualizationV2Builder axisCenteredZero(Boolean axisCenteredZero) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisCenteredZero = axisCenteredZero;
        return this;
    }
    
    public CandlestickVisualizationV2Builder barAlignment(BarAlignment barAlignment) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).barAlignment = barAlignment;
        return this;
    }
    
    public CandlestickVisualizationV2Builder barWidthFactor(Double barWidthFactor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).barWidthFactor = barWidthFactor;
        return this;
    }
    
    public CandlestickVisualizationV2Builder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
    StackingConfig stackingResource = stacking.build();
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).stacking = stackingResource;
        return this;
    }
    
    public CandlestickVisualizationV2Builder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
    HideSeriesConfig hideFromResource = hideFrom.build();
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).hideFrom = hideFromResource;
        return this;
    }
    
    public CandlestickVisualizationV2Builder insertNulls(BoolOrFloat64 insertNulls) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).insertNulls = insertNulls;
        return this;
    }
    
    public CandlestickVisualizationV2Builder spanNulls(BoolOrFloat64 spanNulls) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).spanNulls = spanNulls;
        return this;
    }
    
    public CandlestickVisualizationV2Builder fillBelowTo(String fillBelowTo) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).fillBelowTo = fillBelowTo;
        return this;
    }
    
    public CandlestickVisualizationV2Builder pointSymbol(String pointSymbol) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).pointSymbol = pointSymbol;
        return this;
    }
    
    public CandlestickVisualizationV2Builder axisBorderShow(Boolean axisBorderShow) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).axisBorderShow = axisBorderShow;
        return this;
    }
    
    public CandlestickVisualizationV2Builder barMaxWidth(Double barMaxWidth) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.defaults == null) {
			this.internal.spec.fieldConfig.defaults = new com.grafana.foundation.dashboardv2.FieldConfig();
		}
		if (this.internal.spec.fieldConfig.defaults.custom == null) {
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.candlestick.FieldConfig();
		}
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).barMaxWidth = barMaxWidth;
        return this;
    }
    
    public CandlestickVisualizationV2Builder overrideByName(String name,List<DynamicConfigValue> properties) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.overrides == null) {
			this.internal.spec.fieldConfig.overrides = new LinkedList<>();
		}
    MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byName";
        matcherConfig.options = name;
    Dashboardv2FieldConfigSourceOverrides dashboardv2FieldConfigSourceOverrides = new Dashboardv2FieldConfigSourceOverrides();
        dashboardv2FieldConfigSourceOverrides.matcher = matcherConfig;
        dashboardv2FieldConfigSourceOverrides.properties = properties;
        this.internal.spec.fieldConfig.overrides.add(dashboardv2FieldConfigSourceOverrides);
        return this;
    }
    
    public CandlestickVisualizationV2Builder overrideByRegexp(String regexp,List<DynamicConfigValue> properties) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.overrides == null) {
			this.internal.spec.fieldConfig.overrides = new LinkedList<>();
		}
    MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byRegexp";
        matcherConfig.options = regexp;
    Dashboardv2FieldConfigSourceOverrides dashboardv2FieldConfigSourceOverrides = new Dashboardv2FieldConfigSourceOverrides();
        dashboardv2FieldConfigSourceOverrides.matcher = matcherConfig;
        dashboardv2FieldConfigSourceOverrides.properties = properties;
        this.internal.spec.fieldConfig.overrides.add(dashboardv2FieldConfigSourceOverrides);
        return this;
    }
    
    public CandlestickVisualizationV2Builder overrideByFieldType(String fieldType,List<DynamicConfigValue> properties) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.overrides == null) {
			this.internal.spec.fieldConfig.overrides = new LinkedList<>();
		}
    MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byType";
        matcherConfig.options = fieldType;
    Dashboardv2FieldConfigSourceOverrides dashboardv2FieldConfigSourceOverrides = new Dashboardv2FieldConfigSourceOverrides();
        dashboardv2FieldConfigSourceOverrides.matcher = matcherConfig;
        dashboardv2FieldConfigSourceOverrides.properties = properties;
        this.internal.spec.fieldConfig.overrides.add(dashboardv2FieldConfigSourceOverrides);
        return this;
    }
    
    public CandlestickVisualizationV2Builder overrideByQuery(String queryRefId,List<DynamicConfigValue> properties) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.fieldConfig == null) {
			this.internal.spec.fieldConfig = new com.grafana.foundation.dashboardv2.FieldConfigSource();
		}
		if (this.internal.spec.fieldConfig.overrides == null) {
			this.internal.spec.fieldConfig.overrides = new LinkedList<>();
		}
    MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byFrameRefID";
        matcherConfig.options = queryRefId;
    Dashboardv2FieldConfigSourceOverrides dashboardv2FieldConfigSourceOverrides = new Dashboardv2FieldConfigSourceOverrides();
        dashboardv2FieldConfigSourceOverrides.matcher = matcherConfig;
        dashboardv2FieldConfigSourceOverrides.properties = properties;
        this.internal.spec.fieldConfig.overrides.add(dashboardv2FieldConfigSourceOverrides);
        return this;
    }
    public VizConfigKind build() {
        return this.internal;
    }
}
