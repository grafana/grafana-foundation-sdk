<?php

namespace Grafana\Foundation\Datagrid;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Datagrid\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Datagrid\OptionsBuilder())',
        ];
            if ($input->selectedSeries !== 0) {
    
        
    $buffer = 'selectedSeries(';
        $arg0 =\var_export($input->selectedSeries, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

