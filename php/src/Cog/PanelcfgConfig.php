<?php

namespace Grafana\Foundation\Cog;

final class PanelcfgConfig
{
    public readonly string $identifier;

    /**
     * @var (callable(array<string, mixed>): object)|null
     */
    public $optionsFromArray;

    /**
     * @var (callable(array<string, mixed>): object)|null
     */
    public $fieldConfigFromArray;

    /**
     * @param (callable(array<string, mixed>): object)|null $optionsFromArray
     * @param (callable(array<string, mixed>): object)|null $fieldConfigFromArray
     */
    public function __construct(string $identifier, ?callable $optionsFromArray = null, ?callable $fieldConfigFromArray = null)
    {
        $this->identifier = $identifier;
        $this->optionsFromArray = $optionsFromArray;
        $this->fieldConfigFromArray = $fieldConfigFromArray;
    }
}
