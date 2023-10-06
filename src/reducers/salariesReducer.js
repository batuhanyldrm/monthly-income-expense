import { FETCH_SALARIES } from "../actions/types";

const salariesReducer = (state = {}, action) => {
    switch (action.type) {
        case FETCH_SALARIES:
            return {...action.payload}
        default:
            return state
    }
}

export default salariesReducer;