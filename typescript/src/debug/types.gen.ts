// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export interface UpdateConfig {
	render: boolean;
	dataChanged: boolean;
	schemaChanged: boolean;
}

export const defaultUpdateConfig = (): UpdateConfig => ({
	render: false,
	dataChanged: false,
	schemaChanged: false,
});

export enum DebugMode {
	Render = "render",
	Events = "events",
	Cursor = "cursor",
	State = "State",
	ThrowError = "ThrowError",
}

export const defaultDebugMode = (): DebugMode => (DebugMode.Render);

export interface Options {
	mode: DebugMode;
	counters?: UpdateConfig;
}

export const defaultOptions = (): Options => ({
	mode: DebugMode.Render,
});

