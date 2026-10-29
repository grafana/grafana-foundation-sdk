<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class AlertRuleStatusConverter
{
    public static function convert(\Grafana\Foundation\Alertingv1beta1\AlertRuleStatus $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv1beta1\AlertRuleStatusBuilder())',
        ];
            if ($input->health !== null) {
    
        
    $buffer = 'health(';
        $arg0 ='\Grafana\Foundation\Alertingv1beta1\AlertRuleHealth::fromValue("'.$input->health.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->state !== null) {
    
        
    $buffer = 'state(';
        $arg0 ='\Grafana\Foundation\Alertingv1beta1\AlertRuleState::fromValue("'.$input->state.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->stateReason !== null) {
    
        
    $buffer = 'stateReason(';
        $arg0 ='\Grafana\Foundation\Alertingv1beta1\AlertRuleStateReason::fromValue("'.$input->stateReason.'")';
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
            if ($input->totals !== null) {
    
        
    $buffer = 'totals(';
        $arg0 = \Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotalsConverter::convert($input->totals);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

