<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Maps regular expressions to replacement text and a color.
 * For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\RegexMap>
 */
class RegexMapBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\RegexMap $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\RegexMap();
    $this->internal->type = "regex";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\RegexMap
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Regular expression to match against and the result to apply when the value matches the regex
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardRegexMapOptions> $options
     */
    public function options(\Grafana\Foundation\Cog\Builder $options): static
    {
        $optionsResource = $options->build();
        $this->internal->options = $optionsResource;
    
        return $this;
    }

}
