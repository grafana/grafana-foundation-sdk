---
title: <span class="badge object-type-interface"></span> BucketScript
---
# <span class="badge object-type-interface"></span> BucketScript

## Definition

```typescript
export interface BucketScript {
	type: "bucket_script";
	pipelineVariables?: elasticsearch.PipelineVariable[];
	id: string;
	settings?: {
		script?: elasticsearch.InlineScript;
	};
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [BucketScriptBuilder](./builder-BucketScriptBuilder.md)
