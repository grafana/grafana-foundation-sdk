<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Rule>
 */
class RuleBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\Rule $internal;

    public function __construct(string $title)
    {
    	$this->internal = new \Grafana\Foundation\Alerting\Rule();
    if (!(strlen($title) >= 1)) {
        throw new \ValueError('strlen($title) must be >= 1');
    }
    if (!(strlen($title) <= 190)) {
        throw new \ValueError('strlen($title) must be <= 190');
    }
    $this->internal->title = $title;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alerting\Rule
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<string, string> $annotations
     */
    public function annotations(array $annotations): static
    {
        $this->internal->annotations = $annotations;
    
        return $this;
    }
    public function condition(string $condition): static
    {
        $this->internal->condition = $condition;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Query>> $data
     */
    public function queries(array $data): static
    {
            $dataResources = [];
            foreach ($data as $r1) {
                    $dataResources[] = $r1->build();
            }
        $this->internal->data = $dataResources;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Query> $data
     */
    public function withQuery(\Grafana\Foundation\Cog\Builder $data): static
    {
        $dataResource = $data->build();
        $this->internal->data[] = $dataResource;
    
        return $this;
    }
    public function execErrState(\Grafana\Foundation\Alerting\RuleExecErrState $execErrState): static
    {
        $this->internal->execErrState = $execErrState;
    
        return $this;
    }
    public function folderUID(string $folderUID): static
    {
        $this->internal->folderUID = $folderUID;
    
        return $this;
    }
    /**
     * The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
     * Before this time has elapsed, the rule is only considered to be Pending.
     */
    public function for(string $for): static
    {
        $this->internal->for = $for;
    
        return $this;
    }
    public function id(int $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    public function isPaused(bool $isPaused): static
    {
        $this->internal->isPaused = $isPaused;
    
        return $this;
    }
    /**
     * @param array<string, string> $labels
     */
    public function labels(array $labels): static
    {
        $this->internal->labels = $labels;
    
        return $this;
    }
    public function noDataState(\Grafana\Foundation\Alerting\RuleNoDataState $noDataState): static
    {
        $this->internal->noDataState = $noDataState;
    
        return $this;
    }
    public function orgID(int $orgID): static
    {
        $this->internal->orgID = $orgID;
    
        return $this;
    }
    public function provenance(string $provenance): static
    {
        $this->internal->provenance = $provenance;
    
        return $this;
    }
    public function ruleGroup(string $ruleGroup): static
    {
        if (!(strlen($ruleGroup) >= 1)) {
            throw new \ValueError('strlen($ruleGroup) must be >= 1');
        }
        if (!(strlen($ruleGroup) <= 190)) {
            throw new \ValueError('strlen($ruleGroup) must be <= 190');
        }
        $this->internal->ruleGroup = $ruleGroup;
    
        return $this;
    }
    public function title(string $title): static
    {
        if (!(strlen($title) >= 1)) {
            throw new \ValueError('strlen($title) must be >= 1');
        }
        if (!(strlen($title) <= 190)) {
            throw new \ValueError('strlen($title) must be <= 190');
        }
        $this->internal->title = $title;
    
        return $this;
    }
    public function uid(string $uid): static
    {
        if (!(strlen($uid) >= 1)) {
            throw new \ValueError('strlen($uid) must be >= 1');
        }
        if (!(strlen($uid) <= 40)) {
            throw new \ValueError('strlen($uid) must be <= 40');
        }
        $this->internal->uid = $uid;
    
        return $this;
    }
    public function updated(string $updated): static
    {
        $this->internal->updated = $updated;
    
        return $this;
    }

}
