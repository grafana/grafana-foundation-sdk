// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Links with references to other dashboards or external resources
export class DashboardLinkBuilder implements cog.Builder<dashboard.DashboardLink> {
    protected readonly internal: dashboard.DashboardLink;

    constructor(title: string) {
        this.internal = dashboard.defaultDashboardLink();
        this.internal.title = title;
    }

    /**
     * Builds the object.
     */
    build(): dashboard.DashboardLink {
        return this.internal;
    }

    // Title to display with the link
    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    // Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
    type(type: dashboard.DashboardLinkType): this {
        this.internal.type = type;
        return this;
    }

    // Icon name to be displayed with the link
    icon(icon: string): this {
        this.internal.icon = icon;
        return this;
    }

    // Tooltip to display when the user hovers their mouse over it
    tooltip(tooltip: string): this {
        this.internal.tooltip = tooltip;
        return this;
    }

    // Link URL. Only required/valid if the type is link
    url(url: string): this {
        this.internal.url = url;
        return this;
    }

    // List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
    tags(tags: string[]): this {
        this.internal.tags = tags;
        return this;
    }

    // If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
    asDropdown(asDropdown: boolean): this {
        this.internal.asDropdown = asDropdown;
        return this;
    }

    // If true, the link will be opened in a new tab
    targetBlank(targetBlank: boolean): this {
        this.internal.targetBlank = targetBlank;
        return this;
    }

    // If true, includes current template variables values in the link as query params
    includeVars(includeVars: boolean): this {
        this.internal.includeVars = includeVars;
        return this;
    }

    // If true, includes current time range in the link as query params
    keepTime(keepTime: boolean): this {
        this.internal.keepTime = keepTime;
        return this;
    }
}
