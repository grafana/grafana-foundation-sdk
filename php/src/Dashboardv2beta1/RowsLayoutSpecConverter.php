<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class RowsLayoutSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\RowsLayoutSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\RowsLayoutSpecBuilder())',
        ];
            if (count($input->rows) >= 1) {
    
        
    $buffer = 'rows(';
        $tmparg0 = [];
        foreach ($input->rows as $arg1) {
        $tmprowsarg1 = \Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowConverter::convert($arg1);
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

