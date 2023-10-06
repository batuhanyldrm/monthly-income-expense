import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';
import { fetchSalaries } from '../actions/salariesActions';

function Salary(props) {

  const {fetchSalaries} = props;

  useEffect(() => {
    fetchSalaries()
  }, []);
  

  return (
    <div>
      Proje
    </div>
  );
}

const mapStateToProps = (state) => ({
});

const mapDispatchToProps = (dispatch) => ({
    fetchSalaries: (data) => {
      dispatch(fetchSalaries(data));
    },
});

export default connect(mapStateToProps,mapDispatchToProps) (Salary);
