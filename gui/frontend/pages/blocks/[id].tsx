import React, {useContext, useState} from 'react';
import Hero from "../../components/Hero";
import {useRouter} from "next/router";
import {Divider} from "@nextui-org/react";
import {Button} from "@nextui-org/button";
import {Card, CardBody} from "@nextui-org/card";
import {NodeContext} from "../index";

const BlockPage = () => {
  const router = useRouter()
  const node = useContext(NodeContext)?.node

  const renderProp = (title: string, content: any) => {
    return (
      <Card>
        <CardBody>
          <h1 className={"text-lg font-bold"}>{title}</h1>
          <span className={"text-success"}>{content}</span>
        </CardBody>
      </Card>
    )
  }

  const customDivider = () => {
    return <Divider className="my-3 w-[40px]"/>
  }

  return (
    <>
      <Hero/>
      <div className="flex flex-col items-center font-mono">
      <div className="grid grid-cols-1 w-fit space-y-5 justify-center align-middle items-center px-10 py-5 font-mono">
        {renderProp("HASH", router.query.hash)}
        {renderProp("PREVIOUS HASH", router.query.prevhash)}
        {renderProp("HEIGHT", router.query.height)}
        {renderProp("TIMESTAMP", router.query.timestamp)}
        {renderProp("NUMBER ONCE", router.query.nonce)}
        {renderProp("PROOF OF WORK", router.query.pow)}
      </div>
      <Button color={"success"}>Go Back</Button>
      </div>
    </>
  );
};

export default BlockPage;
