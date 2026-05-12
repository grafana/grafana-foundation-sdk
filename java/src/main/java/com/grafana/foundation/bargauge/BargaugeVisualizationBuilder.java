// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bargauge;

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
import com.grafana.foundation.common.BarGaugeDisplayMode;
import com.grafana.foundation.common.BarGaugeValueMode;
import com.grafana.foundation.common.BarGaugeNamePlacement;
import com.grafana.foundation.common.BarGaugeSizing;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;

public class BargaugeVisualizationBuilder implements com.grafana.foundation.cog.Builder<VizConfigKind> {
    protected final VizConfigKind internal;
    
    public BargaugeVisualizationBuilder() {
        this.internal = new VizConfigKind();
        this.internal.kind = "VizConfig";
        this.internal.group = "bargauge";
    }
    public BargaugeVisualizationBuilder displayName(String displayName) {
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
    
    public BargaugeVisualizationBuilder displayNameFromDS(String displayNameFromDS) {
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
    
    public BargaugeVisualizationBuilder description(String description) {
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
    
    public BargaugeVisualizationBuilder path(String path) {
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
    
    public BargaugeVisualizationBuilder unit(String unit) {
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
    
    public BargaugeVisualizationBuilder decimals(Double decimals) {
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
    
    public BargaugeVisualizationBuilder min(Double min) {
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
    
    public BargaugeVisualizationBuilder max(Double max) {
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
    
    public BargaugeVisualizationBuilder mappings(List<ValueMapping> mappings) {
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
    
    public BargaugeVisualizationBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds) {
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
    
    public BargaugeVisualizationBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
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
    
    public BargaugeVisualizationBuilder dataLinks(List<Object> links) {
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
    
    public BargaugeVisualizationBuilder actions(List<com.grafana.foundation.cog.Builder<Action>> actions) {
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
    
    public BargaugeVisualizationBuilder noValue(String noValue) {
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
    
    public BargaugeVisualizationBuilder fieldMinMax(Boolean fieldMinMax) {
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
    
    public BargaugeVisualizationBuilder nullValueMode(NullValueMode nullValueMode) {
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
    
    public BargaugeVisualizationBuilder overrides(List<com.grafana.foundation.cog.Builder<Dashboardv2beta1FieldConfigSourceOverrides>> overrides) {
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
    
    public BargaugeVisualizationBuilder override(MatcherConfig matcher,List<DynamicConfigValue> properties) {
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
    
    public BargaugeVisualizationBuilder displayMode(BarGaugeDisplayMode displayMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.spec.options).displayMode = displayMode;
        return this;
    }
    
    public BargaugeVisualizationBuilder valueMode(BarGaugeValueMode valueMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.spec.options).valueMode = valueMode;
        return this;
    }
    
    public BargaugeVisualizationBuilder namePlacement(BarGaugeNamePlacement namePlacement) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.spec.options).namePlacement = namePlacement;
        return this;
    }
    
    public BargaugeVisualizationBuilder showUnfilled(Boolean showUnfilled) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.spec.options).showUnfilled = showUnfilled;
        return this;
    }
    
    public BargaugeVisualizationBuilder sizing(BarGaugeSizing sizing) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.spec.options).sizing = sizing;
        return this;
    }
    
    public BargaugeVisualizationBuilder minVizWidth(Integer minVizWidth) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.spec.options).minVizWidth = minVizWidth;
        return this;
    }
    
    public BargaugeVisualizationBuilder minVizHeight(Integer minVizHeight) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.spec.options).minVizHeight = minVizHeight;
        return this;
    }
    
    public BargaugeVisualizationBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.bargauge.Options();
		}
    VizLegendOptions legendResource = legend.build();
        ((Options) this.internal.spec.options).legend = legendResource;
        return this;
    }
    
    public BargaugeVisualizationBuilder reduceOptions(com.grafana.foundation.cog.Builder<ReduceDataOptions> reduceOptions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.bargauge.Options();
		}
    ReduceDataOptions reduceOptionsResource = reduceOptions.build();
        ((Options) this.internal.spec.options).reduceOptions = reduceOptionsResource;
        return this;
    }
    
    public BargaugeVisualizationBuilder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.bargauge.Options();
		}
    VizTextDisplayOptions textResource = text.build();
        ((Options) this.internal.spec.options).text = textResource;
        return this;
    }
    
    public BargaugeVisualizationBuilder maxVizHeight(Integer maxVizHeight) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.spec.options).maxVizHeight = maxVizHeight;
        return this;
    }
    
    public BargaugeVisualizationBuilder orientation(VizOrientation orientation) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.bargauge.Options();
		}
        ((Options) this.internal.spec.options).orientation = orientation;
        return this;
    }
    
    public BargaugeVisualizationBuilder overrideByName(String name,List<DynamicConfigValue> properties) {
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
    
    public BargaugeVisualizationBuilder overrideByRegexp(String regexp,List<DynamicConfigValue> properties) {
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
    
    public BargaugeVisualizationBuilder overrideByFieldType(String fieldType,List<DynamicConfigValue> properties) {
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
    
    public BargaugeVisualizationBuilder overrideByQuery(String queryRefId,List<DynamicConfigValue> properties) {
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
