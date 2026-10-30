<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class NoDataState implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, NoDataState>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function noData(): self
    {
        if (!isset(self::$instances["NoData"])) {
            self::$instances["NoData"] = new self("NoData");
        }

        return self::$instances["NoData"];
    }

    public static function ok(): self
    {
        if (!isset(self::$instances["Ok"])) {
            self::$instances["Ok"] = new self("Ok");
        }

        return self::$instances["Ok"];
    }

    public static function alerting(): self
    {
        if (!isset(self::$instances["Alerting"])) {
            self::$instances["Alerting"] = new self("Alerting");
        }

        return self::$instances["Alerting"];
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
        if ($value === "NoData") {
            return self::noData();
        }

        if ($value === "Ok") {
            return self::ok();
        }

        if ($value === "Alerting") {
            return self::alerting();
        }

        if ($value === "KeepLast") {
            return self::keepLast();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum NoDataState");
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

