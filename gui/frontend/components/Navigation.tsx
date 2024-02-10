import React, {useState} from "react";
import { Tabs, Tab } from "@nextui-org/react";
import { MdWallet } from "react-icons/md";
import { GrSend } from "react-icons/gr";
import { AiOutlineTransaction } from "react-icons/ai";
import { TbPick } from "react-icons/tb";
import Wallet from "./Wallet";
import Send from "./Send";

export default function Navigation() {
    const [selected, setSelected] = useState<string | number>("wallet")

    const walletLabel = () => {
        return <div className="flex items-center space-x-2">
            <MdWallet className="text-lg"/>
            <span>Wallet</span>
        </div>
    }

    const sendLabel = () => {
        return <div className="flex items-center space-x-2">
            <GrSend className=""/>
            <span>Send</span>
        </div>
    }

    const transactionsLabel = () => {
        return <div className="flex items-center space-x-2">
            <AiOutlineTransaction className="text-lg"/>
            <span>Transactions</span>
        </div>
    }

    const minerLabel = () => {
        return <div className="flex items-center space-x-2">
            <TbPick className="text-lg"/>
            <span>Mine</span>
        </div>
    }

    let tabs = [
        {
            id: "wallet",
            label: walletLabel(),
            content: Wallet,
        },
        {
            id: "send",
            label: sendLabel(),
            content: Send,
        },
        {
            id: "transactions",
            label: transactionsLabel(),
            content: "Hello",
        },
        {
            id: "miner",
            label: minerLabel(),
            content: "Hello",
        }
    ];

    return (
        <div className="flex w-full flex-col px-10 py-5">
            <Tabs aria-label="Dynamic tabs"
                  className="justify-center"
                  items={tabs}
                  variant={"light"}
                  selectedKey={selected}
                  onSelectionChange={(string) => setSelected(string)}
                  size={"lg"}
                  color={"success"}>
                {(item) => (
                  <Tab key={item.id} title={item.label}>
                      {React.createElement(item.content)}
                  </Tab>
                )}
            </Tabs>
        </div>
    );
}
