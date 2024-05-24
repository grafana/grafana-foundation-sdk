// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"errors"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboard] = (*DashboardBuilder)(nil)

type DashboardBuilder struct {
	internal        *Dashboard
	errors          map[string]cog.BuildErrors
	currentY        uint32
	currentX        uint32
	lastPanelHeight uint32
}

func NewDashboardBuilder(title string) *DashboardBuilder {
	resource := &Dashboard{}
	builder := &DashboardBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Title = &title

	return builder
}

func (builder *DashboardBuilder) Build() (Dashboard, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Dashboard", err)...)
	}

	if len(errs) != 0 {
		return Dashboard{}, errs
	}

	return *builder.internal, nil
}

// Unique numeric identifier for the dashboard.
// `id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.
func (builder *DashboardBuilder) Id(id int64) *DashboardBuilder {
	builder.internal.Id = &id

	return builder
}

// Unique dashboard identifier that can be generated by anyone. string (8-40)
func (builder *DashboardBuilder) Uid(uid string) *DashboardBuilder {
	builder.internal.Uid = &uid

	return builder
}

// Title of dashboard.
func (builder *DashboardBuilder) Title(title string) *DashboardBuilder {
	builder.internal.Title = &title

	return builder
}

// Description of dashboard.
func (builder *DashboardBuilder) Description(description string) *DashboardBuilder {
	builder.internal.Description = &description

	return builder
}

// This property should only be used in dashboards defined by plugins.  It is a quick check
// to see if the version has changed since the last time.
func (builder *DashboardBuilder) Revision(revision int64) *DashboardBuilder {
	builder.internal.Revision = &revision

	return builder
}

// ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal
func (builder *DashboardBuilder) GnetId(gnetId string) *DashboardBuilder {
	builder.internal.GnetId = &gnetId

	return builder
}

// Tags associated with dashboard.
func (builder *DashboardBuilder) Tags(tags []string) *DashboardBuilder {
	builder.internal.Tags = tags

	return builder
}

// Theme of dashboard.
// Default value: dark.
func (builder *DashboardBuilder) Style(style DashboardStyle) *DashboardBuilder {
	builder.internal.Style = style

	return builder
}

// Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
func (builder *DashboardBuilder) Timezone(timezone string) *DashboardBuilder {
	builder.internal.Timezone = &timezone

	return builder
}

// Whether a dashboard is editable or not.
func (builder *DashboardBuilder) Editable() *DashboardBuilder {
	builder.internal.Editable = true

	return builder
}

// Whether a dashboard is editable or not.
func (builder *DashboardBuilder) Readonly() *DashboardBuilder {
	builder.internal.Editable = false

	return builder
}

// Configuration of dashboard cursor sync behavior.
// Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).
func (builder *DashboardBuilder) Tooltip(graphTooltip DashboardCursorSync) *DashboardBuilder {
	builder.internal.GraphTooltip = graphTooltip

	return builder
}

// Time range for dashboard.
// Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.
func (builder *DashboardBuilder) Time(from string, to string) *DashboardBuilder {
	if builder.internal.Time == nil {
		builder.internal.Time = &struct {
			From string `json:"from"`
			To   string `json:"to"`
		}{}
	}
	builder.internal.Time.From = from
	if builder.internal.Time == nil {
		builder.internal.Time = &struct {
			From string `json:"from"`
			To   string `json:"to"`
		}{}
	}
	builder.internal.Time.To = to

	return builder
}

