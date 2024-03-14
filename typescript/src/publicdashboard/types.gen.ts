// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export interface PublicDashboard {
	// Unique public dashboard identifier
	uid: string;
	// Dashboard unique identifier referenced by this public dashboard
	dashboardUid: string;
	// Unique public access token
	accessToken?: string;
	// Flag that indicates if the public dashboard is enabled
	isEnabled: boolean;
	// Flag that indicates if annotations are enabled
	annotationsEnabled: boolean;
	// Flag that indicates if the time range picker is enabled
	timeSelectionEnabled: boolean;
}

export const defaultPublicDashboard = (): PublicDashboard => ({
	uid: "",
	dashboardUid: "",
	isEnabled: false,
	annotationsEnabled: false,
	timeSelectionEnabled: false,
});

