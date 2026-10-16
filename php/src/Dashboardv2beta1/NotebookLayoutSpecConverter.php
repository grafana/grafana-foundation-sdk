<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class NotebookLayoutSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutSpecBuilder())',
        ];
            if (count($input->cells) >= 1) {
    
        
    $buffer = 'cells(';
        $tmparg0 = [];
        foreach ($input->cells as $arg1) {
        $tmpcellsarg1 = \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemKindConverter::convert($arg1);
        $tmparg0[] = $tmpcellsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

