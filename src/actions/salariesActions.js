import { getSalaries } from "../api/salaryApi"
import { FETCH_SALARIES } from "./types"

export const fetchSalaries = () => async (
    dispatch
) => {
    const resp = await getSalaries()
    dispatch({
        type: FETCH_SALARIES,
        payload: resp.data
    })
}