<?php

namespace Grafana\Foundation\Cog;

final class DataqueryConfig
{
    public readonly string $identifier;

    /**
     * @var callable(array<string, mixed>): Dataquery
     */
    public $fromArray;

    /**
     * @var (callable(Dataquery): string)|null
     */
    public $convert;
    /**
     * @var (callable(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind): string)|null
     */
    public $convertv2;

    /**
     * @param callable(array<string, mixed>): Dataquery $fromArray
     * @param (callable(Dataquery): string)|null $convert 
     * @param (callable(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind): string)|null $convertv2
     */
    public function __construct(string $identifier, callable $fromArray, ?callable $convert = null, ?callable $convertv2 = null)
    {
        $this->identifier = $identifier;
        $this->fromArray = $fromArray;
        $this->convert = $convert;
        $this->convertv2 = $convertv2;
    }
}
