---
title: <span class="badge object-type-interface"></span> XYSeriesConfig
---
# <span class="badge object-type-interface"></span> XYSeriesConfig

## Definition

```typescript
export interface XYSeriesConfig {
	name?: {
		fixed?: string;
	};
	frame?: {
		matcher: xychart.MatcherConfig;
	};
	x?: {
		matcher: xychart.MatcherConfig;
	};
	y?: {
		matcher: xychart.MatcherConfig;
	};
	color?: {
		matcher: xychart.MatcherConfig;
	};
	size?: {
		matcher: xychart.MatcherConfig;
	};
}

```
## See also

 * <span class="badge builder"></span> [XYSeriesConfigBuilder](./builder-XYSeriesConfigBuilder.md)
