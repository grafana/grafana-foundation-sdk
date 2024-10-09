<?php

namespace Grafana\Foundation\Alerting;

final class ContactPointConverter
{
    public static function convert(\Grafana\Foundation\Alerting\ContactPoint $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\ContactPointBuilder())',
        ];
            if ($input->disableResolveMessage !== null) {
    
        
    $buffer = 'disableResolveMessage(';
        $arg0 =\var_export($input->disableResolveMessage, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->name !== null && $input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->provenance !== null && $input->provenance !== "") {
    
        
    $buffer = 'provenance(';
        $arg0 =\var_export($input->provenance, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->settings !== null) {
    
        
    $buffer = 'settings(';
        $arg0 =\var_export($input->settings, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Alerting\ContactPointType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->uid !== null && $input->uid !== "") {
    
        
    $buffer = 'uid(';
        $arg0 =\var_export($input->uid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

