package com.grafana.custompanel;

import com.grafana.foundation.cog.Builder;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.dashboard.*;

import java.util.LinkedList;
import java.util.List;

public class CustomPanelBuilder implements Builder<Panel> {
    private final Panel internal;

    public CustomPanelBuilder() {
        this.internal = new Panel();
        this.internal.type = "custom-panel";
        this.transparent(true);
        this.height(9);
        this.span(12);
    }

    public CustomPanelBuilder makeItBeautiful() {
        if (this.internal.options == null) {
            this.internal.options = new CustomPanelOptions();
        }

        CustomPanelOptions options = (CustomPanelOptions) this.internal.options;
        options.setMakeItBeautiful(true);

        this.internal.options = options;
        return this;
    }

    public CustomPanelBuilder id(Integer id) {
        this.internal.id = id;
        return this;
    }

    public CustomPanelBuilder targets(com.grafana.foundation.cog.Builder<List<Dataquery>> targets) {
        this.internal.targets = (List)targets.build();
        return this;
    }

    public CustomPanelBuilder withTarget(com.grafana.foundation.cog.Builder<Dataquery> targets) {
        if (this.internal.targets == null) {
            this.internal.targets = new LinkedList();
        }

        this.internal.targets.add((Dataquery)targets.build());
        return this;
    }

    public CustomPanelBuilder title(String title) {
        this.internal.title = title;
        return this;
    }

    public CustomPanelBuilder description(String description) {
        this.internal.description = description;
        return this;
    }

    public CustomPanelBuilder transparent(Boolean transparent) {
        this.internal.transparent = transparent;
        return this;
    }

