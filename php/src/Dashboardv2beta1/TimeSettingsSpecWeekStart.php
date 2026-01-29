<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class TimeSettingsSpecWeekStart implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TimeSettingsSpecWeekStart>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function saturday(): self
    {
        if (!isset(self::$instances["Saturday"])) {
            self::$instances["Saturday"] = new self("saturday");
        }

        return self::$instances["Saturday"];
    }

    public static function monday(): self
    {
        if (!isset(self::$instances["Monday"])) {
            self::$instances["Monday"] = new self("monday");
        }

        return self::$instances["Monday"];
    }

    public static function sunday(): self
    {
        if (!isset(self::$instances["Sunday"])) {
            self::$instances["Sunday"] = new self("sunday");
        }

        return self::$instances["Sunday"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "saturday") {
            return self::saturday();
        }

        if ($value === "monday") {
            return self::monday();
        }

        if ($value === "sunday") {
            return self::sunday();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TimeSettingsSpecWeekStart");
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

