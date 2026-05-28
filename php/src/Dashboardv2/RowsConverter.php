<?php

namespace Grafana\Foundation\Dashboardv2;

final class RowsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\RowsLayoutKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\RowsBuilder())',
        ];
            if (count($input->spec->rows) >= 1) {
    
        
    $buffer = 'rows(';
        $tmparg0 = [];
        foreach ($input->spec->rows as $arg1) {
        $tmprowsarg1 = \Grafana\Foundation\Dashboardv2\RowConverter::convert($arg1);
        $tmparg0[] = $tmprowsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

