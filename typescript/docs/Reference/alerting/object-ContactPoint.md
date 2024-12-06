---
title: <span class="badge object-type-interface"></span> ContactPoint
---
# <span class="badge object-type-interface"></span> ContactPoint

EmbeddedContactPoint is the contact point type that is used

by grafanas embedded alertmanager implementation.

## Definition

```typescript
export interface ContactPoint {
	disableResolveMessage?: boolean;
	// Name is used as grouping key in the UI. Contact points with the
	// same name will be grouped in the UI.
	name?: string;
	provenance?: string;
	settings: alerting.Json;
	type: "alertmanager" | " dingding" | " discord" | " email" | " googlechat" | " kafka" | " line" | " opsgenie" | " pagerduty" | " pushover" | " sensugo" | " slack" | " teams" | " telegram" | " threema" | " victorops" | " webhook" | " wecom";
	// UID is the unique identifier of the contact point. The UID can be
	// set by the user.
	uid?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [ContactPointBuilder](./builder-ContactPointBuilder.md)
