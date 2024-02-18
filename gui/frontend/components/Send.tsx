import React, {useContext, useEffect, useState} from 'react'
import {Input} from "@nextui-org/input";
import { IoMdArrowRoundDown } from "react-icons/io";
import {Button} from "@nextui-org/button";
import { IoSend } from "react-icons/io5";
import {GetAddresses, SendTransaction} from "../wailsjs/wailsjs/go/gui/App";
import {Autocomplete, AutocompleteItem} from "@nextui-org/react";
import {NodeContext} from "../pages";

const Send = () => {
  const node = useContext(NodeContext)!.node
  const [from, setFrom] = useState("")
  const [to, setTo] = useState("")
  const [amount, setAmount] = useState(0)
  const [sent, setSent] = useState(false)
  const [selectedKey, setSelectedKey] = React.useState<React.Key | null>(null);

  // Get JSON of addresses
  const [addresses, setAddresses] = useState([""])
  useEffect(() => {
    const fetchAddresses = () => {
      try {
        GetAddresses(node)
          .then(result => {
            setAddresses(result)
          })
          .catch(err => {
            console.log(err)
          })
      } catch (e) {
        console.error(e)
      }
    }
    fetchAddresses()
  }, [])

  // Send Functionality
  const sendData = () => {
    try {
      SendTransaction(from, to, amount, node, true)
        .then(() => {
          setSent(true)
        })
        .catch(e => {
          console.log(e)
        })
    } catch (e) {
      console.error(e)
    }
  }

  // Update autocomplete
  let items: {key: number, address: string}[] = []
  for (let i = 0; i < addresses.length ; i++) {
    items.push({key: i, address: addresses[i]})
  }

  return (
    <>
      {/* Sender */}
      <div className="flex justify-center font-mono">
        <Autocomplete
          variant={"bordered"}
          defaultItems={items}
          label="Sender Address"
          className="w-1/3"
          onSelectionChange={(key) => {setSelectedKey(key)}}
          onInputChange={(value) => {setFrom(value)}}
        >
          {(item) => <AutocompleteItem key={item.key}>{item.address}</AutocompleteItem>}
        </Autocomplete>
      </div>

      {/* Arrows */}
      <div className="py-4 flex justify-center">
        <IoMdArrowRoundDown className="text-3xl text-success drop-shadow-2xl"/>
      </div>

      {/* Recipient */}
      <div className="flex justify-center font-mono">
        <Autocomplete
          variant={"bordered"}
          defaultItems={items}
          label="Recipient Address"
          className="w-1/3"
          onSelectionChange={(key) => {setSelectedKey(key)}}
          onInputChange={(value) => {setTo(value)}}
        >
          {(item) => <AutocompleteItem key={item.key}>{item.address}</AutocompleteItem>}
        </Autocomplete>
      </div>

      {/* Amount */}
      <div className="flex justify-center font-mono pt-5">
        <Input
          label="Price"
          placeholder="0.00"
          variant={"bordered"}
          className="w-1/6"
          onChange={e => {setAmount(parseInt(e.target.value))}}
          startContent={
            <div className="pointer-events-none flex items-center">
              <span className="text-default-400 text-small">$</span>
            </div>
          }
          endContent={
            <div className="flex items-center">
              <label className="sr-only" htmlFor="currency">
                Currency
              </label>
            </div>
          }
          type="number"
        />
      </div>

      {/* Button */}
      <div className="flex justify-center p-6">
        <Button onClick={() => {sendData(); console.log(to)}} endContent={<IoSend/>} radius="lg" color={"success"} className="text-black shadow-lg">
          Send!
        </Button>
      </div>
    </>
  )
}

export default Send