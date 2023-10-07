import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';
import { fetchSalaries } from '../actions/salariesActions';
import { Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from '@material-ui/core';
import SalaryListItem from './SalaryListItem';

function Salary(props) {

  const {fetchSalaries, salaries} = props;

  useEffect(() => {
    fetchSalaries()
  }, []);
  
console.log(salaries, "www")
  return (
    <div>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 5 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>Salary</TableCell>
              <TableCell align="left">Debit</TableCell>
              <TableCell align="right">Money Gain</TableCell>
              {/* <TableCell align="right">Amount</TableCell>
              <TableCell align="right">Delete</TableCell>
              <TableCell align="right">Edit</TableCell> */}
            </TableRow>
          </TableHead>
          <TableBody>
            {/* {salaries && salaries.map((salary, index) => (
              <SalaryListItem
              salary={salary}
              index={index}
              key={salary.id + "" + index}
              />
            ))} */}
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  );
}

const mapStateToProps = (state) => ({
  salaries: state.salaries
});

const mapDispatchToProps = (dispatch) => ({
    fetchSalaries: (data) => {
      dispatch(fetchSalaries(data));
    },
});

export default connect(mapStateToProps,mapDispatchToProps) (Salary);
