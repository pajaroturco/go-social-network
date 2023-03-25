import React from "react";
import { Modal, ModalBody } from "react-bootstrap";
import LogoWhite from "../../../assets/png/logo-white.png";

import "./BasicModal.scss";

export default function BasicModal(props) {
  const { show, setShow, children } = props;
  return (
    <Modal
      className="basic-modal"
      show={show}
      onHide={() => setShow(false)}
      centered
      size="lg"
    >
      <Modal.Header>
        <Modal.Title>
          <img src={LogoWhite} alt="twitter" />
        </Modal.Title>
      </Modal.Header>
      <ModalBody>{children}</ModalBody>
    </Modal>
  );
}
