<?php

namespace Grafana\Foundation\Preferences;

final class CookiePreferencesConverter
{
    public static function convert(\Grafana\Foundation\Preferences\CookiePreferences $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Preferences\CookiePreferencesBuilder())',
        ];
            if ($input->analytics !== null) {
    
        
    $buffer = 'analytics(';
        $arg0 =\var_export($input->analytics, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->performance !== null) {
    
        
    $buffer = 'performance(';
        $arg0 =\var_export($input->performance, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->functional !== null) {
    
        
    $buffer = 'functional(';
        $arg0 =\var_export($input->functional, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

