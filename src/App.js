
import './App.css';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';
import logo from './resources/People.jpeg';
import React from 'react';
import axios from 'axios';
import TextField from '@material-ui/core/TextField';
import Alert from '@material-ui/lab/Alert';
import CustomersTable from './components/CustomersTable';

const initialstate = {
  name: '',
  email: '',
  phonenumbers: '',
  success: '',
  button: 0,
  mapOfCustomer: []
  
}

class App extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = initialstate;

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }


  handleChange = event => {
    if (event.target.id === 'name') {
      this.setState({ name: event.target.value });
  } else if (event.target.id === 'email') {
      this.setState({ email: event.target.value });
  } else if (event.target.id === 'phonenumbers') {
      this.setState({ phonenumbers: event.target.value });
  }
 
  }


  handleSubmit(event) {
    console.log(event.target.id);
    var allcustomers = [];
    var insertion;
    alert("Form was submitted");

            const user = {
              name: this.state.name,
              email: this.state.email,
              phonenumbers: this.state.phonenumbers
            };
             
            (this.state.button === 1) ? insertion ="true" : insertion ="false";

            axios({
              method: 'POST',
              baseURL:`http://localhost:8080/`, 
              params: {'name': this.state.name,'email': this.state.email,
                     'phonenumbers': this.state.phonenumbers, 'insertion': insertion
                },
          }).then(res => {
          
          console.log(res.data);
          if(res.data === true) {
            console.log("The record is inserted successfully ! ")
            this.setState({ success : true });
          } else if(res.data === false) {
            this.setState({ success : false });
          } else  {
            
            var responseMap = new Map();
            responseMap = res.data;
          
           
            
            Object.keys(responseMap).map((customerName,i) => {
               console.log(responseMap[i]);
               allcustomers[i] = responseMap[i];
            })
            
            this.setState({mapOfCustomer: allcustomers})
          }
        })
    
    event.preventDefault();
  }

  render() {
  return (
    
    <div className="App">
         <img src={logo} className="" alt="logo" />
        <Grid
        container
        direction="column"
        justifyContent="center"
        alignItems="center"
       >
        <form onSubmit={this.handleSubmit}>
          <Grid direction="column" container spacing={4}>
               <Grid item xs>
                  <TextField
                  id="name"
                  label="Name"
                  multiline
                  maxRows={1}
                  onChange={this.handleChange}
                />
                </Grid>
                <Grid item xs>
                <TextField
                    id="email"
                    label="Email"
                    multiline
                    maxRows={1}
                    onChange={this.handleChange}
                  />
                </Grid>
                <Grid item xs>
                    <TextField
                          id="phonenumbers"
                          label="Phone Numbers"
                          placeholder="XXXXXXXXXX"
                          multiline
                          variant="filled"
                          onChange = {this.handleChange}
                        />
                
              </Grid>
              <Button id="submit" style={{ marginTop:`20px` }} variant="contained" color="secondary" type="submit" onClick={() => this.state.button = 1}>
                Submit
              </Button>
              <Button id="records" style={{ marginTop:`20px` }} variant="contained" color="primary" type="submit" onClick={() => this.state.button = 2}>
                All Records
              </Button>
              <Grid>
              {this.state.success ? <div>
                <Alert variant="outlined" severity="success">
                Successfull customer record inserted into Database! Click on All Records to View all records ! 
                          </Alert>

              </div> : null }
              {this.state.success === false ? <div><Alert severity="error">Error in inserting customer record ! Contact Administrator</Alert></div>:null}
          </Grid>
          </Grid>
        </form>
      </Grid>
         
     <div>
     {this.state.mapOfCustomer ? <CustomersTable customer= {this.state.mapOfCustomer}></CustomersTable> : null }
     </div>
    </div> 
  );
}
}
export default App;
