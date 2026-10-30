<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class RecordingRuleConverter
{
    public static function convert(\Grafana\Foundation\Alertingv1beta1\RecordingRuleSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv1beta1\RecordingRuleBuilder())',
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
        $arg0 = \Grafana\Foundation\Alertingv1beta1\IntervalTriggerConverter::convert($input->trigger);
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
            if ($input->metric !== "") {
    
        
    $buffer = 'metric(';
        $arg0 =\var_export($input->metric, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'expressions(';
        $arg0 = "[";
        foreach ($input->expressions as $key => $arg1) {
            $tmpexpressionsarg1 = \Grafana\Foundation\Alertingv1beta1\ExpressionConverter::convert($arg1);
            $arg0 .= "\t".var_export($key, true)." => $tmpexpressionsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->targetDatasourceUID !== "") {
    
        
    $buffer = 'targetDatasourceUID(';
        $arg0 =\var_export($input->targetDatasourceUID, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

