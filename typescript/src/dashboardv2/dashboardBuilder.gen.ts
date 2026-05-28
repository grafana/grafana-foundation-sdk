// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';
import * as resource from '../resource';

export class DashboardBuilder implements cog.Builder<dashboardv2.Dashboard> {
    protected readonly internal: dashboardv2.Dashboard;

    constructor(title: string) {
        this.internal = dashboardv2.defaultDashboard();
        this.internal.title = title;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.Dashboard {
        return this.internal;
    }

    annotations(annotations: cog.Builder<dashboardv2.AnnotationQueryKind>[]): this {
        const annotationsResources = annotations.map(builder1 => builder1.build());
        this.internal.annotations = annotationsResources;
        return this;
    }

    // Configuration of dashboard cursor sync behavior.
    // "Off" for no shared crosshair or tooltip (default).
    // "Crosshair" for shared crosshair.
    // "Tooltip" for shared crosshair AND shared tooltip.
    cursorSync(cursorSync: dashboardv2.DashboardCursorSync): this {
        this.internal.cursorSync = cursorSync;
        return this;
    }

    // Description of dashboard.
    description(description: string): this {
        this.internal.description = description;
        return this;
    }

    // Whether a dashboard is editable or not.
    editable(editable: boolean): this {
        this.internal.editable = editable;
        return this;
    }

    elements(elements: Record<string, cog.Builder<dashboardv2.Element>>): this {
        const elementsResource = (function() {
            let results1: Record<string, dashboardv2.Element> = {};
            for (const key1 in elements) {
                const val1 = elements[key1];
                results1[key1] = val1.build();
            }
            return results1;
        }());
        this.internal.elements = elementsResource;
        return this;
    }

    element(key: string,element: cog.Builder<dashboardv2.Element>): this {
        if (!this.internal.elements) {
            this.internal.elements = {};
        }
        const elementResource = element.build();
        this.internal.elements[key] = elementResource;
        return this;
    }

    layout(layout: cog.Builder<dashboardv2.GridLayoutKind> | cog.Builder<dashboardv2.RowsLayoutKind> | cog.Builder<dashboardv2.AutoGridLayoutKind> | cog.Builder<dashboardv2.TabsLayoutKind>): this {
        const layoutResource = layout.build();
        this.internal.layout = layoutResource;
        return this;
    }

    // Links with references to other dashboards or external websites.
    links(links: cog.Builder<dashboardv2.DashboardLink>[]): this {
        const linksResources = links.map(builder1 => builder1.build());
        this.internal.links = linksResources;
        return this;
    }

    // When set to true, the dashboard will redraw panels at an interval matching the pixel width.
    // This will keep data "moving left" regardless of the query refresh rate. This setting helps
    // avoid dashboards presenting stale live data.
    liveNow(liveNow: boolean): this {
        this.internal.liveNow = liveNow;
        return this;
    }

    // When set to true, the dashboard will load all panels in the dashboard when it's loaded.
    preload(preload: boolean): this {
        this.internal.preload = preload;
        return this;
    }

    // Plugins only. The version of the dashboard installed together with the plugin.
    // This is used to determine if the dashboard should be updated when the plugin is updated.
    revision(revision: number): this {
        this.internal.revision = revision;
        return this;
    }

    // Tags associated with dashboard.
    tags(tags: string[]): this {
        this.internal.tags = tags;
        return this;
    }

    timeSettings(timeSettings: cog.Builder<dashboardv2.TimeSettingsSpec>): this {
        const timeSettingsResource = timeSettings.build();
        this.internal.timeSettings = timeSettingsResource;
        return this;
    }

    // Title of dashboard.
    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    // Configured template variables.
    variables(variables: cog.Builder<dashboardv2.VariableKind>[]): this {
        const variablesResources = variables.map(builder1 => builder1.build());
        this.internal.variables = variablesResources;
        return this;
    }

    // Configured template variables.
    variable(variable: cog.Builder<dashboardv2.VariableKind>): this {
        if (!this.internal.variables) {
            this.internal.variables = [];
        }
        const variableResource = variable.build();
        this.internal.variables.push(variableResource);
        return this;
    }

    preferences(preferences: cog.Builder<dashboardv2.Preferences>): this {
        const preferencesResource = preferences.build();
        this.internal.preferences = preferencesResource;
        return this;
    }
}

/**
 * Creates a resource manifest from a Dashboard.
 */
export function manifest(name: string,spec: cog.Builder<dashboardv2.Dashboard>): cog.Builder<resource.Manifest> {
	const builder = new resource.ManifestBuilder();
	builder.apiVersion("dashboard.grafana.app/v2");
	builder.kind("Dashboard");
	builder.metadata(resource.named(name));
	builder.spec(spec.build());

	return builder;
}

