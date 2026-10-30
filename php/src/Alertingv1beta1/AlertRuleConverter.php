<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class AlertRuleConverter
{
    public static function convert(\Grafana\Foundation\Alertingv1beta1\AlertRuleSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv1beta1\AlertRuleBuilder())',
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
            if ($input->annotations !== null) {
    
        
    $buffer = 'annotations(';
        $arg0 = "[";
        foreach ($input->annotations as $key => $arg1) {
            $tmpannotationsarg1 =\var_export($arg1, true);
            $arg0 .= "\t".var_export($key, true)." => $tmpannotationsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->for !== null && $input->for !== "") {
    
        
    $buffer = 'for(';
        $arg0 =\var_export($input->for, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->keepFiringFor !== null && $input->keepFiringFor !== "") {
    
        
    $buffer = 'keepFiringFor(';
        $arg0 =\var_export($input->keepFiringFor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->missingSeriesEvalsToResolve !== null) {
    
        
    $buffer = 'missingSeriesEvalsToResolve(';
        $arg0 =\var_export($input->missingSeriesEvalsToResolve, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'noDataState(';
        $arg0 ='\Grafana\Foundation\Alertingv1beta1\NoDataState::fromValue("'.$input->noDataState.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'execErrState(';
        $arg0 ='\Grafana\Foundation\Alertingv1beta1\ExecErrState::fromValue("'.$input->execErrState.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'notificationSettings(';
        switch (true) {
            case $input->notificationSettings instanceof \Grafana\Foundation\Alertingv1beta1\SimplifiedRouting:
                $disjunctionnotificationSettings = \Grafana\Foundation\Alertingv1beta1\SimplifiedRoutingConverter::convert($input->notificationSettings);
                $arg0 = $disjunctionnotificationSettings;
                break;
            case $input->notificationSettings instanceof \Grafana\Foundation\Alertingv1beta1\NamedRoutingTree:
                $disjunctionnotificationSettings = \Grafana\Foundation\Alertingv1beta1\NamedRoutingTreeConverter::convert($input->notificationSettings);
                $arg0 = $disjunctionnotificationSettings;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
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
    
    
            if ($input->panelRef !== null) {
    
        
    $buffer = 'panelRef(';
        $arg0 = \Grafana\Foundation\Alertingv1beta1\PanelRefConverter::convert($input->panelRef);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

