import { Button, Card, Popconfirm, Tree } from 'antd'
import React, { useEffect, useState } from 'react'
import { useLogsContext } from 'src/contexts/LogsContext'
import ConfigLogRuleModal from '../../ConfigLogRuleModal'
import { MdAdd, MdDeleteOutline, MdModeEdit } from 'react-icons/md'
import './index.css'
import { deleteLogRuleApi } from 'src/api/logs'
import { showToast } from 'src/utils/toast'

const LogRuleList = () => {
  const { logRules, tableInfo, updateTableInfo, getLogTableInfo, updateLoading } = useLogsContext()
  const [treeData, setTreeData] = useState([])
  const [modalVisible, setModalVisible] = useState(false)
  const [selectedKeys, setSelectedKeys] = useState([])
  const [logRuleInfo, setLogRuleInfo] = useState(null)
  const editLogRule = (rule) => {
    setLogRuleInfo(rule)
    setModalVisible(true)
  }
  const deleteLogRule = (rule) => {
    updateLoading(true)
    deleteLogRuleApi({
      dataBase: rule.dataBase,
      parseName: rule.parseName,
      tableName: rule.tableName,
    }).then((res) => {
      showToast({
        title: '删除日志规则成功',
        color: 'success',
      })
      getLogTableInfo()
    })
  }
  const titleRender = (nodeData) => {
    return (
      <div className="logRuleItem">
        <div className="flex flex-col">
          <div>{nodeData.title}</div>
          <div className="text-xs text-gray-400">{nodeData.parseInfo}</div>
        </div>
        {!nodeData.isDefault && (
          <div className="action">
            <Button
              color="primary"
              variant="filled"
              icon={<MdModeEdit />}
              className="mr-2"
              size="small"
              onClick={(e) => {
                e.stopPropagation()
                editLogRule(nodeData)
              }}
            ></Button>
            <Popconfirm
              title={
                <>
                  是否确定删除名为“<span className="font-bold ">{nodeData.parseName}</span>
                  ”的日志规则
                </>
              }
              onConfirm={(e) => {
                e.stopPropagation()
                deleteLogRule(nodeData)
              }}
              okText="确定"
              cancelText="取消"
            >
              <Button
                size="small"
                color="danger"
                variant="filled"
                icon={<MdDeleteOutline />}
                onClick={(e) => {
                  e.stopPropagation()
                }}
              ></Button>
            </Popconfirm>
          </div>
        )}
      </div>
    )
  }
  useEffect(() => {
    setTreeData(
      logRules.map((rule, index) => ({
        key: rule.dataBase + rule.tableName,
        title: rule.parseName,
        ...rule,
        isDefault: index === 0,
      })),
    )
  }, [logRules])
  const onSelect = (selectedKeys, { selectedNodes }) => {
    updateTableInfo({
      dataBase: selectedNodes[0].dataBase,
      tableName: selectedNodes[0].tableName,
      cluster: '',
      parseName: selectedNodes[0].parseName,
    })
  }
  useEffect(() => {
    setSelectedKeys([tableInfo.dataBase + tableInfo.tableName])
  }, [tableInfo])

  return (
    <Card
      className="overflow-y-auto h-1/2 w-full overflow-x-hidden"
      title={
        <>
          <span>日志解析规则</span>
          <div className="flex flex-row">
            {/* <AiOutlineInfoCircle size={16} className="ml-1" /> */}
            <span className="text-xs text-gray-400">点击规则查询对应服务的日志</span>
          </div>
        </>
      }
      classNames={{
        body: 'p-0 pr-2',
      }}
      style={{ display: 'flex', flexDirection: 'column' }} // 设置 Card 的高度，使用 flexbox
      bodyStyle={{ flexGrow: 1, overflow: 'auto' }}
      extra={
        <Button
          type="primary"
          size="small"
          icon={<MdAdd size={20} />}
          onClick={() => setModalVisible(true)}
          className="flex-grow-0 flex-shrink-0"
        >
          {/* <span className="text-xs"></span> */}
        </Button>
      }
    >
      <Tree
        selectedKeys={selectedKeys}
        onSelect={onSelect}
        // onCheck={onCheck}
        treeData={treeData}
        titleRender={titleRender}
        className="pr-3 h-full"
        blockNode
      />
      <ConfigLogRuleModal
        modalVisible={modalVisible}
        closeModal={() => {
          setLogRuleInfo(null)
          setModalVisible(false)
        }}
        logRuleInfo={logRuleInfo}
      />
    </Card>
  )
}

export default LogRuleList