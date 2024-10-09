<?php

namespace Grafana\Foundation\Alerting;

final class NotificationTemplateConverter
{
    public static function convert(\Grafana\Foundation\Alerting\NotificationTemplate $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\NotificationTemplateBuilder())',
        ];
            if ($input->name !== null && $input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->provenance !== "") {
    
        
    $buffer = 'provenance(';
        $arg0 =\var_export($input->provenance, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->template !== null && $input->template !== "") {
    
        
    $buffer = 'template(';
        $arg0 =\var_export($input->template, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

