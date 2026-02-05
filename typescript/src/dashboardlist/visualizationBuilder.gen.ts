// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as dashboardlist from '../dashboardlist';

export class VisualizationBuilder implements cog.Builder<dashboardv2beta1.VizConfigKind> {
    protected readonly internal: dashboardv2beta1.VizConfigKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultVizConfigKind();
        this.internal.kind = "VizConfig";
        this.internal.group = "dashlist";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.VizConfigKind {
        return this.internal;
    }

    // The display value for this field.  This supports template variables blank is auto
    displayName(displayName: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.displayName = displayName;
        return this;
    }

    // This can be used by data sources that return and explicit naming structure for values and labels
    // When this property is configured, this value is used rather than the default naming strategy.
    displayNameFromDS(displayNameFromDS: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.displayNameFromDS = displayNameFromDS;
        return this;
    }

    // Human readable field metadata
    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.description = description;
        return this;
    }

    // An explicit path to the field in the datasource.  When the frame meta includes a path,
    // This will default to `${frame.meta.path}/${field.name}
    // 
    // When defined, this value can be used as an identifier within the datasource scope, and
    // may be used to update the results
    path(path: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.path = path;
        return this;
    }

    // Unit a field should use. The unit you select is applied to all fields except time.
    // You can use the units ID availables in Grafana or a custom unit.
    // Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
    // As custom unit, you can use the following formats:
    // `suffix:<suffix>` for custom unit that should go after value.
    // `prefix:<prefix>` for custom unit that should go before value.
    // `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
    // `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
    // `count:<unit>` for a custom count unit.
    // `currency:<unit>` for custom a currency unit.
    unit(unit: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.unit = unit;
        return this;
    }

    // Specify the number of decimals Grafana includes in the rendered value.
    // If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
    // For example 1.1234 will display as 1.12 and 100.456 will display as 100.
    // To display all decimals, set the unit to `String`.
    decimals(decimals: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.decimals = decimals;
        return this;
    }

    // The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
    min(min: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.min = min;
        return this;
    }

    // The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
    max(max: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.max = max;
        return this;
    }

    // Convert input values into a display string
    mappings(mappings: dashboardv2beta1.ValueMapping[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.mappings = mappings;
        return this;
    }

    // Map numeric values to states
    thresholds(thresholds: cog.Builder<dashboardv2beta1.ThresholdsConfig>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        const thresholdsResource = thresholds.build();
        this.internal.spec.fieldConfig.defaults.thresholds = thresholdsResource;
        return this;
    }

    // Panel color configuration
    colorScheme(color: cog.Builder<dashboardv2beta1.FieldColor>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        const colorResource = color.build();
        this.internal.spec.fieldConfig.defaults.color = colorResource;
        return this;
    }

    // The behavior when clicking on a result
    dataLinks(links: any[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.links = links;
        return this;
    }

    // Define interactive HTTP requests that can be triggered from data visualizations.
    actions(actions: cog.Builder<dashboardv2beta1.Action>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        const actionsResources = actions.map(builder1 => builder1.build());
        this.internal.spec.fieldConfig.defaults.actions = actionsResources;
        return this;
    }

    // Alternative to empty string
    noValue(noValue: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.noValue = noValue;
        return this;
    }

    // Overrides are the options applied to specific fields overriding the defaults.
    overrides(overrides: {
	matcher: dashboardv2beta1.MatcherConfig;
	properties: dashboardv2beta1.DynamicConfigValue[];
}[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        this.internal.spec.fieldConfig.overrides = overrides;
        return this;
    }

    // Overrides are the options applied to specific fields overriding the defaults.
    override(override: {
	matcher: dashboardv2beta1.MatcherConfig;
	properties: dashboardv2beta1.DynamicConfigValue[];
}): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.overrides) {
            this.internal.spec.fieldConfig.overrides = [];
        }
        this.internal.spec.fieldConfig.overrides.push(override);
        return this;
    }

    // Adds override rules for a specific field, referred to by its name.
    overrideByName(name: string,properties: dashboardv2beta1.DynamicConfigValue[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.overrides) {
            this.internal.spec.fieldConfig.overrides = [];
        }
        this.internal.spec.fieldConfig.overrides.push({
        matcher: {
        id: "byName",
        options: name,
    },
        properties: properties,
    });
        return this;
    }

    // Adds override rules for the fields whose name match the given regexp.
    overrideByRegexp(regexp: string,properties: dashboardv2beta1.DynamicConfigValue[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.overrides) {
            this.internal.spec.fieldConfig.overrides = [];
        }
        this.internal.spec.fieldConfig.overrides.push({
        matcher: {
        id: "byRegexp",
        options: regexp,
    },
        properties: properties,
    });
        return this;
    }

    // Adds override rules for all the fields of the given type.
    overrideByFieldType(fieldType: string,properties: dashboardv2beta1.DynamicConfigValue[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.overrides) {
            this.internal.spec.fieldConfig.overrides = [];
        }
        this.internal.spec.fieldConfig.overrides.push({
        matcher: {
        id: "byType",
        options: fieldType,
    },
        properties: properties,
    });
        return this;
    }

    overrideByQuery(queryRefId: string,properties: dashboardv2beta1.DynamicConfigValue[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.overrides) {
            this.internal.spec.fieldConfig.overrides = [];
        }
        this.internal.spec.fieldConfig.overrides.push({
        matcher: {
        id: "byFrameRefID",
        options: queryRefId,
    },
        properties: properties,
    });
        return this;
    }

    keepTime(keepTime: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = dashboardlist.defaultOptions();
        }
        this.internal.spec.options.keepTime = keepTime;
        return this;
    }

    includeVars(includeVars: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = dashboardlist.defaultOptions();
        }
        this.internal.spec.options.includeVars = includeVars;
        return this;
    }

    showStarred(showStarred: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = dashboardlist.defaultOptions();
        }
        this.internal.spec.options.showStarred = showStarred;
        return this;
    }

    showRecentlyViewed(showRecentlyViewed: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = dashboardlist.defaultOptions();
        }
        this.internal.spec.options.showRecentlyViewed = showRecentlyViewed;
        return this;
    }

    showSearch(showSearch: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = dashboardlist.defaultOptions();
        }
        this.internal.spec.options.showSearch = showSearch;
        return this;
    }

    showHeadings(showHeadings: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = dashboardlist.defaultOptions();
        }
        this.internal.spec.options.showHeadings = showHeadings;
        return this;
    }

    showFolderNames(showFolderNames: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = dashboardlist.defaultOptions();
        }
        this.internal.spec.options.showFolderNames = showFolderNames;
        return this;
    }

    maxItems(maxItems: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = dashboardlist.defaultOptions();
        }
        this.internal.spec.options.maxItems = maxItems;
        return this;
    }

    query(query: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = dashboardlist.defaultOptions();
        }
        this.internal.spec.options.query = query;
        return this;
    }

    tags(tags: string[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = dashboardlist.defaultOptions();
        }
        this.internal.spec.options.tags = tags;
        return this;
    }

    // folderId is deprecated, and migrated to folderUid on panel init
    folderId(folderId: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = dashboardlist.defaultOptions();
        }
        this.internal.spec.options.folderId = folderId;
        return this;
    }

    folderUID(folderUID: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = dashboardlist.defaultOptions();
        }
        this.internal.spec.options.folderUID = folderUID;
        return this;
    }
}

