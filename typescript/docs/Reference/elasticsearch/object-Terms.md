---
title: <span class="badge object-type-interface"></span> Terms
---
# <span class="badge object-type-interface"></span> Terms

## Definition

```typescript
export interface Terms {
	field?: string;
	id: string;
	type: "terms";
	settings?: {
		order?: elasticsearch.TermsOrder;
		size?: string;
		min_doc_count?: string;
		orderBy?: string;
		missing?: string;
	};
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [TermsBuilder](./builder-TermsBuilder.md)
