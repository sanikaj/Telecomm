
import Grid from '@material-ui/core/Grid';
import React from 'react';
import  './Customers.css';

class CustomersTable extends React.Component {
   

    render () { return (
        this.props ? <React.Fragment>
            {Object.keys(this.props.customer).length === 0 ? <h2>All Customer Records </h2> : null}
        { Object.keys(this.props.customer).map((customerName,i)=> {
          return (
           
  
                    <div style={{ marginTop:`40px` }}>
                       
                        <Grid
                            container
                            spacing={12}
                            alignItems="center"
                            justifyContent="center"
                            
                        >
                        <table >
                        
                        <tr>
                            <th>ID</th>
                            <th>Name</th>
                            <th>Email</th>
                            <th>Created Date</th>
                            <th>Phone Numbers</th>
                        </tr>
                      
                       
                        <tr>
                            <td>{this.props.customer[i].CustomerId}</td>
                            <td>{this.props.customer[i].Name}</td>
                            <td>{this.props.customer[i].Email}</td>
                            <td>{this.props.customer[i].CreatedDate}</td>
                            <td>{this.props.customer[i].Phones}</td>
                        </tr>
                      
                            </table>
                      </Grid>
                      </div>
                    
          );
        })}
      </React.Fragment> : null
    );
  }
}

export default CustomersTable;