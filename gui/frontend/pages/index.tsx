import type { NextPage } from 'next'
import Navigation from "../components/Navigation";
import React, {createContext, useState} from "react";
import Refresh from "../components/Refresh";

type NodeContextType = {
  node: string;
  setNode: React.Dispatch<React.SetStateAction<string>>
}

type LoginContextType = {
  isLoggedIn: boolean;
  setIsLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
}

type NavContextType = {
  selectedTab: string | number;
  setSelectedTab: React.Dispatch<React.SetStateAction<string | number>>
}

export const NodeContext = createContext<null | NodeContextType>(null);
export const LoginContext = createContext<null | LoginContextType>(null);
export const NavContext = createContext<null | NavContextType>(null);

const Home: NextPage = () => {

  return (
    <>
      <Navigation />
      <Refresh />
    </>
  )
}

export default Home
