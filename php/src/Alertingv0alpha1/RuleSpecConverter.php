<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class RuleSpecConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\RuleSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\RuleSpecBuilder())',
        ];
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->paused !== null) {
    
        
    $buffer = 'paused(';
        $arg0 =\var_export($input->paused, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'trigger(';
        $arg0 = \Grafana\Foundation\Alertingv0alpha1\IntervalTriggerConverter::convert($input->trigger);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->labels !== null) {
    
        
    $buffer = 'labels(';
        $arg0 = "[";
        foreach ($input->labels as $key => $arg1) {
            $tmplabelsarg1 =\var_export($arg1, true);
            $arg0 .= "\t".var_export($key, true)." => $tmplabelsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'expressions(';
        $arg0 = "[";
        foreach ($input->expressions as $key => $arg1) {
            $tmpexpressionsarg1 = \Grafana\Foundation\Alertingv0alpha1\ExpressionConverter::convert($arg1);
            $arg0 .= "\t".var_export($key, true)." => $tmpexpressionsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