// Configuration of the time picker shown at the top of a dashboard.
func (builder *DashboardBuilder) Timepicker(timepicker cog.Builder[TimePicker]) *DashboardBuilder {
	timepickerResource, err := timepicker.Build()
	if err != nil {
		builder.errors["timepicker"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Timepicker = &timepickerResource

	return builder
}

// The month that the fiscal year starts on.  0 = January, 11 = December
func (builder *DashboardBuilder) FiscalYearStartMonth(fiscalYearStartMonth uint8) *DashboardBuilder {
	if !(fiscalYearStartMonth < 12) {
		builder.errors["fiscalYearStartMonth"] = cog.MakeBuildErrors("fiscalYearStartMonth", errors.New("fiscalYearStartMonth must be < 12"))
		return builder
	}
	builder.internal.FiscalYearStartMonth = &fiscalYearStartMonth

	return builder
}

// When set to true, the dashboard will redraw panels at an interval matching the pixel width.
// This will keep data "moving left" regardless of the query refresh rate. This setting helps
// avoid dashboards presenting stale live data
func (builder *DashboardBuilder) LiveNow(liveNow bool) *DashboardBuilder {
	builder.internal.LiveNow = &liveNow

	return builder
}

// Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
func (builder *DashboardBuilder) WeekStart(weekStart string) *DashboardBuilder {
	builder.internal.WeekStart = &weekStart

	return builder
}

// Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
func (builder *DashboardBuilder) Refresh(string string) *DashboardBuilder {
	if builder.internal.Refresh == nil {
		builder.internal.Refresh = &StringOrBool{}
	}
	builder.internal.Refresh.String = &string

	return builder
}

// Version of the dashboard, incremented each time the dashboard is updated.
func (builder *DashboardBuilder) Version(version uint32) *DashboardBuilder {
	builder.internal.Version = &version

	return builder
}

func (builder *DashboardBuilder) WithPanel(panel cog.Builder[Panel]) *DashboardBuilder {
	panelResource, err := panel.Build()
	if err != nil {
		builder.errors["panels"] = err.(cog.BuildErrors)
		return builder
	}

	if panelResource.GridPos == nil {
		panelResource.GridPos = &GridPos{}
	}
	// The panel either has no position set, or it is the first panel of the dashboard.
	// In that case, we position it on the grid
	if panelResource.GridPos.X == 0 && panelResource.GridPos.Y == 0 {
		panelResource.GridPos.X = builder.currentX
		panelResource.GridPos.Y = builder.currentY
	}
	builder.internal.Panels = append(builder.internal.Panels, PanelOrRowPanel{
		Panel: &panelResource,
	})

	// Prepare the coordinates for the next panel
	builder.currentX += panelResource.GridPos.W
	builder.lastPanelHeight = max(builder.lastPanelHeight, panelResource.GridPos.H)

	// Check for grid width overflow?
	if builder.currentX >= 24 {
		builder.currentX = 0
		builder.currentY += builder.lastPanelHeight
		builder.lastPanelHeight = 0
	}

	return builder
}

func (builder *DashboardBuilder) WithRow(rowPanel cog.Builder[RowPanel]) *DashboardBuilder {
	rowPanelResource, err := rowPanel.Build()
	if err != nil {
		builder.errors["panels"] = err.(cog.BuildErrors)
		return builder
	}

	// Position the row on the grid
	rowPanelResource.GridPos = &GridPos{
		X: 0, // beginning of the line
		Y: builder.currentY + builder.lastPanelHeight,

		H: 1,
		W: 24, // full width
	}
	builder.internal.Panels = append(builder.internal.Panels, PanelOrRowPanel{
		RowPanel: &rowPanelResource,
	})

	// Reset the state for the next row
	builder.currentX = 0
	builder.currentY = rowPanelResource.GridPos.Y + 1
	builder.lastPanelHeight = 0

	// Position the row's panels on the grid
	for _, panel := range rowPanelResource.Panels {
		// The panel either has no position set, or it is the first panel of the dashboard.
		// In that case, we position it on the grid
		if panel.GridPos.X == 0 && panel.GridPos.Y == 0 {
			panel.GridPos.X = builder.currentX
			panel.GridPos.Y = builder.currentY
		}

		// Prepare the coordinates for the next panel
		builder.currentX += panel.GridPos.W
		builder.lastPanelHeight = max(builder.lastPanelHeight, panel.GridPos.H)

		// Check for grid width overflow?
		if builder.currentX >= 24 {
			builder.currentX = 0
			builder.currentY += builder.lastPanelHeight
			builder.lastPanelHeight = 0
		}
	}

	return builder
}

// Configured template variables
func (builder *DashboardBuilder) Variables(list []cog.Builder[VariableModel]) *DashboardBuilder {
	listResources := make([]VariableModel, 0, len(list))
	for _, r1 := range list {
		listDepth1, err := r1.Build()
		if err != nil {
			builder.errors["templating.list"] = err.(cog.BuildErrors)
			return builder
		}
		listResources = append(listResources, listDepth1)
	}
	builder.internal.Templating.List = listResources

	return builder
}

// Configured template variables
func (builder *DashboardBuilder) WithVariable(list cog.Builder[VariableModel]) *DashboardBuilder {
	listResource, err := list.Build()
	if err != nil {
		builder.errors["templating.list"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Templating.List = append(builder.internal.Templating.List, listResource)

	return builder
}

// Contains the list of annotations that are associated with the dashboard.
// Annotations are used to overlay event markers and overlay event tags on graphs.
// Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
// See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
func (builder *DashboardBuilder) Annotations(list []cog.Builder[AnnotationQuery]) *DashboardBuilder {
	listResources := make([]AnnotationQuery, 0, len(list))
	for _, r1 := range list {
		listDepth1, err := r1.Build()
		if err != nil {
			builder.errors["annotations.list"] = err.(cog.BuildErrors)
			return builder
		}
		listResources = append(listResources, listDepth1)
	}
	builder.internal.Annotations.List = listResources

	return builder
}

// Contains the list of annotations that are associated with the dashboard.
// Annotations are used to overlay event markers and overlay event tags on graphs.
// Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
// See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
func (builder *DashboardBuilder) Annotation(list cog.Builder[AnnotationQuery]) *DashboardBuilder {
	listResource, err := list.Build()
	if err != nil {
		builder.errors["annotations.list"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Annotations.List = append(builder.internal.Annotations.List, listResource)

	return builder
}

// Links with references to other dashboards or external websites.
func (builder *DashboardBuilder) Links(links []cog.Builder[DashboardLink]) *DashboardBuilder {
	linksResources := make([]DashboardLink, 0, len(links))
	for _, r1 := range links {
		linksDepth1, err := r1.Build()
		if err != nil {
			builder.errors["links"] = err.(cog.BuildErrors)
			return builder
		}
		linksResources = append(linksResources, linksDepth1)
	}
	builder.internal.Links = linksResources

	return builder
}

// Links with references to other dashboards or external websites.
func (builder *DashboardBuilder) Link(links cog.Builder[DashboardLink]) *DashboardBuilder {
	linksResource, err := links.Build()
	if err != nil {
		builder.errors["links"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Links = append(builder.internal.Links, linksResource)

	return builder
}

// Snapshot options. They are present only if the dashboard is a snapshot.
func (builder *DashboardBuilder) Snapshot(snapshot cog.Builder[Snapshot]) *DashboardBuilder {
	snapshotResource, err := snapshot.Build()
	if err != nil {
		builder.errors["snapshot"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Snapshot = &snapshotResource

	return builder
}

func (builder *DashboardBuilder) applyDefaults() {
	builder.Timezone("browser")
	builder.Tooltip(0)
	builder.FiscalYearStartMonth(0)
}
