<?php

namespace Grafana\Foundation\Xychart;

/**
 * NOTE: (copied from dashboard_kind.cue, since not exported)
 * Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
 * It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\MatcherConfig>
 */
class MatcherConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Xychart\MatcherConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Xychart\MatcherConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Xychart\MatcherConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The matcher id. This is used to find the matcher implementation from registry.
     */
    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    /**
     * The matcher options. This is specific to the matcher implementation.
     * @param mixed $options
     */
    public function options( $options): static
    {
        $this->internal->options = $options;
    
        return $this;
    }

}
