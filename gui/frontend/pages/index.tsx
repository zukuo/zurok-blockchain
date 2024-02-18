import type { NextPage } from 'next'
import Navigation from "../components/Navigation";
import Hero from "../components/Hero";
import React, {createContext, useState} from "react";
import Login from "../components/Login";

type NodeContextType = {
  node: string;
  setNode: React.Dispatch<React.SetStateAction<string>>
}

type LoginContextType = {
  isLoggedIn: boolean;
  setIsLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
}

export const NodeContext = createContext<null | NodeContextType>(null);
export const LoginContext = createContext<null | LoginContextType>(null);

const Home: NextPage = () => {
  const [node, setNode] = useState("")
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  return (
    <>
      <NodeContext.Provider value={{node: node, setNode: setNode}}>
        <Hero/>
        {!isLoggedIn ? <LoginContext.Provider value={{isLoggedIn: isLoggedIn, setIsLoggedIn: setIsLoggedIn}}>
          <Login />
        </LoginContext.Provider> : null}
        {isLoggedIn ? <Navigation/> : null}
      </NodeContext.Provider>
    </>
  )
}

export default Home
