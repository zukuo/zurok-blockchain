import React, {useContext, useEffect, useState} from 'react';
import {useRouter} from "next/router";
import {Button} from "@nextui-org/button";
import {Card, CardBody} from "@nextui-org/card";
import {NodeContext} from "../index";
import {gui} from "../../wailsjs/wailsjs/go/models";
import {GetTransactions} from "../../wailsjs/wailsjs/go/gui/App";
import {element} from "prop-types";

const TransactionPage = () => {
  const router = useRouter()

  const node = useContext(NodeContext)!.node

  let txsArr: gui.transactions[] = []
  const [transactions, setTransactions] = useState(txsArr)

  useEffect(() => {
    const fetchTransactions = () => {
      try {
        GetTransactions(node)
          .then(result => {
            setTransactions(result)
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

  const renderBigCard = (title: string, content: any) => {
    return (
      <Card>
        <CardBody>
          <h1 className={"text-medium font-bold pb-2"}>{title}</h1>
          <span className={"text-success text-small"}>{content}</span>
        </CardBody>
      </Card>
    )
  }

  const id = parseInt(router.query.id as string) - 1

  const colorWithItem = (title: string, content: any) => {
    return (
      <>
        <span className={"text-foreground text-small"}>{title}</span>
        <span className={"text-success text-small pb-1"}>{content}</span>
      </>
    )
  }

  const checker = (element: any) => {
    if (element == "") {
      return "N/A"
    } else {
      return element
    }
  }

  const renderInputs = () => {
    return (
      <Card>
        <CardBody>
          <h1 className={"text-medium font-bold pb-2"}>INPUTS</h1>
          {transactions[id]?.inputs.map((txin, index) => (
            <>
              <div className={(((index + 1) % 2 == 0) && (index != 0)) ? "pb-3" : undefined}></div>
              <Card key={index}>
                <CardBody>
                  <h1 className={"text-medium font-bold pb-1"}>INPUT {index}</h1>
                  {colorWithItem("Txid:", checker(txin.txid))}
                  {colorWithItem("Vout:", checker(txin.vout))}
                  {colorWithItem("Public Key:", checker(txin.pubkey))}
                  {colorWithItem("Signature:", checker(txin.signature))}
                </CardBody>
              </Card>
            </>
          ))}
        </CardBody>
      </Card>
    )
  }

  const renderOutputs = () => {
    return (
      <Card>
        <CardBody>
          <h1 className={"text-medium font-bold pb-2"}>OUTPUTS</h1>
          {transactions[id]?.outputs.map((txout, index) => (
            <>
              <div className={(((index + 1) % 2 == 0) && (index != 0)) ? "pb-3" : undefined}></div>
              <Card key={index}>
                <CardBody>
                  <h1 className={"text-medium font-bold pb-1"}>OUTPUT {index}</h1>
                  {colorWithItem("Value:", checker(txout.value))}
                  {colorWithItem("Public Key Hash:", checker(txout.pubkeyhash))}
                </CardBody>
              </Card>
            </>
          ))}
        </CardBody>
      </Card>
    )
  }

  return (
    <>
      <div className="flex flex-col items-center font-mono">
        <div className="grid grid-cols-1 space-y-5 max-w-3xl break-all justify-center align-middle items-center px-10 py-5 font-mono">
          {renderProp(`BLOCK ${router.query.height}`, router.query.block)}
          {renderProp("TRANSACTION ID", router.query.transaction)}
          {renderProp("AMOUNT", router.query.amount)}
          {renderInputs()}
          {renderOutputs()}
        </div>
        <Button onClick={() => router.back()} color={"success"}>Go Back</Button>
      </div>
    </>
  );
};

export default TransactionPage;
