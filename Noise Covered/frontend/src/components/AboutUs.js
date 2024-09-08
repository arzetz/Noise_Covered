import React from "react";
import "./AboutUs.css";
import { useState, useEffect } from "react";
import { Col, Container, Row } from "react-bootstrap";
import { Link, Router } from "react-router-dom";
function AboutUs() {
  const [isFirstEffectDone, setIsFirstEffectDone] = useState(false);
  const [displayedFirstText, setDisplayedFirstText] = useState("");
  const [isFading, setIsFading] = useState(false);
  const [isHidden, setIsHidden] = useState(false);
  const [time, setTime] = useState(50);
  const firstText = "К то мы?";
  const secondText = `В  2024 году был создан розничный магазин пластинок под названием NOISE COVERED.
Вдохновлённая идеей профессиональной записи каверов, наша команда
нашла большое количество артистов, кто мог бы заниматься созданием исполнений
и их ресейлом вместе с нами.

Данная площадка - место собрания уникальной коллекции творческих
людей, стать частью которой вы можете просто  приобретя интересующую вас
запись или выставив на продажу свой товар.
`;
  const thirdText = `\nХотите развиваться с нами?`;
  const fourthTextHome = `На главную /`;
  const fourthTextCatalogue = `Да!`;

  useEffect(() => {
    let currentIndex = 0;
    let interval;

    const delayTimeout = setTimeout(() => {
      interval = setInterval(() => {
        if (currentIndex < firstText.length - 1) {
          setDisplayedFirstText((prev) => prev + firstText[currentIndex]);
          currentIndex++;
        } else {
          clearInterval(interval);
          setTimeout(() => {
            setIsFading(true);
            setTimeout(() => {
              setIsFirstEffectDone(true);
              setIsHidden(true);
            }, 2000);
          }, 2000);
        }
      }, time);
    }, 1000);

    return () => {
      clearInterval(interval);
      clearTimeout(delayTimeout);
    };
  }, [time]);

  return (
    <div className="background-container_about_us">
      <Container
        className={`${
          isHidden ? "about_us_container" : "about_us_container_centered"
        }`}
      >
        <span
          className={`custom-font1 ${isFading ? "fade-out" : ""} ${
            isHidden ? "fade-out-hidden" : ""
          }`}
        >
          {displayedFirstText}
        </span>
        <span
          className={`custom-font2 ${isHidden ? "fade-in" : "not-displaying"}`}
          style={{ display: isHidden ? "block" : "none" }}
        >
          {secondText}
        </span>
        <span
          className={`custom-font3 ${isHidden ? "fade-in" : "not-displaying"}`}
          style={{ display: isHidden ? "block" : "none" }}
        >
          {thirdText}
        </span>
        <Container className="links_container d-flex justify-content-center">
        <div className={`background-image_da ${isHidden ? "fade-in" : "not-displaying"}`} style={{ display: isHidden ? "block" : "none" }}></div>
            <Link to="/" className="link">
              <span
                className={`span_container custom-font1 ${
                  isHidden ? "fade-in" : "not-displaying"
                }`}
                style={{ display: isHidden ? "block" : "none" }}
              >
                {fourthTextHome}
              </span>
            </Link>

          <Link to="/" className="link">
            <span
              className={`span_container custom-font1 ${
                isHidden ? "fade-in" : "not-displaying"
              }`}
              style={{ display: isHidden ? "block" : "none" }}
            >
              {fourthTextCatalogue}
            </span>
          </Link>

        </Container>
      </Container>
    </div>
  );
}
export default AboutUs;
