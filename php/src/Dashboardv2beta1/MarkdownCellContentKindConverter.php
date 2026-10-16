<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class MarkdownCellContentKindConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentKindBuilder())',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentSpecConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

