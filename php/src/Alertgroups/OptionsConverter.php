<?php

namespace Grafana\Foundation\Alertgroups;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Alertgroups\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertgroups\OptionsBuilder())',
        ];
            if ($input->labels !== "") {
    
        
    $buffer = 'labels(';
        $arg0 =\var_export($input->labels, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->alertmanager !== "") {
    
        
    $buffer = 'alertmanager(';
        $arg0 =\var_export($input->alertmanager, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'expandAll(';
        $arg0 =\var_export($input->expandAll, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

