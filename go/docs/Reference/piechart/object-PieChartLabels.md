---
title: <span class="badge object-type-enum"></span> PieChartLabels
---
# <span class="badge object-type-enum"></span> PieChartLabels

Select labels to display on the pie chart.

 - Name - The series or field name.

 - Percent - The percentage of the whole.

 - Value - The raw numerical value.

## Definition

```go
type PieChartLabels string
const (
	PieChartLabelsName PieChartLabels = "name"
	PieChartLabelsValue PieChartLabels = "value"
	PieChartLabelsPercent PieChartLabels = "percent"
)

```
