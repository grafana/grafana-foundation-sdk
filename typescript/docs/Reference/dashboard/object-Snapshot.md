---
title: <span class="badge object-type-interface"></span> Snapshot
---
# <span class="badge object-type-interface"></span> Snapshot

A dashboard snapshot shares an interactive dashboard publicly.

It is a read-only version of a dashboard, and is not editable.

It is possible to create a snapshot of a snapshot.

Grafana strips away all sensitive information from the dashboard.

Sensitive information stripped: queries (metric, template,annotation) and panel links.

## Definition

```typescript
export interface Snapshot {
	// Time when the snapshot was created
	created: string;
	// Time when the snapshot expires, default is never to expire
	expires: string;
	// Is the snapshot saved in an external grafana instance
	external: boolean;
	// external url, if snapshot was shared in external grafana instance
	externalUrl: string;
	// Unique identifier of the snapshot
	id: number;
	// Optional, defined the unique key of the snapshot, required if external is true
	key: string;
	// Optional, name of the snapshot
	name: string;
	// org id of the snapshot
	orgId: number;
	// last time when the snapshot was updated
	updated: string;
	// url of the snapshot, if snapshot was shared internally
	url?: string;
	// user id of the snapshot creator
	userId: number;
	dashboard?: dashboard.Dashboard;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [SnapshotBuilder](./builder-SnapshotBuilder.md)
