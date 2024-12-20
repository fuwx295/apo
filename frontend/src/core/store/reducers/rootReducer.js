import { combineReducers } from 'redux'
import { persistReducer } from 'redux-persist'
import timeRangePersistConfig from '../persist/timeRangePersistConfig'
import timeRangeReducer from './timeRangeReducer'
import settingPersistConfig from '../persist/settingPresistConfig'
import settingReducer from './settingReducer'
import topologyPresistConfig from '../persist/topologyPresistConfig'
import topologyReducer from './topologyReducer'
import urlParamsReducer from './urlParamsReducer'
import urlParamsPresistConfig from '../persist/urlParamsPresistConfig'
import groupLabelReducer from './groupLabelReducer'
import groupLabelPresistConfig from '../persist/groupLabelPresistConfig'
import userReducer from './userReducer'
import userPersistConfig from '../persist/userPresistConfig'
import personalizedSettingReducer from './personalizedSettingReducer'
import personalizedSettingPersistConfig from '../persist/personalizedSettingPersistConfig'

const rootReducer = combineReducers({
  timeRange: persistReducer(timeRangePersistConfig, timeRangeReducer),
  settingReducer: persistReducer(settingPersistConfig, settingReducer),
  topologyReducer: persistReducer(topologyPresistConfig, topologyReducer),
  urlParamsReducer: persistReducer(urlParamsPresistConfig, urlParamsReducer),
  groupLabelReducer: persistReducer(groupLabelPresistConfig, groupLabelReducer),
  userReducer: persistReducer(userPersistConfig, userReducer),
  personalizedSetting: persistReducer(personalizedSettingPersistConfig, personalizedSettingReducer)
})

export default rootReducer
