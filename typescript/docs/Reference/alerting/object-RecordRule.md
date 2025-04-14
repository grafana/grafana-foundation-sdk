---
title: <span class="badge object-type-interface"></span> RecordRule
---
# <span class="badge object-type-interface"></span> RecordRule

## Definition

```typescript
export interface RecordRule {
	// Which expression node should be used as the input for the recorded metric.
	from: string;
	// Name of the recorded metric.
	metric: string;
	// Which data source should be used to write the output of the recording rule, specified by UID.
	target_datasource_uid?: string;
}

```
## See also

 * <span class="badge builder"></span> [RecordRuleBuilder](./builder-RecordRuleBuilder.md)
