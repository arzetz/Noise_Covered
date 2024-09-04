import React from "react";
import "./HomePage.css"; // Импортируйте файл стилей
import { Container, Row, Col, Nav } from "react-bootstrap";

function HomePage() {
  return (
    <div className="background-container">
      <Container className="centered-container main-header">
        <Row>
          <Col lg={12} className="main-header">
            <span className="main">NOISE | COVE<span className="RED">RED</span></span>
          </Col>
          </Row>
      </Container>
      <Container
        className="centered-container"
      >
        <Row>
          <Col className="my-nav">
            <Nav>
              <Nav.Item>
                <Nav.Link href="/"><span className="menu">ГЛАВНАЯ</span></Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link href="/lol"><span className="menu">О НАС</span></Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link href="#"><span className="menu">КАТАЛОГ</span></Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link href="#"><span className="menu">КОНТАКТЫ</span></Nav.Link>
              </Nav.Item>
            </Nav>
          </Col>     
        </Row>
      </Container>
      <Container>
        <Row>
          <Col>
          <span className="phone-text">МАГАЗИН ВИНИЛОВЫХ ПЛАСТИНОК</span>
          </Col>
          </Row>
      </Container>

    </div>
  );
}

export default HomePage;
