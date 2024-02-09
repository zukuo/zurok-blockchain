import '../styles/globals.css'
import type { AppProps } from 'next/app'
import {NextUIProvider} from '@nextui-org/react'
import Navigation from "../components/Navigation";

function MyApp({ Component, pageProps }: AppProps) {
  return (
      <NextUIProvider className={"bg-transparent dark"}>
        <Component {...pageProps} />
      </NextUIProvider>
  )
}

export default MyApp
