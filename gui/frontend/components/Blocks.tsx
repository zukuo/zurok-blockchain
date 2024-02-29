import React, {useContext, useEffect, useMemo, useState} from 'react';
import {NodeContext} from "../pages";
import {GetBlockInfos} from "../wailsjs/wailsjs/go/gui/App";
import {
  getKeyValue,
  Pagination,
  Spinner,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from "@nextui-org/react";
import Link from "next/link";

const Blocks = () => {
  const node = useContext(NodeContext)!.node

  let blocksArr: {
    key: number,
    hash: string,
    prevhash: string,
    height: number,
    timestamp: string,
    nonce: number,
    pow: boolean,
  }[] = []
  const [blocks, setBlocks] = useState(blocksArr)

  useEffect(() => {
    const fetchBlocks = () => {
      try {
        GetBlockInfos(node)
          .then(result => {
            setBlocks(result)
          })
          .catch(err => {
            console.log(err)
          })
      } catch (e) {
        console.error(e)
      }
    }
    fetchBlocks()
  }, [blocks])

  const loadingState =  blocks?.length === 0 ? "loading" : "idle"

  // Table Pagination
  const [page, setPage] = useState(1);
  const rowsPerPage = 10;
  const pages = Math.ceil(blocks.length / rowsPerPage);
  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;
    return blocks.slice(start, end);
  }, [page, blocks]);

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
          <TableColumn key="key">ID</TableColumn>
          <TableColumn key="hash">Block</TableColumn>
        </TableHeader>
        <TableBody items={items} loadingContent={<Spinner/>} loadingState={loadingState}>
          {(item) => (
            <TableRow key={item.key}>
              {(columnKey) => (
                <TableCell className="m-auto text-center">
                  {columnKey == "hash"
                    ? <Link className={"text-small text-foreground hover:text-success transition"} href={{
                      pathname: "/blocks/[id]",
                      query: {
                        id: item.key,
                        hash: item.hash,
                        prevhash: item.prevhash,
                        height: item.height,
                        timestamp: item.timestamp,
                        nonce: item.nonce,
                        pow: item.pow,
                      }}}
                      as={`/blocks/${item.key}`}>
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

export default Blocks;
