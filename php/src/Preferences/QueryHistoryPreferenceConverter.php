<?php

namespace Grafana\Foundation\Preferences;

final class QueryHistoryPreferenceConverter
{
    public static function convert(\Grafana\Foundation\Preferences\QueryHistoryPreference $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Preferences\QueryHistoryPreferenceBuilder())',
        ];
            if ($input->homeTab !== null && $input->homeTab !== "") {
    
        
    $buffer = 'homeTab(';
        $arg0 =\var_export($input->homeTab, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

