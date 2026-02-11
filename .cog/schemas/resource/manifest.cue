package resource

DashboardV2Beta1: "dashboard.grafana.app/v2beta1"
DashboardV2Alpha1: "dashboard.grafana.app/v2alpha1"
DashboardKind: "Dashboard"

Manifest: {
    apiVersion: string
    kind: string
    metadata: #Metadata
    spec: _
}

#Metadata: {
    name: string
    namespace?: string
    labels?: [string]: string
    annotations?: [string]: string
    uid?: string
    resourceVersion?: string
    generation?: int64
    creationTimestamp?: string
    updateTimestamp?: string
    deletionTimestamp?: string
}
