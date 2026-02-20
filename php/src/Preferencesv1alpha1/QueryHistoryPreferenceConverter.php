<?php

namespace Grafana\Foundation\Preferencesv1alpha1;

final class QueryHistoryPreferenceConverter
{
    public static function convert(\Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreference $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreferenceBuilder())',
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

