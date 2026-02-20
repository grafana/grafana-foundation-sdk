# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import preferencesv1alpha1
from ..builders import resource as resource_builder
from ..models import resource


class QueryHistoryPreference(cogbuilder.Builder[preferencesv1alpha1.QueryHistoryPreference]):
    _internal: preferencesv1alpha1.QueryHistoryPreference

    def __init__(self) -> None:
        self._internal = preferencesv1alpha1.QueryHistoryPreference()

    def build(self) -> preferencesv1alpha1.QueryHistoryPreference:
        """
        Builds the object.
        """
        return self._internal    
    
    def home_tab(self, home_tab: str) -> typing.Self:    
        """
        one of: '' | 'query' | 'starred';
        """
            
        self._internal.home_tab = home_tab
    
        return self
    


class CookiePreferences(cogbuilder.Builder[preferencesv1alpha1.CookiePreferences]):
    _internal: preferencesv1alpha1.CookiePreferences

    def __init__(self) -> None:
        self._internal = preferencesv1alpha1.CookiePreferences()

    def build(self) -> preferencesv1alpha1.CookiePreferences:
        """
        Builds the object.
        """
        return self._internal    
    
    def analytics(self, analytics: object) -> typing.Self:    
        self._internal.analytics = analytics
    
        return self
    
    def performance(self, performance: object) -> typing.Self:    
        self._internal.performance = performance
    
        return self
    
    def functional(self, functional: object) -> typing.Self:    
        self._internal.functional = functional
    
        return self
    


class NavbarPreference(cogbuilder.Builder[preferencesv1alpha1.NavbarPreference]):
    _internal: preferencesv1alpha1.NavbarPreference

    def __init__(self) -> None:
        self._internal = preferencesv1alpha1.NavbarPreference()

    def build(self) -> preferencesv1alpha1.NavbarPreference:
        """
        Builds the object.
        """
        return self._internal    
    
    def bookmark_urls(self, bookmark_urls: list[str]) -> typing.Self:    
        self._internal.bookmark_urls = bookmark_urls
    
        return self
    


class Preferences(cogbuilder.Builder[preferencesv1alpha1.Preferences]):
    _internal: preferencesv1alpha1.Preferences

    def __init__(self) -> None:
        self._internal = preferencesv1alpha1.Preferences()

    def build(self) -> preferencesv1alpha1.Preferences:
        """
        Builds the object.
        """
        return self._internal    
    
    def home_dashboard_uid(self, home_dashboard_uid: str) -> typing.Self:    
        """
        UID for the home dashboard
        """
            
        self._internal.home_dashboard_uid = home_dashboard_uid
    
        return self
    
    def timezone(self, timezone: str) -> typing.Self:    
        """
        The timezone selection
        TODO: this should use the timezone defined in common
        """
            
        self._internal.timezone = timezone
    
        return self
    
    def week_start(self, week_start: str) -> typing.Self:    
        """
        day of the week (sunday, monday, etc)
        """
            
        self._internal.week_start = week_start
    
        return self
    
    def theme(self, theme: str) -> typing.Self:    
        """
        light, dark, empty is default
        """
            
        self._internal.theme = theme
    
        return self
    
    def language(self, language: str) -> typing.Self:    
        """
        Selected language (beta)
        """
            
        self._internal.language = language
    
        return self
    
    def regional_format(self, regional_format: str) -> typing.Self:    
        """
        Selected locale (beta)
        """
            
        self._internal.regional_format = regional_format
    
        return self
    
    def query_history(self, query_history: cogbuilder.Builder[preferencesv1alpha1.QueryHistoryPreference]) -> typing.Self:    
        """
        Explore query history preferences
        """
            
        query_history_resource = query_history.build()
        self._internal.query_history = query_history_resource
    
        return self
    
    def cookie_preferences(self, cookie_preferences: cogbuilder.Builder[preferencesv1alpha1.CookiePreferences]) -> typing.Self:    
        """
        Cookie preferences
        """
            
        cookie_preferences_resource = cookie_preferences.build()
        self._internal.cookie_preferences = cookie_preferences_resource
    
        return self
    
    def navbar(self, navbar: cogbuilder.Builder[preferencesv1alpha1.NavbarPreference]) -> typing.Self:    
        """
        Navigation preferences
        """
            
        navbar_resource = navbar.build()
        self._internal.navbar = navbar_resource
    
        return self
    

"""
Creates a resource manifest from Preferences.
"""
def manifest(name: str, preferences: cogbuilder.Builder[preferencesv1alpha1.Preferences]) -> cogbuilder.Builder[resource.Manifest]:
    builder = resource_builder.Manifest()
    builder.api_version("preferences.grafana.app/v1alpha1")
    builder.kind("Preferences")
    builder.metadata(resource_builder.named(name))
    builder.spec(preferences.build())

    return builder
