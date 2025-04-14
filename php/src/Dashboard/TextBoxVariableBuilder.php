<?php

namespace Grafana\Foundation\Dashboard;

/**
 * A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\VariableModel>
 */
class TextBoxVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\VariableModel $internal;

    public function __construct(string $name)
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\VariableModel();
    $this->internal->name = $name;
    $this->internal->type = \Grafana\Foundation\Dashboard\VariableType::textbox();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\VariableModel
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Name of variable
     */
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

    /**
     * Optional display name
     */
    public function label(string $label): static
    {
        $this->internal->label = $label;
    
        return $this;
    }

    /**
     * Visibility configuration for the variable
     */
    public function hide(\Grafana\Foundation\Dashboard\VariableHide $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }

    /**
     * Description of variable. It can be defined but `null`.
     */
    public function description(string $description): static
    {
        $this->internal->description = $description;
    
        return $this;
    }

    /**
     * Query used to fetch values for a variable
     * @param string|array<string, mixed> $query
     */
    public function defaultValue( $query): static
    {
        $this->internal->query = $query;
    
        return $this;
    }

    /**
     * Shows current selected variable text/value on the dashboard
     */
    public function current(\Grafana\Foundation\Dashboard\VariableOption $current): static
    {
        $this->internal->current = $current;
    
        return $this;
    }

    /**
     * Options that can be selected for a variable.
     * @param array<\Grafana\Foundation\Dashboard\VariableOption> $options
     */
    public function options(array $options): static
    {
        $this->internal->options = $options;
    
        return $this;
    }

}
