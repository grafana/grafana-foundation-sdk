// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.datagrid;

import com.grafana.foundation.dashboardv2beta1.VizConfigKind;
import java.util.List;
import com.grafana.foundation.dashboardv2beta1.ValueMapping;
import com.grafana.foundation.dashboardv2beta1.ThresholdsConfig;
import com.grafana.foundation.dashboardv2beta1.FieldColor;
import com.grafana.foundation.dashboardv2beta1.Action;
import java.util.LinkedList;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides;
import com.grafana.foundation.dashboardv2beta1.MatcherConfig;
import com.grafana.foundation.dashboardv2beta1.DynamicConfigValue;

public class DatagridVizConfigKindBuilder implements com.grafana.foundation.cog.Builder<VizConfigKind> {
    protected final VizConfigKind internal;
    
    public DatagridVizConfigKindBuilder() {
        this.internal = new VizConfigKind();
        this.internal.kind = "VizConfig";
        this.internal.group = "datagrid";
    }
    public DatagridVizConfigKindBuilder displayName(String displayName) {
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
    
    public DatagridVizConfigKindBuilder displayNameFromDS(String displayNameFromDS) {
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
    
    public DatagridVizConfigKindBuilder description(String description) {
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
    
    public DatagridVizConfigKindBuilder path(String path) {
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
    
    public DatagridVizConfigKindBuilder unit(String unit) {
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
    
    public DatagridVizConfigKindBuilder decimals(Double decimals) {
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
    
    public DatagridVizConfigKindBuilder min(Double min) {
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
    
    public DatagridVizConfigKindBuilder max(Double max) {
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
    
    public DatagridVizConfigKindBuilder mappings(List<ValueMapping> mappings) {
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
    
    public DatagridVizConfigKindBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds) {
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
    
    public DatagridVizConfigKindBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
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
    
    public DatagridVizConfigKindBuilder dataLinks(List<Object> links) {
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
    
    public DatagridVizConfigKindBuilder actions(List<com.grafana.foundation.cog.Builder<Action>> actions) {
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
    
    public DatagridVizConfigKindBuilder noValue(String noValue) {
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
    
    public DatagridVizConfigKindBuilder overrides(List<com.grafana.foundation.cog.Builder<Dashboardv2beta1FieldConfigSourceOverrides>> overrides) {
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
    
    public DatagridVizConfigKindBuilder override(MatcherConfig matcher,List<DynamicConfigValue> properties) {
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
    
    public DatagridVizConfigKindBuilder overrideByName(String name,List<DynamicConfigValue> properties) {
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
    
    public DatagridVizConfigKindBuilder overrideByRegexp(String regexp,List<DynamicConfigValue> properties) {
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
    
    public DatagridVizConfigKindBuilder overrideByFieldType(String fieldType,List<DynamicConfigValue> properties) {
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
    
    public DatagridVizConfigKindBuilder overrideByQuery(String queryRefId,List<DynamicConfigValue> properties) {
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
    
    public DatagridVizConfigKindBuilder selectedSeries(Integer selectedSeries) {
        if (!(selectedSeries >= 0)) {
            throw new IllegalArgumentException("selectedSeries must be >= 0");
        }
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.datagrid.Options();
		}
        ((Options) this.internal.spec.options).selectedSeries = selectedSeries;
        return this;
    }
    public VizConfigKind build() {
        return this.internal;
    }
}
