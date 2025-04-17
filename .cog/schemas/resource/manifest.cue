package resource

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