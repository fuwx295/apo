import React, { useEffect, useState } from 'react'
import AreaChart from '../Chart/LineAreaChart'
import { BsArrowDown, BsArrowUp } from 'react-icons/bs'
import { MetricsLineChartColor } from 'src/constants'
import { convertTime } from 'src/utils/time'

const ArrowComponent = ({ value }) => {
  if (value === null || value === 0) {
    return null
  }

  return value < 0 ? <BsArrowDown color={'#24d160'} /> : <BsArrowUp color={'#ff3366'} />
}

const TempCell = (props) => {
  const { data, compare, type, timeRange } = props
  const [displayValue, setDisplayValue] = useState('')
  // const ArrowIcon = (props) =>{
  //   return props.type === 'up' ?
  // }
  const color = MetricsLineChartColor[type]
  const displayRatio = (value) => {
    if (value === null) {
      return '--'
    }
    let newValue = parseFloat((Math.floor(value * 100) / 100).toString())
    if (value < -99999) {
      newValue = '< -99999'
    } else if (0 > value && value > -0.01) {
      newValue = '> -0.01'
    } else if (0 < value && value < 0.01) {
      newValue = '< 0.01'
    } else if (value > 99999) {
      newValue = '> 99999'
    }
    newValue += '%'
    return newValue
  }
  useEffect(() => {
    if (data) {
      if (data.value === null) {
        setDisplayValue('N/A')
        return
      } else {
        let value = parseFloat((Math.floor(data.value * 100) / 100).toString())
        switch (type) {
          case 'latency':
            let convertValue = Math.floor((data.value / 1000) * 100) / 100
            if (data.value > 0 && data.value < 10) {
              value = '< 0.01 ms'
            } else {
              value = parseFloat(convertValue.toString()) + 'ms'
            }

            break
          case 'errorRate':
            if (data.value > 0 && data.value < 0.01) {
              value = '< 0.01%'
            } else {
              value += '%'
            }
            break
          case 'tps':
            if (data.value > 0 && data.value < 0.01) {
              value = '< 0.01'
            }
            value += `次/分`

            break
          case 'logs':
            if (data.value > 0 && data.value < 0.01) {
              value = '< 0.01'
            }
            value += `个`
            break
        }
        setDisplayValue(value)
      }
    }
  }, [type, data])
  return (
    data && (
      <div className="flex items-center flex-row flex-wrap justify-center pr-5">
        <div className="flex mr-1 flex-1 justify-end">
          {/* eslint-disable-next-line react/prop-types */}
          {displayValue}
        </div>

        <div className="felx h-[30px] items-center flex-1">
          <AreaChart color={color} data={data.chartData} timeRange={timeRange} />
        </div>

        <div className="h-full text-[10px] text-neutral-400 w-full">
          <div className="flex p-1 justify-center items-center">
            <span className="flex-1 text-right">日同比 </span>
            <span className="flex-1 inline-flex  justify-center items-center">
              {/* {data.ratio.dayOverDay !== null
                ? parseFloat(data.ratio.dayOverDay.toFixed(2)) + '%'
                : '--'} */}
              {displayRatio(data.ratio.dayOverDay)}
              <ArrowComponent value={data.ratio.dayOverDay} />
            </span>
          </div>
          <div className="flex p-1 justify-center items-center">
            <span className="flex-1 text-right">周同比 </span>
            <span className="flex-1 inline-flex  justify-center items-center">
              {displayRatio(data.ratio.weekOverDay)}{' '}
              {/* {data.ratio.weekOverDay !== null
                ? parseFloat(data.ratio.weekOverDay.toFixed(2)) + '%'
                : '--'} */}
              <ArrowComponent value={data.ratio.weekOverDay} />
            </span>
          </div>
        </div>
      </div>
    )
  )
}

export default TempCell