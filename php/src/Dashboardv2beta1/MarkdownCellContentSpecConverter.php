<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class MarkdownCellContentSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentSpecBuilder())',
        ];
            if ($input->text !== "") {
    
        
    $buffer = 'text(';
        $arg0 =\var_export($input->text, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

