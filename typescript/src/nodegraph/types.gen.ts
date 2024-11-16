// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export interface ArcOption {
	// Field from which to get the value. Values should be less than 1, representing fraction of a circle.
	field?: string;
	// The color of the arc.
	color?: string;
}

export const defaultArcOption = (): ArcOption => ({
});

export interface NodeOptions {
	// Unit for the main stat to override what ever is set in the data frame.
	mainStatUnit?: string;
	// Unit for the secondary stat to override what ever is set in the data frame.
	secondaryStatUnit?: string;
	// Define which fields are shown as part of the node arc (colored circle around the node).
	arcs?: ArcOption[];
}

export const defaultNodeOptions = (): NodeOptions => ({
});

export interface EdgeOptions {
	// Unit for the main stat to override what ever is set in the data frame.
	mainStatUnit?: string;
	// Unit for the secondary stat to override what ever is set in the data frame.
	secondaryStatUnit?: string;
}

export const defaultEdgeOptions = (): EdgeOptions => ({
});

export enum ZoomMode {
	Cooperative = "cooperative",
	Greedy = "greedy",
}

export const defaultZoomMode = (): ZoomMode => (ZoomMode.Cooperative);

export interface Options {
	nodes?: NodeOptions;
	edges?: EdgeOptions;
	// How to handle zoom/scroll events in the node graph
	zoomMode?: ZoomMode;
}

export const defaultOptions = (): Options => ({
});

