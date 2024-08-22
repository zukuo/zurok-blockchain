import '../styles/globals.css'
import type { AppProps } from 'next/app'
import {NextUIProvider} from '@nextui-org/react'
import {ThemeProvider as NextThemesProvider} from "next-themes";
import {NodeContext, LoginContext, NavContext} from "./index";
import React, {useState} from "react";
import Hero from "../components/Hero";
import Login from "../components/Login";

function MyApp({ Component, pageProps }: AppProps) {
  const [node, setNode] = useState("")
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [selectedTab, setSelectedTab] = useState<string | number>("wallet")

  return (
    <NextUIProvider className={"bg-transparent"}>
      <NextThemesProvider attribute="class" defaultTheme="dark">
        <NodeContext.Provider value={{node: node, setNode: setNode}}>
          <LoginContext.Provider value={{isLoggedIn: isLoggedIn, setIsLoggedIn: setIsLoggedIn}}>
            <NavContext.Provider value={{selectedTab: selectedTab, setSelectedTab: setSelectedTab}}>
              <Hero/>
              {!isLoggedIn ? <Login /> : <Component {...pageProps} />}
            </NavContext.Provider>
          </LoginContext.Provider>
        </NodeContext.Provider>
      </NextThemesProvider>
    </NextUIProvider>
  )
}

export default MyApp
