<?php

namespace Grafana\Foundation\Alerting;

final class RuleConverter
{
    public static function convert(\Grafana\Foundation\Alerting\Rule $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\RuleBuilder('.\var_export($input->title, true).'))',
        ];
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
            if ($input->condition !== "") {
    
        
    $buffer = 'condition(';
        $arg0 =\var_export($input->condition, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->data) >= 1) {
    
        
    $buffer = 'queries(';
        $tmparg0 = [];
        foreach ($input->data as $arg1) {
        $tmpdataarg1 = \Grafana\Foundation\Alerting\QueryConverter::convert($arg1);
        $tmparg0[] = $tmpdataarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'execErrState(';
        $arg0 ='\Grafana\Foundation\Alerting\RuleExecErrState::fromValue("'.$input->execErrState.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->folderUID !== "") {
    
        
    $buffer = 'folderUID(';
        $arg0 =\var_export($input->folderUID, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->for !== "") {
    
        
    $buffer = 'for(';
        $arg0 =\var_export($input->for, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->id !== null) {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->isPaused !== null) {
    
        
    $buffer = 'isPaused(';
        $arg0 =\var_export($input->isPaused, true);
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
    $buffer = 'noDataState(';
        $arg0 ='\Grafana\Foundation\Alerting\RuleNoDataState::fromValue("'.$input->noDataState.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'orgID(';
        $arg0 =\var_export($input->orgID, true);
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
            if ($input->ruleGroup !== "") {
    
        
    $buffer = 'ruleGroup(';
        $arg0 =\var_export($input->ruleGroup, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
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
            if ($input->updated !== null) {
    
        
    $buffer = 'updated(';
        $arg0 =\var_export($input->updated, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

