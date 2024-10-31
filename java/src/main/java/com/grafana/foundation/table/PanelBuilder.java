// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.table;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import com.grafana.foundation.common.TableSortByFieldState;
import com.grafana.foundation.common.TableFooterOptions;
import com.grafana.foundation.common.TableCellHeight;
import com.grafana.foundation.dashboard.Panel;
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
    this.internal.type = "table";
        this.transparent(false);
        this.height(9);
        this.span(12);
        this.frameIndex(0.0);
        this.showHeader(true);
        this.showTypeIcons(false);
        TableFooterOptions.Builder tableFooterOptionsResource = new TableFooterOptions.Builder();
        tableFooterOptionsResource.show(false);
        tableFooterOptionsResource.countRows(false);
        this.footer(tableFooterOptionsResource);
        this.cellHeight(TableCellHeight.SM);
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
    public PanelBuilder frameIndex(Double frameIndex) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.table.Options();
		}
        com.grafana.foundation.table.Options optionsResource = (com.grafana.foundation.table.Options) this.internal.options;
        optionsResource.frameIndex = frameIndex;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder showHeader(Boolean showHeader) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.table.Options();
		}
        com.grafana.foundation.table.Options optionsResource = (com.grafana.foundation.table.Options) this.internal.options;
        optionsResource.showHeader = showHeader;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder showTypeIcons(Boolean showTypeIcons) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.table.Options();
		}
        com.grafana.foundation.table.Options optionsResource = (com.grafana.foundation.table.Options) this.internal.options;
        optionsResource.showTypeIcons = showTypeIcons;
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder sortBy(com.grafana.foundation.cog.Builder<List<TableSortByFieldState>> sortBy) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.table.Options();
		}
        com.grafana.foundation.table.Options optionsResource = (com.grafana.foundation.table.Options) this.internal.options;
        optionsResource.sortBy = sortBy.build();
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder footer(com.grafana.foundation.cog.Builder<TableFooterOptions> footer) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.table.Options();
		}
        com.grafana.foundation.table.Options optionsResource = (com.grafana.foundation.table.Options) this.internal.options;
        optionsResource.footer = footer.build();
    this.internal.options = optionsResource;
        return this;
    }
    public PanelBuilder cellHeight(TableCellHeight cellHeight) {
		if (this.internal.options == null) {
			this.internal.options = new com.grafana.foundation.table.Options();
		}
        com.grafana.foundation.table.Options optionsResource = (com.grafana.foundation.table.Options) this.internal.options;
        optionsResource.cellHeight = cellHeight;
    this.internal.options = optionsResource;
        return this;
    }
    
    public Panel build() {
        return this.internal;
    }
}
