---
title: <span class="badge object-type-interface"></span> AnnotationEventFieldMapping
---
# <span class="badge object-type-interface"></span> AnnotationEventFieldMapping

Annotation event field mapping. Defines how to map a data frame field to an annotation event field.

## Definition

```typescript
export interface AnnotationEventFieldMapping {
	// Source type for the field value
	source?: string;
	// Constant value to use when source is "text"
	value?: string;
	// Regular expression to apply to the field value
	regex?: string;
}

```
## See also

 * <span class="badge builder"></span> [AnnotationEventFieldMappingBuilder](./builder-AnnotationEventFieldMappingBuilder.md)
