<?php

namespace Grafana\Foundation\Preferencesv1alpha1;

final class PreferencesConverter
{
    public static function convert(\Grafana\Foundation\Preferencesv1alpha1\Preferences $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Preferencesv1alpha1\PreferencesBuilder())',
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
            if ($input->regionalFormat !== null && $input->regionalFormat !== "") {
    
        
    $buffer = 'regionalFormat(';
        $arg0 =\var_export($input->regionalFormat, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryHistory !== null) {
    
        
    $buffer = 'queryHistory(';
        $arg0 = \Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreferenceConverter::convert($input->queryHistory);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->cookiePreferences !== null) {
    
        
    $buffer = 'cookiePreferences(';
        $arg0 = \Grafana\Foundation\Preferencesv1alpha1\CookiePreferencesConverter::convert($input->cookiePreferences);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->navbar !== null) {
    
        
    $buffer = 'navbar(';
        $arg0 = \Grafana\Foundation\Preferencesv1alpha1\NavbarPreferenceConverter::convert($input->navbar);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

