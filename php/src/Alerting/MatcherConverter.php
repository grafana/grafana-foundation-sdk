<?php

namespace Grafana\Foundation\Alerting;

final class MatcherConverter
{
    public static function convert(\Grafana\Foundation\Alerting\Matcher $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\MatcherBuilder())',
        ];
            if ($input->name !== null && $input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->type !== null) {
    
        
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Alerting\MatchType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->value !== null && $input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

