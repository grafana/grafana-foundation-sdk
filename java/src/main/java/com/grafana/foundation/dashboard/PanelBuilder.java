// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;
import com.grafana.foundation.cog.variants.Dataquery;
import java.util.LinkedList;

public class PanelBuilder<T extends PanelBuilder<T>> implements com.grafana.foundation.cog.Builder<Panel> {
    protected final Panel internal;
    
    public PanelBuilder() {
        this.internal = new Panel();
    }
    public T type(String type) {
        if (!(type.length() >= 1)) {
            throw new IllegalArgumentException("type.length() must be >= 1");
        }
    this.internal.type = type;
        return (T) this;
    }
    
    public T id(Integer id) {
    this.internal.id = id;
        return (T) this;
    }
    
    public T pluginVersion(String pluginVersion) {
    this.internal.pluginVersion = pluginVersion;
        return (T) this;
    }
    
    public T tags(List<String> tags) {
    this.internal.tags = tags;
        return (T) this;
    }
    
    public T targets(com.grafana.foundation.cog.Builder<List<Dataquery>> targets) {
    this.internal.targets = targets.build();
        return (T) this;
    }
    
    public T withTarget(com.grafana.foundation.cog.Builder<Dataquery> target) {
		if (this.internal.targets == null) {
			this.internal.targets = new LinkedList<>();
		}
    this.internal.targets.add(target.build());
        return (T) this;
    }
    
    public T title(String title) {
    this.internal.title = title;
        return (T) this;
    }
    
    public T description(String description) {
    this.internal.description = description;
        return (T) this;
    }
    
    public T transparent(Boolean transparent) {
    this.internal.transparent = transparent;
        return (T) this;
    }
    
    public T datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return (T) this;
    }
    
    public T gridPos(GridPos gridPos) {
    this.internal.gridPos = gridPos;
        return (T) this;
    }
    
    public T height(Integer h) {
        if (!(h > 0)) {
            throw new IllegalArgumentException("h must be > 0");
        }
		if (this.internal.gridPos == null) {
			this.internal.gridPos = new com.grafana.foundation.dashboard.GridPos();
		}
    this.internal.gridPos.h = h;
        return (T) this;
    }
    
    public T span(Integer w) {
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
        return (T) this;
    }
    
    public T links(com.grafana.foundation.cog.Builder<List<DashboardLink>> links) {
    this.internal.links = links.build();
        return (T) this;
    }
    
    public T repeat(String repeat) {
    this.internal.repeat = repeat;
        return (T) this;
    }
    
    public T repeatDirection(PanelRepeatDirection repeatDirection) {
    this.internal.repeatDirection = repeatDirection;
        return (T) this;
    }
    
    public T repeatPanelId(Long repeatPanelId) {
    this.internal.repeatPanelId = repeatPanelId;
        return (T) this;
    }
    
    public T maxDataPoints(Double maxDataPoints) {
    this.internal.maxDataPoints = maxDataPoints;
        return (T) this;
    }
    
    public T transformations(List<DataTransformerConfig> transformations) {
    this.internal.transformations = transformations;
        return (T) this;
    }
    
    public T withTransformation(DataTransformerConfig transformation) {
		if (this.internal.transformations == null) {
			this.internal.transformations = new LinkedList<>();
		}
    this.internal.transformations.add(transformation);
        return (T) this;
    }
    
    public T interval(String interval) {
    this.internal.interval = interval;
        return (T) this;
    }
    
    public T timeFrom(String timeFrom) {
    this.internal.timeFrom = timeFrom;
        return (T) this;
    }
    
    public T timeShift(String timeShift) {
    this.internal.timeShift = timeShift;
        return (T) this;
    }
    
    public T libraryPanel(LibraryPanelRef libraryPanel) {
    this.internal.libraryPanel = libraryPanel;
        return (T) this;
    }
    
    public T options(Object options) {
    this.internal.options = options;
        return (T) this;
    }
    
    public T fieldConfig(FieldConfigSource fieldConfig) {
    this.internal.fieldConfig = fieldConfig;
        return (T) this;
    }
    
    public T displayName(String displayName) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    this.internal.fieldConfig.defaults.displayName = displayName;
        return (T) this;
    }
    
    public T unit(String unit) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    this.internal.fieldConfig.defaults.unit = unit;
        return (T) this;
    }
    
    public T decimals(Double decimals) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    this.internal.fieldConfig.defaults.decimals = decimals;
        return (T) this;
    }
    
    public T min(Double min) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    this.internal.fieldConfig.defaults.min = min;
        return (T) this;
    }
    
    public T max(Double max) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    this.internal.fieldConfig.defaults.max = max;
        return (T) this;
    }
    
    public T mappings(List<ValueMapping> mappings) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    this.internal.fieldConfig.defaults.mappings = mappings;
        return (T) this;
    }
    
    public T thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    this.internal.fieldConfig.defaults.thresholds = thresholds.build();
        return (T) this;
    }
    
    public T colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    this.internal.fieldConfig.defaults.color = color.build();
        return (T) this;
    }
    
    public T dataLinks(com.grafana.foundation.cog.Builder<List<DashboardLink>> links) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    this.internal.fieldConfig.defaults.links = links.build();
        return (T) this;
    }
    
    public T noValue(String noValue) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    this.internal.fieldConfig.defaults.noValue = noValue;
        return (T) this;
    }
    
    public T custom(Object custom) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.defaults == null) {
			this.internal.fieldConfig.defaults = new com.grafana.foundation.dashboard.FieldConfig();
		}
    this.internal.fieldConfig.defaults.custom = custom;
        return (T) this;
    }
    
    public T defaults(FieldConfig defaults) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
    this.internal.fieldConfig.defaults = defaults;
        return (T) this;
    }
    
    public T overrides(com.grafana.foundation.cog.Builder<List<DashboardFieldConfigSourceOverrides>> overrides) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
    this.internal.fieldConfig.overrides = overrides.build();
        return (T) this;
    }
    
    public T withOverride(com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides> override) {
		if (this.internal.fieldConfig == null) {
			this.internal.fieldConfig = new com.grafana.foundation.dashboard.FieldConfigSource();
		}
		if (this.internal.fieldConfig.overrides == null) {
			this.internal.fieldConfig.overrides = new LinkedList<>();
		}
    this.internal.fieldConfig.overrides.add(override.build());
        return (T) this;
    }
    public Panel build() {
        return this.internal;
    }
}
