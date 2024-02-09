import type { NextPage } from 'next'
import Navigation from "../components/Navigation";
import Hero from "../components/Hero";

const Home: NextPage = () => {

  return (
    <>
      <Hero/>
      {/*<div className="p-10 mx-auto flex gap-4 items-center">*/}
        <Navigation/>
      {/*</div>*/}
    </>
  )
}

export default Home
