<?php

namespace Grafana\Foundation\Playlistv0alpha1;

final class ItemType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ItemType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function dashboardByTag(): self
    {
        if (!isset(self::$instances["DashboardByTag"])) {
            self::$instances["DashboardByTag"] = new self("dashboard_by_tag");
        }

        return self::$instances["DashboardByTag"];
    }

    public static function dashboardByUid(): self
    {
        if (!isset(self::$instances["DashboardByUid"])) {
            self::$instances["DashboardByUid"] = new self("dashboard_by_uid");
        }

        return self::$instances["DashboardByUid"];
    }

    public static function dashboardById(): self
    {
        if (!isset(self::$instances["DashboardById"])) {
            self::$instances["DashboardById"] = new self("dashboard_by_id");
        }

        return self::$instances["DashboardById"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "dashboard_by_tag") {
            return self::dashboardByTag();
        }

        if ($value === "dashboard_by_uid") {
            return self::dashboardByUid();
        }

        if ($value === "dashboard_by_id") {
            return self::dashboardById();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ItemType");
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

