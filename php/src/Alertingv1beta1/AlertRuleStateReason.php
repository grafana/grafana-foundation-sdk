<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class AlertRuleStateReason implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, AlertRuleStateReason>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function evaluated(): self
    {
        if (!isset(self::$instances["Evaluated"])) {
            self::$instances["Evaluated"] = new self("Evaluated");
        }

        return self::$instances["Evaluated"];
    }

    public static function keepLast(): self
    {
        if (!isset(self::$instances["KeepLast"])) {
            self::$instances["KeepLast"] = new self("KeepLast");
        }

        return self::$instances["KeepLast"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "Evaluated") {
            return self::evaluated();
        }

        if ($value === "KeepLast") {
            return self::keepLast();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum AlertRuleStateReason");
    }

    public function jsonSerialize(): string
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}

