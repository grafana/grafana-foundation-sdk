// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as librarypanel from '../librarypanel';
import * as common from '../common';
import * as dashboard from '../dashboard';

// Dashboard panels are the basic visualization building blocks.
export class PanelModelBuilder implements cog.Builder<librarypanel.PanelModel> {
    protected readonly internal: librarypanel.PanelModel;

    constructor() {
        this.internal = librarypanel.defaultPanelModel();
    }

    /**
     * Builds the object.
     */
    build(): librarypanel.PanelModel {
        return this.internal;
    }

    // The panel plugin type id. This is used to find the plugin to display the panel.
    type(type: string): this {
        if (!(type.length >= 1)) {
            throw new Error("type.length must be >= 1");
        }
        this.internal.type = type;
        return this;
    }

    // The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
    pluginVersion(pluginVersion: string): this {
        this.internal.pluginVersion = pluginVersion;
        return this;
    }

    // Depends on the panel plugin. See the plugin documentation for details.
    targets(targets: cog.Builder<cog.Dataquery>[]): this {
        const targetsResources = targets.map(builder1 => builder1.build());
        this.internal.targets = targetsResources;
        return this;
    }

    // Panel title.
    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    // Panel description.
    description(description: string): this {
        this.internal.description = description;
        return this;
    }

    // Whether to display the panel without a background.
    transparent(transparent: boolean): this {
        this.internal.transparent = transparent;
        return this;
    }

    // The datasource used in all targets.
    datasource(datasource: common.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }

    // Panel links.
    links(links: cog.Builder<dashboard.DashboardLink>[]): this {
        const linksResources = links.map(builder1 => builder1.build());
        this.internal.links = linksResources;
        return this;
    }

    // Name of template variable to repeat for.
    repeat(repeat: string): this {
        this.internal.repeat = repeat;
        return this;
    }

    // Direction to repeat in if 'repeat' is set.
    // `h` for horizontal, `v` for vertical.
    repeatDirection(repeatDirection: "h" | "v"): this {
        this.internal.repeatDirection = repeatDirection;
        return this;
    }

    // Option for repeated panels that controls max items per row
    // Only relevant for horizontally repeated panels
    maxPerRow(maxPerRow: number): this {
        this.internal.maxPerRow = maxPerRow;
        return this;
    }

    // The maximum number of data points that the panel queries are retrieving.
    maxDataPoints(maxDataPoints: number): this {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }

    // List of transformations that are applied to the panel data before rendering.
    // When there are multiple transformations, Grafana applies them in the order they are listed.
    // Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
    transformations(transformations: dashboard.DataTransformerConfig[]): this {
        this.internal.transformations = transformations;
        return this;
    }

    // The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
    // This value must be formatted as a number followed by a valid time
    // identifier like: "40s", "3d", etc.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    interval(interval: string): this {
        this.internal.interval = interval;
        return this;
    }

    // Overrides the relative time range for individual panels,
    // which causes them to be different than what is selected in
    // the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
    // time periods or days on the same dashboard.
    // The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
    // `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
    // Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    timeFrom(timeFrom: string): this {
        this.internal.timeFrom = timeFrom;
        return this;
    }

    // Overrides the time range for individual panels by shifting its start and end relative to the time picker.
    // For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
    // Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    timeShift(timeShift: string): this {
        this.internal.timeShift = timeShift;
        return this;
    }

    // Controls if the timeFrom or timeShift overrides are shown in the panel header
    hideTimeOverride(hideTimeOverride: boolean): this {
        this.internal.hideTimeOverride = hideTimeOverride;
        return this;
    }

    // Compare the current time range with a previous period
    // For example "1d" to compare current period but shifted back 1 day
    timeCompare(timeCompare: string): this {
        this.internal.timeCompare = timeCompare;
        return this;
    }

    // Sets panel queries cache timeout.
    cacheTimeout(cacheTimeout: string): this {
        this.internal.cacheTimeout = cacheTimeout;
        return this;
    }

    // Overrides the data source configured time-to-live for a query cache item in milliseconds
    queryCachingTTL(queryCachingTTL: number): this {
        this.internal.queryCachingTTL = queryCachingTTL;
        return this;
    }

    // It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
    options(options: any): this {
        this.internal.options = options;
        return this;
    }

    // Field options allow you to change how the data is displayed in your visualizations.
    fieldConfig(fieldConfig: dashboard.FieldConfigSource): this {
        this.internal.fieldConfig = fieldConfig;
        return this;
    }
}

