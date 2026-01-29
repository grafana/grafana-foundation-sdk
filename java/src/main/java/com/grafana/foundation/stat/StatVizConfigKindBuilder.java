// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.stat;

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
import com.grafana.foundation.common.BigValueGraphMode;
import com.grafana.foundation.common.BigValueColorMode;
import com.grafana.foundation.common.BigValueJustifyMode;
import com.grafana.foundation.common.BigValueTextMode;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.PercentChangeColorMode;
import com.grafana.foundation.common.VizOrientation;

public class StatVizConfigKindBuilder implements com.grafana.foundation.cog.Builder<VizConfigKind> {
    protected final VizConfigKind internal;
    
    public StatVizConfigKindBuilder() {
        this.internal = new VizConfigKind();
        this.internal.kind = "VizConfig";
        this.internal.group = "stat";
    }
    public StatVizConfigKindBuilder displayName(String displayName) {
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
    
    public StatVizConfigKindBuilder displayNameFromDS(String displayNameFromDS) {
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
    
    public StatVizConfigKindBuilder description(String description) {
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
    
    public StatVizConfigKindBuilder path(String path) {
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
    
    public StatVizConfigKindBuilder unit(String unit) {
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
    
    public StatVizConfigKindBuilder decimals(Double decimals) {
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
    
    public StatVizConfigKindBuilder min(Double min) {
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
    
    public StatVizConfigKindBuilder max(Double max) {
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
    
    public StatVizConfigKindBuilder mappings(List<ValueMapping> mappings) {
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
    
    public StatVizConfigKindBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds) {
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
    
    public StatVizConfigKindBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
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
    
    public StatVizConfigKindBuilder dataLinks(List<Object> links) {
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
    
    public StatVizConfigKindBuilder actions(List<com.grafana.foundation.cog.Builder<Action>> actions) {
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
    
    public StatVizConfigKindBuilder noValue(String noValue) {
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
    
    public StatVizConfigKindBuilder fieldMinMax(Boolean fieldMinMax) {
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
    
    public StatVizConfigKindBuilder nullValueMode(NullValueMode nullValueMode) {
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
    
    public StatVizConfigKindBuilder overrides(List<com.grafana.foundation.cog.Builder<Dashboardv2beta1FieldConfigSourceOverrides>> overrides) {
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
    
    public StatVizConfigKindBuilder override(String systemRef,MatcherConfig matcher,List<DynamicConfigValue> properties) {
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
        dashboardv2beta1FieldConfigSourceOverrides.systemRef = systemRef;
        dashboardv2beta1FieldConfigSourceOverrides.matcher = matcher;
        dashboardv2beta1FieldConfigSourceOverrides.properties = properties;
        this.internal.spec.fieldConfig.overrides.add(dashboardv2beta1FieldConfigSourceOverrides);
        return this;
    }
    
    public StatVizConfigKindBuilder overrideByName(String name,List<DynamicConfigValue> properties) {
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
    
    public StatVizConfigKindBuilder overrideByRegexp(String regexp,List<DynamicConfigValue> properties) {
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
    
    public StatVizConfigKindBuilder overrideByFieldType(String fieldType,List<DynamicConfigValue> properties) {
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
    
    public StatVizConfigKindBuilder overrideByQuery(String queryRefId,List<DynamicConfigValue> properties) {
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
    
    public StatVizConfigKindBuilder graphMode(BigValueGraphMode graphMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.stat.Options();
		}
        ((Options) this.internal.spec.options).graphMode = graphMode;
        return this;
    }
    
    public StatVizConfigKindBuilder colorMode(BigValueColorMode colorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.stat.Options();
		}
        ((Options) this.internal.spec.options).colorMode = colorMode;
        return this;
    }
    
    public StatVizConfigKindBuilder justifyMode(BigValueJustifyMode justifyMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.stat.Options();
		}
        ((Options) this.internal.spec.options).justifyMode = justifyMode;
        return this;
    }
    
    public StatVizConfigKindBuilder textMode(BigValueTextMode textMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.stat.Options();
		}
        ((Options) this.internal.spec.options).textMode = textMode;
        return this;
    }
    
    public StatVizConfigKindBuilder wideLayout(Boolean wideLayout) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.stat.Options();
		}
        ((Options) this.internal.spec.options).wideLayout = wideLayout;
        return this;
    }
    
    public StatVizConfigKindBuilder showPercentChange(Boolean showPercentChange) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.stat.Options();
		}
        ((Options) this.internal.spec.options).showPercentChange = showPercentChange;
        return this;
    }
    
    public StatVizConfigKindBuilder reduceOptions(com.grafana.foundation.cog.Builder<ReduceDataOptions> reduceOptions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.stat.Options();
		}
    ReduceDataOptions reduceOptionsResource = reduceOptions.build();
        ((Options) this.internal.spec.options).reduceOptions = reduceOptionsResource;
        return this;
    }
    
    public StatVizConfigKindBuilder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.stat.Options();
		}
    VizTextDisplayOptions textResource = text.build();
        ((Options) this.internal.spec.options).text = textResource;
        return this;
    }
    
    public StatVizConfigKindBuilder percentChangeColorMode(PercentChangeColorMode percentChangeColorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.stat.Options();
		}
        ((Options) this.internal.spec.options).percentChangeColorMode = percentChangeColorMode;
        return this;
    }
    
    public StatVizConfigKindBuilder orientation(VizOrientation orientation) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.stat.Options();
		}
        ((Options) this.internal.spec.options).orientation = orientation;
        return this;
    }
    public VizConfigKind build() {
        return this.internal;
    }
}
