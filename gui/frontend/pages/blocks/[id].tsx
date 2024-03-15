import React from 'react';
import {useRouter} from "next/router";
import {Button} from "@nextui-org/button";
import {Card, CardBody} from "@nextui-org/card";
import {Tab, Tabs} from "@nextui-org/react";
import {AiOutlineTransaction} from "react-icons/ai";
import Transactions from "../../components/Transactions";
import {HiOutlineCube} from "react-icons/hi";

const BlockPage = () => {
  const router = useRouter()
  const height = parseInt(router.query.height as string)

  const infoLabel = () => {
    return <div className="flex items-center space-x-2">
      <HiOutlineCube className="text-lg"/>
      <span>Info</span>
    </div>
  }

  const transactionsLabel = () => {
    return <div className="flex items-center space-x-2">
      <AiOutlineTransaction className="text-md"/>
      <span>Transactions</span>
    </div>
  }

  const renderProp = (title: string, content: any) => {
    return (
      <Card>
        <CardBody>
          <h1 className={"text-medium font-bold"}>{title}</h1>
          <span className={"text-success text-small"}>{content}</span>
        </CardBody>
      </Card>
    )
  }

  const renderInfo = () => {
    return (
      <div className="flex flex-col items-center font-mono">
        <div className="grid grid-cols-1 w-fit space-y-5 justify-center align-middle items-center px-10 font-mono">
          {renderProp("HASH", router.query.hash)}
          {renderProp("PREVIOUS HASH", router.query.prevhash)}
          {renderProp("HEIGHT", router.query.height)}
          {renderProp("TIMESTAMP", router.query.timestamp)}
          {renderProp("NUMBER ONCE", router.query.nonce)}
          {renderProp("PROOF OF WORK", router.query.pow)}
        </div>
      </div>
    )
  }

  let tabs = [
    {
      id: "Info",
      label: infoLabel(),
      content: renderInfo(),
    },
    {
      id: "Transactions",
      label: transactionsLabel(),
      content: Transactions(height),
    },
  ]

  return (
    <>
      <div className="flex w-full flex-col px-10 py-5">
        <Tabs aria-label="Blocks Tabs"
              items={tabs}
              className={"justify-center"}
              color={"success"}
              variant={"light"}
              size={"lg"}>
          {(item) => (
            <Tab key={item.id} title={item.label}>
              {item.content}
            </Tab>
          )}
        </Tabs>
      </div>

      <div className="flex justify-center">
        <Button onClick={() => router.back()} color={"success"}>Go Back</Button>
      </div>
    </>

)
  ;
};

export default BlockPage;
