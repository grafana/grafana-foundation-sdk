# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class QueryHistoryPreference:
    # one of: '' | 'query' | 'starred';
    home_tab: typing.Optional[str]

    def __init__(self, home_tab: typing.Optional[str] = None) -> None:
        self.home_tab = home_tab

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.home_tab is not None:
            payload["homeTab"] = self.home_tab
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "homeTab" in data:
            args["home_tab"] = data["homeTab"]        

        return cls(**args)


class NavbarPreference:
    bookmark_urls: list[str]

    def __init__(self, bookmark_urls: typing.Optional[list[str]] = None) -> None:
        self.bookmark_urls = bookmark_urls if bookmark_urls is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "bookmarkUrls": self.bookmark_urls,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "bookmarkUrls" in data:
            args["bookmark_urls"] = data["bookmarkUrls"]        

        return cls(**args)


class Preferences:
    # Explicit home URL (NOTE: this can only be modified in the system settings)
    home_url: typing.Optional[str]
    # UID for the home dashboard
    home_dashboard_uid: typing.Optional[str]
    # The timezone selection
    timezone: typing.Optional[str]
    # day of the week (sunday, monday, etc)
    week_start: typing.Optional[str]
    # user interface theme
    theme: typing.Optional[str]
    # Selected language
    language: typing.Optional[str]
    # Explore query history preferences
    query_history: typing.Optional['QueryHistoryPreference']
    # Navigation preferences
    navbar: typing.Optional['NavbarPreference']

    def __init__(self, home_url: typing.Optional[str] = None, home_dashboard_uid: typing.Optional[str] = None, timezone: typing.Optional[str] = None, week_start: typing.Optional[str] = None, theme: typing.Optional[str] = None, language: typing.Optional[str] = None, query_history: typing.Optional['QueryHistoryPreference'] = None, navbar: typing.Optional['NavbarPreference'] = None) -> None:
        self.home_url = home_url
        self.home_dashboard_uid = home_dashboard_uid
        self.timezone = timezone
        self.week_start = week_start
        self.theme = theme
        self.language = language
        self.query_history = query_history
        self.navbar = navbar

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.home_url is not None:
            payload["homeURL"] = self.home_url
        if self.home_dashboard_uid is not None:
            payload["homeDashboardUID"] = self.home_dashboard_uid
        if self.timezone is not None:
            payload["timezone"] = self.timezone
        if self.week_start is not None:
            payload["weekStart"] = self.week_start
        if self.theme is not None:
            payload["theme"] = self.theme
        if self.language is not None:
            payload["language"] = self.language
        if self.query_history is not None:
            payload["queryHistory"] = self.query_history
        if self.navbar is not None:
            payload["navbar"] = self.navbar
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "homeURL" in data:
            args["home_url"] = data["homeURL"]
        if "homeDashboardUID" in data:
            args["home_dashboard_uid"] = data["homeDashboardUID"]
        if "timezone" in data:
            args["timezone"] = data["timezone"]
        if "weekStart" in data:
            args["week_start"] = data["weekStart"]
        if "theme" in data:
            args["theme"] = data["theme"]
        if "language" in data:
            args["language"] = data["language"]
        if "queryHistory" in data:
            args["query_history"] = QueryHistoryPreference.from_json(data["queryHistory"])
        if "navbar" in data:
            args["navbar"] = NavbarPreference.from_json(data["navbar"])        

        return cls(**args)


PreferencesApiVersion: typing.Literal["preferences.grafana.app/v1alpha1"] = "preferences.grafana.app/v1alpha1"


PreferencesKind: typing.Literal["Preferences"] = "Preferences"



