import React, {useEffect, useState} from 'react'
import {Input} from "@nextui-org/input";
import {MdSwapVert} from "react-icons/md";
import {Button} from "@nextui-org/button";
import { IoSend } from "react-icons/io5";
import {GetAddresses} from "../wailsjs/wailsjs/go/gui/App";
import {Autocomplete, AutocompleteItem} from "@nextui-org/react";

const Send = () => {
  const node = "3000"

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

  let items: {key: number, address: string}[] = []
  for (let i = 0; i < addresses.length ; i++) {
    items.push({key: i, address: addresses[i]})
  }

  return (
    <>
      {/* Sender */}
      <div className="flex justify-center font-mono">
        <Autocomplete
          defaultItems={items}
          label="Sender Wallet"
          placeholder="Search an address"
          className="w-1/3"
        >
          {(item) => <AutocompleteItem key={item.key}>{item.address}</AutocompleteItem>}
        </Autocomplete>
      </div>

      {/* Arrows */}
      <div className="py-4 flex justify-center">
        <MdSwapVert className="text-3xl text-success drop-shadow-2xl"/>
      </div>

      {/* Recipient */}
      <div className="flex justify-center font-mono">
        <Autocomplete
          defaultItems={items}
          label="Recipient Wallet"
          placeholder="Search an address"
          className="w-1/3"
        >
          {(item) => <AutocompleteItem key={item.key}>{item.address}</AutocompleteItem>}
        </Autocomplete>
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