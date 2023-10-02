import axios from "axios";
const RESPAPI = 'http://localhost:8080'

export const getSalaries = async () => {
    const resp = await axios.get(`${RESPAPI}/salaries`)
    return resp;
}

export const getSalary = async (id) => {
    const resp = await axios.get(`${RESPAPI}/salary/${id}`)
    return resp;
}

export const postSalary = async (data) => {
    const resp = await axios.post(`${RESPAPI}/salary`, {
        debit: data.debit,
        salary: data.salary,
        moneyGain: data.moneyGain
    })
    return resp;
}
