import React, {useEffect, useState} from "react";
import {Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, getKeyValue} from "@nextui-org/react";
import {GetAddressesWithBalances} from "../wailsjs/wailsjs/go/gui/App"

const Wallet = () => {
  const node = "3000"
  let addressBalances: { key: number; address: string; balance: number; }[]  = []

  const [wallets, setWallets] = useState(addressBalances)

  useEffect(() => {
    const fetchWallets = () => {
      try {
        GetAddressesWithBalances(node)
          .then(result => {
            setWallets(result)
          })
          .catch(err => {
            console.log(err)
          })
      } catch (e) {
        console.error(e)
      }
    }
    fetchWallets()
  }, [])

  return (
    <div className="flex justify-center">
      <Table aria-label="Wallet Addresses with Balances" className="w-2/3 text-gray-300">
        <TableHeader>
          <TableColumn key="address">Address</TableColumn>
          <TableColumn key="balance" align={"center"} width={150}>Balance</TableColumn>
        </TableHeader>
        <TableBody items={wallets}>
          {(item) => (
            <TableRow key={item.key}>
              {(columnKey) => <TableCell>{getKeyValue(item, columnKey)}</TableCell>}
            </TableRow>
          )}
        </TableBody>
        {/*<TableBody emptyContent={"No wallets found."}>{[]}</TableBody>*/}
      </Table>
    </div>
  )
}

export default Wallet