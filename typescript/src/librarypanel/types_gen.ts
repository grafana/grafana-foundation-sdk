// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';


export interface LibraryPanel {
	// Folder UID
	folderUid?: string;
	// Library element UID
	uid: string;
	// Panel name (also saved in the model)
	name: string;
	// Panel description
	description?: string;
	// The panel type (from inside the model)
	type: string;
	// Dashboard version when this was saved (zero if unknown)
	schemaVersion?: number;
	// panel version, incremented each time the dashboard is updated.
	version: number;
	// TODO: should be the same panel schema defined in dashboard
	// Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
	model: {
		// The panel plugin type id. This is used to find the plugin to display the panel.
		type: string;
		// The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
		pluginVersion?: string;
		// Depends on the panel plugin. See the plugin documentation for details.
		targets?: cog.Dataquery[];
		// Panel title.
		title?: string;
		// Panel description.
		description?: string;
		// Whether to display the panel without a background.
		transparent?: boolean;
		// The datasource used in all targets.
		datasource?: dashboard.DataSourceRef;
		// Panel links.
		links?: dashboard.DashboardLink[];
		// Name of template variable to repeat for.
		repeat?: string;
		// Direction to repeat in if 'repeat' is set.
		// `h` for horizontal, `v` for vertical.
		repeatDirection?: "h" | "v";
		// Option for repeated panels that controls max items per row
		// Only relevant for horizontally repeated panels
		maxPerRow?: number;
		// The maximum number of data points that the panel queries are retrieving.
		maxDataPoints?: number;
		// List of transformations that are applied to the panel data before rendering.
		// When there are multiple transformations, Grafana applies them in the order they are listed.
		// Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
		transformations?: dashboard.DataTransformerConfig[];
		// The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
		// This value must be formatted as a number followed by a valid time
		// identifier like: "40s", "3d", etc.
		// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
		interval?: string;
		// Overrides the relative time range for individual panels,
		// which causes them to be different than what is selected in
		// the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
		// time periods or days on the same dashboard.
		// The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
		// `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
		// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
		// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
		timeFrom?: string;
		// Overrides the time range for individual panels by shifting its start and end relative to the time picker.
		// For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
		// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
		// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
		timeShift?: string;
		// Controls if the timeFrom or timeShift overrides are shown in the panel header
		hideTimeOverride?: boolean;
		// It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
		options?: any;
		// Field options allow you to change how the data is displayed in your visualizations.
		fieldConfig?: dashboard.FieldConfigSource;
	};
	// Object storage metadata
	meta?: LibraryElementDTOMeta;
}

export const defaultLibraryPanel = (): LibraryPanel => ({
	uid: "",
	name: "",
	type: "",
	version: 0,
	model: {
	type: "",
	transparent: false,
	repeatDirection: "h",
},
});

export interface LibraryElementDTOMetaUser {
	id: number;
	name: string;
	avatarUrl: string;
}

export const defaultLibraryElementDTOMetaUser = (): LibraryElementDTOMetaUser => ({
	id: 0,
	name: "",
	avatarUrl: "",
});

export interface LibraryElementDTOMeta {
	folderName: string;
	folderUid: string;
	connectedDashboards: number;
	created: string;
	updated: string;
	createdBy: LibraryElementDTOMetaUser;
	updatedBy: LibraryElementDTOMetaUser;
}

export const defaultLibraryElementDTOMeta = (): LibraryElementDTOMeta => ({
	folderName: "",
	folderUid: "",
	connectedDashboards: 0,
	created: "",
	updated: "",
	createdBy: defaultLibraryElementDTOMetaUser(),
	updatedBy: defaultLibraryElementDTOMetaUser(),
});

