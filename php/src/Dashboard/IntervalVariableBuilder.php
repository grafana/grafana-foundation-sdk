<?php

namespace Grafana\Foundation\Dashboard;

/**
 * A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\VariableModel>
 */
class IntervalVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\VariableModel $internal;

    public function __construct(string $name)
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\VariableModel();
    $this->internal->name = $name;
    $this->internal->type = \Grafana\Foundation\Dashboard\VariableType::interval();
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
    public function values( $query): static
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
     * Allow custom values to be entered in the variable
     */
    public function allowCustomValue(bool $allowCustomValue): static
    {
        $this->internal->allowCustomValue = $allowCustomValue;
    
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

    /**
     * Additional static options for query variable
     * @param array<\Grafana\Foundation\Dashboard\VariableOption> $staticOptions
     */
    public function staticOptions(array $staticOptions): static
    {
        $this->internal->staticOptions = $staticOptions;
    
        return $this;
    }

    /**
     * Ordering of static options in relation to options returned from data source for query variable
     */
    public function staticOptionsOrder(\Grafana\Foundation\Dashboard\VariableModelStaticOptionsOrder $staticOptionsOrder): static
    {
        $this->internal->staticOptionsOrder = $staticOptionsOrder;
    
        return $this;
    }

    /**
     * Dynamically calculates interval by dividing time range by the count specified.
     */
    public function auto(bool $auto): static
    {
        $this->internal->auto = $auto;
    
        return $this;
    }

    /**
     * The minimum threshold below which the step count intervals will not divide the time.
     */
    public function minInterval(string $autoMin): static
    {
        $this->internal->autoMin = $autoMin;
    
        return $this;
    }

    /**
     * How many times the current time range should be divided to calculate the value, similar to the Max data points query option.
     * For example, if the current visible time range is 30 minutes, then the auto interval groups the data into 30 one-minute increments.
     */
    public function stepCount(int $autoCount): static
    {
        if (!($autoCount > 0)) {
            throw new \ValueError('$autoCount must be > 0');
        }
        $this->internal->autoCount = $autoCount;
    
        return $this;
    }

    public function definition(string $definition): static
    {
        $this->internal->definition = $definition;
    
        return $this;
    }

}
