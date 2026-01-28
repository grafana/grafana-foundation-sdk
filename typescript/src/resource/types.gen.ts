// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export const DashboardV2Beta1 = "dashboard.grafana.app/v2beta1";

export const DashboardV2Alpha1 = "dashboard.grafana.app/v2alpha1";

export const DashboardKind = "Dashboard";

export interface Manifest {
	apiVersion: string;
	kind: string;
	metadata: Metadata;
	spec: any;
}

export const defaultManifest = (): Manifest => ({
	apiVersion: "",
	kind: "",
	metadata: defaultMetadata(),
	spec: {},
});

export interface Metadata {
	name: string;
	namespace?: string;
	labels?: Record<string, string>;
	annotations?: Record<string, string>;
	uid?: string;
	resourceVersion?: string;
	generation?: number;
	creationTimestamp?: string;
	updateTimestamp?: string;
	deletionTimestamp?: string;
}

export const defaultMetadata = (): Metadata => ({
	name: "",
});