    public CustomPanelBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }

    public CustomPanelBuilder gridPos(GridPos gridPos) {
        this.internal.gridPos = gridPos;
        return this;
    }

    public CustomPanelBuilder height(Integer h) {
        if (h <= 0) {
            throw new IllegalArgumentException("h must be > 0");
        } else {
            if (this.internal.gridPos == null) {
                this.internal.gridPos = new GridPos();
            }

            this.internal.gridPos.h = h;
            return this;
        }
    }

    public CustomPanelBuilder span(Integer w) {
        if (w <= 0) {
            throw new IllegalArgumentException("w must be > 0");
        } else if (w > 24) {
            throw new IllegalArgumentException("w must be <= 24");
        } else {
            if (this.internal.gridPos == null) {
                this.internal.gridPos = new GridPos();
            }

            this.internal.gridPos.w = w;
            return this;
        }
    }

    public CustomPanelBuilder links(com.grafana.foundation.cog.Builder<List<DashboardLink>> links) {
        this.internal.links = (List)links.build();
        return this;
    }

    public CustomPanelBuilder repeat(String repeat) {
        this.internal.repeat = repeat;
        return this;
    }

    public CustomPanelBuilder repeatDirection(PanelRepeatDirection repeatDirection) {
        this.internal.repeatDirection = repeatDirection;
        return this;
    }

    public CustomPanelBuilder maxPerRow(Double maxPerRow) {
        this.internal.maxPerRow = maxPerRow;
        return this;
    }

    public CustomPanelBuilder maxDataPoints(Double maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }

    public CustomPanelBuilder transformations(List<DataTransformerConfig> transformations) {
        this.internal.transformations = transformations;
        return this;
    }

    public CustomPanelBuilder withTransformation(DataTransformerConfig transformations) {
        if (this.internal.transformations == null) {
            this.internal.transformations = new LinkedList();
        }

        this.internal.transformations.add(transformations);
        return this;
    }

    public CustomPanelBuilder interval(String interval) {
        this.internal.interval = interval;
        return this;
    }

    public CustomPanelBuilder timeFrom(String timeFrom) {
        this.internal.timeFrom = timeFrom;
        return this;
    }

    public CustomPanelBuilder timeShift(String timeShift) {
        this.internal.timeShift = timeShift;
        return this;
    }

    public CustomPanelBuilder hideTimeOverride(Boolean hideTimeOverride) {
        this.internal.hideTimeOverride = hideTimeOverride;
        return this;
    }

    public CustomPanelBuilder libraryPanel(LibraryPanelRef libraryPanel) {
        this.internal.libraryPanel = libraryPanel;
        return this;
    }

    public CustomPanelBuilder cacheTimeout(String cacheTimeout) {
        this.internal.cacheTimeout = cacheTimeout;
        return this;
    }

    public CustomPanelBuilder queryCachingTTL(Double queryCachingTTL) {
        this.internal.queryCachingTTL = queryCachingTTL;
        return this;
    }

    public CustomPanelBuilder displayName(String displayName) {
        if (this.internal.fieldConfig == null) {
            this.internal.fieldConfig = new FieldConfigSource();
        }

        if (this.internal.fieldConfig.defaults == null) {
            this.internal.fieldConfig.defaults = new FieldConfig();
        }

        this.internal.fieldConfig.defaults.displayName = displayName;
        return this;
    }

    public CustomPanelBuilder unit(String unit) {
        if (this.internal.fieldConfig == null) {
            this.internal.fieldConfig = new FieldConfigSource();
        }

        if (this.internal.fieldConfig.defaults == null) {
            this.internal.fieldConfig.defaults = new FieldConfig();
        }

        this.internal.fieldConfig.defaults.unit = unit;
        return this;
    }

    public CustomPanelBuilder decimals(Double decimals) {
        if (this.internal.fieldConfig == null) {
            this.internal.fieldConfig = new FieldConfigSource();
        }

        if (this.internal.fieldConfig.defaults == null) {
            this.internal.fieldConfig.defaults = new FieldConfig();
        }

        this.internal.fieldConfig.defaults.decimals = decimals;
        return this;
    }

    public CustomPanelBuilder min(Double min) {
        if (this.internal.fieldConfig == null) {
            this.internal.fieldConfig = new FieldConfigSource();
        }

        if (this.internal.fieldConfig.defaults == null) {
            this.internal.fieldConfig.defaults = new FieldConfig();
        }

        this.internal.fieldConfig.defaults.min = min;
        return this;
    }

    public CustomPanelBuilder max(Double max) {
        if (this.internal.fieldConfig == null) {
            this.internal.fieldConfig = new FieldConfigSource();
        }

        if (this.internal.fieldConfig.defaults == null) {
            this.internal.fieldConfig.defaults = new FieldConfig();
        }

        this.internal.fieldConfig.defaults.max = max;
        return this;
    }

    public CustomPanelBuilder mappings(List<ValueMapping> mappings) {
        if (this.internal.fieldConfig == null) {
            this.internal.fieldConfig = new FieldConfigSource();
        }

        if (this.internal.fieldConfig.defaults == null) {
            this.internal.fieldConfig.defaults = new FieldConfig();
        }

        this.internal.fieldConfig.defaults.mappings = mappings;
        return this;
    }

    public CustomPanelBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds) {
        if (this.internal.fieldConfig == null) {
            this.internal.fieldConfig = new FieldConfigSource();
        }

        if (this.internal.fieldConfig.defaults == null) {
            this.internal.fieldConfig.defaults = new FieldConfig();
        }

        this.internal.fieldConfig.defaults.thresholds = (ThresholdsConfig)thresholds.build();
        return this;
    }

    public CustomPanelBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color) {
        if (this.internal.fieldConfig == null) {
            this.internal.fieldConfig = new FieldConfigSource();
        }

        if (this.internal.fieldConfig.defaults == null) {
            this.internal.fieldConfig.defaults = new FieldConfig();
        }

        this.internal.fieldConfig.defaults.color = (FieldColor)color.build();
        return this;
    }

    public CustomPanelBuilder noValue(String noValue) {
        if (this.internal.fieldConfig == null) {
            this.internal.fieldConfig = new FieldConfigSource();
        }

        if (this.internal.fieldConfig.defaults == null) {
            this.internal.fieldConfig.defaults = new FieldConfig();
        }

        this.internal.fieldConfig.defaults.noValue = noValue;
        return this;
    }

    public CustomPanelBuilder overrides(com.grafana.foundation.cog.Builder<List<DashboardFieldConfigSourceOverrides>> overrides) {
        if (this.internal.fieldConfig == null) {
            this.internal.fieldConfig = new FieldConfigSource();
        }

        this.internal.fieldConfig.overrides = (List)overrides.build();
        return this;
    }

    public CustomPanelBuilder withOverride(com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides> overrides) {
        if (this.internal.fieldConfig == null) {
            this.internal.fieldConfig = new FieldConfigSource();
        }

        if (this.internal.fieldConfig.overrides == null) {
            this.internal.fieldConfig.overrides = new LinkedList();
        }

        this.internal.fieldConfig.overrides.add((DashboardFieldConfigSourceOverrides)overrides.build());
        return this;
    }

    public Panel build() {
        return this.internal;
    }
}
