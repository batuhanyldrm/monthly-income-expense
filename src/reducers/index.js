import {combineReducers} from "redux";
import salariesReducer from "./salariesReducer";

const reducers = combineReducers({
    salaries: salariesReducer
});

export default reducers;