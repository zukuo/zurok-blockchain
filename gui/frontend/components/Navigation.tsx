import React, {useContext} from "react";
import { Tabs, Tab } from "@nextui-org/react";
import { MdWallet } from "react-icons/md";
import {GrCubes, GrSend} from "react-icons/gr";
import {TbPick} from "react-icons/tb";
import Wallet from "./Wallet";
import Send from "./Send";
import Blocks from "./Blocks";
import {NavContext} from "../pages";
import {LuHardDriveDownload} from "react-icons/lu";
import Receive from "./Receive";
import Mine from "./Mine";

export default function Navigation() {
    const navContext = useContext(NavContext)

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

    const blocksLabel = () => {
        return <div className="flex items-center space-x-2">
            <GrCubes className="text-md"/>
            <span>Blocks</span>
        </div>
    }

    const receiveLabel = () => {
        return <div className="flex items-center space-x-2">
            <LuHardDriveDownload className="text-lg"/>
            <span>Receive</span>
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
            id: "blocks",
            label: blocksLabel(),
            content: Blocks,
        },
        {
            id: "receive",
            label: receiveLabel(),
            content: Receive,
        },
        {
            id: "miner",
            label: minerLabel(),
            content: Mine,
        }
    ];

    return (
        <div className="flex w-full flex-col px-10 py-5">
            <Tabs aria-label="Dynamic tabs"
                  className="justify-center"
                  items={tabs}
                  variant={"light"}
                  selectedKey={navContext?.selectedTab}
                  onSelectionChange={(string) => navContext?.setSelectedTab(string)}
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
