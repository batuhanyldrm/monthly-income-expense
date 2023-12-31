import React from "react";
import { IconButton, TableCell, TableRow } from "@material-ui/core";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';

const SalaryListItem = (props) => {

  const {
    salary
  } = props;
  
  return (
    <>
    {salary.users && salary.users.map((user) =>
      <TableRow
        key={salary.id}
        sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
      >
        <TableCell component="th" scope="row" align="left"> {user.name}</TableCell>
        <TableCell component="th" scope="row" align="left"> {user.email}</TableCell>
        <TableCell component="th" scope="row" align="left"> {salary.salary}</TableCell>
        <TableCell align="left">{salary.debit}</TableCell>
        <TableCell align="center">{salary.moneyGain}</TableCell>
        <TableCell align="right">
        <IconButton
        /* onClick={()=>deleteProduct(product.id)} */
        >
            <DeleteIcon/>
        </IconButton>
        </TableCell>
        <TableCell align="right">
        <IconButton
        /* onClick={()=>handleEdit(product.id)} */
        >
            <EditIcon/>
        </IconButton>
        </TableCell>
      </TableRow>
      )}
    </>
  )
};

export default SalaryListItem;
