<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class RecordingRuleStatusConverter
{
    public static function convert(\Grafana\Foundation\Alertingv1beta1\RecordingRuleStatus $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv1beta1\RecordingRuleStatusBuilder())',
        ];
            if ($input->health !== null) {
    
        
    $buffer = 'health(';
        $arg0 ='\Grafana\Foundation\Alertingv1beta1\RecordingRuleHealth::fromValue("'.$input->health.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lastEvaluationTime !== null) {
    
        
    $buffer = 'lastEvaluationTime(';
        $arg0 =\var_export($input->lastEvaluationTime, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->evaluationDuration !== null) {
    
        
    $buffer = 'evaluationDuration(';
        $arg0 =\var_export($input->evaluationDuration, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lastError !== null && $input->lastError !== "") {
    
        
    $buffer = 'lastError(';
        $arg0 =\var_export($input->lastError, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

