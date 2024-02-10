import React from 'react'
import {Input} from "@nextui-org/input";
import {MdSwapVert} from "react-icons/md";
import {Button} from "@nextui-org/button";
import { IoSend } from "react-icons/io5";

const Send = () => {
  return (
    <>
      {/* Sender */}
      <div className="flex justify-center font-mono">
        <Input type="email" label="Sender Address" className="w-1/3"/>
      </div>

      {/* Arrows */}
      <div className="py-4 flex justify-center">
        <MdSwapVert className="text-3xl text-success drop-shadow-2xl"/>
      </div>

      {/* Recipient */}
      <div className="flex justify-center font-mono">
        <Input type="email" label="Recipient Address" className="w-1/3"/>
      </div>

      {/* Button */}
      <div className="flex justify-center p-6">
        <Button endContent={<IoSend/>} radius="lg" color={"success"} className="text-black shadow-lg">
          Send!
        </Button>
      </div>
    </>
  )
}

export default Send