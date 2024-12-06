<?php

namespace Grafana\Foundation\Preferences;

final class PreferencesConverter
{
    public static function convert(\Grafana\Foundation\Preferences\Preferences $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Preferences\PreferencesBuilder())',
        ];
            if ($input->homeDashboardUID !== null && $input->homeDashboardUID !== "") {
    
        
    $buffer = 'homeDashboardUID(';
        $arg0 =\var_export($input->homeDashboardUID, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timezone !== null && $input->timezone !== "") {
    
        
    $buffer = 'timezone(';
        $arg0 =\var_export($input->timezone, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->weekStart !== null && $input->weekStart !== "") {
    
        
    $buffer = 'weekStart(';
        $arg0 =\var_export($input->weekStart, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->theme !== null && $input->theme !== "") {
    
        
    $buffer = 'theme(';
        $arg0 =\var_export($input->theme, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->language !== null && $input->language !== "") {
    
        
    $buffer = 'language(';
        $arg0 =\var_export($input->language, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryHistory !== null) {
    
        
    $buffer = 'queryHistory(';
        $arg0 = \Grafana\Foundation\Preferences\QueryHistoryPreferenceConverter::convert($input->queryHistory);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->cookiePreferences !== null) {
    
        
    $buffer = 'cookiePreferences(';
        $arg0 = \Grafana\Foundation\Preferences\CookiePreferencesConverter::convert($input->cookiePreferences);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->navbar !== null) {
    
        
    $buffer = 'navbar(';
        $arg0 = \Grafana\Foundation\Preferences\NavbarPreferenceConverter::convert($input->navbar);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

