---
title: <span class="badge object-type-interface"></span> MatcherConfig
---
# <span class="badge object-type-interface"></span> MatcherConfig

NOTE: (copied from dashboard_kind.cue, since not exported)

Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.

It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.

## Definition

```typescript
export interface MatcherConfig {
	// The matcher id. This is used to find the matcher implementation from registry.
	id: string;
	// The matcher options. This is specific to the matcher implementation.
	options?: any;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [MatcherConfigBuilder](./builder-MatcherConfigBuilder.md)
