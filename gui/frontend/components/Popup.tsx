import React from 'react';
import {Button, Modal, ModalBody, ModalContent, ModalFooter, ModalHeader, useDisclosure} from "@nextui-org/react";

const Popup = (title: string, content: string, isOpen: boolean, onOpenChange: () => void) => {
  return (
    <Modal isOpen={isOpen} onOpenChange={onOpenChange} backdrop={"blur"}>
      <ModalContent>
        {(onClose) => (
          <>
            <ModalHeader className="flex flex-col gap-1">{title}</ModalHeader>
            <ModalBody>{content}</ModalBody>
            <ModalFooter>
              <Button color="danger" variant="light" onPress={onClose}>
                Close
              </Button>
            </ModalFooter>
          </>
        )}
      </ModalContent>
    </Modal>
  );
};

export default Popup;
