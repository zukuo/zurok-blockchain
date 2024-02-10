import React from 'react'
import {Input} from "@nextui-org/input";
import {MdSwapVert} from "react-icons/md";

const Send = () => {
  return (
    <div className="flex flex-col justify-center">
      <div>
        <Input type="email" label="Sender Address" className="justify-center"/>
      </div>
      <div className="py-4">
        <MdSwapVert className="text-3xl text-gray-700"/>
      </div>
    </div>
  )
}

export default Send