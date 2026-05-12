// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

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
import com.grafana.foundation.common.HeatmapCalculationOptions;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.HideSeriesConfig;

public class HeatmapVisualizationV2Builder implements com.grafana.foundation.cog.Builder<VizConfigKind> {
    protected final VizConfigKind internal;
    
    public HeatmapVisualizationV2Builder() {
        this.internal = new VizConfigKind();
        this.internal.kind = "VizConfig";
        this.internal.group = "heatmap";
    }
    public HeatmapVisualizationV2Builder displayName(String displayName) {
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
    
    public HeatmapVisualizationV2Builder displayNameFromDS(String displayNameFromDS) {
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
    
    public HeatmapVisualizationV2Builder description(String description) {
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
    
    public HeatmapVisualizationV2Builder path(String path) {
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
    
    public HeatmapVisualizationV2Builder unit(String unit) {
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
    
    public HeatmapVisualizationV2Builder decimals(Double decimals) {
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
    
    public HeatmapVisualizationV2Builder min(Double min) {
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
    
    public HeatmapVisualizationV2Builder max(Double max) {
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
    
    public HeatmapVisualizationV2Builder mappings(List<ValueMapping> mappings) {
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
    
    public HeatmapVisualizationV2Builder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds) {
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
    
    public HeatmapVisualizationV2Builder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
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
    
    public HeatmapVisualizationV2Builder dataLinks(List<Object> links) {
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
    
    public HeatmapVisualizationV2Builder actions(List<com.grafana.foundation.cog.Builder<Action>> actions) {
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
    
    public HeatmapVisualizationV2Builder noValue(String noValue) {
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
    
    public HeatmapVisualizationV2Builder fieldMinMax(Boolean fieldMinMax) {
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
    
    public HeatmapVisualizationV2Builder nullValueMode(NullValueMode nullValueMode) {
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
    
    public HeatmapVisualizationV2Builder overrides(List<com.grafana.foundation.cog.Builder<Dashboardv2FieldConfigSourceOverrides>> overrides) {
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
    
    public HeatmapVisualizationV2Builder override(MatcherConfig matcher,List<DynamicConfigValue> properties) {
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
    
    public HeatmapVisualizationV2Builder calculate(Boolean calculate) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.spec.options).calculate = calculate;
        return this;
    }
    
    public HeatmapVisualizationV2Builder calculation(com.grafana.foundation.cog.Builder<HeatmapCalculationOptions> calculation) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
    HeatmapCalculationOptions calculationResource = calculation.build();
        ((Options) this.internal.spec.options).calculation = calculationResource;
        return this;
    }
    
    public HeatmapVisualizationV2Builder color(com.grafana.foundation.cog.Builder<HeatmapColorOptions> color) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
    HeatmapColorOptions colorResource = color.build();
        ((Options) this.internal.spec.options).color = colorResource;
        return this;
    }
    
    public HeatmapVisualizationV2Builder filterValues(com.grafana.foundation.cog.Builder<FilterValueRange> filterValues) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
    FilterValueRange filterValuesResource = filterValues.build();
        ((Options) this.internal.spec.options).filterValues = filterValuesResource;
        return this;
    }
    
    public HeatmapVisualizationV2Builder rowsFrame(com.grafana.foundation.cog.Builder<RowsHeatmapOptions> rowsFrame) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
    RowsHeatmapOptions rowsFrameResource = rowsFrame.build();
        ((Options) this.internal.spec.options).rowsFrame = rowsFrameResource;
        return this;
    }
    
    public HeatmapVisualizationV2Builder showValue(VisibilityMode showValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.spec.options).showValue = showValue;
        return this;
    }
    
    public HeatmapVisualizationV2Builder cellGap(Integer cellGap) {
        if (!(cellGap >= 0)) {
            throw new IllegalArgumentException("cellGap must be >= 0");
        }
        if (!(cellGap <= 25)) {
            throw new IllegalArgumentException("cellGap must be <= 25");
        }
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.spec.options).cellGap = cellGap;
        return this;
    }
    
    public HeatmapVisualizationV2Builder cellRadius(Float cellRadius) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.spec.options).cellRadius = cellRadius;
        return this;
    }
    
    public HeatmapVisualizationV2Builder cellValues(com.grafana.foundation.cog.Builder<CellValues> cellValues) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
    CellValues cellValuesResource = cellValues.build();
        ((Options) this.internal.spec.options).cellValues = cellValuesResource;
        return this;
    }
    
    public HeatmapVisualizationV2Builder yAxis(com.grafana.foundation.cog.Builder<YAxisConfig> yAxis) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
    YAxisConfig yAxisResource = yAxis.build();
        ((Options) this.internal.spec.options).yAxis = yAxisResource;
        return this;
    }
    
    public HeatmapVisualizationV2Builder legend(com.grafana.foundation.cog.Builder<HeatmapLegend> legend) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
    HeatmapLegend legendResource = legend.build();
        ((Options) this.internal.spec.options).legend = legendResource;
        return this;
    }
    
    public HeatmapVisualizationV2Builder tooltip(com.grafana.foundation.cog.Builder<HeatmapTooltip> tooltip) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
    HeatmapTooltip tooltipResource = tooltip.build();
        ((Options) this.internal.spec.options).tooltip = tooltipResource;
        return this;
    }
    
    public HeatmapVisualizationV2Builder exemplars(com.grafana.foundation.cog.Builder<ExemplarConfig> exemplars) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
    ExemplarConfig exemplarsResource = exemplars.build();
        ((Options) this.internal.spec.options).exemplars = exemplarsResource;
        return this;
    }
    
    public HeatmapVisualizationV2Builder selectionMode(HeatmapSelectionMode selectionMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.VizConfigSpec();
		}
		if (this.internal.spec.options == null) {
			this.internal.spec.options = new com.grafana.foundation.heatmap.Options();
		}
        ((Options) this.internal.spec.options).selectionMode = selectionMode;
        return this;
    }
    
    public HeatmapVisualizationV2Builder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
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
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.heatmap.FieldConfig();
		}
    ScaleDistributionConfig scaleDistributionResource = scaleDistribution.build();
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).scaleDistribution = scaleDistributionResource;
        return this;
    }
    
    public HeatmapVisualizationV2Builder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
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
			this.internal.spec.fieldConfig.defaults.custom = new com.grafana.foundation.heatmap.FieldConfig();
		}
    HideSeriesConfig hideFromResource = hideFrom.build();
        ((FieldConfig) this.internal.spec.fieldConfig.defaults.custom).hideFrom = hideFromResource;
        return this;
    }
    
    public HeatmapVisualizationV2Builder overrideByName(String name,List<DynamicConfigValue> properties) {
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
    
    public HeatmapVisualizationV2Builder overrideByRegexp(String regexp,List<DynamicConfigValue> properties) {
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
    
    public HeatmapVisualizationV2Builder overrideByFieldType(String fieldType,List<DynamicConfigValue> properties) {
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
    
    public HeatmapVisualizationV2Builder overrideByQuery(String queryRefId,List<DynamicConfigValue> properties) {
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
