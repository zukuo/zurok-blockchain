import React, {useContext} from 'react';
import {Button} from "@nextui-org/button";
import {RxReload} from "react-icons/rx";
import {StartNode} from "../wailsjs/wailsjs/go/gui/App";
import {NodeContext} from "../pages";

const Refresh = () => {
  const node = useContext(NodeContext)!.node

  return (
    <div className="flex justify-center">
      <Button onClick={() => StartNode(node, "")} isIconOnly color="success" aria-label="Refresh" radius="full">
        <RxReload className="text-xl" />
      </Button>
    </div>
  );
};

export default Refresh;
