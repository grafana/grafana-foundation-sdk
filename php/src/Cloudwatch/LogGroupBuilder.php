<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\LogGroup>
 */
class LogGroupBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Cloudwatch\LogGroup $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Cloudwatch\LogGroup();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Cloudwatch\LogGroup
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * ARN of the log group
     */
    public function arn(string $arn): static
    {
        $this->internal->arn = $arn;
    
        return $this;
    }
    /**
     * Name of the log group
     */
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    /**
     * AccountId of the log group
     */
    public function accountId(string $accountId): static
    {
        $this->internal->accountId = $accountId;
    
        return $this;
    }
    /**
     * Label of the log group
     */
    public function accountLabel(string $accountLabel): static
    {
        $this->internal->accountLabel = $accountLabel;
    
        return $this;
    }

}
