// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bargauge;

import com.grafana.foundation.common.BarGaugeDisplayMode;
import com.grafana.foundation.common.BarGaugeValueMode;
import com.grafana.foundation.common.BarGaugeNamePlacement;
import com.grafana.foundation.common.BarGaugeSizing;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
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

public class PanelBuilder implements com.grafana.foundation.cog.Builder<Panel> {
    private Panel internal;

    public PanelBuilder() {
        this.internal = new Panel();
    this.internal.type = "bargauge";
        this.transparent(false);
        this.height(9);
        this.span(12);
        this.displayMode(BarGaugeDisplayMode.GRADIENT);
        this.valueMode(BarGaugeValueMode.COLOR);
        this.namePlacement(BarGaugeNamePlacement.AUTO);
        this.showUnfilled(true);
        this.sizing(BarGaugeSizing.AUTO);
        this.minVizWidth(8);
        this.minVizHeight(16);
        this.maxVizHeight(300);
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
    public PanelBuilder displayMode(BarGaugeDisplayMode displayMode) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        com.grafana.foundation.bargauge.Options optionsResource = (com.grafana.foundation.bargauge.Options) this.internal.options;
        optionsResource.displayMode = displayMode;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder valueMode(BarGaugeValueMode valueMode) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        com.grafana.foundation.bargauge.Options optionsResource = (com.grafana.foundation.bargauge.Options) this.internal.options;
        optionsResource.valueMode = valueMode;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder namePlacement(BarGaugeNamePlacement namePlacement) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        com.grafana.foundation.bargauge.Options optionsResource = (com.grafana.foundation.bargauge.Options) this.internal.options;
        optionsResource.namePlacement = namePlacement;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder showUnfilled(Boolean showUnfilled) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        com.grafana.foundation.bargauge.Options optionsResource = (com.grafana.foundation.bargauge.Options) this.internal.options;
        optionsResource.showUnfilled = showUnfilled;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder sizing(BarGaugeSizing sizing) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        com.grafana.foundation.bargauge.Options optionsResource = (com.grafana.foundation.bargauge.Options) this.internal.options;
        optionsResource.sizing = sizing;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder minVizWidth(Integer minVizWidth) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        com.grafana.foundation.bargauge.Options optionsResource = (com.grafana.foundation.bargauge.Options) this.internal.options;
        optionsResource.minVizWidth = minVizWidth;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder minVizHeight(Integer minVizHeight) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        com.grafana.foundation.bargauge.Options optionsResource = (com.grafana.foundation.bargauge.Options) this.internal.options;
        optionsResource.minVizHeight = minVizHeight;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder reduceOptions(com.grafana.foundation.cog.Builder<ReduceDataOptions> reduceOptions) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        com.grafana.foundation.bargauge.Options optionsResource = (com.grafana.foundation.bargauge.Options) this.internal.options;
        optionsResource.reduceOptions = reduceOptions.build();
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        com.grafana.foundation.bargauge.Options optionsResource = (com.grafana.foundation.bargauge.Options) this.internal.options;
        optionsResource.text = text.build();
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder maxVizHeight(Integer maxVizHeight) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        com.grafana.foundation.bargauge.Options optionsResource = (com.grafana.foundation.bargauge.Options) this.internal.options;
        optionsResource.maxVizHeight = maxVizHeight;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder orientation(VizOrientation orientation) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.bargauge.Options();
		}
        com.grafana.foundation.bargauge.Options optionsResource = (com.grafana.foundation.bargauge.Options) this.internal.options;
        optionsResource.orientation = orientation;
    this.internal.options = optionsResource;
        return this;
    }
    
    public Panel build() {
        return this.internal;
    }
}
