<?php

namespace Grafana\Foundation\Preferences;

/**
 * Spec defines user, team or org Grafana preferences
 * swagger:model Preferences
 */
class Preferences implements \JsonSerializable
{
    /**
     * UID for the home dashboard
     */
    public ?string $homeDashboardUID;

    /**
     * The timezone selection
     * TODO: this should use the timezone defined in common
     */
    public ?string $timezone;

    /**
     * day of the week (sunday, monday, etc)
     */
    public ?string $weekStart;

    /**
     * light, dark, empty is default
     */
    public ?string $theme;

    /**
     * Selected language (beta)
     */
    public ?string $language;

    /**
     * Explore query history preferences
     */
    public ?\Grafana\Foundation\Preferences\QueryHistoryPreference $queryHistory;

    /**
     * Cookie preferences
     */
    public ?\Grafana\Foundation\Preferences\CookiePreferences $cookiePreferences;

    /**
     * Navigation preferences
     */
    public ?\Grafana\Foundation\Preferences\NavbarPreference $navbar;

    /**
     * @param string|null $homeDashboardUID
     * @param string|null $timezone
     * @param string|null $weekStart
     * @param string|null $theme
     * @param string|null $language
     * @param \Grafana\Foundation\Preferences\QueryHistoryPreference|null $queryHistory
     * @param \Grafana\Foundation\Preferences\CookiePreferences|null $cookiePreferences
     * @param \Grafana\Foundation\Preferences\NavbarPreference|null $navbar
     */
    public function __construct(?string $homeDashboardUID = null, ?string $timezone = null, ?string $weekStart = null, ?string $theme = null, ?string $language = null, ?\Grafana\Foundation\Preferences\QueryHistoryPreference $queryHistory = null, ?\Grafana\Foundation\Preferences\CookiePreferences $cookiePreferences = null, ?\Grafana\Foundation\Preferences\NavbarPreference $navbar = null)
    {
        $this->homeDashboardUID = $homeDashboardUID;
        $this->timezone = $timezone;
        $this->weekStart = $weekStart;
        $this->theme = $theme;
        $this->language = $language;
        $this->queryHistory = $queryHistory;
        $this->cookiePreferences = $cookiePreferences;
        $this->navbar = $navbar;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{homeDashboardUID?: string, timezone?: string, weekStart?: string, theme?: string, language?: string, queryHistory?: mixed, cookiePreferences?: mixed, navbar?: mixed} $inputData */
        $data = $inputData;
        return new self(
            homeDashboardUID: $data["homeDashboardUID"] ?? null,
            timezone: $data["timezone"] ?? null,
            weekStart: $data["weekStart"] ?? null,
            theme: $data["theme"] ?? null,
            language: $data["language"] ?? null,
            queryHistory: isset($data["queryHistory"]) ? (function($input) {
    	/** @var array{homeTab?: string} */
    $val = $input;
    	return \Grafana\Foundation\Preferences\QueryHistoryPreference::fromArray($val);
    })($data["queryHistory"]) : null,
            cookiePreferences: isset($data["cookiePreferences"]) ? (function($input) {
    	/** @var array{analytics?: mixed, performance?: mixed, functional?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Preferences\CookiePreferences::fromArray($val);
    })($data["cookiePreferences"]) : null,
            navbar: isset($data["navbar"]) ? (function($input) {
    	/** @var array{bookmarkUrls?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Preferences\NavbarPreference::fromArray($val);
    })($data["navbar"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->homeDashboardUID)) {
            $data["homeDashboardUID"] = $this->homeDashboardUID;
        }
        if (isset($this->timezone)) {
            $data["timezone"] = $this->timezone;
        }
        if (isset($this->weekStart)) {
            $data["weekStart"] = $this->weekStart;
        }
        if (isset($this->theme)) {
            $data["theme"] = $this->theme;
        }
        if (isset($this->language)) {
            $data["language"] = $this->language;
        }
        if (isset($this->queryHistory)) {
            $data["queryHistory"] = $this->queryHistory;
        }
        if (isset($this->cookiePreferences)) {
            $data["cookiePreferences"] = $this->cookiePreferences;
        }
        if (isset($this->navbar)) {
            $data["navbar"] = $this->navbar;
        }
        return $data;
    }
}
