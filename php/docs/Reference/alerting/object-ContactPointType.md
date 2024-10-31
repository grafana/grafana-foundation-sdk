---
title: <span class="badge object-type-enum"></span> ContactPointType
---
# <span class="badge object-type-enum"></span> ContactPointType

## Definition

```php
final class ContactPointType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ContactPointType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function alertmanager(): self
    {
        if (!isset(self::$instances["Alertmanager"])) {
            self::$instances["Alertmanager"] = new self("alertmanager");
        }

        return self::$instances["Alertmanager"];
    }

    public static function dingding(): self
    {
        if (!isset(self::$instances["Dingding"])) {
            self::$instances["Dingding"] = new self(" dingding");
        }

        return self::$instances["Dingding"];
    }

    public static function discord(): self
    {
        if (!isset(self::$instances["Discord"])) {
            self::$instances["Discord"] = new self(" discord");
        }

        return self::$instances["Discord"];
    }

    public static function email(): self
    {
        if (!isset(self::$instances["Email"])) {
            self::$instances["Email"] = new self(" email");
        }

        return self::$instances["Email"];
    }

    public static function googlechat(): self
    {
        if (!isset(self::$instances["Googlechat"])) {
            self::$instances["Googlechat"] = new self(" googlechat");
        }

        return self::$instances["Googlechat"];
    }

    public static function kafka(): self
    {
        if (!isset(self::$instances["Kafka"])) {
            self::$instances["Kafka"] = new self(" kafka");
        }

        return self::$instances["Kafka"];
    }

    public static function line(): self
    {
        if (!isset(self::$instances["Line"])) {
            self::$instances["Line"] = new self(" line");
        }

        return self::$instances["Line"];
    }

    public static function opsgenie(): self
    {
        if (!isset(self::$instances["Opsgenie"])) {
            self::$instances["Opsgenie"] = new self(" opsgenie");
        }

        return self::$instances["Opsgenie"];
    }

    public static function pagerduty(): self
    {
        if (!isset(self::$instances["Pagerduty"])) {
            self::$instances["Pagerduty"] = new self(" pagerduty");
        }

        return self::$instances["Pagerduty"];
    }

    public static function pushover(): self
    {
        if (!isset(self::$instances["Pushover"])) {
            self::$instances["Pushover"] = new self(" pushover");
        }

        return self::$instances["Pushover"];
    }

    public static function sensugo(): self
    {
        if (!isset(self::$instances["Sensugo"])) {
            self::$instances["Sensugo"] = new self(" sensugo");
        }

        return self::$instances["Sensugo"];
    }

    public static function slack(): self
    {
        if (!isset(self::$instances["Slack"])) {
            self::$instances["Slack"] = new self(" slack");
        }

        return self::$instances["Slack"];
    }

    public static function teams(): self
    {
        if (!isset(self::$instances["Teams"])) {
            self::$instances["Teams"] = new self(" teams");
        }

        return self::$instances["Teams"];
    }

    public static function telegram(): self
    {
        if (!isset(self::$instances["Telegram"])) {
            self::$instances["Telegram"] = new self(" telegram");
        }

        return self::$instances["Telegram"];
    }

    public static function threema(): self
    {
        if (!isset(self::$instances["Threema"])) {
            self::$instances["Threema"] = new self(" threema");
        }

        return self::$instances["Threema"];
    }

    public static function victorops(): self
    {
        if (!isset(self::$instances["Victorops"])) {
            self::$instances["Victorops"] = new self(" victorops");
        }

        return self::$instances["Victorops"];
    }

    public static function webhook(): self
    {
        if (!isset(self::$instances["Webhook"])) {
            self::$instances["Webhook"] = new self(" webhook");
        }

        return self::$instances["Webhook"];
    }

    public static function wecom(): self
    {
        if (!isset(self::$instances["Wecom"])) {
            self::$instances["Wecom"] = new self(" wecom");
        }

        return self::$instances["Wecom"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "alertmanager") {
            return self::alertmanager();
        }

        if ($value === " dingding") {
            return self::dingding();
        }

        if ($value === " discord") {
            return self::discord();
        }

        if ($value === " email") {
            return self::email();
        }

        if ($value === " googlechat") {
            return self::googlechat();
        }

        if ($value === " kafka") {
            return self::kafka();
        }

        if ($value === " line") {
            return self::line();
        }

        if ($value === " opsgenie") {
            return self::opsgenie();
        }

        if ($value === " pagerduty") {
            return self::pagerduty();
        }

        if ($value === " pushover") {
            return self::pushover();
        }

        if ($value === " sensugo") {
            return self::sensugo();
        }

        if ($value === " slack") {
            return self::slack();
        }

        if ($value === " teams") {
            return self::teams();
        }

        if ($value === " telegram") {
            return self::telegram();
        }

        if ($value === " threema") {
            return self::threema();
        }

        if ($value === " victorops") {
            return self::victorops();
        }

        if ($value === " webhook") {
            return self::webhook();
        }

        if ($value === " wecom") {
            return self::wecom();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ContactPointType");
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

```
