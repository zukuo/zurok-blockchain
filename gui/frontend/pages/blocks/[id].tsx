import React, {useContext} from 'react';
import {useRouter} from "next/router";
import {Button} from "@nextui-org/button";
import {Card, CardBody} from "@nextui-org/card";

const BlockPage = () => {
  const router = useRouter()

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

  return (
    <>
      <div className="flex flex-col items-center font-mono">
        <div className="grid grid-cols-1 w-fit space-y-5 justify-center align-middle items-center px-10 py-5 font-mono">
          {renderProp("HASH", router.query.hash)}
          {renderProp("PREVIOUS HASH", router.query.prevhash)}
          {renderProp("HEIGHT", router.query.height)}
          {renderProp("TIMESTAMP", router.query.timestamp)}
          {renderProp("NUMBER ONCE", router.query.nonce)}
          {renderProp("PROOF OF WORK", router.query.pow)}
        </div>
        <Button onClick={() => router.back()} color={"success"}>Go Back</Button>
      </div>
    </>
  );
};

export default BlockPage;
