import { combineReducers } from "redux";
import { persistReducer } from "redux-persist";
import timeRangePersistConfig from "../persist/timeRangePersistConfig";
import timeRangeReducer from "./timeRangeReducer";
import settingPersistConfig from "../persist/settingPresistConfig";
import settingReducer from "./settingReducer";

const rootReducer = combineReducers({
    timeRange: persistReducer(timeRangePersistConfig, timeRangeReducer),
    settingReducer: persistReducer(settingPersistConfig, settingReducer),
});
  
  export default rootReducer;