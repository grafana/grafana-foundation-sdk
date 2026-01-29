<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Determine if the variable shows on dashboard
 * Accepted values are `dontHide` (show label and value), `hideLabel` (show value only), `hideVariable` (show nothing), `inControlsMenu` (show in a drop-down menu).
 */
final class VariableHide implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, VariableHide>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function dontHide(): self
    {
        if (!isset(self::$instances["dontHide"])) {
            self::$instances["dontHide"] = new self("dontHide");
        }

        return self::$instances["dontHide"];
    }

    public static function hideLabel(): self
    {
        if (!isset(self::$instances["hideLabel"])) {
            self::$instances["hideLabel"] = new self("hideLabel");
        }

        return self::$instances["hideLabel"];
    }

    public static function hideVariable(): self
    {
        if (!isset(self::$instances["hideVariable"])) {
            self::$instances["hideVariable"] = new self("hideVariable");
        }

        return self::$instances["hideVariable"];
    }

    public static function inControlsMenu(): self
    {
        if (!isset(self::$instances["inControlsMenu"])) {
            self::$instances["inControlsMenu"] = new self("inControlsMenu");
        }

        return self::$instances["inControlsMenu"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "dontHide") {
            return self::dontHide();
        }

        if ($value === "hideLabel") {
            return self::hideLabel();
        }

        if ($value === "hideVariable") {
            return self::hideVariable();
        }

        if ($value === "inControlsMenu") {
            return self::inControlsMenu();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VariableHide");
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

