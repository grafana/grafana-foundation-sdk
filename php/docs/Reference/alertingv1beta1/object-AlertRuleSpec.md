---
title: <span class="badge object-type-class"></span> AlertRuleSpec
---
# <span class="badge object-type-class"></span> AlertRuleSpec

## Definition

```php
class AlertRuleSpec implements \JsonSerializable
{
    public string $title;

    public ?bool $paused;

    public \Grafana\Foundation\Alertingv1beta1\IntervalTrigger $trigger;

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    /**
     * @var array<string, string>|null
     */
    public ?array $annotations;

    public ?string $for;

    public ?string $keepFiringFor;

    public ?int $missingSeriesEvalsToResolve;

    public \Grafana\Foundation\Alertingv1beta1\NoDataState $noDataState;

    public \Grafana\Foundation\Alertingv1beta1\ExecErrState $execErrState;

    /**
     * @var \Grafana\Foundation\Alertingv1beta1\SimplifiedRouting|\Grafana\Foundation\Alertingv1beta1\NamedRoutingTree
     */
    public $notificationSettings;

    /**
     * @var array<string, \Grafana\Foundation\Alertingv1beta1\Expression>
     */
    public array $expressions;

    public ?\Grafana\Foundation\Alertingv1beta1\PanelRef $panelRef;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [AlertRuleBuilder](./builder-AlertRuleBuilder.md)
