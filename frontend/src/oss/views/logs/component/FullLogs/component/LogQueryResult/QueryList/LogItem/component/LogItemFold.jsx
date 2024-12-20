import { useLogsContext } from 'src/core/contexts/LogsContext';
import LogValueTag from './LogValueTag';
import LogKeyTag from './LogKeyTag';
import { useMemo } from 'react';

const LogItemFold = ({ tags, fields }) => {
  const { tableInfo, displayFields } = useLogsContext();

  // 计算过滤后的 tags
  const filteredTags = useMemo(() => {
    return Object.entries(tags).filter(([key, value]) => {
      return (
        displayFields.includes(key) &&  // 检查是否在显示字段中
        value !== '' &&  // 确保值不为空
        key !== (tableInfo?.timeField || 'timestamp') &&  // 排除时间字段
        typeof value !== 'object'  // 确保值不是对象
      );
    });
  }, [tags, displayFields, tableInfo?.timeField]);

  // 计算过滤后的 fields
  const filteredFields = useMemo(() => {
    return fields ? Object.entries(fields).filter(([key, value]) => displayFields.includes(key)) : [];
  }, [fields, displayFields]);

  return (
    <>
      {/* 渲染 tags */}
      <div className="text-ellipsis text-wrap flex" style={{ display: '-webkit-box' }}>
        {filteredTags.map(([key, value]) => (
          <LogValueTag key={key} objKey={key} value={String(value)} />
        ))}
      </div>

      {/* 渲染 fields */}
      <div className="text-ellipsis text-wrap flex flex-col overflow-hidden">
        {filteredFields.map(([key, value]) => (
          <div key={key}>
            <LogKeyTag title={key} description={value} />
          </div>
        ))}
      </div>
    </>
  );
};

export default LogItemFold;
