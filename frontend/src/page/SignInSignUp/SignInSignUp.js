import React, { useState } from "react";
import { Container, Row, Col, Button } from "react-bootstrap";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faSearch,
  faUsers,
  faComment,
} from "@fortawesome/free-solid-svg-icons";
import BasicModal from "../../components/Modal/BasicModal/BasicModal";
import LogoWhite from "../../assets/png/logo-white.png";
import LogoBlue from "../../assets/png/logo.png";

import "./SignInSignUp.scss";

export default function SignInSignUp() {
  const [showModal, setShowModal] = useState(true);
  const [contentModal, setcontentModal] = useState(null);

  return (
    <div>
      <Container className="signInSignUp" fluid>
        <Row>
          <LeftComponent />
          <RightComponent />
        </Row>
      </Container>
      <BasicModal show={showModal} setShow={setShowModal}>
        <div>
          <h2>Modal Content</h2>
        </div>
      </BasicModal>
    </div>
  );
}

function LeftComponent() {
  return (
    <Col className="signInSignUp__left" xs={6}>
      <img src={LogoBlue} alt="twittor" />
      <div>
        <h2>
          <FontAwesomeIcon icon={faSearch} />
          Sigue lo que te interesa
        </h2>
        <h2>
          <FontAwesomeIcon icon={faUsers} />
          Enterate de que esta hablando la gente
        </h2>
        <h2>
          <FontAwesomeIcon icon={faComment} />
          Unete a la conversacion
        </h2>
      </div>
    </Col>
  );
}

function RightComponent() {
  return (
    <Col className="signInSignUp__right" xs={6}>
      <div>
        <img src={LogoWhite} alt="twittor" />
        <h2>Mira lo que esta pasando en el mundo en este momento</h2>
        <h3>Unete a Twittor hoy mismo</h3>
        <Button variant="primary">Registrate</Button>
        <Button variant="outline-primary">Inicia Sesion</Button>
      </div>
    </Col>
  );
}
