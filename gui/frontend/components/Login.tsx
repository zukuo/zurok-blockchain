import React, {useContext} from 'react';
import {LoginContext, NodeContext} from "../pages";
import {Input} from "@nextui-org/input";
import {Button} from "@nextui-org/button";
import {FiLogIn} from "react-icons/fi";

const Login = () => {
  const nodeContext = useContext(NodeContext)
  const loginContext = useContext(LoginContext)

  return (
    <>
      {/* Node Input */}
      <div className="flex justify-center font-mono pt-6">
        <Input variant={"bordered"}
               label="Node Number"
               className={"w-1/3"}
               onChange={(value) => {
                 nodeContext?.setNode(value.target.value)
               }}/>
      </div>

      {/* Button */}
      {/* TODO: If the NODE doesn't exist don't login */}
      <div className="flex justify-center p-6">
        <Button endContent={<FiLogIn/>}
                radius="lg"
                color={"success"}
                onClick={() => loginContext?.setIsLoggedIn(true)}>
          Enter
        </Button>
      </div>
    </>
  )
}

export default Login;