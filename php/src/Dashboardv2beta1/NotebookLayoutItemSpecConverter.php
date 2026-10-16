<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class NotebookLayoutItemSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpecBuilder())',
        ];
            
    
        {
    $buffer = 'element(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\ElementReferenceConverter::convert($input->element);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'source(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpecSource::fromValue("'.$input->source.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->collapsed !== null) {
    
        
    $buffer = 'collapsed(';
        $arg0 =\var_export($input->collapsed, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

