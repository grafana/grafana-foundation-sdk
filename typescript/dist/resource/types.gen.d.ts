export declare const DashboardV2Beta1 = "dashboard.grafana.app/v2beta1";
export declare const DashboardV2Alpha1 = "dashboard.grafana.app/v2alpha1";
export declare const DashboardKind = "Dashboard";
export interface Manifest {
    apiVersion: string;
    kind: string;
    metadata: Metadata;
    spec: any;
}
export declare const defaultManifest: () => Manifest;
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
export declare const defaultMetadata: () => Metadata;
