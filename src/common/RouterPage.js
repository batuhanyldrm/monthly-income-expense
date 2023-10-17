import { BrowserRouter, Route, Routes} from "react-router-dom";
import React from 'react'
import Salary from "../components/Salary";

const RouterPage = () => {
  return (
    <BrowserRouter>
        <Routes>
          <Route path="/" element={<Salary component={Salary} title={"Salary"} />}></Route>
        </Routes>
          
    </BrowserRouter>
  )
}

export default RouterPage