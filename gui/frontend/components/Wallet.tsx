import React, {useEffect, useMemo, useState} from "react"
import {
  Table,
  TableHeader,
  TableColumn,
  TableBody,
  TableRow,
  TableCell,
  getKeyValue,
  Spinner,
  Pagination
} from "@nextui-org/react";
import {GetAddressesWithBalances} from "../wailsjs/wailsjs/go/gui/App"
import {Button} from "@nextui-org/button"
import { FaPlus } from "react-icons/fa";

const Wallet = () => {
  const node = "3000"
  let addressBalances: { key: number; address: string; balance: number; }[]  = []

  // Get JSON of key, address, balance
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

  // Check if it is loading data
  const loadingState =  wallets?.length === 0 ? "loading" : "idle"

  // Table Pagination
  const [page, setPage] = useState(1);
  const rowsPerPage = 10;
  const pages = Math.ceil(wallets.length / rowsPerPage);
  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;
    return wallets.slice(start, end);
  }, [page, wallets]);

  return (
    <>
      {/* Table */}
      <div className="flex justify-center">
        <Table aria-label="Wallet Addresses with Balances"
               className="max-h-[489px] min-w-[477px] max-w-[477px] font-mono"
               color={"success"}
               bottomContent={pages > 1 ?
                 <div className="flex w-full justify-center">
                   <Pagination
                     isCompact
                     showControls
                     showShadow
                     loop
                     color="success"
                     initialPage={1}
                     page={page}
                     total={pages}
                     onChange={(page) => setPage(page)}
                   />
                 </div> : null
               }>
          <TableHeader>
            <TableColumn key="address">Address</TableColumn>
            <TableColumn key="balance" align={"center"} width={150} className="text-center">Balance</TableColumn>
          </TableHeader>
          <TableBody items={items} loadingContent={<Spinner/>} loadingState={loadingState}>
            {(item) => (
              <TableRow key={item.key}>
                {(columnKey) => <TableCell className="m-auto text-center">{getKeyValue(item, columnKey)}</TableCell>}
              </TableRow>
            )}
          </TableBody>
          {/*<TableBody emptyContent={"No wallets found."}>{[]}</TableBody>*/}
        </Table>
      </div>

      {/* Button */}
      <div className="flex justify-center p-7">
        <Button endContent={<FaPlus/>} radius="lg" color={"success"} className="text-black shadow-lg">
          New Wallet
        </Button>
      </div>
    </>
  )
}

export default Wallet