<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\Dashboard>
 */
class DashboardBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\Dashboard $internal;
    private int $currentY;
    private int $currentX;
    private int $lastPanelHeight;

    public function __construct(string $title)
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\Dashboard();
    $this->internal->title = $title;
        $this->currentY = 0;
        $this->currentX = 0;
        $this->lastPanelHeight = 0;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\Dashboard
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Unique numeric identifier for the dashboard.
     * `id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.
     */
    public function id(int $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }

    /**
     * Unique dashboard identifier that can be generated by anyone. string (8-40)
     */
    public function uid(string $uid): static
    {
        $this->internal->uid = $uid;
    
        return $this;
    }

    /**
     * Title of dashboard.
     */
    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

    /**
     * Description of dashboard.
     */
    public function description(string $description): static
    {
        $this->internal->description = $description;
    
        return $this;
    }

    /**
     * This property should only be used in dashboards defined by plugins.  It is a quick check
     * to see if the version has changed since the last time.
     */
    public function revision(int $revision): static
    {
        $this->internal->revision = $revision;
    
        return $this;
    }

    /**
     * ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal
     */
    public function gnetId(string $gnetId): static
    {
        $this->internal->gnetId = $gnetId;
    
        return $this;
    }

    /**
     * Tags associated with dashboard.
     * @param array<string> $tags
     */
    public function tags(array $tags): static
    {
        $this->internal->tags = $tags;
    
        return $this;
    }

    /**
     * Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
     */
    public function timezone(string $timezone): static
    {
        $this->internal->timezone = $timezone;
    
        return $this;
    }

    /**
     * Whether a dashboard is editable or not.
     */
    public function editable(): static
    {
        $this->internal->editable = true;
    
        return $this;
    }

    /**
     * Whether a dashboard is editable or not.
     */
    public function readonly(): static
    {
        $this->internal->editable = false;
    
        return $this;
    }

    /**
     * Configuration of dashboard cursor sync behavior.
     * Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).
     */
    public function tooltip(\Grafana\Foundation\Dashboard\DashboardCursorSync $graphTooltip): static
    {
        $this->internal->graphTooltip = $graphTooltip;
    
        return $this;
    }

    /**
     * Time range for dashboard.
     * Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.
     */
    public function time(string $from, string $to): static
    {    
        if ($this->internal->time === null) {
            $this->internal->time = new \Grafana\Foundation\Dashboard\DashboardDashboardTime();
        }
        assert($this->internal->time instanceof \Grafana\Foundation\Dashboard\DashboardDashboardTime);
        $this->internal->time->from = $from;
        $this->internal->time->to = $to;
    
        return $this;
    }

    /**
     * Configuration of the time picker shown at the top of a dashboard.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\TimePickerConfig> $timepicker
     */
    public function timepicker(\Grafana\Foundation\Cog\Builder $timepicker): static
    {
        $timepickerResource = $timepicker->build();
        $this->internal->timepicker = $timepickerResource;
    
        return $this;
    }

    /**
     * The month that the fiscal year starts on.  0 = January, 11 = December
     */
    public function fiscalYearStartMonth(int $fiscalYearStartMonth): static
    {
        if (!($fiscalYearStartMonth < 12)) {
            throw new \ValueError('$fiscalYearStartMonth must be < 12');
        }
        $this->internal->fiscalYearStartMonth = $fiscalYearStartMonth;
    
        return $this;
    }

    /**
     * When set to true, the dashboard will redraw panels at an interval matching the pixel width.
     * This will keep data "moving left" regardless of the query refresh rate. This setting helps
     * avoid dashboards presenting stale live data
     */
    public function liveNow(bool $liveNow): static
    {
        $this->internal->liveNow = $liveNow;
    
        return $this;
    }

    /**
     * Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
     */
    public function weekStart(string $weekStart): static
    {
        $this->internal->weekStart = $weekStart;
    
        return $this;
    }

    /**
     * Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
     */
    public function refresh(string $refresh): static
    {
        $this->internal->refresh = $refresh;
    
        return $this;
    }

    /**
     * Version of the dashboard, incremented each time the dashboard is updated.
     */
    public function version(int $version): static
    {
        $this->internal->version = $version;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\Panel> $panel
     */
    public function withPanel(\Grafana\Foundation\Cog\Builder $panel): static
    {    
        if ($this->internal->panels === null) {
            $this->internal->panels = [];
        }
        
        $panelResource = $panel->build();
    
        if ($panelResource->gridPos === null) {
            $panelResource->gridPos = new \Grafana\Foundation\Dashboard\GridPos();
        }
        // The panel either has no position set, or it is the first panel of the dashboard.
        // In that case, we position it on the grid
        if ($panelResource->gridPos->x === 0 && $panelResource->gridPos->y === 0) {
    	    $panelResource->gridPos->x = $this->currentX;
    	    $panelResource->gridPos->y = $this->currentY;
        }
        $this->internal->panels[] = $panelResource;
    
        // Prepare the coordinates for the next panel
        $this->currentX += $panelResource->gridPos->w;
        $this->lastPanelHeight = max($this->lastPanelHeight, $panelResource->gridPos->h);
    
        // Check for grid width overflow?
        if ($this->currentX >= 24) {
            $this->currentX = 0;
            $this->currentY += $this->lastPanelHeight;
            $this->lastPanelHeight = 0;
        }
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\RowPanel> $rowPanel
     */
    public function withRow(\Grafana\Foundation\Cog\Builder $rowPanel): static
    {    
        if ($this->internal->panels === null) {
            $this->internal->panels = [];
        }
        
        $rowPanelResource = $rowPanel->build();
    
        // Position the row on the grid
        if ($rowPanelResource->gridPos === null || ($rowPanelResource->gridPos->x === 0 && $rowPanelResource->gridPos->y === 0)) {
            $rowPanelResource->gridPos = new \Grafana\Foundation\Dashboard\GridPos(
                x: 0, // beginning of the line
                y: $this->currentY + $this->lastPanelHeight,
    
                h: 1,
                w: 24, // full width
            );
        }
        $this->internal->panels[] = $rowPanelResource;
    
        // Reset the state for the next row
        $this->currentX = 0;
        $this->currentY = $rowPanelResource->gridPos->y + 1;
        $this->lastPanelHeight = 0;
    
        // Position the row's panels on the grid
        foreach ($rowPanelResource->panels as $panel) {
            if ($panel->gridPos === null) {
                $panel->gridPos = new \Grafana\Foundation\Dashboard\GridPos();
            }
    
            // The panel either has no position set, or it is the first panel of the dashboard.
            // In that case, we position it on the grid
            if ($panel->gridPos->x === 0 && $panel->gridPos->y === 0) {
                $panel->gridPos->x = $this->currentX;
                $panel->gridPos->y = $this->currentY;
            }
    
            // Prepare the coordinates for the next panel
            $this->currentX += $panel->gridPos->w;
            $this->lastPanelHeight = max($this->lastPanelHeight, $panel->gridPos->h);
    
            // Check for grid width overflow?
            if ($this->currentX >= 24) {
                $this->currentX = 0;
                $this->currentY += $this->lastPanelHeight;
                $this->lastPanelHeight = 0;
            }
        }
    
        return $this;
    }

    /**
     * Configured template variables
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\VariableModel>> $variables
     */
    public function variables(array $variables): static
    {
            $variablesResources = [];
            foreach ($variables as $r1) {
                    $variablesResources[] = $r1->build();
            }
        $this->internal->templating->list = $variablesResources;
    
        return $this;
    }

    /**
     * Configured template variables
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\VariableModel> $variable
     */
    public function withVariable(\Grafana\Foundation\Cog\Builder $variable): static
    {    
        if ($this->internal->templating->list === null) {
            $this->internal->templating->list = [];
        }
        
        $variableResource = $variable->build();
        $this->internal->templating->list[] = $variableResource;
    
        return $this;
    }

    /**
     * Contains the list of annotations that are associated with the dashboard.
     * Annotations are used to overlay event markers and overlay event tags on graphs.
     * Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
     * See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationQuery>> $annotations
     */
    public function annotations(array $annotations): static
    {
            $annotationsResources = [];
            foreach ($annotations as $r1) {
                    $annotationsResources[] = $r1->build();
            }
        $this->internal->annotations->list = $annotationsResources;
    
        return $this;
    }

    /**
     * Contains the list of annotations that are associated with the dashboard.
     * Annotations are used to overlay event markers and overlay event tags on graphs.
     * Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
     * See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationQuery> $annotation
     */
    public function annotation(\Grafana\Foundation\Cog\Builder $annotation): static
    {    
        if ($this->internal->annotations->list === null) {
            $this->internal->annotations->list = [];
        }
        
        $annotationResource = $annotation->build();
        $this->internal->annotations->list[] = $annotationResource;
    
        return $this;
    }

    /**
     * Links with references to other dashboards or external websites.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardLink>> $links
     */
    public function links(array $links): static
    {
            $linksResources = [];
            foreach ($links as $r1) {
                    $linksResources[] = $r1->build();
            }
        $this->internal->links = $linksResources;
    
        return $this;
    }

    /**
     * Links with references to other dashboards or external websites.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardLink> $link
     */
    public function link(\Grafana\Foundation\Cog\Builder $link): static
    {    
        if ($this->internal->links === null) {
            $this->internal->links = [];
        }
        
        $linkResource = $link->build();
        $this->internal->links[] = $linkResource;
    
        return $this;
    }

    /**
     * Snapshot options. They are present only if the dashboard is a snapshot.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\Snapshot> $snapshot
     */
    public function snapshot(\Grafana\Foundation\Cog\Builder $snapshot): static
    {
        $snapshotResource = $snapshot->build();
        $this->internal->snapshot = $snapshotResource;
    
        return $this;
    }

    /**
     * When set to true, the dashboard will load all panels in the dashboard when it's loaded.
     */
    public function preload(bool $preload): static
    {
        $this->internal->preload = $preload;
    
        return $this;
    }

}
