import React, {useContext, useEffect, useMemo, useState} from 'react';
import {NodeContext} from "../pages";
import {GetTransactions} from "../wailsjs/wailsjs/go/gui/App";
import {
  getKeyValue,
  Pagination,
  Spinner,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow
} from "@nextui-org/react";
import Link from "next/link";
import {gui} from "../wailsjs/wailsjs/go/models";

const Transactions = (blockHeight: number = 0) => {
  const node = useContext(NodeContext)!.node

  let txsArr: gui.transactions[] = []
  const [transactions, setTransactions] = useState(txsArr)

  useEffect(() => {
    const fetchTransactions = () => {
      try {
        GetTransactions(node)
          .then(result => {
            let filteredResults = result.filter((tx) =>
              tx.height == blockHeight
            );
            setTransactions(filteredResults)
          })
          .catch(err => {
            console.log(err)
          })
      } catch (e) {
        console.error(e)
      }
    }
    fetchTransactions()
  }, [transactions])


  const loadingState =  transactions?.length === 0 ? "loading" : "idle"

  // Table Pagination
  const [page, setPage] = useState(1);
  const rowsPerPage = 10;
  const pages = Math.ceil(transactions.length / rowsPerPage);
  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;
    return transactions.slice(start, end);
  }, [page, transactions]);

  return (
    <div className="flex justify-center">
      <Table aria-label="Blocks"
             className="font-mono max-w-fit"
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
          {/*<TableColumn key="height">Block</TableColumn>*/}
          <TableColumn key="transaction">Transaction</TableColumn>
          <TableColumn key="amount">Amount</TableColumn>
        </TableHeader>
        <TableBody items={items} loadingContent={<Spinner/>} loadingState={loadingState}>
          {(item) => (
            <TableRow key={item.key}>
              {(columnKey) => (
                <TableCell className="m-auto text-center">
                  {(columnKey == "transaction")
                    ? <Link className={"text-small text-foreground hover:text-success transition"} href={{
                      pathname: "/transactions/[id]",
                      query: {
                        id: item.key,
                        transaction: item.transaction,
                        amount: item.amount,
                        block: item.block,
                        height: item.height,
                      }
                    }}
                            as={`/transactions/${item.key}`}>
                      {getKeyValue(item, columnKey)}
                    </Link>
                    : getKeyValue(item, columnKey)}
                </TableCell>
              )}
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
};

export default Transactions;
