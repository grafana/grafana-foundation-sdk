// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

export class DashboardBuilder implements cog.Builder<dashboard.Dashboard> {
    protected readonly internal: dashboard.Dashboard;
    private currentY: number = 0;
    private currentX: number = 0;
    private lastPanelHeight: number = 0;

    constructor(title: string) {
        this.internal = dashboard.defaultDashboard();
        this.internal.title = title;
    }

    build(): dashboard.Dashboard {
        return this.internal;
    }

    // Unique numeric identifier for the dashboard.
    // `id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.
    id(id: number | null): this {
        this.internal.id = id;
        return this;
    }

    // Unique dashboard identifier that can be generated by anyone. string (8-40)
    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }

    // Title of dashboard.
    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    // Description of dashboard.
    description(description: string): this {
        this.internal.description = description;
        return this;
    }

    // This property should only be used in dashboards defined by plugins.  It is a quick check
    // to see if the version has changed since the last time.
    revision(revision: number): this {
        this.internal.revision = revision;
        return this;
    }

    // ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal
    gnetId(gnetId: string): this {
        this.internal.gnetId = gnetId;
        return this;
    }

    // Tags associated with dashboard.
    tags(tags: string[]): this {
        this.internal.tags = tags;
        return this;
    }

    // Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
    timezone(timezone: string): this {
        this.internal.timezone = timezone;
        return this;
    }

    // Whether a dashboard is editable or not.
    editable(): this {
        this.internal.editable = true;
        return this;
    }

    // Whether a dashboard is editable or not.
    readonly(): this {
        this.internal.editable = false;
        return this;
    }

    // Configuration of dashboard cursor sync behavior.
    // Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).
    tooltip(graphTooltip: dashboard.DashboardCursorSync): this {
        this.internal.graphTooltip = graphTooltip;
        return this;
    }

    // Time range for dashboard.
    // Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.
    time(time: {
	from: string;
	to: string;
}): this {
        this.internal.time = time;
        return this;
    }

    // Configuration of the time picker shown at the top of a dashboard.
    timepicker(timepicker: cog.Builder<dashboard.TimePickerConfig>): this {
        const timepickerResource = timepicker.build();
        this.internal.timepicker = timepickerResource;
        return this;
    }

    // The month that the fiscal year starts on.  0 = January, 11 = December
    fiscalYearStartMonth(fiscalYearStartMonth: number): this {
        if (!(fiscalYearStartMonth < 12)) {
            throw new Error("fiscalYearStartMonth must be < 12");
        }
        this.internal.fiscalYearStartMonth = fiscalYearStartMonth;
        return this;
    }

    // When set to true, the dashboard will redraw panels at an interval matching the pixel width.
    // This will keep data "moving left" regardless of the query refresh rate. This setting helps
    // avoid dashboards presenting stale live data
    liveNow(liveNow: boolean): this {
        this.internal.liveNow = liveNow;
        return this;
    }

    // Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
    weekStart(weekStart: string): this {
        this.internal.weekStart = weekStart;
        return this;
    }

    // Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
    refresh(refresh: string | false): this {
        this.internal.refresh = refresh;
        return this;
    }

    // Version of the dashboard, incremented each time the dashboard is updated.
    version(version: number): this {
        this.internal.version = version;
        return this;
    }

    withPanel(panel: cog.Builder<dashboard.Panel>): this {
        if (!this.internal.panels) {
            this.internal.panels = [];
        }
        const panelResource = panel.build();

		if (!panelResource.gridPos) {
			panelResource.gridPos = dashboard.defaultGridPos();
		}

		// The panel either has no position set, or it is the first panel of the dashboard.
		// In that case, we position it on the grid
		if (panelResource.gridPos.x == 0 && panelResource.gridPos.y == 0) {
			panelResource.gridPos.x = this.currentX;
			panelResource.gridPos.y = this.currentY;
		}
        this.internal.panels.push(panelResource);

		// Prepare the coordinates for the next panel
		this.currentX += panelResource.gridPos.w;
		this.lastPanelHeight = Math.max(this.lastPanelHeight, panelResource.gridPos.h);

		// Check for grid width overflow?
		if (this.currentX >= 24) {
			this.currentX = 0;
			this.currentY += this.lastPanelHeight;
			this.lastPanelHeight = 0;
		}
        return this;
    }

    withRow(rowPanel: cog.Builder<dashboard.RowPanel>): this {
        if (!this.internal.panels) {
            this.internal.panels = [];
        }
        const rowPanelResource = rowPanel.build();

		// Position the row on the grid
		if (!rowPanelResource.gridPos || (rowPanelResource.gridPos.x == 0 && rowPanelResource.gridPos.y == 0)) {
			rowPanelResource.gridPos = {
				x: 0, // beginning of the line
				y: this.currentY + this.lastPanelHeight,

				h: 1,
				w: 24, // full width
			};
		}
        this.internal.panels.push(rowPanelResource);

		// Reset the state for the next row
		this.currentX = 0;
		this.currentY = rowPanelResource.gridPos.y + 1;
		this.lastPanelHeight = 0;

		// Position the row's panels on the grid
		rowPanelResource.panels.forEach(panel => {
			if (!panel.gridPos) {
				panel.gridPos = dashboard.defaultGridPos();
			}

			// The panel either has no position set, or it is the first panel of the dashboard.
			// In that case, we position it on the grid
			if (panel.gridPos.x == 0 && panel.gridPos.y == 0) {
				panel.gridPos.x = this.currentX;
				panel.gridPos.y = this.currentY;
			}

			// Prepare the coordinates for the next panel
			this.currentX += panel.gridPos.w;
			this.lastPanelHeight = Math.max(this.lastPanelHeight, panel.gridPos.h);

			// Check for grid width overflow?
			if (this.currentX >= 24) {
				this.currentX = 0;
				this.currentY += this.lastPanelHeight;
				this.lastPanelHeight = 0;
			}
		});
        return this;
    }

    // Configured template variables
    variables(list: cog.Builder<dashboard.VariableModel>[]): this {
        if (!this.internal.templating) {
            this.internal.templating = {
};
        }
        const listResources = list.map(builder1 => builder1.build());
        this.internal.templating.list = listResources;
        return this;
    }

    // Configured template variables
    withVariable(list: cog.Builder<dashboard.VariableModel>): this {
        if (!this.internal.templating) {
            this.internal.templating = {
};
        }
        if (!this.internal.templating.list) {
            this.internal.templating.list = [];
        }
        const listResource = list.build();
        this.internal.templating.list.push(listResource);
        return this;
    }

    // Contains the list of annotations that are associated with the dashboard.
    // Annotations are used to overlay event markers and overlay event tags on graphs.
    // Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
    // See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
    annotations(list: cog.Builder<dashboard.AnnotationQuery>[]): this {
        if (!this.internal.annotations) {
            this.internal.annotations = dashboard.defaultAnnotationContainer();
        }
        const listResources = list.map(builder1 => builder1.build());
        this.internal.annotations.list = listResources;
        return this;
    }

    // Contains the list of annotations that are associated with the dashboard.
    // Annotations are used to overlay event markers and overlay event tags on graphs.
    // Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
    // See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
    annotation(list: cog.Builder<dashboard.AnnotationQuery>): this {
        if (!this.internal.annotations) {
            this.internal.annotations = dashboard.defaultAnnotationContainer();
        }
        if (!this.internal.annotations.list) {
            this.internal.annotations.list = [];
        }
        const listResource = list.build();
        this.internal.annotations.list.push(listResource);
        return this;
    }

    // Links with references to other dashboards or external websites.
    links(links: cog.Builder<dashboard.DashboardLink>[]): this {
        const linksResources = links.map(builder1 => builder1.build());
        this.internal.links = linksResources;
        return this;
    }

    // Links with references to other dashboards or external websites.
    link(links: cog.Builder<dashboard.DashboardLink>): this {
        if (!this.internal.links) {
            this.internal.links = [];
        }
        const linksResource = links.build();
        this.internal.links.push(linksResource);
        return this;
    }

    // Snapshot options. They are present only if the dashboard is a snapshot.
    snapshot(snapshot: cog.Builder<dashboard.Snapshot>): this {
        const snapshotResource = snapshot.build();
        this.internal.snapshot = snapshotResource;
        return this;
    }
}
