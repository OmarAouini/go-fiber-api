import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import {Switch, Route} from 'react-router-dom';
import Home from './components/Home';
import Restreamers from './components/Restreamers';

function App() {
  return (
    <div className="App">
      <Switch>
              <Route exact path={["/", "/home"]} component={Home}/>
              <Route path='/restreamers' component={Restreamers} />
              <Route path='/about'  />
      </Switch>
    </div>
  );
}

export default App;
