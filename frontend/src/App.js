import React, { useState } from 'react';
import SignInSignUp from './page/SignInSignUp';


export default function App() {

  const [user, setUser] = useState({name: 'John'});

  if (user) {
    return (
      <div className="App">
        <SignInSignUp />
      </div>
    );
  } else {
    return (
      <div className="App">
        <h1>Hi Stranger</h1>
      </div>
    );
  }

}
