import '../styles/globals.css'
import type { AppProps } from 'next/app'
import {NextUIProvider} from '@nextui-org/react'
import {ThemeProvider as NextThemesProvider} from "next-themes";
import {NodeContext, LoginContext} from "./index";
import React, {useState} from "react";
import Hero from "../components/Hero";
import Login from "../components/Login";

function MyApp({ Component, pageProps }: AppProps) {
  const [node, setNode] = useState("")
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  return (
    <NextUIProvider className={"bg-transparent"}>
      <NextThemesProvider attribute="class" defaultTheme="dark">
        <NodeContext.Provider value={{node: node, setNode: setNode}}>
          <LoginContext.Provider value={{isLoggedIn: isLoggedIn, setIsLoggedIn: setIsLoggedIn}}>
            <Hero/>
            {!isLoggedIn ? <Login /> : <Component {...pageProps} />}
          </LoginContext.Provider>
        </NodeContext.Provider>
      </NextThemesProvider>
    </NextUIProvider>
  )
}

export default MyApp
