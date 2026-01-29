<?php

namespace Grafana\Foundation\Grafanapyroscope;

final class QueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Grafanapyroscope\QueryBuilder())',
        ];
            if ($input->version !== "" && $input->version !== "v0") {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasourceConverter::convert($input->datasource);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Grafanapyroscope\Dataquery && $input->spec->labelSelector !== "" && $input->spec->labelSelector !== "{}") {
    
        
    $buffer = 'labelSelector(';
        $arg0 =\var_export($input->spec->labelSelector, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Grafanapyroscope\Dataquery && $input->spec->spanSelector !== null && count($input->spec->spanSelector) >= 1) {
    
        
    $buffer = 'spanSelector(';
        $tmparg0 = [];
        foreach ($input->spec->spanSelector as $arg1) {
        $tmpspanSelectorarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpspanSelectorarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Grafanapyroscope\Dataquery && $input->spec->profileTypeId !== "") {
    
        
    $buffer = 'profileTypeId(';
        $arg0 =\var_export($input->spec->profileTypeId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Grafanapyroscope\Dataquery && count($input->spec->groupBy) >= 1) {
    
        
    $buffer = 'groupBy(';
        $tmparg0 = [];
        foreach ($input->spec->groupBy as $arg1) {
        $tmpgroupByarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpgroupByarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Grafanapyroscope\Dataquery && $input->spec->limit !== null) {
    
        
    $buffer = 'limit(';
        $arg0 =\var_export($input->spec->limit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Grafanapyroscope\Dataquery && $input->spec->maxNodes !== null) {
    
        
    $buffer = 'maxNodes(';
        $arg0 =\var_export($input->spec->maxNodes, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Grafanapyroscope\Dataquery && $input->spec->refId !== null && $input->spec->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->spec->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Grafanapyroscope\Dataquery && $input->spec->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->spec->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Grafanapyroscope\Dataquery && $input->spec->queryType !== null && $input->spec->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->spec->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Grafanapyroscope\Dataquery && $input->spec->datasource !== null) {
    
        
    $buffer = 'oldDatasource(';
        $arg0 ='(new \Grafana\Foundation\Common\DataSourceRef('.(($input->spec->datasource->type !== null) ? 'type: '.\var_export($input->spec->datasource->type, true).', ' : '').''.(($input->spec->datasource->uid !== null) ? 'uid: '.\var_export($input->spec->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

